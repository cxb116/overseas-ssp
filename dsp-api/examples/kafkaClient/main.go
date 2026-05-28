package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// Worker 数，越大吞吐越高，建议 = CPU 核心 * 4
const WorkerCount = 32

type ConsumerGroupHandler struct {
	workerPool chan *sarama.ConsumerMessage
}

func NewConsumerGroupHandler() *ConsumerGroupHandler {
	h := &ConsumerGroupHandler{
		workerPool: make(chan *sarama.ConsumerMessage, 100000), // 高缓冲避免阻塞
	}
	h.startWorkers()
	return h
}

// 启动多 worker 并发处理消息
func (h *ConsumerGroupHandler) startWorkers() {
	for i := 0; i < WorkerCount; i++ {
		go func(id int) {
			for msg := range h.workerPool {
				processKafkaMessage(id, msg)
			}
		}(i)
	}
}

// Kafka 回调：新 session 开启
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Kafka 回调：session 结束
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// 核心方法：每条消息来了这里
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.workerPool <- msg          // 投递给 worker
		session.MarkMessage(msg, "") // 标记位移（异步，不阻塞）
	}
	return nil
}

// 实际处理消息逻辑（你业务要改这里）
func processKafkaMessage(workerID int, msg *sarama.ConsumerMessage) {
	// TODO: 替换成你的业务逻辑
	fmt.Printf("[Worker %d] Topic=%s Partition=%d Offset=%d Value=%s\n",
		workerID, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
}

// --------------- 启动 ConsumerGroup ------------------

func StartKafkaConsumerGroup(brokers []string, groupID string, topics []string) error {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.MaxVersion
	cfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	cfg.Consumer.Group.Session.Timeout = 30 * 1000 * 1000_000 // 30s
	cfg.ChannelBufferSize = 40960

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		return err
	}

	handler := NewConsumerGroupHandler()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if err := consumerGroup.Consume(ctx, topics, handler); err != nil {
				log.Println("Error from consumer:", err)
			}
		}
	}()

	// 监听退出信号
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		cancel()
	}()

	wg.Wait()
	return consumerGroup.Close()
}

func main() {
	brokers := []string{
		"115.190.91.254:9092",
		"117.72.67.186:9092",
	}
	StartKafkaConsumerGroup(
		brokers,
		"total-consumer-group",
		[]string{"topicReqTotalDay", "topicSspData", "topicDspData"},
	)

	StartKafkaConsumerGroup(
		brokers,
		"req_budget-consumer-group",
		[]string{"topicReqTotalDay", "topicSspData", "topicDspData"},
	)

	StartKafkaConsumerGroup(
		brokers,
		"traffic-consumer-group",
		[]string{"topicReqTotalDay", "topicSspData", "topicDspData"},
	)

}
