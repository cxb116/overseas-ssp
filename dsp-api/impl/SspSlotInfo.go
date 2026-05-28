package impl

type SspSlotInfo struct {
	Id              int64 `json:"id"`         // 给媒体分派的广告位id // 匹配指标
	MediaId         int64 `json:"media_id"`   // 媒体公司Id
	AppId           int64 `json:"app_id"`     // app应用Id
	OsType          int   `json:"os_type"`    // 操作系统类型，1=Android，2=iOS
	AdTypeId        int   `json:"ad_type_id"` // 媒体方广告类型
	AdSceneId       int   `json:"ad_scene_id"`
	SspPayType      int   `json:"ssp_pay_type"`     // 下游媒体结算方式，1=分成，2=RTB
	SspDealRatio    int   `json:"ssp_deal_ratio"`   // 下游媒体分成系数，0到100，单位%  （只有分成有的系数）
	Height          int   `json:"height"`           // 广告位高
	Width           int   `json:"width"`            // 广告位宽
	InteractionType int   `json:"interaction_type"` // 交互类型是否支持，1：打开网页，2：deeplink， 3：直接下载应用；4: 广点通; 5 小程序跳转 6,应用商店下载，7 快应用

	AppName     string `json:"app_name"`     // App 名称
	PkgName     string `json:"pkg_name"`     // 包名称
	DownloadUrl string `json:"download_url"` // 下载地址

}

func (SspSlotInfo) TableName() string {
	return "ssp_slot_info"
}
