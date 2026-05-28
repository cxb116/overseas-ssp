package impl

//type RequestMetrics struct {
//	Type        int    // 0 请求过滤  1 ssp 请求， 2 请求dsp
//	RequestId   string // 唯一ID，用于链路追踪
//	SspSlotId   int64  // 媒体广告位
//	DspSlotCode string // 预算广告位
//
//	// 结果
//	Status      int    // 最终状态码
//	ErrorReason string // 错误原因
//
//	// 标记位
//	IsBudgetMatched bool // 是否匹配到预算
//	IsFiltered      bool // 是否被过滤
//	IsTimeout       bool // 是否超时
//	// 新增：事件发生时间戳（毫秒）
//	EventTs int64
//}
//
//type ResponseMetrics struct {
//	Type        int    // 1 响应ssp， 2 响应dsp
//	RequestId   string // 唯一ID，用于链路追踪
//	SspSlotId   int64  // 媒体广告位
//	DspSlotCode string // 预算广告位
//
//	// 结果
//	Status      int    // 最终状态码
//	ErrorReason string // 错误原因
//
//	// 标记位
//	IsFiltered bool // 是否被过滤
//	IsTimeout  bool // 是否超时
//	//  新增：事件发生时间戳（毫秒）
//	EventTs int64
//}

type DspRequestMetrics struct {
	RequestId   string // 唯一ID，用于链路追踪
	SspSlotId   int64  // 媒体广告位ID
	DspSlotId   int64  // 预算位广告位ID（可能为空）
	DspSlotCode string // 预算广告位编码（可能为空）

	// 统计字段
	SspReq        int // 请求次数（req_pv）
	SspRes        int // 返回次数（ret_pv）
	DspReq        int
	DspRes        int
	SspReqDiscard int // 丢弃次数（discard）
	SspResDiscard int // 媒体请求丢弃
	DspReqDiscard int // 预算请求丢
	DspResDiscard int
	EventTs       int64 // 事件时间戳：YYYYMMDDHHmm

	ReqCount int64 //总请求次数
}
