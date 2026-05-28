package impl

//import (
//	"context"
//	"time"
//
//	"github.com/cxb116/DSP/constant"
//	"github.com/cxb116/DSP/internal/logger"
//	"github.com/cxb116/DSP/internal/utils"
//)
//
//// HandleBidRequestDirect 直连处理广告请求（不走 worker queue）
//func HandleBidRequestDirect(ctx context.Context, bidReq *BidRequest, timeout time.Duration) BidResponse {
//	if bidReq == nil {
//		return BidResponse{Res: constant.REQ_CODE_DSP_ERROR}
//	}
//
//	reqCtx := &RequestContext{
//		BidRequest:  bidReq,
//		BidResponse: &BidResponse{},
//		Timeout:     timeout,
//		Context:     ctx,
//	}
//
//	// 广告位ID 为空直接返回
//	if reqCtx.BidRequest.SlotId == 0 {
//		reqCtx.BidResponse.Res = constant.REQ_CODE_SLOT_ID
//		return BidResponse{Res: -1}
//	}
//
//	// 验证这个请求id 是否存在
//	sspSlotInfo := DspSlotInitMapsHandler.GetSspSlotInfo(reqCtx.BidRequest.SlotId)
//	if sspSlotInfo == nil {
//		reqCtx.BidResponse.Res = constant.REQ_CODE_SLOT_ID
//		return BidResponse{Res: -2}
//	}
//
//	// 请求去匹配预算
//	dspSlotIndexManager := DspSlotInitMapsHandler.MatchBudgetHandler(reqCtx.BidRequest)
//	if dspSlotIndexManager == nil || len(dspSlotIndexManager.DspSlotInfoMaps) == 0 {
//		logger.ErrorLog.Log().Msgf("流量广告位id没有匹配倒数据,请检查媒体广告位ID:%d", reqCtx.BidRequest.SlotId)
//		reqCtx.BidResponse.Res = constant.REQ_CODE_DSP_NOT_SLOT_ID
//		return BidResponse{Res: -2}
//	}
//
//	reqCtx.DspSlotManager = dspSlotIndexManager
//	reqCtx.DspSlotManager.SspSlotInfo = *sspSlotInfo
//
//	// 过滤不合格的数据
//	resCodeErr := ProcessHandler(reqCtx, dspSlotIndexManager)
//	if resCodeErr == constant.REQ_CODE_DEVICE_ERR {
//		logger.ErrorLog.Log().Msgf("请求参数异常,下游媒体ID: %d", reqCtx.BidRequest.SlotId)
//		reqCtx.BidResponse.Res = constant.REQ_CODE_DEVICE_ERR
//
//		met := DspRequestMetrics{
//			SspSlotId:     reqCtx.BidRequest.SlotId,
//			DspSlotCode:   reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][0].DspSlotCode,
//			DspSlotId:     reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][0].Id,
//			SspReq:        1,
//			SspReqDiscard: 1,
//			EventTs:       utils.Get10MinTime(time.Now()),
//		}
//		ResponseMessageMetrics(met)
//		return BidResponse{Res: resCodeErr}
//	}
//
//	//for i := range reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId] {
//	//	if i == 0 {
//	//		met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//	//			SspSlotId:   reqCtx.BidRequest.SlotId,
//	//			DspSlotCode: reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][i].DspSlotCode,
//	//			DspSlotId:   reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][i].Id,
//	//			SspReq:      1,
//	//			ReqCount:    1,
//	//			EventTs:     utils.Get10MinTime(time.Now()),
//	//		}
//	//		ResponseMessageMetrics(met) // 统计请求次数
//	//		continue
//	//	}
//	//	met := DspRequestMetrics{
//	//		SspSlotId:   reqCtx.BidRequest.SlotId,
//	//		DspSlotCode: reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][i].DspSlotCode,
//	//		DspSlotId:   reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][i].Id,
//	//		SspReq:      1,
//	//		EventTs:     utils.Get10MinTime(time.Now()),
//	//	}
//	//	ResponseMessageMetrics(met)
//	//}
//
//	localReqCtx := &RequestContext{
//		BidRequest:     reqCtx.BidRequest,
//		DspSlotManager: reqCtx.DspSlotManager,
//		Context:        ctx,
//	}
//
//	// 核心：使用 fanout 风格并发聚合
//	return DspDispatchDocumentManager(ctx, reqCtx.DspSlotManager.DspSlotInfoMaps[reqCtx.BidRequest.SlotId][0], localReqCtx)
//}
