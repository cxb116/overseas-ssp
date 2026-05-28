package impl

//
//import (
//	"time"
//
//	"github.com/cxb116/DSP/constant"
//	"github.com/cxb116/DSP/internal/logger"
//	"github.com/cxb116/DSP/internal/utils"
//)
//
//var GlobalWorkerChannelHandler *WorkerChannelHandler
//
//func InitGlobalWorkerChannelHandler(workerPoolSize int, maxWorkerTaskLen int) {
//	ensureMetricsDispatcher()
//	GlobalWorkerChannelHandler = NewWorkerChannelHandler(workerPoolSize, maxWorkerTaskLen)
//	GlobalWorkerChannelHandler.StartWorkerPool()
//}
//
//type WorkerChannelHandler struct {
//	WorkerPoolSize   int                  // worker 池子大小
//	MaxWorkerTaskLen int                  // worker 管道队列长度
//	TaskQueue        chan *RequestContext // 共享任务队列
//}
//
//// NewWorkerChannelHandler 初始化工作池
//func NewWorkerChannelHandler(workerPoolSize int, maxWorkerTaskLen int) *WorkerChannelHandler {
//	return &WorkerChannelHandler{
//		WorkerPoolSize:   workerPoolSize,
//		MaxWorkerTaskLen: maxWorkerTaskLen,
//		TaskQueue:        make(chan *RequestContext, maxWorkerTaskLen),
//	}
//}
//
//// StartWorkerPool 启动工作池
//func (this *WorkerChannelHandler) StartWorkerPool() {
//	for i := 0; i < this.WorkerPoolSize; i++ {
//		go this.OnWorkerHandle(i)
//	}
//}
//
//func (this *WorkerChannelHandler) SendRequestToTaskQueue(ctx *RequestContext) bool {
//	if ctx == nil {
//		return false
//	}
//	ctx.EnqueueAt = time.Now()
//
//	select {
//	case <-ctx.Context.Done():
//		return false
//	case this.TaskQueue <- ctx:
//		return true
//	default:
//		logger.ErrorLog.Error().Msgf("worker queue is full, request dropped before dispatch")
//		return false
//	}
//}
//
//func (this *WorkerChannelHandler) OnWorkerHandle(workerId int) {
//	for reqContext := range this.TaskQueue {
//		if reqContext == nil || reqContext.BidRequest == nil {
//			logger.ErrorLog.Error().Msgf("reqContext is nil or request payload is missing")
//			continue
//		}
//		this.doRequestDispatch(reqContext, workerId)
//	}
//}
//
//// doRequestDispatch 执行读取资源位的请求
//func (this *WorkerChannelHandler) doRequestDispatch(reqContext *RequestContext, workerId int) {
//	defer func() {
//		if err := recover(); err != nil {
//			logger.ErrorLog.Log().Msgf("doRequestDispatch err: workerId %d, err: %v", workerId, err)
//			//if reqContext != nil && reqContext.BidResponse != nil {
//			//	req := BidResponse{Res: 0}
//			//	select {
//			//	case reqContext.ResponseChan <- &req:
//			//	case <-reqContext.Context.Done():
//			//
//			//	default:
//			//	}
//			//}
//		}
//	}()
//
//	if reqContext.Context.Err() != nil {
//		return
//	}
//
//	// 广告位ID 为空直接返回
//	if reqContext.BidRequest.SlotId == 0 {
//		reqContext.BidResponse.Res = constant.REQ_CODE_SLOT_ID
//		this.transResponseChan(BidResponse{Res: -1}, reqContext)
//		return
//	}
//
//	//验证这个请求id 是否存在
//	sspSlotInfo := DspSlotInitMapsHandler.GetSspSlotInfo(reqContext.BidRequest.SlotId)
//	if sspSlotInfo == nil {
//		reqContext.BidResponse.Res = constant.REQ_CODE_SLOT_ID
//		this.transResponseChan(BidResponse{Res: -2}, reqContext)
//		return
//	}
//
//	// 请求去匹配预算
//	dspSlotIndexManager := DspSlotInitMapsHandler.MatchBudgetHandler(reqContext.BidRequest)
//	if reqContext.Context.Err() != nil {
//		return
//	}
//	if dspSlotIndexManager == nil || len(dspSlotIndexManager.DspSlotInfoMaps) == 0 {
//		logger.ErrorLog.Log().Msgf("流量广告位id没有匹配倒数据,请检查媒体广告位ID:%d", reqContext.BidRequest.SlotId)
//
//		reqContext.BidResponse.Res = constant.REQ_CODE_DSP_NOT_SLOT_ID
//		this.transResponseChan(BidResponse{Res: -2}, reqContext) // TODO 没有匹配到预算,不需要统计
//		return
//	}
//	reqContext.DspSlotManager = dspSlotIndexManager
//
//	//过滤不合格的数据
//	res_code_err := ProcessHandler(reqContext, dspSlotIndexManager)
//	if reqContext.Context.Err() != nil {
//		return
//	}
//	if res_code_err == constant.REQ_CODE_DEVICE_ERR {
//		logger.ErrorLog.Log().Msgf("请求参数异常,下游媒体ID: %d", reqContext.BidRequest.SlotId)
//		reqContext.BidResponse.Res = constant.REQ_CODE_DEVICE_ERR
//
//		met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//			SspSlotId:     reqContext.BidRequest.SlotId,                                                           // 媒体广告位
//			DspSlotCode:   reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][0].DspSlotCode, // 预算广告位
//			DspSlotId:     reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][0].Id,          // 预算广告位id
//			SspReq:        1,
//			SspReqDiscard: 1,
//			EventTs:       utils.Get10MinTime(time.Now()),
//		}
//		ResponseMessageMetrics(met)
//		this.transResponseChan(BidResponse{Res: res_code_err}, reqContext)
//		return
//	}
//
//	for i := range reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId] {
//		if i == 0 {
//			met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//				SspSlotId:   reqContext.BidRequest.SlotId,                                                           // 媒体广告位
//				DspSlotCode: reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][i].DspSlotCode, // 预算广告位
//				DspSlotId:   reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][i].Id,          // 预算广告位id
//				SspReq:      1,
//				ReqCount:    1,
//				EventTs:     utils.Get10MinTime(time.Now()),
//			}
//			ResponseMessageMetrics(met) // 统计请求次数
//			continue
//		}
//
//		met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//			SspSlotId:   reqContext.BidRequest.SlotId,                                                           // 媒体广告位
//			DspSlotCode: reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][i].DspSlotCode, // 预算广告位
//			DspSlotId:   reqContext.DspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][i].Id,          // 预算广告位id
//			SspReq:      1,
//			EventTs:     utils.Get10MinTime(time.Now()),
//		}
//		ResponseMessageMetrics(met) // 统计请求次数
//	}
//
//	bidResponse := DspMatchExchange(reqContext)
//
//	this.transResponseChan(bidResponse, reqContext)
//}
//
//// 管道响应数据
//func (this *WorkerChannelHandler) transResponseChan(bidResponse BidResponse, ctx *RequestContext) {
//	select {
//	case ctx.ResponseChan <- &bidResponse:
//		//if bidResponse.Res == constant.REQ_CODE_SUC { //成功, 有填充,响应下游媒体
//		//
//		//	met.SspRes = 1
//		//	ResponseMessageMetrics(met)
//		//} else if bidResponse.Res == constant.REQ_CODE_DSP_TIMEOUT { // 预算超时 || 上游异常
//		//	logger.ErrorLog.Log().Msgf("上游预算超时,媒体ID:%d", ctx.BidRequest.SlotId)
//		//	met.DspResDiscard = 1
//		//	ResponseMessageMetrics(met)
//		//} else if bidResponse.Res == constant.REQ_CODE_DSP_ERROR { //dsp 异常
//		//	logger.ErrorLog.Log().Msgf("上游预算异常,媒体ID:%d", ctx.BidRequest.SlotId)
//		//
//		//} else if bidResponse.Res == constant.REQ_CODE_DSP_NOT_FLOORPRICE { //  低于底价
//		//	met.DspResDiscard = 1
//		//	ResponseMessageMetrics(met)
//		//} else if bidResponse.Res == constant.REQ_CODE_DEVICE_ERR || bidResponse.Res == constant.REQ_CODE_DSP_NOT_SLOT_ID { // 流量参数错误
//		//	ResponseMessageMetrics(met)
//		//}
//	case <-ctx.Context.Done():
//	}
//}
//
//func ResponseMessageMetrics(met DspRequestMetrics) {
//	if met.SspSlotId == 0 {
//		return
//	}
//	dispatchMetrics(met)
//}
