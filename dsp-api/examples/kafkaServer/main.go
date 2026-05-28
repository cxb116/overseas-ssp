package main

import (
	"github.com/cxb116/DSP/internal/client"
	"github.com/cxb116/DSP/internal/logger"
	"time"
)

func main() {

	client.InitKafkaAsyncProducer([]string{"115.190.91.254:9092"})

	for {

		time.Sleep(10 * time.Second)
		client.KafkaSyncSend("TimeTopicRequest", "key", []byte("111"))
		logger.BusinessLog.Info().Msgf("kafka 启动")
	}

}
