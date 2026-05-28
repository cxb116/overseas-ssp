package impl

import (
	"context"
	"sync"
	"time"
)

var requestContextPool = sync.Pool{
	New: func() any {
		return &RequestContext{}
	},
}

type RequestContext struct {
	BidRequest  *BidRequest
	BidResponse *BidResponse
	// 绑定数据统计
	DspRequestMetrics []*DspRequestMetrics // 统计kafka 请求次数
	Timeout           time.Duration

	ResponseChan chan *BidResponse
	Cancel       context.CancelFunc
	Context      context.Context
}

func GetRequestContext() *RequestContext {
	ctx := requestContextPool.Get().(*RequestContext)

	ctx.BidRequest = &BidRequest{}
	ctx.BidResponse = &BidResponse{}

	ctx.ResponseChan = make(chan *BidResponse, 1)

	ctx.Timeout = 4 * time.Second

	ctx.Context, ctx.Cancel =
		context.WithTimeout(context.Background(), ctx.Timeout)

	return ctx
}

func PutRequestContext(ctx *RequestContext) {
	if ctx == nil {
		return
	}

	if ctx.Cancel != nil {
		ctx.Cancel()
	}

	ctx.BidRequest = nil
	ctx.BidResponse = nil
	ctx.DspRequestMetrics = nil
	ctx.Context = nil
	ctx.Cancel = nil

	select {
	case <-ctx.ResponseChan:
	default:
	}

	requestContextPool.Put(ctx)
}
