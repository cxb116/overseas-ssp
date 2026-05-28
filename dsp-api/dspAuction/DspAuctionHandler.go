package dspAuction

import (
	"math"

	"github.com/cxb116/DSP/constant"
	"github.com/cxb116/DSP/impl"
)

// RTB竞价
// 参数1， 返回给下游的出价，参数2 给上游的成交价
func SetAuctionPrice(reqCxt *impl.RequestContext, dspSlotInfo *impl.DspSlotInfo, price float64) (float64, float64) {

	//ssp := reqCxt.DspSlotManager.SspSlotInfo
	//dsp := dspSlotInfo
	//
	//dsp.AuctionPrice = price // 成交价
	//
	//// RTB -> RTB   // 如果下游是rtb 的话，那就需要出我方自己的利润然后再去返回下游价格
	//if dsp.DspPayType == constant.PAY_PRICE_RTB && ssp.SspPayType == constant.PAY_PRICE_RTB {
	//
	//	if dsp.DspLaunch.DspPayTypeRatio > 0 {
	//		ratio := float64(dsp.DspLaunch.DspPayTypeRatio) / 100
	//		bidPrice := price * ratio
	//
	//		dsp.BidPrice = ceilOneDecimal(bidPrice)
	//		return ceilOneDecimal(bidPrice), price
	//	} else {
	//		// 如果没有设置，默认就是 ratio 是 70%
	//		bidPrice := price * 0.7
	//		dsp.BidPrice = ceilOneDecimal(bidPrice)
	//
	//		return ceilOneDecimal(bidPrice), price
	//	}
	//
	//}
	//
	//// RTB -> 分成   上游RTB下游分成,只有成交价
	//if dsp.DspPayType == constant.PAY_PRICE_RTB && ssp.SspPayType == constant.PAY_PRICE_COMPOSITION {
	//	dsp.BidPrice = 0
	//	return 0, price
	//}
	//
	//// 其他情况全部归零上下游都是分成
	//dsp.BidPrice = 0
	//dsp.AuctionPrice = 0
	return 0, 0
}

// 低价过滤
// false 不过滤，true 为过滤
func FloorPriceFilter(reqCxt *impl.RequestContext, dspSlotInfo *impl.DspSlotInfo, price float64) bool {

	dsp := dspSlotInfo
	if dsp.DspPayType == constant.PAY_PRICE_RTB {
		if dsp.DspLaunch.FloorPrice <= price {
			return false
		}
		return true
	}
	return false
}

// 设置底价
// 如果下游是rtb,我方也是rtb,就需要走这个逻辑,  下游是rtb 上游必须是rtb
func SetBidFloorPrice(reqCxt *impl.RequestContext, dspSlotInfo impl.DspSlotInfo) float64 {
	//if reqCxt.DspSlotManager.SspSlotInfo.SspPayType == constant.PAY_PRICE_RTB { // 下游是rtb,上游必须是rtb
	//	if reqCxt.BidRequest.Bidfloor <= 0 {
	//		return 1
	//	}
	//	dealRatio := 0.0
	//	if dspSlotInfo.DspLaunch.DspPayTypeRatio == 0 {
	//		dealRatio = 70
	//	} else {
	//		dealRatio = float64(dspSlotInfo.DspLaunch.DspPayTypeRatio)
	//	}
	//	if dspSlotInfo.DspPayType == constant.PAY_PRICE_RTB {
	//		floorPrice := reqCxt.BidRequest.Bidfloor / dealRatio * 100 // 将值切到1.7 后一位
	//		return ceilOneDecimal(floorPrice)
	//	}
	//}
	//// 下游是分成，上游是rtb
	//if dspSlotInfo.DspPayType == constant.PAY_PRICE_RTB {
	//	return float64(dspSlotInfo.DspLaunch.FloorPrice)
	//}
	//return float64(dspSlotInfo.DspLaunch.FloorPrice)

	return 0

}

func ceilOneDecimal(num float64) float64 {
	return math.Ceil(num*10) / 10
}
