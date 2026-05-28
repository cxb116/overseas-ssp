package impl

import "github.com/cxb116/DSP/constant"

// 请求参数流程处理
func ProcessHandler(reqContext *RequestContext, dspSlotManager *DspSlotManager) int {
	device := reqContext.BidRequest.Device

	osType := dspSlotManager.DspSlotInfoMaps[reqContext.BidRequest.SlotId][0].OsType

	if reqContext.BidRequest.Ip == "" || device.Ua == "" {
		return constant.REQ_CODE_DEVICE_ERR
	}

	// OS 校验
	if device.Plt != osType {
		return constant.REQ_CODE_DEVICE_ERR
	}

	switch osType {

	case constant.REQ_IOS:
		hasIDFA := device.Idfa != ""

		hasValidCAID := false
		for _, caid := range device.Caids {
			if caid.Caid != "" {
				hasValidCAID = true
				break
			}
		}

		if !hasIDFA && !hasValidCAID {
			return constant.REQ_CODE_DEVICE_ERR
		}

	case constant.REQ_ANDROID:
		// 至少一个存在
		if device.Oaid == "" && device.Aaid == "" {
			return constant.REQ_CODE_DEVICE_ERR
		}
	}

	return constant.REQ_CODE_DEVICE_SUC
}
