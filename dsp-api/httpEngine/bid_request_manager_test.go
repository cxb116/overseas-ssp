package httpEngine

import (
	"context"
	"testing"
	"time"

	"github.com/cxb116/DSP/constant"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/impl"
	"github.com/cxb116/DSP/internal/config"
	"github.com/cxb116/DSP/internal/logger"
)

func TestWaitForWinnerSetsTimeoutCode(t *testing.T) {
	global.EngineConfig.Log = config.Log{LogLevel: "error"}
	logger.InitDefaultLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	reqCtx := &impl.RequestContext{
		BidResponse:  &impl.BidResponse{},
		ResponseChan: make(chan *impl.BidResponse, 1),
		Context:      ctx,
		Cancel:       cancel,
	}

	time.Sleep(20 * time.Millisecond)

	res := waitForWinner(reqCtx)
	if res == nil {
		t.Fatal("expected timeout response")
	}
	if res.Res != constant.REQ_CODE_DSP_TIMEOUT {
		t.Fatalf("expected timeout code %d, got %d", constant.REQ_CODE_DSP_TIMEOUT, res.Res)
	}
}
