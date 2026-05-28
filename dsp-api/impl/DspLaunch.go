package impl

type DspLaunch struct {
	Id        int64 `json:"id"`          // 权重表id
	SspSlotId int64 `json:"ssp_slot_id"` //  流量广告位Id
	DspSlotId int64 `json:"dsp_slot_id"` //  预算广告位id

	TrafficWeight   int     `json:"traffic_weight"`   // 权重 int
	LaunchStrategy  int     `json:"launch_strategy"`  // 投放策略
	FloorPrice      float64 `json:"floor_price"`      // 底价
	IpLimit         int     `json:"ip_limit"`         // ip 限流次数
	TrackSchwarz    string  `json:"track_schwarz"`    // 上报黑名单
	LogCaptureAt    int64   `json:"log_capture_at"`   // 捕获日志时长
	LaunchTime      int     `json:"launch_time"`      // 投放时段
	CrowdDirection  int     `json:"crowd_direction"`  // 人群定向 1 不限制，2 定向，3排除
	RegionDirection int     `json:"region_direction"` // 地域定向 1 不限制，2 定向，3排除
	BrandDirection  int     `json:"brand_direction"`  // 品牌定向 1 不限制，2 定向，3排除
	PkgTransfer     int     `json:"pkg_transfer"`     // 包透传  0 关闭,1 开启
	PriceTransfer   int     `json:"price_transfer"`   // 底价透传 0 关闭 1 开启
	Indexs          int     `json:"indexs"`           // 相同归为1个流量需要请求的多个预算归类
	//dsp_pay_type_ratio
	DspPayTypeRatio int `json:"dsp_pay_type_ratio"` // 上游rtb的情况下是有数据的
}

func (DspLaunch) TableName() string {
	return "dsp_launch"
}
