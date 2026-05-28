package dsp

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/cxb116/DSP/constant"
	"github.com/cxb116/DSP/document/dspStruct/huanxi"
	"github.com/cxb116/DSP/dspAuction"
	"github.com/cxb116/DSP/impl"
	"github.com/cxb116/DSP/internal/logger"
	"github.com/cxb116/DSP/internal/utils"
)

func HuanXiDocument(dspCode int64) {
	impl.DspRegister(dspCode, func(reqContext *impl.RequestContext) impl.DspHandler {
		return &HuanXiDsp{
			DspCode:        dspCode,
			RequestContext: nil,
			DspSlotInfo:    nil,
		}
	})
}

type HuanXiDsp struct {
	DspCode        int64
	RequestContext *impl.RequestContext
	DspSlotInfo    *impl.DspSlotInfo
	FloorPrice     float64 // 下游rtb,
}

func (this *HuanXiDsp) SetDspRequestContext(reqCtx *impl.RequestContext, dspSlotInfo *impl.DspSlotInfo) {
	this.RequestContext = reqCtx
	this.DspSlotInfo = dspSlotInfo
}

func (this *HuanXiDsp) ReturnSelf() impl.DspHandler {
	return this
}

func (this *HuanXiDsp) DspRequestHandler() impl.BidResponse {
	bidRequest := this.RequestContext.BidRequest
	bidResponse := impl.BidResponse{}

	//floorPrice := dspAuction.SetBidFloorPrice(this.RequestContext, *this.DspSlotInfo)

	var caid string
	var caidVer string
	if bidRequest.Device.Caids != nil && len(bidRequest.Device.Caids) > 0 {
		caid = bidRequest.Device.Caids[0].Caid
		caidVer = bidRequest.Device.Caids[0].Caid_ver
	}
	dspCode := this.DspSlotInfo.DspSlotCode
	dspCodeInt64 := StringToInt64(dspCode)
	AppIdInt64 := StringToInt64(this.DspSlotInfo.DspAppId)
	req := huanxi.HuanXiReq{
		App: huanxi.HuanXiReqApp{
			An:  bidRequest.App.An,
			Pkg: bidRequest.App.Pkg,
			Vc:  bidRequest.App.Vc,
		},
		Device: huanxi.HuanXiReqDevice{
			Plt:      bidRequest.Device.Plt,
			Dvt:      bidRequest.Device.Dvt,
			Ov:       bidRequest.Device.Ov,
			Dpi:      bidRequest.Device.Dpi,
			Ppi:      bidRequest.Device.Ppi,
			Density:  bidRequest.Device.Density,
			Swidth:   bidRequest.Device.Swidth,
			Sheight:  bidRequest.Device.Sheight,
			Vendor:   bidRequest.Device.Vendor,
			Mdl:      bidRequest.Device.Mdl,
			Brd:      bidRequest.Device.Brd,
			Country:  bidRequest.Device.Country,
			Lg:       bidRequest.Device.Lg,
			Net:      bidRequest.Device.Net,
			Opt:      bidRequest.Device.Opt,
			Dso:      bidRequest.Device.Dso,
			Mac:      bidRequest.Device.Mac,
			Serialno: bidRequest.Device.Serialno,
			Aid:      bidRequest.Device.Aid,
			Imei:     bidRequest.Device.Imei,
			Imei_md5: bidRequest.Device.Imei_md5,
			Imei2:    bidRequest.Device.Imei,
			Oaid:     bidRequest.Device.Oaid,
			Icc:      bidRequest.Device.Icc,
			Iccid:    bidRequest.Device.Iccid,
			Idfa:     bidRequest.Device.Idfa,
			Idfv:     bidRequest.Device.Idfv,
			Openudid: bidRequest.Device.Openudid,
			Caids: []huanxi.Caids{
				{
					Caid:     caid,
					Caid_ver: caidVer,
				},
			},
			Bssid:            bidRequest.Device.Bssid,
			Ssid:             bidRequest.Device.Ssid,
			Wifi_mac:         bidRequest.Device.Wifi_mac,
			Hms:              bidRequest.Device.Hms,
			Hag:              bidRequest.Device.Hag,
			Hardware_machine: bidRequest.Device.Hardware_machine,
			Hardware_model:   bidRequest.Device.Hardware_model,
			Device_name:      bidRequest.Device.Device_name,
			Sys_compiling_ts: bidRequest.Device.Sys_compiling_ts,
			Init_ts:          bidRequest.Device.Init_ts,
			Startup_ts:       bidRequest.Device.Startup_ts,
			Upgrade_ts:       bidRequest.Device.Upgrade_ts,
			Timezone:         bidRequest.Device.Timezone,
			Memory:           bidRequest.Device.Memory,
			Hard_disk:        bidRequest.Device.Hard_disk,
			Cpu_freq:         bidRequest.Device.Cpu_freq,
			Cpu_cnt:          bidRequest.Device.Cpu_cnt,
			Idfa_policy:      bidRequest.Device.Idfa_policy,
			Battery_status:   bidRequest.Device.Battery_status,
			Battery_power:    bidRequest.Device.Battery_power,
			Boot_mark:        bidRequest.Device.Boot_mark,
			Update_mark:      bidRequest.Device.Update_mark,
			Packages:         bidRequest.Device.Packages,
			Ua:               bidRequest.Device.Ua,
			Paid:             bidRequest.Device.Paid,
			Aaid:             bidRequest.Device.Aaid,
		},
		Geo: huanxi.HuanXiReqGeo{
			Lon:  bidRequest.Geo.Lon,
			Lat:  bidRequest.Geo.Lat,
			Addr: bidRequest.Geo.Addr,
		},
		Adw:     bidRequest.AdW,
		Adh:     bidRequest.AdH,
		SlotId:  dspCodeInt64,
		AppId:   AppIdInt64,
		Ip:      bidRequest.Ip,
		Ipv6:    bidRequest.Ipv6,
		Ishttps: bidRequest.Ishttps,
		//BidFloor: floorPrice,
	}

	requestJson, err := json.Marshal(&req)
	if err != nil {
		logger.Log.Info().Msg("Huanxi json 转换失败")
		return impl.BidResponse{Res: -20, Message: "json转化失败"}
	}

	if this.DspSlotInfo.DspLaunch.LogCaptureAt > time.Now().Unix() {
		logger.BusinessLog.Info().Msgf("%s ### DSP_HUANXI req catch ### 预算位ID: %d, 流量ID: %d,", string(requestJson), this.DspSlotInfo.Id, this.RequestContext.BidRequest.SlotId)
	}

	huanxiUrl := this.DspSlotInfo.DspCompany.Url
	reqCtx := this.RequestContext.Context
	resByte, b, isTimeout := impl.SendOnJSONRequestWithContext(reqCtx, requestJson, huanxiUrl, impl.ClientPool)
	if b == false {
		bidResponse.Res = constant.REQ_CODE_DSP_ERROR
		logger.Log.Error().Msgf("Huanxi 请求失败")
		return impl.BidResponse{Res: constant.REQ_CODE_DSP_ERROR, Message: "Huanxi 请求失败"}
	}

	// 超时处理
	if isTimeout == true {
		bidResponse.Res = constant.REQ_CODE_DSP_TIMEOUT
		logger.ErrorLog.Error().Msgf("Huanxi 超时")
		return impl.BidResponse{Res: constant.REQ_CODE_DSP_TIMEOUT, Message: "请求超时"}
	}

	var huanxiRes huanxi.HuanXiRes
	err = json.Unmarshal(resByte, &huanxiRes)
	if err != nil {
		logger.ErrorLog.Error().Msgf("HuanXiRes Unmarshal Err: %v  预算位ID: %d, 流量ID: %d,", err, this.DspSlotInfo.Id, this.RequestContext.BidRequest.SlotId)
		bidResponse.Res = constant.REQ_CODE_DSP_ERROR
		return impl.BidResponse{Res: constant.REQ_CODE_DSP_ERROR, Message: "预算转换失败"}
	}

	logger.Log.Info().Msgf("json: huanxiRes:", huanxiRes)
	if huanxiRes.Ad == nil {
		bidResponse.Res = constant.REQ_CODE_DSP_NIL
		logger.ErrorLog.Error().Msgf("HuanXiRes Ad nil: 预算位ID: %d, 流量ID: %d", this.DspSlotInfo.Id, this.RequestContext.BidRequest.SlotId)
		return impl.BidResponse{Res: 0, Message: "无填充"}
	}

	if this.DspSlotInfo.DspLaunch.LogCaptureAt > time.Now().Unix() {
		marshal, _ := json.Marshal(huanxiRes)
		logger.BusinessLog.Info().Msgf("%s ### DSP_HUANXI res catch ### 预算位ID: %d, 流量ID: %d,", string(marshal), this.DspSlotInfo.Id, this.RequestContext.BidRequest.SlotId)
	}

	//huanxiRes.Ad.Adext.Price = 1.4
	//// 验证完以后首先去检查底价，物料是否匹配
	//filterFloorPriceBool := dspAuction.FloorPriceFilter(this.RequestContext, this.DspSlotInfo, huanxiRes.Ad.Adext.Price)
	//if filterFloorPriceBool == true {
	//	bidResponse.Res = constant.REQ_CODE_DSP_NOT_FLOORPRICE
	//	logger.ErrorLog.Info().Msgf("HuanXiRes 出价太低 %d : 预算位ID: %d, 流量ID: %d ", huanxiRes.Ad.Adext.Price, this.DspSlotInfo.Id, this.RequestContext.BidRequest.SlotId)
	//	return impl.BidResponse{Res: constant.REQ_CODE_DSP_NOT_FLOORPRICE, Message: "预算出价太低"}
	//}

	price, winPrice := dspAuction.SetAuctionPrice(this.RequestContext, this.DspSlotInfo, huanxiRes.Ad.Adext.Price)

	bidResponse.Ad.Adext.Price = price
	bidResponse.Ad.Adw = huanxiRes.Ad.Adw
	bidResponse.Ad.Adh = huanxiRes.Ad.Adh
	bidResponse.Ad.Img = huanxiRes.Ad.Img
	bidResponse.Ad.Img2 = huanxiRes.Ad.Img2
	bidResponse.Ad.Img3 = huanxiRes.Ad.Img3
	bidResponse.Ad.Guidepic = huanxiRes.Ad.Guidepic
	bidResponse.Ad.Title = huanxiRes.Ad.Title

	for _, trackUrl := range huanxiRes.Ad.Click {
		bidResponse.Ad.Click = append(bidResponse.Ad.Click, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Imp {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Imp = append(bidResponse.Ad.Imp, trackUrl)
	}

	bidResponse.Ad.Act = huanxiRes.Ad.Act
	bidResponse.Ad.Lpg = huanxiRes.Ad.Lpg
	bidResponse.Ad.Dplink = huanxiRes.Ad.Dplink
	bidResponse.Ad.Ulk = huanxiRes.Ad.Ulk
	bidResponse.Ad.Qak = huanxiRes.Ad.Qak
	bidResponse.Ad.Logo = huanxiRes.Ad.Logo
	bidResponse.Ad.Icon = huanxiRes.Ad.Icon
	bidResponse.Ad.Html = huanxiRes.Ad.Html
	bidResponse.Ad.Adext.Pkg = huanxiRes.Ad.Adext.Pkg

	for _, trackUrl := range huanxiRes.Ad.Adext.Nurl {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Nurl = append(bidResponse.Ad.Adext.Nurl, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Lurl {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Lurl = append(bidResponse.Ad.Adext.Lurl, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Carurl {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Carurl = append(bidResponse.Ad.Adext.Carurl, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Downbegin {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Downbegin = append(bidResponse.Ad.Adext.Downbegin, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Downsucc {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Downsucc = append(bidResponse.Ad.Adext.Downsucc, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Installsucc {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Installsucc = append(bidResponse.Ad.Adext.Installsucc, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Installbegin {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Installbegin = append(bidResponse.Ad.Adext.Installbegin, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Appactive {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Appactive = append(bidResponse.Ad.Adext.Appactive, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Ktbegin {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Ktbegin = append(bidResponse.Ad.Adext.Ktbegin, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Ktfail {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Ktfail = append(bidResponse.Ad.Adext.Ktfail, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Adext.Kt {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Adext.Kt = append(bidResponse.Ad.Adext.Kt, trackUrl)
	}

	bidResponse.Ad.Adext.Currency = huanxiRes.Ad.Adext.Currency
	bidResponse.Ad.Videoext.Vurl = huanxiRes.Ad.Videoext.Vurl
	bidResponse.Ad.Videoext.Duration = huanxiRes.Ad.Videoext.Duration
	bidResponse.Ad.Videoext.Keep = huanxiRes.Ad.Videoext.Keep
	bidResponse.Ad.Videoext.Vhtml = huanxiRes.Ad.Videoext.Vhtml
	bidResponse.Ad.Videoext.Lpic = huanxiRes.Ad.Videoext.Lpic
	bidResponse.Ad.Videoext.Endcard_html = huanxiRes.Ad.Videoext.Endcard_html
	bidResponse.Ad.Videoext.Endcard_img = huanxiRes.Ad.Videoext.Endcard_img

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_first_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_first_urls = append(bidResponse.Ad.Videoext.Video_first_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_mid_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_mid_urls = append(bidResponse.Ad.Videoext.Video_mid_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_third_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_third_urls = append(bidResponse.Ad.Videoext.Video_third_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_begin_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_begin_urls = append(bidResponse.Ad.Videoext.Video_begin_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_end_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_end_urls = append(bidResponse.Ad.Videoext.Video_end_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_mute_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_mute_urls = append(bidResponse.Ad.Videoext.Video_mute_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_unmute_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_unmute_urls = append(bidResponse.Ad.Videoext.Video_unmute_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_skip_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_skip_urls = append(bidResponse.Ad.Videoext.Video_skip_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_close_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_close_urls = append(bidResponse.Ad.Videoext.Video_close_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_pause_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_pause_urls = append(bidResponse.Ad.Videoext.Video_pause_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_resume_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_resume_urls = append(bidResponse.Ad.Videoext.Video_resume_urls, trackUrl)
	}
	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_replay_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_replay_urls = append(bidResponse.Ad.Videoext.Video_replay_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_fullscreen_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_fullscreen_urls = append(bidResponse.Ad.Videoext.Video_fullscreen_urls, trackUrl)
	}

	for _, trackUrl := range huanxiRes.Ad.Videoext.Video_exit_fullscreen_urls {
		trackUrl = HuanXiReplaceWithParams(trackUrl, winPrice)
		bidResponse.Ad.Videoext.Video_exit_fullscreen_urls = append(bidResponse.Ad.Videoext.Video_exit_fullscreen_urls, trackUrl)
	}

	bidResponse.Res = constant.REQ_CODE_SUC // 一定要这这个否则没有上报链接
	logger.Log.Info().Msgf("发布的欢喜测试----------------------------")
	return bidResponse
}

func (this *HuanXiDsp) GetDspCode() int64 {
	return this.DspCode
}

func HuanXiReplaceWithParams(url string, winPrice float64) string {

	// 宏替换
	replacer := strings.NewReplacer(
		"__WIN_PRICE__", utils.FloatToString(winPrice),
	)

	return replacer.Replace(url)
}

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}
	return i
}
