package zhonghui

type ZhongHuiReq struct {
	Id      string `json:"id,omitempty"`      //Y 请求id，开发者⾃⾏⽣成，需保证其唯⼀性
	Version string `json:"version,omitempty"` //Y 接⼝版本号，如1.0
	Imp     Imp    `json:"imp,omitempty"`     //Object Y ⼴告位曝光信息
	App     App    `json:"app,omitempty"`     //Object Y 应⽤信息
	Device  Device `json:"device,omitempty"`  //Object Y 设备信息
	User    User   `json:"user,omitempty"`    //Object N ⽤户信息
	Tmax    int    `json:"tmax,omitempty"`    //N 超时时间，单位毫秒，默认500ms
}

type Imp struct {
	Tagid    string `json:"tagid,omitempty"`    //Y ⼴告位ID，由平台提供
	W        int    `json:"w,omitempty"`        //Y ⼴告位宽度
	H        int    `json:"bidfloor,omitempty"` //N 千次展示底价，(单位：分)； 竞价模式必填且值⼤于0
	Bidfloor int    `json:"bidfloor,omitempty"`
	Dplink   int    `json:"dplink,omitempty"` //N 是否⽀持deeplink，0-不⽀持 1-⽀持（默认0）
	Ulink    int    `json:"ulink,omitempty"`  //N 是否⽀持universallinks，0不⽀持 1⽀持（默认0
}

type App struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Bundle string `json:"bundle,omitempty"`
	Ver    string `json:"ver,omitempty"`
}
type Device struct {
	Ua              string  `json:"ua,omitempty"`          //Y 浏览器UserAgent字符串
	DeviceType      int     `json:"device_type,omitempty"` //Y 设备类型0.未知1.⼿机2.平板3.TV
	Os              int     `json:"os,omitempty"`          //操 作系统1.Android2.iOS
	Osv             string  `json:"osv,omitempty"`         // 系统版本
	Osl             int     `json:"osl,omitempty"`         // 安卓系系统版本系统级别
	Make            string  `json:"make,omitempty"`        // 设备制造商如“Apple”
	Model           string  `json:"model,omitempty"`       // 设备型号
	W               int     `json:"w,omitempty"`           // 屏幕宽度
	H               int     `json:"h,omitempty"`           // 屏幕高度
	Ppi             int     `json:"ppi,omitempty"`         // 像素密度
	Dpi             float64 `json:"dpi,omitempty"`         // 物理像素密度
	Density         string  `json:"density,omitempty"`     // 屏幕分辨率
	Imei            string  `json:"imei,omitempty"`        // imei 安卓设备10
	Imeimd5         string  `json:"imeimd5,omitempty"`     // imei md5值
	Oaid            string  `json:"oaid,omitempty"`        // oaid
	Oaidmd5         string  `json:"oaidmd5,omitempty"`
	Dpid            string  `json:"dpid,omitempty"`
	Dpidmd5         string  `json:"dpidmd5,omitempty"`
	Mac             string  `json:"mac,omitempty"`
	Idfa            string  `json:"idfa,omitempty"`
	Idfamd5         string  `json:"idfamd5,omitempty"`
	Ip              string  `json:"ip,omitempty"`
	Ipv6            string  `json:"ipv6,omitempty"`
	Carrier         int     `json:"carrier,omitempty"`
	Network         int     `json:"network,omitempty"`
	Imsi            string  `json:"imsi,omitempty"`
	Orientation     int     `json:"orientation,omitempty"`
	DeviceName      string  `json:"device_name,omitempty"`
	SyscmpTime      string  `json:"syscmp_time,omitempty"`
	DeviceInitTime  string  `json:"device_init_time,omitempty"`
	StartupTime     string  `json:"startup_time,omitempty"`
	UpdateTime      string  `json:"update_time,omitempty"`
	Hmscore         string  `json:"hms_core,omitempty"`
	AppstoreVer     string  `json:"appstore_ver,omitempty"`
	AppstoreVerCode string  `json:"appstore_ver_code,omitempty"`
	HardwareModel   string  `json:"hardware_model,omitempty"`
	HardwareMachine string  `json:"hardware_machine,omitempty"`
	Country         string  `json:"country,omitempty"`
	Language        string  `json:"language,omitempty"`
	Timezone        string  `json:"timezone,omitempty"`
	CpuNum          int64   `json:"cpu_num,omitempty"`
	DiskTotal       int64   `json:"disk_total,omitempty"`
	MemTotal        int64   `json:"mem_total,omitempty"`
	BootMark        string  `json:"boot_mark,omitempty"`
	UpdateMark      string  `json:"update_mark,omitempty"`
	Paid            string  `json:"paid,omitempty"`
	Lon             float64 `json:"lon,omitempty"`
	Lat             float64 `json:"lat,omitempty"`
	Aaid            string  `json:"aaid,omitempty"`
	Alidid          string  `json:"alid,omitempty"`
	Aliuid          string  `json:"aluid,omitempty"`
	Caids           []Caids `json:"caids,omitempty"`
}

type Caids struct {
	Caid    string `json:"caid,omitempty"`
	Version string `json:"version,omitempty"`
}

type User struct {
	Id                string   `json:"id,omitempty"`
	Yob               int      `json:"yob,omitempty"`
	Gender            string   `json:"gender,omitempty"`
	Keywords          []string `json:"keywords,omitempty"`
	InstallAppPkgList []string `json:"install_app_pkg_list,omitempty"`
}

type Response struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data *DataRes `json:"data"`
}

type DataRes struct {
	Title          string        `json:"title"`
	Desc           string        `json:"desc"`
	Imgs           []ImgRes      `json:"imgs"`
	Logo           ImgRes        `json:"logo"`
	Action         int           `json:"action"`
	CreativeType   int           `json:"creative_type"`
	Clickurl       string        `json:"clickurl"`
	Clickurl2      string        `json:"clickurl2"`
	Deeplink       string        `json:"deeplink"`
	Universalurl   string        `json:"universalurl"`
	MarketUrl      string        `json:"market_url"`
	WxId           string        `json:"wx_id"`
	WxPath         string        `json:"wx_path"`
	Video          VideoRes      `json:"video"`
	App            AppRes        `json:"app"`
	Price          int           `json:"price"`
	Nurl           string        `json:"nurl"`
	Tracking       []TrackingRes `json:"tracking"`
	ClickReportUrl []string      `json:"click_report_url"`
}

type ImgRes struct {
	W   int    `json:"w"`
	H   int    `json:"h"`
	Url string `json:"url"`
}

type VideoRes struct {
	Url         string `json:"url"`
	W           int    `json:"w"`
	H           int    `json:"h"`
	Duration    int    `json:"duration"`
	SkipMinTime int    `json:"skip_min_time"`
	CoverUrl    string `json:"cover_url"`
	EndCardUrl  string `json:"end_card_url"`
	EndCardHtml string `json:"end_card_html"`
	EndTitle    string `json:"end_title"`
	EndDesc     string `json:"end_desc"`
}

type AppRes struct {
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Bundle     string `json:"bundle"`
	Size       int    `json:"size"`
	DownUrl    string `json:"down_url"`
	Version    string `json:"version"`
	Author     string `json:"author"`
	Permission string `json:"permission"`
	Privacy    string `json:"privacy"`
	Introduced string `json:"introduced"`
}

type TrackingRes struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}
