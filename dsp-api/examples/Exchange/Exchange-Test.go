package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// BidderRequest 表示一个竞价方请求
// BidderName：竞价方名称
// ReqCount：该 bidder 需要发起多少个 HTTP 请求
// ShouldPanic：用于模拟异常
type BidderRequest struct {
	BidderName  string
	ReqCount    int
	ShouldPanic bool
}

// EngineConfig：高并发场景下的关键控制项
type EngineConfig struct {
	MaxBidderParallel     int           // 单次请求里，最多并行执行多少个 bidder
	GlobalHTTPMaxInFlight int           // 全进程（示例里是全请求）HTTP 并发上限
	PerBidderMaxInFlight  int           // 单 bidder 并发上限
	PerBidderWorker       int           // 单 bidder 内部 worker 数
	BidderTimeout         time.Duration // 单 bidder 预算
}

// bidResponseWrapper：用于聚合单个 bidder 的返回结果
type bidResponseWrapper struct {
	bidder          string   // bidder 名称
	adapterSeatBids []string // 成功返回的 bid 列表
	errs            []error  // 错误列表
}

// httpCallInfo：模拟一次 HTTP 请求返回
type httpCallInfo struct {
	body string // 返回内容
	err  error  // 错误
}

// BidderAdapter：一个 bidder 的适配器
type BidderAdapter struct {
	name     string        // bidder 名称
	connPool chan struct{} // 用 channel 模拟连接池（信号量）
}

// getAllBids：全局入口，负责并发执行所有 bidder 请求
// 关键点：
// 1. bidder fan-out 但有上限（MaxBidderParallel）
// 2. bidder deadline（BidderTimeout）
// 3. 父级 ctx 到期立即提前收敛，不再硬等所有 bidder
func getAllBids(
	ctx context.Context, // 全局 context，用于超时控制
	bidderRequests []BidderRequest, // 所有 bidder 请求
	adapterMap map[string]*BidderAdapter, // bidder -> adapter 映射
	cfg EngineConfig,
	globalHTTPLimiter chan struct{},
) map[string][]string {

	// 最终结果：bidder -> bids
	adapterBids := make(map[string][]string, len(bidderRequests))

	// 用 buffered channel 收集所有 bidder 结果
	// buffer = bidder 数量，避免 goroutine 阻塞
	chBids := make(chan *bidResponseWrapper, len(bidderRequests))

	maxBidderParallel := cfg.MaxBidderParallel
	if maxBidderParallel <= 0 {
		maxBidderParallel = len(bidderRequests)
	}
	bidderLimiter := make(chan struct{}, maxBidderParallel)

	// 遍历所有 bidder request
	for _, bidder := range bidderRequests {

		// recoverSafely 包装 bidder 执行函数（防 panic）
		bidderRunner := recoverSafely(
			bidderRequests, // 所有 bidder（用于打印 debug）
			func(br BidderRequest) {
				// 单次请求里的 bidder fan-out 限流，防止高峰并发扩散
				if !acquireSlot(ctx, bidderLimiter) {
					chBids <- &bidResponseWrapper{
						bidder: br.BidderName,
						errs:   []error{ctx.Err()},
					}
					return
				}
				defer releaseSlot(bidderLimiter)

				// 模拟 panic（测试容错能力）
				if br.ShouldPanic {
					panic("mock panic in bidder")
				}

				adapter, ok := adapterMap[br.BidderName]
				if !ok {
					chBids <- &bidResponseWrapper{
						bidder: br.BidderName,
						errs:   []error{fmt.Errorf("bidder adapter not found: %s", br.BidderName)},
					}
					return
				}

				// bidder 预算，避免慢 bidder 拖垮整体 RT
				bidderCtx := ctx
				cancel := func() {}
				if cfg.BidderTimeout > 0 {
					bidderCtx, cancel = context.WithTimeout(ctx, cfg.BidderTimeout)
				}
				defer cancel()

				// 调用具体 adapter 发起 bid 请求
				seatBids, errs := adapter.requestBid(bidderCtx, br, cfg.PerBidderWorker, globalHTTPLimiter)

				// 将结果写入 channel
				chBids <- &bidResponseWrapper{
					bidder:          br.BidderName,
					adapterSeatBids: seatBids,
					errs:            errs,
				}
			},
			chBids, // 结果 channel
		)

		// 每个 bidder 启一个 goroutine 执行
		go bidderRunner(bidder)
	}

	// 收集返回结果，到父级 deadline 就提前返回
	pending := len(bidderRequests)
	for pending > 0 {
		select {
		case brw := <-chBids:
			pending--

			// 如果有错误，打印 warning
			if len(brw.errs) > 0 {
				fmt.Printf("[WARN] bidder=%s errors=%v\n", brw.bidder, brw.errs)
			}

			// 如果有 bid 数据，写入最终结果 map
			if len(brw.adapterSeatBids) > 0 {
				adapterBids[brw.bidder] =
					append(adapterBids[brw.bidder], brw.adapterSeatBids...)
			}

		case <-ctx.Done():
			fmt.Printf("[WARN] getAllBids early return: %v\n", ctx.Err())
			return adapterBids
		}
	}

	// 返回所有 bidder 的 bid 结果
	return adapterBids
}

// recoverSafely：对 bidder 执行函数做 panic 捕获
func recoverSafely(
	bidderRequests []BidderRequest, // 全部 bidder（用于 debug）
	inner func(BidderRequest), // 实际执行逻辑
	chBids chan *bidResponseWrapper, // 结果 channel
) func(BidderRequest) {

	// 返回一个包装函数（闭包）
	return func(bidderRequest BidderRequest) {

		// defer 捕获 panic
		defer func() {

			// 如果发生 panic
			if r := recover(); r != nil {

				// 收集所有 bidder 名称（用于 debug 输出）
				allBidders := make([]string, 0, len(bidderRequests))
				for _, b := range bidderRequests {
					allBidders = append(allBidders, b.BidderName)
				}

				// 打印 panic 信息
				fmt.Printf("[PANIC] bidder=%s recovered=%v all=%s\n",
					bidderRequest.BidderName,
					r,
					strings.Join(allBidders, ","))

				// 向 channel 写 fallback 结果（避免阻塞主流程）
				chBids <- &bidResponseWrapper{
					bidder: bidderRequest.BidderName,
					errs:   []error{fmt.Errorf("panic recovered: %v", r)},
				}
			}
		}()

		// 执行真正逻辑
		inner(bidderRequest)
	}
}

// requestBid：单个 bidder 的请求逻辑（内部并发 HTTP）
func (bidder *BidderAdapter) requestBid(
	ctx context.Context,
	bidderRequest BidderRequest,
	maxWorkers int,
	globalHTTPLimiter chan struct{},
) ([]string, []error) {
	if bidderRequest.ReqCount <= 0 {
		return nil, nil
	}

	// 构造请求列表
	reqData := make([]int, bidderRequest.ReqCount)
	for i := 0; i < bidderRequest.ReqCount; i++ {
		reqData[i] = i + 1
	}

	workerNum := maxWorkers
	if workerNum <= 0 {
		workerNum = 1
	}
	if workerNum > len(reqData) {
		workerNum = len(reqData)
	}

	// 固定 worker + 任务队列，避免 ReqCount 大时创建过多 goroutine
	reqChannel := make(chan int, len(reqData))
	responseChannel := make(chan *httpCallInfo, len(reqData))
	for _, reqID := range reqData {
		reqChannel <- reqID
	}
	close(reqChannel)

	for i := 0; i < workerNum; i++ {
		go func() {
			for reqID := range reqChannel {
				httpInfo := bidder.doRequest(ctx, reqID, globalHTTPLimiter)
				select {
				case responseChannel <- httpInfo:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// 聚合结果
	bids := make([]string, 0, len(reqData))
	errs := make([]error, 0, 1)

	// 收集所有请求结果（固定次数）
	for i := 0; i < len(reqData); i++ {
		var httpInfo *httpCallInfo
		select {
		case httpInfo = <-responseChannel:
		case <-ctx.Done():
			errs = append(errs, ctx.Err())
			return bids, errs
		}

		// 如果请求失败
		if httpInfo.err != nil {
			errs = append(errs, httpInfo.err)
			continue
		}

		// 成功则加入 bids
		bids = append(bids, httpInfo.body)
	}

	return bids, errs
}

// doRequest：模拟 HTTP 请求 + 连接池控制
func (bidder *BidderAdapter) doRequest(
	ctx context.Context,
	reqID int,
	globalHTTPLimiter chan struct{},
) *httpCallInfo {
	// 固定顺序获取 slot，避免不同 goroutine 间死锁
	if !acquireSlot(ctx, globalHTTPLimiter) {
		return &httpCallInfo{err: ctx.Err()}
	}
	defer releaseSlot(globalHTTPLimiter)

	if !acquireSlot(ctx, bidder.connPool) {
		return &httpCallInfo{err: ctx.Err()}
	}
	defer releaseSlot(bidder.connPool)

	// 模拟随机延迟（30~150ms）
	latency := time.Duration(30+rand.Intn(120)) * time.Millisecond

	// 模拟 HTTP 等待或 context cancel
	select {

	// 正常返回
	case <-time.After(latency):
		return &httpCallInfo{
			body: fmt.Sprintf("%s-bid-%d", bidder.name, reqID),
		}

	// 超时取消
	case <-ctx.Done():
		return &httpCallInfo{err: ctx.Err()}
	}
}

func acquireSlot(ctx context.Context, sem chan struct{}) bool {
	if sem == nil {
		return true
	}
	select {
	case sem <- struct{}{}:
		return true
	case <-ctx.Done():
		return false
	}
}

func releaseSlot(sem chan struct{}) {
	if sem == nil {
		return
	}
	<-sem
}

// main：入口函数，模拟整个竞价流程
func main() {

	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 高性能骨架配置
	cfg := EngineConfig{
		MaxBidderParallel:     3,
		GlobalHTTPMaxInFlight: 8,
		PerBidderMaxInFlight:  2,
		PerBidderWorker:       4,
		BidderTimeout:         220 * time.Millisecond,
	}

	globalHTTPLimiter := make(chan struct{}, cfg.GlobalHTTPMaxInFlight)

	// 初始化 bidder adapter（每个 2 并发限制）
	adapterMap := map[string]*BidderAdapter{
		"appnexus": {name: "appnexus", connPool: make(chan struct{}, cfg.PerBidderMaxInFlight)},
		"rubicon":  {name: "rubicon", connPool: make(chan struct{}, cfg.PerBidderMaxInFlight)},
		"openx":    {name: "openx", connPool: make(chan struct{}, cfg.PerBidderMaxInFlight)},
	}

	// 构造 bidder 请求列表
	bidderRequests := []BidderRequest{
		{BidderName: "appnexus", ReqCount: 3},
		{BidderName: "rubicon", ReqCount: 1}, // sync fast path
		{BidderName: "openx", ReqCount: 4},
		{BidderName: "openx", ReqCount: 2, ShouldPanic: true}, // panic 测试
	}

	// 设置全局超时 400ms
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	// 记录开始时间
	start := time.Now()

	// 执行所有 bidder 请求
	adapterBids := getAllBids(ctx, bidderRequests, adapterMap, cfg, globalHTTPLimiter)

	// 输出耗时
	fmt.Printf("elapsed=%s\n", time.Since(start))

	// 输出最终结果
	fmt.Printf("final bids=%v\n", adapterBids)
}
