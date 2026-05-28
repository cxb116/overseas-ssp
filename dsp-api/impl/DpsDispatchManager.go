package impl

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/cxb116/DSP/constant"
	"github.com/cxb116/DSP/internal/logger"
	"github.com/cxb116/DSP/internal/utils"

	"github.com/cxb116/DSP/global"
)

//func DspMatchExchange(reqCtx *RequestContext) BidResponse {
//	bidRequest := ExchangeHandler(reqCtx)
//	return bidRequest
//}

// DspDispatchDocumentManager 请求资源位获取广告信息
func DspDispatchDocumentManager(ctx context.Context, dspSlotInfo DspSlotInfo, reqCtx *RequestContext) BidResponse {
	registryHandler := GetDspDoc(reqCtx, dspSlotInfo)
	if registryHandler == nil {
		return BidResponse{Res: constant.REQ_CODE_DSP_NOT_MATCH}
	}
	registryHandler.SetDspRequestContext(reqCtx, &dspSlotInfo) // 设置对象上下文
	//StatisticsDspRequestHandler(*reqCtx.BidRequest, dspSlotInfo, 1, 0, 0)
	response := registryHandler.DspRequestHandler()

	if response.Res == constant.REQ_CODE_SUC {
		imsUrl := CreateImsReportTalkMetrics(ctx, reqCtx, &dspSlotInfo)
		clkUrl := CreateClkReportTalkMetrics(ctx, reqCtx, &dspSlotInfo)
		response.Ad.Click = append(response.Ad.Click, clkUrl)
		response.Ad.Imp = append(response.Ad.Imp, imsUrl)
		//StatisticsDspRequestHandler(*reqCtx.BidRequest, dspSlotInfo, 0, 1, 0)
	} else { // 排除正常的都需要去吧没有填充的，丢去的都要统计一下
		//统计次数
		if response.Res == constant.REQ_CODE_DSP_TIMEOUT || response.Res == constant.REQ_CODE_DSP_NOT_FLOORPRICE {
			//StatisticsDspRequestHandler(*reqCtx.BidRequest, dspSlotInfo, 0, 0, 1)
		} else {
			//StatisticsDspRequestHandler(*reqCtx.BidRequest, dspSlotInfo, 0, 0, 0)
		}
	}
	return response
}

//func StatisticsDspRequestHandler(bidRequest BidRequest, dspSlotInfo DspSlotInfo, dspReq, dspRes, dspResDisCard int) {
//	met := DspRequestMetrics{ // 唯一ID，用于链路追踪
//		SspSlotId:     bidRequest.SlotId,       // 媒体广告位
//		DspSlotCode:   dspSlotInfo.DspSlotCode, // 预算广告位
//		DspSlotId:     dspSlotInfo.Id,          // 预算广告位id
//		DspRes:        dspRes,
//		DspReq:        dspReq,
//		DspResDiscard: dspResDisCard,
//		EventTs:       utils.Get10MinTime(time.Now()),
//	}
//	ResponseMessageMetrics(met)
//}

type DspHandlerCreator func(request *RequestContext) DspHandler

type DspHandler interface {

	// 设置预算reqContext
	SetDspRequestContext(reqCtx *RequestContext, dspSlotInfo *DspSlotInfo)

	GetDspCode() int64

	ReturnSelf() DspHandler
	//// 请求转换
	DspRequestHandler() BidResponse
	//发送请求
	//SendDspRequest() error
	// dsp响应转换
	//DspResponseHandler() error
	// 获取预算 DspCode

}

var dspRegistryMaps = make(map[int64]DspHandlerCreator) // map 不会写入所有不用枷锁

func DspRegister(dspCodeIndex int64, creator DspHandlerCreator) {
	if _, exist := dspRegistryMaps[dspCodeIndex]; !exist {
		dspRegistryMaps[dspCodeIndex] = creator
	}
}

// 获取预算对接文档
func GetDspDoc(reqContext *RequestContext, info DspSlotInfo) DspHandler {
	if info.CompanyId != 0 {
		company := GetDspCompany(info.CompanyId)
		if creator, exist := dspRegistryMaps[company.DspCode]; exist {
			return creator(reqContext)
		}
	}

	return nil
}

// 组装上报链接
func CreateImsReportTalkMetrics(ctx context.Context, reqCon *RequestContext, dspSlotInfo *DspSlotInfo) string {
	if ctx != nil && ctx.Err() != nil {
		return ""
	}
	var CaidStr string
	if len(reqCon.BidRequest.Device.Caids) > 0 {
		CaidStr = reqCon.BidRequest.Device.Caids[0].Caid
	}

	metricsLink := ImsReportTalkMetrics{
		SspSlotId:    reqCon.BidRequest.SlotId,
		DspSlotId:    dspSlotInfo.Id,
		DspSlotCode:  dspSlotInfo.DspSlotCode,
		Ip:           reqCon.BidRequest.Ip,
		Ua:           reqCon.BidRequest.Device.Ua,
		DspPayType:   dspSlotInfo.DspPayType,   //  上游结算方式
		AuctionPrice: dspSlotInfo.AuctionPrice, //  上游成交价
		//SspPayType:   reqCon.DspSlotManager.SspSlotInfo.SspPayType, //  下游结算方式
		BidPrice:  dspSlotInfo.BidPrice,
		Caid:      CaidStr,
		Idfa:      reqCon.BidRequest.Device.Idfa,
		Oaid:      reqCon.BidRequest.Device.Oaid,
		AndroidId: reqCon.BidRequest.Device.Aid,
		Imei:      reqCon.BidRequest.Device.Imei,
		Os:        reqCon.BidRequest.Device.Plt,
		AuctionTs: utils.TimeNowUnix(),
		EventTs:   utils.Get10MinTime(time.Now()),
		//Nurls:     reqCon.BidResponse.Seat[i].Nurls,
		//Lurls:     reqCon.BidResponse.Seat[i].Lurls,
	}

	marshal, err := json.Marshal(&metricsLink)
	if err != nil {
		logger.ErrorLog.Error().Msgf("CreateDspIms json 转换失败")
		return ""
	}
	if ctx != nil && ctx.Err() != nil {
		return ""
	}

	toStringUrl := base64.RawURLEncoding.EncodeToString(marshal)

	var imsBuf strings.Builder
	imsBuf.Grow(1500)
	imsBuf.WriteString(global.EngineConfig.Ims)
	imsBuf.WriteString("rid=")
	imsBuf.WriteString(utils.Int64ToString(reqCon.BidRequest.SlotId))
	imsBuf.WriteString("&pkg=")
	imsBuf.WriteString(reqCon.BidRequest.App.Pkg)
	imsBuf.WriteString("&source=")
	imsBuf.WriteString(dspSlotInfo.Source)
	imsBuf.WriteString("&data=")
	imsBuf.WriteString(toStringUrl)
	imsBuf.WriteString(global.EngineConfig.ImsLink)

	return imsBuf.String()
}

// 组装点击
func CreateClkReportTalkMetrics(ctx context.Context, reqCon *RequestContext, dspSlotInfo *DspSlotInfo) string {
	if ctx != nil && ctx.Err() != nil {
		return ""
	}
	var CaidStr string
	if len(reqCon.BidRequest.Device.Caids) > 0 {
		CaidStr = reqCon.BidRequest.Device.Caids[0].Caid
	}

	metricsLink := ClkReportTalkMetrics{
		EventTs:     utils.Get10MinTime(time.Now()),
		SspSlotId:   reqCon.BidRequest.SlotId,
		DspSlotId:   dspSlotInfo.Id,
		DspSlotCode: dspSlotInfo.DspSlotCode,
		Ip:          reqCon.BidRequest.Ip,
		//AuctionPrice: dspSlotInfo.AuctionPrice,
		Caid:      CaidStr,
		Idfa:      reqCon.BidRequest.Device.Idfa,
		Oaid:      reqCon.BidRequest.Device.Oaid,
		Os:        reqCon.BidRequest.Device.Plt,
		AuctionTs: utils.TimeNowUnix(),
	}

	marshal, err := json.Marshal(&metricsLink)
	if err != nil {
		logger.ErrorLog.Error().Msgf("CreateDspClk json 转换失败")
		return ""
	}
	if ctx != nil && ctx.Err() != nil {
		return ""
	}

	toStringUrl := base64.RawURLEncoding.EncodeToString(marshal)

	var clkBuf strings.Builder
	clkBuf.Grow(1500)
	clkBuf.WriteString(global.EngineConfig.Cls)
	clkBuf.WriteString("rid=")
	clkBuf.WriteString(utils.Int64ToString(reqCon.BidRequest.SlotId))
	clkBuf.WriteString("&pkg=")
	clkBuf.WriteString(reqCon.BidRequest.App.Pkg)
	clkBuf.WriteString("&creativeId=Ut12LsI")
	clkBuf.WriteString("&source=")
	clkBuf.WriteString(dspSlotInfo.Source)
	clkBuf.WriteString("&data=")
	clkBuf.WriteString(toStringUrl)
	clkBuf.WriteString(global.EngineConfig.ClkLink)

	return clkBuf.String()
}
