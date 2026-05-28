package client

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/logger"
)

var (
	AsyncProducer sarama.AsyncProducer
	KafkaAdmin    sarama.ClusterAdmin
	KafkaConfig   *sarama.Config
)

const defaultKafkaSendTimeout = 100 * time.Millisecond

// KafkaStats Kafka 发送统计
type KafkaStats struct {
	TotalSent    int64 // 发送尝试总次数
	SuccessCount int64 // 成功入队/回调成功次数
	ErrorCount   int64 // 异步错误次数
	TimeoutCount int64 // 发送超时次数
}

var (
	Stats = &KafkaStats{}
)

func InitKafkaAsyncProducer(brokers []string) error {
	if len(brokers) == 0 {
		return errors.New("empty kafka brokers")
	}

	cfg := sarama.NewConfig()

	// ===== 生产性能配置 =====
	cfg.Producer.RequiredAcks = sarama.WaitForLocal // 只等待 leader ack，吞吐优先
	cfg.Producer.Compression = sarama.CompressionSnappy
	cfg.Producer.Flush.Frequency = resolveFlushFrequency()
	cfg.Producer.Flush.Bytes = 256 * 1024
	cfg.Producer.Flush.Messages = resolveFlushMessages()
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = global.EngineConfig.Kafka.ReturnSuccess
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.Retry.Max = 3
	cfg.Producer.Retry.Backoff = 100 * time.Millisecond
	cfg.ChannelBufferSize = 100000
	cfg.Metadata.RefreshFrequency = 30 * time.Second
	cfg.Version = sarama.V2_8_0_0

	KafkaConfig = cfg
	ResetKafkaStats()

	var err error
	AsyncProducer, err = sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		return err
	}

	KafkaAdmin, err = sarama.NewClusterAdmin(brokers, cfg)
	if err != nil {
		if AsyncProducer != nil {
			_ = AsyncProducer.Close()
			AsyncProducer = nil
		}
		return err
	}

	// 异步错误监听：必须消费，避免阻塞内部通道
	go func() {
		for err := range AsyncProducer.Errors() {
			atomic.AddInt64(&Stats.ErrorCount, 1)
			logger.ErrorLog.Error().Msgf("Kafka async error: %v", err)
		}
	}()

	// 成功回调只有在显式开启时才消费并统计，避免无意义开销
	if cfg.Producer.Return.Successes {
		go func() {
			for range AsyncProducer.Successes() {
				atomic.AddInt64(&Stats.SuccessCount, 1)
			}
		}()
	}

	return nil
}

func EnsureTopic(topic string, partitions int32, replicas int16) error {
	topics, err := KafkaAdmin.ListTopics()
	if err != nil {
		return err
	}

	if _, exists := topics[topic]; exists {
		return nil
	}

	detail := &sarama.TopicDetail{
		NumPartitions:     partitions,
		ReplicationFactor: replicas,
	}

	err = KafkaAdmin.CreateTopic(topic, detail, false)
	if err != nil {
		return err
	}

	return nil
}

// InitTopics 初始化所需 topic（在程序启动时调用一次）
func InitTopics() error {
	if KafkaAdmin == nil {
		logger.ErrorLog.Error().Msgf("KafkaAdmin not initialized, skipping topic initialization")
		return nil
	}

	topics := []struct {
		name       string
		partitions int32
		replicas   int16
	}{
		{
			name:       global.EngineConfig.Kafka.TopicReqData,
			partitions: 3,
			replicas:   1,
		},
		{
			name:       global.EngineConfig.Kafka.TopicResData,
			partitions: 3,
			replicas:   1,
		},
	}

	for _, t := range topics {
		if t.name == "" {
			continue
		}

		if err := EnsureTopic(t.name, t.partitions, t.replicas); err != nil {
			logger.ErrorLog.Error().Msgf("Failed to ensure topic %s: %v", t.name, err)
			return err
		}
		logger.Log.Log().Msgf("Kafka topic ensured: %s (partitions=%d, replicas=%d)", t.name, t.partitions, t.replicas)
	}

	return nil
}

func KafkaAsyncSend(topic, key string, msg []byte) error {
	return kafkaSend(topic, key, msg)
}

func KafkaSyncSend(topic, key string, msg []byte) error {
	// 兼容旧接口：仍走异步高吞吐链路，不阻塞主流程
	return kafkaSend(topic, key, msg)
}

func kafkaSend(topic, key string, msg []byte) error {
	if AsyncProducer == nil {
		atomic.AddInt64(&Stats.ErrorCount, 1)
		return fmt.Errorf("kafka producer not initialized")
	}

	atomic.AddInt64(&Stats.TotalSent, 1)

	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}

	select {
	case AsyncProducer.Input() <- message:
		// 未开启成功回调时，以入队成功作为成功统计，减少通道开销
		if KafkaConfig != nil && !KafkaConfig.Producer.Return.Successes {
			atomic.AddInt64(&Stats.SuccessCount, 1)
		}
		return nil
	default:
	}

	timer := time.NewTimer(defaultKafkaSendTimeout)
	defer timer.Stop()
	select {
	case AsyncProducer.Input() <- message:
		if KafkaConfig != nil && !KafkaConfig.Producer.Return.Successes {
			atomic.AddInt64(&Stats.SuccessCount, 1)
		}
		return nil
	case <-timer.C:
		atomic.AddInt64(&Stats.TimeoutCount, 1)
		return nil
	}
}

// GetProducerStatus 获取 Producer 状态
func GetProducerStatus() map[string]interface{} {
	if AsyncProducer == nil {
		return map[string]interface{}{
			"status": "not_initialized",
		}
	}

	return map[string]interface{}{
		"status": "running",
		"config": map[string]interface{}{
			"channelBufferSize": KafkaConfig.ChannelBufferSize,
			"flushFrequency":    KafkaConfig.Producer.Flush.Frequency,
			"flushMessages":     KafkaConfig.Producer.Flush.Messages,
			"flushBytes":        KafkaConfig.Producer.Flush.Bytes,
			"requiredAcks":      KafkaConfig.Producer.RequiredAcks,
			"returnSuccesses":   KafkaConfig.Producer.Return.Successes,
		},
	}
}

// GetKafkaStats 获取 Kafka 统计信息
func GetKafkaStats() *KafkaStats {
	return &KafkaStats{
		TotalSent:    atomic.LoadInt64(&Stats.TotalSent),
		SuccessCount: atomic.LoadInt64(&Stats.SuccessCount),
		ErrorCount:   atomic.LoadInt64(&Stats.ErrorCount),
		TimeoutCount: atomic.LoadInt64(&Stats.TimeoutCount),
	}
}

// LogKafkaStats 打印 Kafka 统计信息到日志
func LogKafkaStats() {
	stats := GetKafkaStats()
	if stats.TotalSent <= 0 {
		return
	}

	successRate := float64(stats.SuccessCount) / float64(stats.TotalSent) * 100
	logger.BusinessLog.Info().Msgf(
		"Kafka stats: total=%d success=%d error=%d timeout=%d successRate=%.2f%%",
		stats.TotalSent, stats.SuccessCount, stats.ErrorCount, stats.TimeoutCount, successRate,
	)
}

// ResetKafkaStats 重置 Kafka 统计信息
func ResetKafkaStats() {
	atomic.StoreInt64(&Stats.TotalSent, 0)
	atomic.StoreInt64(&Stats.SuccessCount, 0)
	atomic.StoreInt64(&Stats.ErrorCount, 0)
	atomic.StoreInt64(&Stats.TimeoutCount, 0)
}

func resolveFlushFrequency() time.Duration {
	if global.EngineConfig.Kafka.Frequency > 0 {
		return global.EngineConfig.Kafka.Frequency
	}
	return 100 * time.Millisecond
}

func resolveFlushMessages() int {
	if global.EngineConfig.Kafka.Messages > 0 {
		return global.EngineConfig.Kafka.Messages
	}
	return 500
}
