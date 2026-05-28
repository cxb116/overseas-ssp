package huanxi

type HuanXiReq struct {
	App      HuanXiReqApp    `json:"app,omitempty"`      // Required  App 应用对象
	Device   HuanXiReqDevice `json:"device,omitempty"`   // Required  Device 设备对象
	Geo      HuanXiReqGeo    `json:"geo,omitempty"`      // Required  Geo Geo 对象
	Adw      int             `json:"adw,omitempty"`      // Required  广告宽（单位 px）
	Adh      int             `json:"adh,omitempty"`      // Required  广告高（单位 px）
	SlotId   int64           `json:"slotid,omitempty"`   // Required  数字欢喜广告平台申请的广告位 ID
	AppId    int64           `json:"appid,omitempty"`    // Required  数字欢喜广告平台申请的应用 ID
	Ip       string          `json:"ip,omitempty"`       // Required  客户端公网 IP 地址
	Ipv6     string          `json:"ipv6,omitempty"`     // Optional  客户端公网 IPV6 地址
	Ishttps  int             `json:"ishttps,omitempty"`  // Optional  是否支持 HTTPS 0: 否 1: 是
	BidFloor float64         `json:"bidfloor,omitempty"` // Optional  底价，单位：分/CPM，货币单位 RMB，RTB竞价结算方式时必填
	//Bidid    string          `json:"bidid,omitempty"`    // Reuired  请求 ID，保证唯一性，方便链路排查
}

type HuanXiReqApp struct {
	An  string `json:"an,omitempty"`  //Required      应用名称
	Pkg string `json:"pkg,omitempty"` //Required      应用包名
	Ver string `json:"ver,omitempty"` //Required      应用版本(例如：1.0.0)
	Vc  string `json:"vc,omitempty"`  //Required      协议版本号(默认 1.2.2)
}

type HuanXiReqDevice struct {
	Plt              int      `json:"plt,omitempty"`              // Required     系统平台1:Android，2:iOS，3:windows，4:mac，5:linux
	Dvt              int      `json:"dvt,omitempty"`              // Required     Device type 0:未知，1:手机，2: pad，3: PC
	Ov               string   `json:"ov,omitempty"`               //  Required      操作系统版本号
	Dpi              float64  `json:"dpi,omitempty"`              // Required       屏幕项目密度，例：480
	Ppi              float64  `json:"ppi,omitempty"`              // Optional       屏幕每英寸像素数目
	Density          float64  `json:"density,omitempty"`          // Optional       屏幕密度 例：2.0
	Swidth           int      `json:"swidth,omitempty"`           //  Required     屏幕宽
	Sheight          int      `json:"sheight,omitempty"`          // Required     屏幕高
	Vendor           string   `json:"vendor,omitempty"`           //  Required      设备厂商，比如”华为”
	Mdl              string   `json:"mdl,omitempty"`              // Required      Model 机型
	Brd              string   `json:"brd,omitempty"`              // Required      Brand 品牌
	Country          string   `json:"country,omitempty"`          // Optional      国家；示例：CN
	Lg               string   `json:"lg,omitempty"`               //  Optional      Language 语言,示例：zh-CN
	Net              int      `json:"net,omitempty"`              // Required     Network 网络类型  0:Unknown 1:WIIF 2:2G 3:3G 4:4G 5:G
	Opt              int      `json:"opt,omitempty"`              // Required     Operater 运营商代码 46000 = 中国移动 46001 = 中国联通 46002 = 中国移动 46003 = 中国电信 46007 = 中国移动 1 = 未知
	Dso              int      `json:"dso,omitempty"`              // Required     ScreenOrientation 屏幕方向 0: 未知 1: 横屏 2: 竖屏
	Mac              string   `json:"mac,omitempty"`              // Required      设备iMAC地址
	Serialno         string   `json:"serialno,omitempty"`         //    Optional      设备序列号
	Aid              string   `json:"aid,omitempty"`              // Optional      AndroidID，安卓系统必填
	Imei             string   `json:"imei,omitempty"`             //    Required      IMEI值
	Imei_md5         string   `json:"imei_md5,omitempty"`         //    Optional      IMIEMD5值
	Imei2            string   `json:"imei2,omitempty"`            //   Optional      第二个IMEI值
	Oaid             string   `json:"oaid,omitempty"`             //    Optional      安卓匿名标识符
	Icc              string   `json:"icc,omitempty"`              // Optional      IMSI
	Iccid            string   `json:"iccid,omitempty"`            //   Optional      ICCID
	Idfa             string   `json:"idfa,omitempty"`             //    Optional      iOS系统必填
	Idfv             string   `json:"idfv,omitempty"`             //    Optional      iOS设备的IDFV值
	Openudid         string   `json:"openudid,omitempty"`         //    Optional      iOS 设备的openudid值
	Caids            []Caids  `json:"caids,omitempty"`            //    Optional      iOS 设备的caid
	Bssid            string   `json:"bssid,omitempty"`            //   Optional      WIFI BSSID
	Ssid             string   `json:"ssid,omitempty"`             //    Optional      WIFI名称
	Wifi_mac         string   `json:"wifi_mac,omitempty"`         //    Optional      WIFI 的 MAC地址
	Hms              string   `json:"hms,omitempty"`              // Optional      鸿蒙内核版本（华为设备必须）
	Hag              string   `json:"hag,omitempty"`              // Optional      应用市场版本（华为设备必须）
	Hardware_machine string   `json:"hardware_machine,omitempty"` //    Optional      设备machine值；示例：iPhone11,4
	Hardware_model   string   `json:"hardware_model,omitempty"`   //  Optional      设备model值；示例：D211AP
	Device_name      string   `json:"device_name,omitempty"`      // Optional      设备名称
	Sys_compiling_ts string   `json:"sys_compiling_ts,omitempty"` //    Optional      系统编译时间，时间戳，毫秒
	Init_ts          string   `json:"init_ts,omitempty"`          // Optional      设备初始化时间戳，毫秒（iOS 设备必须）
	Startup_ts       string   `json:"startup_ts,omitempty"`       //  Optional      最近启动时间戳，毫秒（iOS 设备必须）
	Upgrade_ts       string   `json:"upgrade_ts,omitempty"`       //  Optional      最近升级时间戳，毫秒（iOS 设备必须）
	Timezone         string   `json:"timezone,omitempty"`         //    Optional      系统当前时区（iOS 设备必须）示例：Asia/Shanghai
	Memory           int      `json:"memory,omitempty"`           //  Optional     物理内存大小，单位：GB（iOS 设备必须）
	Hard_disk        int      `json:"hard_disk,omitempty"`        //   Optional     物理硬盘大小，单位：GB（iOS 设备必须）
	Cpu_cnt          int      `json:"cpu_cnt,omitempty"`          //     Optional     处理器核数（iOS 设备必须）
	Cpu_freq         float64  `json:"cpu_freq,omitempty"`         //    Optional      处理器主频（iOS 设备必须）
	Idfa_policy      int      `json:"idfa_policy,omitempty"`      // Optional     IDFA 授权策略（iOS 设备必须） 0：未确定  1：受限制  2：被拒绝  3：已授权
	Battery_status   int      `json:"battery_status,omitempty"`   //  Optional     电池充电状态（iOS 设备必须）  1：未知状态  2：不在充电  3：正在充电  4：满电状态
	Battery_power    int      `json:"battery_power,omitempty"`    //   Optional     电池电量比例（iOS 设备必须） 例如: 60 代表 60%，80 代表 80%
	Boot_mark        string   `json:"boot_mark,omitempty"`        //   Optional      系统启动标识，原值传输
	Update_mark      string   `json:"update_mark,omitempty"`      //    Optional      系统更新标识，原值传输
	Packages         []string `json:"packages,omitempty"`         //    Optional    string    用户已安装包名列表
	Ua               string   `json:"ua,omitempty"`               //  Required      User-Agent
	App_store_vc     string   `json:"app_store_vc,omitempty"`     //    Optional      OPPO应用商店版本号，OPPO预算时必填
	Paid             string   `json:"paid,omitempty"`             //    Optional      拼多多的paid
	Aaid             string   `json:"aaid,omitempty"`             //    Optional      阿里巴巴匿名设备标识，需集成阿里 SDK 获取
}

type Caids struct {
	Caid     string `json:"caid,omitempty"`
	Caid_ver string `json:"caid_ver,omitempty"`
}

type HuanXiReqGeo struct {
	Lon  float64 `json:"lon,omitempty"`  // Required      Longitude 经度
	Lat  float64 `json:"lat,omitempty"`  // Required      Latitude 纬度
	Addr string  `json:"addr,omitempty"` //    Optional      Address 详细地址
}

type HuanXiRes struct {
	Res   int          `json:"res"`          // 返回状态码，0=成功，-1=失败
	Ad    *HuanXiResAd `json:"ad,omitempty"` // Ad  广告物料
	Bidid string       `json:"bidid,omitempty"`
}

type HuanXiResAd struct {
	Adw      int            `json:"adw,omitempty"`      // 广告宽(px)
	Adh      int            `json:"adh,omitempty"`      // 广告高(px)
	Img      string         `json:"img,omitempty"`      //  广告图主图
	Img2     string         `json:"img2,omitempty"`     //  广告图副图
	Img3     string         `json:"img3,omitempty"`     //  广告图副图
	Guidepic string         `json:"guidepic,omitempty"` //  引导图
	Title    string         `json:"title,omitempty"`    //  广告标题
	Text     string         `json:"text,omitempty"`     //  广告文案
	Click    []string       `json:"click,omitempty"`    //string    点击上报地址
	Imp      []string       `json:"imp,omitempty"`      //string    展示上报地址
	Act      int            `json:"act,omitempty"`      // 交互方式，1=跳转落地页，2=下载，3=跳转落地页下载（比如广点通），4=DeepLink
	Lpg      string         `json:"lpg,omitempty"`      //  广告跳转落地页链接，支持重定向
	Dplink   string         `json:"dplink,omitempty"`   //  deeplink链接，如果不为空，则点击的时候要优先处理，唤醒失败再调用lpg落地页链接
	Ulk      string         `json:"ulk,omitempty"`      //  iOS universal link 通用链接
	Qak      string         `json:"qak,omitempty"`
	Logo     string         `json:"logo,omitempty"`     //  logo图
	Icon     string         `json:"icon,omitempty"`     //  icon图
	Html     string         `json:"html,omitempty"`     //  动态广告代码
	Adext    HuanXiResAdExt `json:"adext,omitempty"`    // AdExt    广告扩展数据对象
	Videoext HuanXiResVideo `json:"videoext,omitempty"` // VideoExt 视频扩展数据
}

type HuanXiResAdExt struct {
	Pkg          string   `json:"pkg,omitempty"`          //  包名
	Nurl         []string `json:"nurl,omitempty"`         //string    竞价成功上报地址
	Lurl         []string `json:"lurl,omitempty"`         //string    竞价失败上报地址
	Carurl       []string `json:"carurl,omitempty"`       //string    点击上报，POST，请参考【附录一】
	Downbegin    []string `json:"downbegin,omitempty"`    //string    开始下载上报地址
	Downsucc     []string `json:"downsucc,omitempty"`     //string    下载成功上报地址
	Installbegin []string `json:"installbegin,omitempty"` //string    开始安装上报地址
	Installsucc  []string `json:"installsucc,omitempty"`  //string    安装成功上报地址
	Appactive    []string `json:"appactive,omitempty"`    //string    下载类广告，下载且安装完成之后,打开应用时上报
	Ktbegin      []string `json:"ktbegin,omitempty"`      //string    deeplink开始唤醒上报地址
	Kt           []string `json:"kt,omitempty"`           //string    deeplink唤醒成功上报地址
	Ktfail       []string `json:"ktfail,omitempty"`       //string    deeplink唤醒失败上报地址
	Price        float64  `json:"price,omitempty"`        //  出价，单位：分/CPM，RTB结算方式时有值
	Currency     string   `json:"currency,omitempty"`     //  币种，人民币=CNY，美元=USD，默认CNY
}
type HuanXiResVideo struct {
	Vurl                       string   `json:"vurl,omitempty"`                       //  视频地址
	Duration                   int      `json:"duration,omitempty"`                   // 播放时长(单位秒)
	Keep                       int      `json:"keep,omitempty"`                       // 持续播放时长后可跳过(单位秒)
	Vhtml                      string   `json:"vhtml,omitempty"`                      //  视频封面HTML代码
	Lpic                       string   `json:"lpic,omitempty"`                       //  视频封面图
	Endcard_img                string   `json:"endcard_img,omitempty"`                //  视频播放完后需要展示的图片地址
	Endcard_html               string   `json:"endcard_html,omitempty"`               //  视频播放完后需要展示的HTML代码
	Video_first_urls           []string `json:"video_first_urls,omitempty"`           //string    视频播放25%时上报地址
	Video_mid_urls             []string `json:"video_mid_urls,omitempty"`             //string    视频播放50%时上报地址
	Video_third_urls           []string `json:"video_third_urls,omitempty"`           //string    视频播放75%时上报地址
	Video_begin_urls           []string `json:"video_begin_urls,omitempty"`           //string    视频开始播放上报地址
	Video_end_urls             []string `json:"video_end_urls,omitempty"`             //string    视频播放完成上报地址
	Video_mute_urls            []string `json:"video_mute_urls,omitempty"`            //string    视频开启静音上报地址
	Video_unmute_urls          []string `json:"video_unmute_urls,omitempty"`          //string    视频关闭静音上报地址
	Video_skip_urls            []string `json:"video_skip_urls,omitempty"`            //string    跳过视频上报地址
	Video_close_urls           []string `json:"video_close_urls,omitempty"`           //string    关闭视频上报地址
	Video_pause_urls           []string `json:"video_pause_urls,omitempty"`           //string    视频暂停时上报地址
	Video_resume_urls          []string `json:"video_resume_urls,omitempty"`          //string    视频暂停再播放时上报地址
	Video_replay_urls          []string `json:"video_replay_urls,omitempty"`          //string    视频重播时上报地址
	Video_fullscreen_urls      []string `json:"video_fullscreen_urls,omitempty"`      //string    视频全屏时上报地址
	Video_exit_fullscreen_urls []string `json:"video_exit_fullscreen_urls,omitempty"` //string    视频退出全屏时上报地址
}
