package impl

type DspSlotInfo struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	AdTypeId        int64  `json:"ad_type_id"`
	AdSceneId       int64  `json:"ad_scene_id"`
	OsType          int    `json:"os_type"`
	DspSlotCode     string `json:"dsp_slot_code"`
	DspAppKey       string `json:"dsp_app_key"`
	DspAppId        string `json:"dsp_app_id"`
	DspAppPkg       string `json:"dsp_app_pkg"`
	DspAppSecret    string `json:"dsp_app_secret"`
	DspAppVer       string `json:"dsp_app_ver"`
	DspAppStoreVer  string `json:"dsp_app_store_ver"`
	PriceEncryptKey string `json:"price_encrypt_key"`
	DspAppStoreLink string `json:"dsp_app_store_link"`
	DspPayType      int    `json:"dsp_pay_type"`
	DspDealRatio    int    `json:"dsp_deal_ratio"`
	CompanyId       int64  `json:"company_id"`

	ProductName string `json:"product_name"`

	// 👇 非数据库字段，必须忽略
	DspCompany DspCompany `gorm:"-"`
	DspLaunch  DspLaunch  `gorm:"-"`

	// 价格处理
	AuctionPrice float64 // 拍卖价

	BidPrice      float64 // 给下游的出价
	Source        string  // 广告来源
	AdvertisersId string  // 广告主Id
	CreativeId    string  // 广告创意Id
}

func (DspSlotInfo) TableName() string {
	return "dsp_slot_info"
}
