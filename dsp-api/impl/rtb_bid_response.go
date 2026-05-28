package impl

// ==================== OpenRTB 2.6 竞价响应对象 ====================

// BidResponse 表示竞价响应
type BidResponse struct {
	// 响应 ID，必须与请求的 ID 匹配
	Id string `json:"id,omitempty"`

	// 座位竞价数组，每个座位包含该席位对所有展示机会的出价
	Seatbid []Seatbid `json:"seatbid,omitempty"`

	// 拍卖类型，应与请求中的 auction type 一致
	Bidid string `json:"bidid,omitempty"`

	// 请求中未出价的展示 ID 数组
	Nbid string `json:"nbi,omitempty"`

	// 扩展字段
	Ext interface{} `json:"ext,omitempty"`

	// 调试信息（仅用于测试）
	Cur string `json:"cur,omitempty"`

	// 货币代码
	Customdata string `json:"customdata,omitempty"`

	// 自定义数据
}

// Seatbid 表示一个出价席位（买方）
type Seatbid struct {
	// 买方 ID
	Bid []Bid `json:"bid,omitempty"`

	// 出价数组
	Seat string `json:"seat,omitempty"`

	// 座位 ID
	Group int32 `json:"group,omitempty"`

	// 扩展字段
	Ext interface{} `json:"ext,omitempty"`

	// 是否分组
}

// Bid 表示对某个展示机会的出价
type Bid struct {
	// 展示 ID，必须与请求中的 imp.id 匹配
	Id string `json:"id,omitempty"`

	// 展示机会的 ID
	Impid string `json:"impid,omitempty"`

	// 出价金额（单位为分）
	Price float64 `json:"price,omitempty"`

	// 出价 ID，唯一标识此出价
	Adid string `json:"adid,omitempty"`

	// 广告 ID
	Nurl string `json:"nurl,omitempty"`

	// 赢标通知 URL
	Burl string `json:"burl,omitempty"`

	// 结算通知 URL
	Lurl string `json:"lurl,omitempty"`

	// 丢标通知 URL
	Adm string `json:"adm,omitempty"`

	// 广告代码（HTML 片段、XML 文档等）
	Adomain []string `json:"adomain,omitempty"`

	// 广告主域名
	Bundle string `json:"bundle,omitempty"`

	// 应用包名
	Iurl string `json:"iurl,omitempty"`

	// 图片 URL
	Cid string `json:"cid,omitempty"`

	// 创意 ID
	Crid string `json:"crid,omitempty"`

	// 创意渲染 ID
	Cat []string `json:"cat,omitempty"`

	// 广告分类数组
	Attr []int32 `json:"attr,omitempty"`

	// 创意属性数组
	Api int32 `json:"api,omitempty"`

	// API 框架
	Protocol int32 `json:"protocol,omitempty"`

	// 协议
	Qagmediarating int32 `json:"qagmediarating,omitempty"`

	// 媒体评级
	Language string `json:"language,omitempty"`

	// 创意语言
	Dealid string `json:"dealid,omitempty"`

	// 交易 ID
	W int32 `json:"w,omitempty"`

	// 广告宽度
	H int32 `json:"h,omitempty"`

	// 广告高度
	Wratio int32 `json:"wratio,omitempty"`

	// 宽度比例
	Hratio int32 `json:"hratio,omitempty"`

	// 高度比例
	Dur int32 `json:"dur,omitempty"`

	// 广告时长（视频/音频）
	Mime string `json:"mime,omitempty"`

	// MIME 类型
	Dealer *Dealer `json:"dealer,omitempty"`

	// 经销商标识符
	Ext interface{} `json:"ext,omitempty"`

	// 扩展字段
}

// Dealer 表示经销商信息
type Dealer struct {
	// 经销商标识符
	Id string `json:"id,omitempty"`

	// 经销商标识符
	Name string `json:"name,omitempty"`

	// 经销商名称
	Ext interface{} `json:"ext,omitempty"`

	// 扩展字段
}
