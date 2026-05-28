package impl

import "sync"

// 最终的预算表在context中使用的
type DspSlotManager struct {
	// 预算
	//DspSlotInfo DspSlotInfo
	// 1个流量需要并发请求多个预算
	DspSlotInfoMaps map[int64][]DspSlotInfo
	DspMutex        sync.RWMutex

	//DspLaunch

	// 媒体
	SspSlotInfo SspSlotInfo
	// 预算公司
	//DspCompany DspCompany
	//// 预算投放
	//DspLaunch DspLaunch
	// 物料数据

	//// 价格处理
	//AuctionPrice float64 // 拍卖价
	//
	//BidPrice      float64 // 给下游的出价
	//Source        string  // 广告来源
	//AdvertisersId string  // 广告主Id
	//CreativeId    string  // 广告创意Id
}
