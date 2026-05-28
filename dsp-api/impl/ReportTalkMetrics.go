package impl

type ImsReportTalkMetrics struct {
	Rid          string  `json:"rid"`          //
	SspSlotId    int64   `json:"sspSlotId"`    // 媒体广告位Id
	DspSlotId    int64   `json:"dspSlotId"`    // 预算位广告位id
	DspSlotCode  string  `json:"dspSlotCode"`  // 预算广告位
	Ip           string  `json:"ip"`           // ip
	Ua           string  `json:"ua"`           // Ua
	AuctionPrice float64 `json:"auctionPrice"` // 成交价
	DspPayType   int     `json:"dspPayType"`   // 上游结算方式 1=分成，2=RTB
	BidPrice     float64 `json:"bidPrice"`     // 上游游底价
	SspPayType   int     `json:"sspPayType"`   // 1=分成，2=RTB
	Caid         string  `json:"caid"`
	Idfa         string  `json:"idfa"`
	Oaid         string  `json:"oaid"`
	AndroidId    string  `json:"androidId"`
	Imei         string  `json:"imei"`
	Os           int     `json:"os"`
	AuctionTs    int64   `json:"auctionTs"` // 成交时间
	//Nurls        []string `json:"nurls"`     // 竞胜利链接
	//Lurls        []string `json:"lurls"`     // 竞争败链接

	WinPrice string `json:"winPrice"` // 下游成交价通过宏替换上报

	UpX   string `json:"up_x"`
	UpY   string `json:"up_y"`
	DownX string `json:"down_x"`
	DownY string `json:"down_y"`
	Ts    int64  `json:"ts"`  //时间戳(毫秒)
	Tts   int64  `json:"tts"` // 时间戳(秒)

	AdvertisersId string `json:"advertisersId"`
	CreativeId    string `json:"creativeId"`
	Source        string `json:"source"`
	Pkg           string `json:"pkg"`
	EventTs       int64  `json:"eventTs"`
}

type ClkReportTalkMetrics struct {
	Rid          string  `json:"rid"`
	SspSlotId    int64   `json:"sspSlotId"`    // 媒体广告位Id
	DspSlotId    int64   `json:"dspSlotId"`    // 预算位广告位id
	DspSlotCode  string  `json:"dspSlotCode"`  // 预算广告位
	Ip           string  `json:"ip"`           // ip
	AuctionPrice float64 `json:"auctionPrice"` // 成交价
	DspPayType   string  `json:"dspPayType"`   // 上游结算方式 1=分成，2=RTB
	Caid         string  `json:"caid"`
	Idfa         string  `json:"idfa"`
	Oaid         string  `json:"oaid"`
	AndroidId    string  `json:"androidId"`
	Imei         string  `json:"imei"`
	Os           int     `json:"os"`
	AuctionTs    int64   `json:"auctionTs"` // 成交时间

	Width    string `json:"width"`
	Height   string `json:"height"`
	WinPrice string `json:"winPrice"` // 下游成交价通过宏替换上报
	UpX      string `json:"up_x"`
	UpY      string `json:"up_y"`
	DownX    string `json:"down_x"`
	DownY    string `json:"down_y"`
	Ts       int64  `json:"ts"` //时间戳(毫秒)

	Pkg           string `json:"pkg"`
	Source        string `json:"source"`
	EventTs       int64  `json:"eventTs"`
	CreativeId    string `json:"creativeId"`
	AdvertisersId string `json:"advertisersId"`
}
