package impl

import (
	"github.com/cxb116/DSP/internal/logger"
	"github.com/panjf2000/ants/v2"
)

var GlobalWorkerChannelHandler *AntsWorkerPool

func InitGlobalWorkerChannelHandler(workerPoolSize int, maxWorkerTaskLen int) {
	GlobalWorkerChannelHandler = NewAntsWorkerPool(workerPoolSize, maxWorkerTaskLen)
}

// AntsWorkerPool 基于 ants 的高性能协程池
type AntsWorkerPool struct {
	pool *ants.Pool
}

// NewAntsWorkerPool 创建 ants 协程池
// workerPoolSize: 协程数量，建议为 CPU 核心数 * 2
// maxWorkerTaskLen: 任务队列大小，0 表示无限制
func NewAntsWorkerPool(workerPoolSize int, maxWorkerTaskLen int) *AntsWorkerPool {
	// 创建 ants 协程池
	pool, err := ants.NewPool(
		workerPoolSize,
		ants.WithPreAlloc(true),                     // 预分配协程
		ants.WithMaxBlockingTasks(maxWorkerTaskLen), // 最大阻塞任务数
		ants.WithNonblocking(false),                 // 允许阻塞
		ants.WithPanicHandler(func(i interface{}) {
			logger.ErrorLog.Error().Msgf("worker panic: %v", i)
		}),
		ants.WithLogger(nil), // 禁用默认日志
	)
	if err != nil {
		logger.Log.Fatal().Msgf("failed to create ants pool: %v", err)
	}

	return &AntsWorkerPool{
		pool: pool,
	}
}

// SendRequestToTaskQueue 发送任务到协程池
// 返回 true 表示成功，false 表示队列满
func (p *AntsWorkerPool) SendRequestToTaskQueue(ctx *RequestContext) bool {
	// 提交任务到协程池
	err := p.pool.Submit(func() {
		p.doRequestDispatch(ctx, 0)
	})

	if err != nil {
		logger.ErrorLog.Warn().Msgf("submit task failed: %v", err)
		return false
	}

	return true
}

// doRequestDispatch 处理任务
func (p *AntsWorkerPool) doRequestDispatch(reqContext *RequestContext, workerId int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Error().Msgf("doRequestDispatch err: workerId %d, err: %v", workerId, err)
		}
	}()

	// nil 检查
	if reqContext == nil || reqContext.BidRequest == nil {
		logger.Log.Warn().Msg("invalid request context")
		return
	}

	// TODO: 实际业务逻辑处理
	logger.Log.Info().Msgf("processing request: %v", reqContext.BidRequest)

	// 使用对象池获取响应
	response := GetBidResponse()
	response.Res = 111
	p.transResponseChan(response, reqContext)

	// 归还对象池
	PutBidResponse(response)
}

// transResponseChan 发送响应
func (p *AntsWorkerPool) transResponseChan(bidResponse *BidResponse, reqContext *RequestContext) {
	if bidResponse == nil || reqContext == nil {
		return
	}
	select {
	case reqContext.ResponseChan <- bidResponse:
	case <-reqContext.Context.Done():
	}
}

// GetStats 获取协程池统计信息
func (p *AntsWorkerPool) GetStats() (running int, waiting int, capacity int, free int) {
	return p.pool.Running(), p.pool.Waiting(), p.pool.Cap(), p.pool.Free()
}

// Release 释放协程池
func (p *AntsWorkerPool) Release() {
	p.pool.Release()
}

// Restart 重启协程池
func (p *AntsWorkerPool) Restart() {
	p.pool.Reboot()
}

// Tune 动态调整协程池大小
func (p *AntsWorkerPool) Tune(size int) {
	p.pool.Tune(size)
}
