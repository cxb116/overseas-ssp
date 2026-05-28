package impl

//import (
//	"context"
//	"fmt"
//	"sync/atomic"
//	"time"
//
//	"github.com/cxb116/DSP/constant"
//	"github.com/cxb116/DSP/internal/logger"
//	"github.com/cxb116/DSP/internal/utils"
//)
//
//// 全局 DSP in-flight 限流器（进程级）
//// 含义：任意时刻，整个进程最多允许 2000 个并发 DSP 外呼在执行。
//// 作用：防止高峰时 goroutine、socket、下游连接数同时失控。
////var globalDSPLimiter = make(chan struct{}, 20000)
//
//type dspMatchResult struct {
//	resp        BidResponse
//	dspSlotInfo DspSlotInfo
//	err         error
//}
//
//type dspDispatchStage uint32
//
//const (
//	dspStageInit dspDispatchStage = iota
//	dspStageWaitingPerReqLimiter
//	dspStageAcquiredPerReqLimiter
//	dspStageWaitingGlobalLimiter
//	dspStageAcquiredGlobalLimiter
//	dspStageDispatching
//	dspStageFinished
//)
//
//func dspStageName(stage dspDispatchStage) string {
//	switch stage {
//	case dspStageWaitingPerReqLimiter:
//		return "waiting_per_req_limiter"
//	case dspStageAcquiredPerReqLimiter:
//		return "acquired_per_req_limiter"
//	case dspStageWaitingGlobalLimiter:
//		return "waiting_global_limiter"
//	case dspStageAcquiredGlobalLimiter:
//		return "acquired_global_limiter"
//	case dspStageDispatching:
//		return "dispatching"
//	case dspStageFinished:
//		return "finished"
//	default:
//		return "init"
//	}
//}
//
//type dspDispatchTrace struct {
//	dspSlotID   int64
//	dspSlotCode string
//	stage       atomic.Uint32
//	startAt     time.Time
//	waitPerReq  atomic.Int64
//	waitGlobal  atomic.Int64
//	dispatchDur atomic.Int64
//}
//
//// DspMatchExchangeHighPerf 使用 Exchange-Test.go 的并发模型处理单次广告请求：
//// 1) fan-out：对命中的多个 DSP 并发发起请求，降低整体等待时间。
//// 2) 分层限流：单请求 fan-out 上限 + 全局 in-flight 上限，防止并发扩散。
//// 3) deadline 收敛：请求到期立即返回当前最优结果，不阻塞主流程。
//// 4) 容错隔离：单个 DSP panic 不影响其他 DSP，也不影响本次请求返回。
//func DspMatchExchangeHighPerf(reqCtx *RequestContext) BidResponse {
//	// 基础参数校验，避免空指针导致整个请求异常。
//	if reqCtx == nil || reqCtx.BidRequest == nil || reqCtx.DspSlotManager == nil {
//		return BidResponse{Res: constant.REQ_CODE_DSP_ERROR}
//	}
//
//	slotID := reqCtx.BidRequest.SlotId
//	// 从预算匹配结果中取出本次请求要并发访问的 DSP 列表。
//	dspSlots := reqCtx.DspSlotManager.DspSlotInfoMaps[slotID]
//	if len(dspSlots) == 0 {
//		return BidResponse{Res: constant.REQ_CODE_DSP_NOT_SLOT_ID}
//	}
//
//	// 单请求 fan-out 上限（默认 6）。
//	// 目的：避免某些广告位绑定过多 DSP 时，一次请求创建过多外呼并发。
//	maxFanout := 100
//	if len(dspSlots) < maxFanout {
//		maxFanout = len(dspSlots)
//	}
//	if maxFanout <= 0 { // 不要出现0个
//		maxFanout = 1
//	}
//
//	// 单请求局部限流器：限制本次请求内部并发。
//	perReqLimiter := make(chan struct{}, maxFanout)
//
//	// 结果通道用 buffer，容量等于 DSP 数量，避免子 goroutine 在回写结果时阻塞。
//	resultCh := make(chan dspMatchResult, len(dspSlots))
//	traces := make([]*dspDispatchTrace, 0, len(dspSlots))
//
//	// fan-out：每个 DSP 启一个 goroutine，请求并行执行。
//	for _, dspSlot := range dspSlots {
//		oneSlot := dspSlotg
//		trace := &dspDispatchTrace{
//			dspSlotID:   oneSlot.Id,
//			dspSlotCode: oneSlot.DspSlotCode,
//			startAt:     time.Now(),
//		}
//		trace.stage.Store(uint32(dspStageInit))
//		traces = append(traces, trace)
//		go func(slot DspSlotInfo) {
//			// 兜底保护：任何单 DSP 执行 panic 都转为错误结果回传，不影响主聚合流程。
//			defer func() {
//				if r := recover(); r != nil {
//					trace.stage.Store(uint32(dspStageFinished))
//					resultCh <- dspMatchResult{
//						resp: BidResponse{Res: constant.REQ_CODE_DSP_ERROR},
//						err:  fmt.Errorf("panic recovered: %v", r),
//					}
//				}
//			}()
//
//			// 第 1 层限流：先抢“单请求并发槽位”。
//			// 若请求已超时，acquire 会失败并快速返回，避免无意义外呼。
//			trace.stage.Store(uint32(dspStageWaitingPerReqLimiter))
//			perReqWaitStart := time.Now()
//			if !acquireDSPToken(reqCtx.Context, perReqLimiter) {
//				trace.waitPerReq.Store(int64(time.Since(perReqWaitStart)))
//				trace.stage.Store(uint32(dspStageFinished))
//				resultCh <- dspMatchResult{
//					resp: BidResponse{Res: constant.REQ_CODE_DSP_TIMEOUT},
//					err:  fmt.Errorf("acquire per-request limiter failed: %w", reqCtx.Context.Err()),
//				}
//				return
//			}
//			perReqWait := time.Since(perReqWaitStart)
//			trace.waitPerReq.Store(int64(perReqWait))
//			trace.stage.Store(uint32(dspStageAcquiredPerReqLimiter))
//			defer releaseDSPToken(perReqLimiter)
//			if perReqWait > 5*time.Millisecond {
//				logger.Log.Info().Msgf(
//					"dsp limiter wait diagnostics, slotID=%d dspSlotId=%d dspSlotCode=%s stage=per_request wait=%s",
//					slotID,
//					slot.Id,
//					slot.DspSlotCode,
//					perReqWait.String(),
//				)
//			}
//
//			// 第 2 层限流：再抢“全局并发槽位”。
//			// 固定顺序（先 per-request 再 global）可避免不同路径间锁顺序不一致问题。
//			globalWaitStart := time.Now()
//			if !acquireDSPToken(reqCtx.Context, globalDSPLimiter) {
//				trace.waitGlobal.Store(int64(time.Since(globalWaitStart)))
//				trace.stage.Store(uint32(dspStageFinished))
//				resultCh <- dspMatchResult{
//					resp: BidResponse{Res: constant.REQ_CODE_DSP_TIMEOUT},
//					err:  fmt.Errorf("acquire global limiter failed: %w", reqCtx.Context.Err()),
//				}
//				return
//			}
//			defer releaseDSPToken(globalDSPLimiter)
//			if globalWait > 5*time.Millisecond {
//				logger.Log.Info().Msgf(
//					"dsp limiter wait diagnostics, slotID=%d dspSlotId=%d dspSlotCode=%s stage=global wait=%s",
//					slotID,
//					slot.Id,
//					slot.DspSlotCode,
//					globalWait.String(),
//				)
//			}
//
//			// 单 DSP 子超时预算。
//			// 子预算不应大于父请求超时，确保不会晚于主请求生命周期。
//			perDSPTimeout := 700 * time.Millisecond
//			if reqCtx.Timeout > 0 && reqCtx.Timeout < perDSPTimeout {
//				perDSPTimeout = reqCtx.Timeout
//			}
//
//			// 每个 DSP 使用独立子 context，便于精准超时控制与取消传播。
//			dspCtx, cancel := context.WithTimeout(reqCtx.Context, perDSPTimeout)
//			defer cancel()
//
//			// 给每个 DSP 创建轻量局部 RequestContext。
//			// 共享只读请求数据，隔离 context，避免并发污染。
//			localReqCtx := &RequestContext{
//				BidRequest:     reqCtx.BidRequest,
//				DspSlotManager: reqCtx.DspSlotManager,
//				Context:        dspCtx,
//			}
//
//
//			resp := DspDispatchDocumentManager(dspCtx, slot, localReqCtx)
//			dispatchDur := time.Since(dispatchStart)
//			trace.dispatchDur.Store(int64(dispatchDur))
//			trace.stage.Store(uint32(dspStageFinished))
//			if dispatchDur > 10*time.Millisecond {
//				logger.Log.Info().Msgf(
//					"dsp dispatch diagnostics, slotID=%d dspSlotId=%d dspSlotCode=%s res=%d dispatch=%s total=%s wait_per_req=%s wait_global=%s",
//					slotID,
//					slot.Id,
//					slot.DspSlotCode,
//					resp.Res,
//					dispatchDur.String(),
//					time.Since(trace.startAt).String(),
//					time.Duration(trace.waitPerReq.Load()).String(),
//					time.Duration(trace.waitGlobal.Load()).String(),
//				)
//			}
//			resultCh <- dspMatchResult{resp: resp, dspSlotInfo: slot}
//
//		}(oneSlot)
//	}
//
//	// 聚合阶段：按“最高价格优先”挑选最佳返回。
//	bestArrResp := []dspMatchResult{}
//	pending := len(dspSlots)
//
//	// 固定收敛循环：
//	// - 正常路径：收满所有 DSP 结果；
//	// - 超时路径：父 ctx 到期立即返回当前已收集的最优结果。
//	for pending > 0 {
//		select {
//		case r := <-resultCh:
//			pending--
//			if r.err != nil {
//				logger.ErrorLog.Warn().Msgf("dsp fanout error, slotID=%d err=%v", slotID, r.err)
//			}
//			bestArrResp = append(bestArrResp, r)
//			//// 只把成功填充
//			//if r.resp.Res == constant.REQ_CODE_SUC {
//			//	met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//			//		SspSlotId:   r.dspSlotInfo.DspLaunch.SspSlotId, // 媒体广告位
//			//		DspSlotCode: r.dspSlotInfo.DspSlotCode,         // 预算广告位
//			//		DspSlotId:   r.dspSlotInfo.Id,                  // 预算广告位id
//			//		SspRes:      1,
//			//		EventTs:     utils.Get10MinTime(time.Now()),
//			//	}
//			//	ResponseMessageMetrics(met)
//			//
//			//	return r.resp
//			//}
//			//return BidResponse{Res: constant.REQ_CODE_DSP_NOT_AD}
//
//		case <-reqCtx.Context.Done():
//			// deadline 到达：优先返回已拿到的最优成功结果；
//			// 若还没有成功填充，则返回超时码。
//
//			if bestArrResp == nil || len(bestArrResp) == 0 {
//				stageCounter := make(map[string]int)
//				traceDetails := make([]string, 0, len(traces))
//				for _, trace := range traces {
//					stage := dspStageName(dspDispatchStage(trace.stage.Load()))
//					stageCounter[stage]++
//					traceDetails = append(traceDetails, fmt.Sprintf(
//						"dspSlotId=%d dspSlotCode=%s stage=%s elapsed=%s waitPerReq=%s waitGlobal=%s dispatch=%s",
//						trace.dspSlotID,
//						trace.dspSlotCode,
//						stage,
//						time.Since(trace.startAt).String(),
//						time.Duration(trace.waitPerReq.Load()).String(),
//						time.Duration(trace.waitGlobal.Load()).String(),
//						time.Duration(trace.dispatchDur.Load()).String(),
//					))
//				}
//				logger.ErrorLog.Log().Msgf("dsp fanout error 超时立即终止请求, slotID=%d", slotID)
//				logger.Log.Info().Msgf(
//					"dsp fanout timeout diagnostics, slotID=%d pending=%d total=%d stages=%v traces=%v err=%v",
//					slotID,
//					pending,
//					len(dspSlots),
//					stageCounter,
//					traceDetails,
//					reqCtx.Context.Err(),
//				)
//				return BidResponse{Res: constant.REQ_CODE_DSP_NOT_AD}
//			} else {
//				stageCounter := make(map[string]int)
//				traceDetails := make([]string, 0, len(traces))
//				for _, trace := range traces {
//					stage := dspStageName(dspDispatchStage(trace.stage.Load()))
//					stageCounter[stage]++
//					traceDetails = append(traceDetails, fmt.Sprintf(
//						"dspSlotId=%d dspSlotCode=%s stage=%s elapsed=%s waitPerReq=%s waitGlobal=%s dispatch=%s",
//						trace.dspSlotID,
//						trace.dspSlotCode,
//						stage,
//						time.Since(trace.startAt).String(),
//						time.Duration(trace.waitPerReq.Load()).String(),
//						time.Duration(trace.waitGlobal.Load()).String(),
//						time.Duration(trace.dispatchDur.Load()).String(),
//					))
//				}
//				logger.Log.Info().Msgf(
//					"dsp fanout partial timeout diagnostics, slotID=%d pending=%d total=%d collected=%d stages=%v traces=%v err=%v",
//					slotID,
//					pending,
//					len(dspSlots),
//					len(bestArrResp),
//					stageCounter,
//					traceDetails,
//					reqCtx.Context.Err(),
//				)
//				for i := range bestArrResp {
//					if bestArrResp[i].resp.Res == constant.REQ_CODE_SUC {
//						if bestArrResp[i].resp.Res == constant.REQ_CODE_SUC {
//							met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//								SspSlotId:   bestArrResp[i].dspSlotInfo.DspLaunch.SspSlotId, // 媒体广告位
//								DspSlotCode: bestArrResp[i].dspSlotInfo.DspSlotCode,         // 预算广告位
//								DspSlotId:   bestArrResp[i].dspSlotInfo.Id,                  // 预算广告位id
//								SspRes:      1,
//								EventTs:     utils.Get10MinTime(time.Now()),
//							}
//							ResponseMessageMetrics(met)
//
//							return bestArrResp[i].resp
//						}
//					}
//					return BidResponse{Res: constant.REQ_CODE_DSP_NOT_AD}
//				}
//			}
//		}
//	}
//
//	// 全部 DSP 都返回后：
//	// - 有成功填充：返回最优；
//	// - 无成功填充：返回默认 nil 结果。
//	for _, resultRes := range bestArrResp {
//		if resultRes.resp.Res == constant.REQ_CODE_SUC {
//			met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//				SspSlotId:   resultRes.dspSlotInfo.DspLaunch.SspSlotId, // 媒体广告位
//				DspSlotCode: resultRes.dspSlotInfo.DspSlotCode,         // 预算广告位
//				DspSlotId:   resultRes.dspSlotInfo.Id,                  // 预算广告位id
//				SspRes:      1,
//				EventTs:     utils.Get10MinTime(time.Now()),
//			}
//			ResponseMessageMetrics(met)
//			return resultRes.resp
//		}
//	}
//	return BidResponse{Res: constant.REQ_CODE_DSP_NOT_AD}
//}
//
//// acquireDSPToken 通过 channel 信号量获取并发槽位。
//// 语义：拿到槽位返回 true；若请求已取消/超时返回 false。
////func acquireDSPToken(ctx context.Context, sem chan struct{}) bool {
////	select {
////	case sem <- struct{}{}:
////		return true
////	case <-ctx.Done():
////		return false
////	}
////}
////
////// releaseDSPToken 释放并发槽位，必须与 acquire 成对使用。
////func releaseDSPToken(sem chan struct{}) {
////	<-sem
////}
