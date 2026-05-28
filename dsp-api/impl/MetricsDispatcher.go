package impl

import (
	"encoding/json"
	"runtime"
	"sync"

	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/client"
	"github.com/cxb116/DSP/internal/logger"
)

const metricsQueueSize = 65536

var (
	metricsDispatchOnce sync.Once
	metricsQueue        chan DspRequestMetrics
)

func ensureMetricsDispatcher() {
	metricsDispatchOnce.Do(func() {
		metricsQueue = make(chan DspRequestMetrics, metricsQueueSize)
		workerCount := runtime.GOMAXPROCS(0)
		if workerCount < 4 {
			workerCount = 4
		}
		for i := 0; i < workerCount; i++ {
			go metricsDispatchWorker()
		}
	})
}

func metricsDispatchWorker() {
	for met := range metricsQueue {
		sendMetricsSync(met)
	}
}

func sendMetricsSync(met DspRequestMetrics) {
	jsonBytes, err := json.Marshal(met)
	if err != nil {
		return
	}
	if err := client.KafkaSyncSend(global.EngineConfig.Kafka.TopicResData, "dsp-consumer", jsonBytes); err != nil {
		logger.ErrorLog.Error().Msgf("kafka metrics send failed: %v", err)
	}
}

func dispatchMetrics(met DspRequestMetrics) {
	ensureMetricsDispatcher()

	select {
	case metricsQueue <- met:
		return
	default:
	}

	// 队列满时同步兜底，避免改变统计语义。
	sendMetricsSync(met)
}
