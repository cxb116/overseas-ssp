package impl

import (
	"testing"
	"time"

	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/config"
	"github.com/cxb116/DSP/internal/logger"
)

func TestSendRequestToTaskQueueReturnsFalseWhenQueueIsFull(t *testing.T) {
	global.EngineConfig.RequestLink = time.Second
	global.EngineConfig.Log = config.Log{LogLevel: "error"}
	logger.InitDefaultLogger()

	handler := NewWorkerChannelHandler(0, 1)

	first := GetRequestContext()
	defer PutRequestContext(first)
	second := GetRequestContext()
	defer PutRequestContext(second)

	if ok := handler.SendRequestToTaskQueue(first); !ok {
		t.Fatal("expected first request to be accepted into queue")
	}

	if ok := handler.SendRequestToTaskQueue(second); ok {
		t.Fatal("expected second request to be rejected when queue is full")
	}
}
