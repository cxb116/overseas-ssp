package impl

import (
	"context"

	"github.com/cxb116/DSP/constant"
)

// 请求初级校验，过滤不合法数据
func RequestCheckLevel1(bidRequest BidRequest) BidResponse {
	// 验证广告位ID 是否存在
	sspSlotInfo := GetSspSlotInfo(bidRequest.SlotId)
	if sspSlotInfo == nil {
		bidResponse := BidResponse{
			Res: constant.REQ_CODE_SLOT_ID,
		}
		return bidResponse
	}

	return BidResponse{Res: 1}
}

func RequestCheckLevel2(bidRequest BidRequest) BidResponse {

	return BidResponse{Res: 1}
}

func BidRequestHandleDirect(bidRequest BidRequest) BidResponse {
	res := RequestCheckLevel1(bidRequest)
	if res.Res < 1 {
		return res
	}

	// 通过请求获取配置预算数据
	ContextLoopDspEventMap := DspSlotGlobalData.MatchBudgetHandler(&bidRequest)
	if ContextLoopDspEventMap == nil {
		return BidResponse{Res: -1}
	}

	if ContextLoopDspEventMap == nil {
		return BidResponse{Res: -1}
	}

	// 第二次校验
	//level2 := RequestCheckLevel2(bidReques)
	//if level2.Res < 0 {
	//	return BidResponse{Res: -1}
	//}
	//
	//reqContext := GetRequestContext()
	//defer PutRequestContext(reqContext)
	//reqContext.BidRequest = &bidRequest
	//response := ExchangeHandlers(reqContext)

	//return BidResponse{Res: 0}

	//reqContext.BidRequest = &bidRequest
	// 往协程池里面加入了
	//if !GlobalWorkerChannelHandler.SendRequestToTaskQueue(reqContext) {
	//	return BidResponse{Res: -1}
	//}

	//waitForWinner(reqContext)
	//return BidResponse{Res: 0}
	return BidResponse{}
}

func ExchangeHandlers(reqCtx *RequestContext) BidResponse {
	//if len(reqCtx.KfHandler.DspSlotIds) == 0 {
	//	return BidResponse{Res: -1}
	//}

	//ctx, cancel := context.WithTimeout(reqCtx.Context, 99660*time.Millisecond)
	//defer cancel()

	//results := make([]CountHandler, 0, len(reqCtx.KfHandler.DspSlotIds))
	//
	//for _, dsp := range reqCtx.KfHandler.DspSlotIds {
	//	wg.Add(1)
	//	go func(dsp int64) {
	//		defer wg.Done()
	//		defer func() {
	//			if r := recover(); r != nil {
	//			}
	//		}()
	//
	//		if !acquireDSPToken(ctx, globalDSPLimiter) {
	//			return
	//		}
	//		defer releaseDSPToken(globalDSPLimiter)
	//
	//		dspSlotInfo := GetDspSlotInfo(dsp)
	//		if dspSlotInfo != nil {
	//			resp := DspDispatchDocumentManager(ctx, *dspSlotInfo, reqCtx)
	//			mu.Lock()
	//			results = append(results, CountHandler{BidResponse: resp})
	//			mu.Unlock()
	//		}
	//	}(dsp)
	//}
	//
	//done := make(chan struct{})
	//go func() {
	//	wg.Wait()
	//	close(done)
	//}()
	//
	//select {
	//case <-done:
	//case <-ctx.Done():
	//	cancel() // 取消 context，让未开始的 goroutine 快速退出
	//}

	return BidResponse{}
}

type CountHandler struct {
	BidResponse BidResponse
	reqCtx      RequestContext
}

// 做策略
func pickBest(nums []CountHandler) BidResponse {
	if len(nums) == 0 {
		return BidResponse{Res: 0}
	}

	if len(nums) == 1 {
		return nums[0].BidResponse
	}

	return BidResponse{Res: 10000}
}

// 全局 DSP in-flight 限流器（进程级）
// 含义：任意时刻，整个进程最多允许 2000 个并发 DSP 外呼在执行。
// 作用：防止高峰时 goroutine、socket、下游连接数同时失控。
var globalDSPLimiter = make(chan struct{}, 2000)

func acquireDSPToken(ctx context.Context, sem chan struct{}) bool {
	select {
	case sem <- struct{}{}:
		return true
	case <-ctx.Done():
		return false
	}
}

// releaseDSPToken 释放并发槽位，必须与 acquire 成对使用。
func releaseDSPToken(sem chan struct{}) {
	<-sem
}
