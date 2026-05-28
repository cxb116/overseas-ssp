package impl

import "sync"

var BidRequestPool = sync.Pool{
	New: func() any {
		return &BidRequest{}
	},
}

// 获取BidRequest 池对象
func GetBidRequest() *BidRequest {
	request := BidRequestPool.Get().(*BidRequest)
	request.Reset()
	return request
}

// 归还BidRequest对象
func PutBidRequest(request *BidRequest) {
	if request != nil {
		BidRequestPool.Put(request)
	}
}

func (this *BidRequest) Reset() {
	this.App.An = ""
	this.App.Pkg = ""
	this.App.Vc = ""
	this.App.Ver = ""
	this.Device.Plt = 0
	this.Device.Dvt = 0
	this.Device.Ov = ""
	this.Device.Dpi = 0
	this.Device.Ppi = 0
	this.Device.Density = 0
	this.Device.Swidth = 0
	this.Device.Sheight = 0
	this.Device.Vendor = ""
	this.Device.Mdl = ""
	this.Device.Brd = ""
	this.Device.Country = ""
	this.Device.Lg = ""
	this.Device.Net = 0
	this.Device.Opt = 0
	this.Device.Dso = 0
	this.Device.Mac = ""
	this.Device.Serialno = ""
	this.Device.Aid = ""
	this.Device.Imei = ""
	this.Device.Imei_md5 = ""
	this.Device.Imei2 = ""
	this.Device.Oaid = ""
	this.Device.Icc = ""
	this.Device.Iccid = ""
	this.Device.Idfa = ""
	this.Device.Idfv = ""
	this.Device.Openudid = ""
	this.Device.Caids = this.Device.Caids[:0]
	this.Device.Bssid = ""
	this.Device.Ssid = ""
	this.Device.Wifi_mac = ""
	this.Device.Hms = ""
	this.Device.Hag = ""
	this.Device.Hardware_machine = ""
	this.Device.Hardware_model = ""
	this.Device.Device_name = ""
	this.Device.Sys_compiling_ts = ""
	this.Device.Init_ts = ""
	this.Device.Startup_ts = ""
	this.Device.Upgrade_ts = ""
	this.Device.Timezone = ""
	this.Device.Memory = 0
	this.Device.Hard_disk = 0
	this.Device.Cpu_cnt = 0
	this.Device.Cpu_freq = 0
	this.Device.Idfa_policy = 0
	this.Device.Battery_status = 0
	this.Device.Battery_power = 0
	this.Device.Boot_mark = ""
	this.Device.Update_mark = ""
	this.Device.Packages = this.Device.Packages[:0]
	this.Device.Ua = ""
	this.Device.App_store_vc = ""
	this.Device.Paid = ""
	this.Device.Aaid = ""
	this.Geo.Lon = 0
	this.Geo.Lat = 0
	this.Geo.Addr = ""

	this.AdW = 0
	this.AdH = 0
	this.SlotId = 0
	this.AppId = 0
	this.Ip = ""
	this.Ipv6 = ""
	this.Ishttps = 0
	this.Bidfloor = 0
}

type BidRequest struct {
	App      ReqApp    `json:"app,omitempty"`
	Device   ReqDevice `json:"device,omitempty"`
	Geo      ReqGeo    `json:"geo,omitempty"`
	AdW      int       `json:"adw,omitempty"`
	AdH      int       `json:"adh,omitempty"`
	SlotId   int64     `json:"slotId,omitempty"`
	AppId    int       `json:"appId,omitempty"`
	Ip       string    `json:"ip,omitempty"`
	Ipv6     string    `json:"ipv6,omitempty"`
	Ishttps  int       `json:"ishttps"`
	Bidfloor float64   `json:"bidfloor,omitempty"`
}
type ReqApp struct {
	An  string `json:"an,omitempty"`
	Pkg string `json:"pkg,omitempty"`
	Ver string `json:"ver,omitempty"`
	Vc  string `json:"vc,omitempty"`
}

type ReqDevice struct {
	Plt              int        `json:"plt,omitempty"`
	Dvt              int        `json:"dvt,omitempty"`
	Ov               string     `json:"ov,omitempty"`
	Dpi              float64    `json:"dpi,omitempty"`
	Ppi              float64    `json:"ppi,omitempty"`
	Density          float64    `json:"density,omitempty"`
	Swidth           int        `json:"swidth,omitempty"`           //Required int屏幕宽
	Sheight          int        `json:"sheight,omitempty"`          //Required int屏幕高
	Vendor           string     `json:"vendor,omitempty"`           //Required string设备厂商，比如”华为”
	Mdl              string     `json:"mdl,omitempty"`              //Required string Model机型
	Brd              string     `json:"brd,omitempty"`              //Required string Brand品牌
	Country          string     `json:"country,omitempty"`          //Optional string国家；示例：CN
	Lg               string     `json:"lg,omitempty"`               //Optional string Language语言,示例：zh-CN
	Net              int        `json:"net"`                        //Required int Network网络类型0:Unknown1:WIIF2:2G3:3G4:4G5:G
	Opt              int        `json:"opt,omitempty"`              //Required int Operater运营商代码46000=中国移动46001=中国联通46002=中国移动46003=中国电信46007=中国移动1=未知
	Dso              int        `json:"dso"`                        //Required int 		ScreenOrientation屏幕方向0:未知1:横屏2:竖屏
	Mac              string     `json:"mac,omitempty"`              //Required string设备iMAC地址
	Serialno         string     `json:"serialno,omitempty"`         //Optional string设备序列号
	Aid              string     `json:"aid,omitempty"`              //Optional string AndroidID，安卓系统必填
	Imei             string     `json:"imei,omitempty"`             //Required string IMEI值
	Imei_md5         string     `json:"imei_md5,omitempty"`         //Optional string IMIEMD5值
	Imei2            string     `json:"imei2,omitempty"`            //Optional string第二个IMEI值
	Oaid             string     `json:"oaid,omitempty"`             //Optional string安卓匿名标识符
	Icc              string     `json:"icc,omitempty"`              //Optional string IMSI
	Iccid            string     `json:"iccid,omitempty"`            //Optional string ICCID
	Idfa             string     `json:"idfa,omitempty"`             //Optional string iOS系统必填
	Idfv             string     `json:"idfv,omitempty"`             //Optional string iOS设备的IDFV值
	Openudid         string     `json:"openudid,omitempty"`         //Optional string iOS设备的openudid值
	Caids            []ReqCaids `json:"caids,omitempty"`            //Optional []ObjectCaid多CAID数组
	Bssid            string     `json:"bssid,omitempty"`            //Optional string WIFIBSSID
	Ssid             string     `json:"ssid,omitempty"`             //Optional string WIFI名称
	Wifi_mac         string     `json:"wifi_mac,omitempty"`         //Optional string WIFI的MAC地址
	Hms              string     `json:"hms,omitempty"`              //Optional string鸿蒙内核版本（华为设备必须）
	Hag              string     `json:"hag,omitempty"`              //	Optional string应用市场版本（华为设备必须）
	Hardware_machine string     `json:"hardware_machine,omitempty"` //	Optional string设备machine值；示例：iPhone11,4
	Hardware_model   string     `json:"hardware_model,omitempty"`   //	Optional string设备model值；示例：D211AP
	Device_name      string     `json:"device_name,omitempty"`      //	Optional string设备名称
	Sys_compiling_ts string     `json:"sys_compiling_ts,omitempty"` //	Optional string系统编译时间，时间戳，毫秒
	Init_ts          string     `json:"init_ts,omitempty"`          //  Optional string设备初始化时间戳，毫秒（iOS设备必须）
	Startup_ts       string     `json:"startup_ts,omitempty"`       //	Optional string最近启动时间戳，毫秒（iOS设备必须）
	Upgrade_ts       string     `json:"upgrade_ts,omitempty"`       //	Optional string最近升级时间戳，毫秒（iOS设备必须）
	Timezone         string     `json:"timezone,omitempty"`         //	Optional string系统当前时区（iOS设备必须）示例：Asia/Shanghai
	Memory           int        `json:"memory,omitempty"`           //	Optional int物理内存大小，单位：GB（iOS设备必须）
	Hard_disk        int        `json:"hard_disk,omitempty"`        //	Optional int物理硬盘大小，单位：GB（iOS设备必须）
	Cpu_cnt          int        `json:"cpu_cnt,omitempty"`          //	Optional int处理器核数（iOS设备必须）
	Cpu_freq         float64    `json:"cpu_freq,omitempty"`         //	Optional double处理器主频（iOS设备必须）
	Idfa_policy      int        `json:"idfa_policy"`                //	Optional int IDFA授权策略（iOS设备必须）0：未确定1：受限制2：被拒绝3：已授权
	Battery_status   int        `json:"battery_status,omitempty"`   //	Optional int电池充电状态（iOS设备必须）1：未知状态2：不在充电3：正在充电4：满电状态
	Battery_power    int        `json:"battery_power,omitempty"`    //	Optional int电池电量比例（iOS设备必须）例如:60代表60%，80代表80%
	Boot_mark        string     `json:"boot_mark,omitempty"`        //	Optional string系统启动标识，原值传输
	Update_mark      string     `json:"update_mark,omitempty"`      //	Optional string系统更新标识，原值传输
	Packages         []string   `json:"packages,omitempty"`         //	Optional []string用户已安装包名列表
	Ua               string     `json:"ua,omitempty"`               //	Required string User-Agent
	App_store_vc     string     `json:"app_store_vc,omitempty"`     //	Optional string OPPO应用商店版本号，OPPO预算时必填
	Paid             string     `json:"paid,omitempty"`             //	Optional string拼多多的paid
	Aaid             string     `json:"aaid,omitempty"`             //	Optional string阿里巴巴匿名设备标识，需集成阿里SDK获取

}

type ReqGeo struct {
	Lon  float64 `json:"lon,omitempty"`
	Lat  float64 `json:"lat,omitempty"`
	Addr string  `json:"addr,omitempty"`
}

type ReqCaids struct {
	Caid     string `json:"caid,omitempty"`
	Caid_ver string `json:"caid_ver,omitempty"`
}
