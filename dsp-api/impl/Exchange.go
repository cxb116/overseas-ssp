package impl

//
//import (
//	"context"
//	"sync"
//	"time"
//
//	"github.com/cxb116/DSP/constant"
//	"github.com/cxb116/DSP/internal/utils"
//)
//
//func ExchangeHandler(reqCtx *RequestContext) BidResponse {
//	parentCtx := context.Background()
//	if reqCtx != nil && reqCtx.Context != nil {
//		parentCtx = reqCtx.Context
//	}
//	timeout := 660 * time.Millisecond
//	if reqCtx != nil && reqCtx.Timeout > 0 && reqCtx.Timeout < timeout {
//		timeout = reqCtx.Timeout
//	}
//	ctx, cancel := context.WithTimeout(parentCtx, timeout)
//	defer cancel()
//
//	dspSlotInfoArr := getDspList(reqCtx)
//	if len(dspSlotInfoArr) == 0 {
//		return BidResponse{Res: 0}
//	}
//
//	var wg sync.WaitGroup
//	respDspArr := make([]DspResult, len(dspSlotInfoArr))
//	respCh := make(chan DspResult, len(dspSlotInfoArr))
//
//	// 检索出有多少个需要并发请求的预算
//	for i, dsp := range dspSlotInfoArr {
//		wg.Add(1)
//		go func(dsp DspSlotInfo, index int) {
//			defer wg.Done()
//
//			resp := DspDispatchDocumentManager(ctx, dsp, reqCtx)
//			result := DspResult{
//				DspSlot: dsp,
//				BidRes:  resp,
//			}
//			respCh <- result
//			respDspArr[index] = result
//		}(dsp, i)
//	}
//
//	// 等待所有请求完成
//	go func() {
//		wg.Wait()
//		close(respCh)
//	}()
//
//	// 等待所有 goroutine 完成
//	wg.Wait()
//
//	return pickBest(respDspArr, reqCtx.BidRequest.SlotId)
//}
//
//// 做策略
//func pickBest(list []DspResult, SspSlotId int64) BidResponse {
//	if len(list) == 0 {
//		return BidResponse{Res: 0}
//	}
//
//	for i := range list {
//		if list[i].BidRes.Res == constant.REQ_CODE_SUC {
//
//			met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//				SspSlotId:   SspSlotId,                   // 媒体广告位
//				DspSlotCode: list[i].DspSlot.DspSlotCode, // 预算广告位
//				DspSlotId:   list[i].DspSlot.Id,          // 预算广告位id
//				//ReqCount:    1,                           // TODO 请求会丢啊
//				SspRes:  1,
//				EventTs: utils.Get10MinTime(time.Now()),
//			}
//			//ResponseMessageMetrics(met)
//			return list[i].BidRes
//		}
//	}
//
//	best := list[0]
//	//if len(list) > 1 {
//	//	GlobalCacheResponseMap.AddCache(SspSlotId, list[1:])
//	//}
//	met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//		SspSlotId:   SspSlotId,                // 媒体广告位
//		DspSlotCode: best.DspSlot.DspSlotCode, // 预算广告位
//		DspSlotId:   best.DspSlot.Id,          // 预算广告位id
//		ReqCount:    1,
//		EventTs:     utils.Get10MinTime(time.Now()),
//	}
//	//ResponseMessageMetrics(met)
//	return best.BidRes
//}
//
//func getDspList(ctx *RequestContext) []DspSlotInfo {
//	dspSlotInfoMaps := ctx.DspSlotManager.DspSlotInfoMaps[ctx.BidRequest.SlotId]
//	return dspSlotInfoMaps
//}
//
//type DspResult struct {
//	BidRes  BidResponse
//	DspSlot DspSlotInfo
//}
