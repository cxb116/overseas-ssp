package dsp_obj

type JunYiMengReq struct {
	Req_id      string    `json:"req_id,omitempty"`      // 请求id，唯⼀标识⼀次⼴告请求；由媒体侧⽣ 成，请确保全局唯⼀
	Api_version string    `json:"api_version,omitempty"` // app API版本号
	App         ReqApp    `json:"app,omitempty"`         // 应用信息
	Slot        ReqSlot   `json:"slot,omitempty"`        // 广告位信息
	Device      ReqDevice `json:"device,omitempty"`      //  设备信息
	Network     Network   `json:"network,omitempty"`     //  网络信息
}

type ReqApp struct {
	Add_id      string `json:"add_id,omitempty"`      //应用ID,由平台提供
	App_version string `json:"app_version,omitempty"` //应用版本号
	App_pkg     string `json:"app_pkg,omitempty"`     //应用包名
	App_name    string `json:"app_name,omitempty"`    //应用名称
}

type ReqSlot struct {
	Slot_id   string `json:"slot_id,omitempty"`   //是 广告位ID，由平台提供
	Slot_w    int    `json:"slot_w,omitempty"`    //是 广告位宽度
	Slot_h    int    `json:"slot_h,omitempty"`    //是 广告位高度
	Bid_type  int    `json:"bid_type"`            //否 竞价类型：1:CPM,2:CPC 当前仅支持CPM,(不填默认为0，代表非竞价)
	Bid_floor int    `json:"bid_floor,omitempty"` //否 底价(分)-竞价时必传
}

type ReqDevice struct {
	Os_type           int       `json:"os_type,omitempty"`           //	是 设备操作系统类型: 1-安卓，2-ios
	Os_version        string    `json:"os_version,omitempty"`        //	是 备操作系统版本。deviceType为1或2时，该字段必填
	Brand             string    `json:"brand,omitempty"`             //	是设备品牌。deviceType为1或2时，该字段必填，iOS系统设备请统⼀填写为Apple
	Model             string    `json:"model,omitempty"`             //	是 设备型号
	Vendor            string    `json:"vendor,omitempty"`            //	是 设备厂商
	Width             int       `json:"width,omitempty"`             //	是 设备屏幕宽度
	Height            int       `json:"height,omitempty"`            //	是 设备屏幕高度
	Density           float64   `json:"density,omitempty"`           //	是屏幕像素密度：每英寸像素,获取方法：安卓：context.getResources().getDisplayMetrics().density iOS：UIScreen.scale
	Ppi               int       `json:"ppi,omitempty"`               //	是 屏幕大小(单位 ppi,每 英寸所有的像素)
	Dpi               int       `json:"dpi,omitempty"`               //	是 屏幕像素密度: 此数据影响广告返回的清晰度，eg:150、440
	Inch              float32   `json:"inch,omitempty"`              //	是 屏幕尺寸: eg:6.4 、6.1
	Ua                string    `json:"ua,omitempty"`                //	是 客户端UserAgent
	Imei              string    `json:"imei,omitempty"`              //	否 设备imei，android手机必填
	Oaid              string    `json:"oaid,omitempty"`              //	否 安卓10以上必填
	Android_id        string    `json:"android_id,omitempty"`        //	否 安卓ID,安卓必填
	Idfa              string    `json:"idfa,omitempty"`              //	否 苹果idfa,苹果必填
	Openudid          string    `json:"openudid,omitempty"`          //	否 苹果openudid,苹果设备必填
	Idfv              string    `json:"idfv,omitempty"`              //	否 苹果设备idfv
	Mac               string    `json:"mac,omitempty"`               //	是 设备mac地址
	Hms_ver_code      string    `json:"hms_ver_code,omitempty"`      // 建议  HMS Core 版本号,华为实现静默安装 依赖的服务，HMS core包名com.huawei.hwid （华为手机必传）
	Hwag_ver          string    `json:"hwag_ver,omitempty"`          //	华为安卓设备的 AG(应用市场)的版 本号保留原始值。华为手机必填
	Boot_mark         string    `json:"boot_mark,omitempty"`         //	系统启动标识,例如：iOS：1623815045.970028 ；安卓：ec7f4f33411a-47bc-8067-744a4e7e0723
	Update_mark       string    `json:"update_mark,omitempty"`       //	系统更新标识，例如：iOS：1581141691.570419583；安卓：1004697.709999999
	Rom_version       string    `json:"rom_version,omitempty"`       //	手机ROM版本
	Sys_comp_time     string    `json:"sys_comp_time,omitempty"`     //  系统编译时间，时间戳，精确到毫秒，如：1545362006000
	App_store_version string    `json:"app_store_version,omitempty"` //  手机自带的应用商店的版本
	Serialno          string    `json:"serialno,omitempty"`          //  系统设备序列号
	Ali_aaid          string    `json:"ali_aaid,omitempty"`          //  阿里巴巴匿名设备标识，需集成阿里SDK 获取，阿里预算必填
	Birth_time        string    `json:"birth_time,omitempty"`        //  系统初始化时间
	Startup_time      string    `json:"startup_time,omitempty"`      //  设备启动时间
	Mb_time           string    `json:"mb_time,omitempty"`           //  系统版本更新时间
	Mem_total         int64     `json:"mem_total,omitempty"`         //  系统总内存空间
	Disk_total        int64     `json:"disk_total,omitempty"`        //  磁盘总空间
	Country_code      string    `json:"country_code,omitempty"`      //  ocal地区，如“CN”，ios需要回传，安卓建议填写该字段
	Local_tz_name     string    `json:"local_tz_name,omitempty"`     //  ocal时区，如"28800"，ios需要回传，安卓建议填写该字段
	Language          string    `json:"language,omitempty"`          //  设备设置的语言：如"zh-Hans-CN" ，ios需要回传，安卓建议填写该字段
	Phone_name        string    `json:"phone_name,omitempty"`        //  设备名称
	Cpu_num           int       `json:"cpu_num,omitempty"`           //  CPU数目，如 4，仅ios需要回传，安卓可不填写该字段
	Cpu_frequency     float64   `json:"cpu_frequency,omitempty"`     //  手机CPU频率，单位:GHzIOS操作系统必传例如:2.2
	Hardware_model    string    `json:"hardware_model,omitempty"`    // ios需要回传设备的 model 值, iOS 必传
	Machine           string    `json:"machine,omitempty"`           // ios需要回传 设备的 machine 值, iOS 必传
	Auth_status       int       `json:"auth_status"`                 // ios需要回传iOS 广告标识授权情况，是否允许获取IDFA： 0：未确定 1：受限制 2：被拒绝 3：授权
	Osl               int       `json:"osl,omitempty"`               // 否 安卓系统版本系统级别
	Paid              string    `json:"paid,omitempty"`              // 否 拼多多的 PAID，如媒体支持可直接传入
	AppList           string    `json:"appList,omitempty"`           // 否 已安装应用包名，多个用逗号隔开
	Geo               DeviceGeo `json:"geo,omitempty"`               // 否 位置信息
	Caids             []Caid    `json:"caids,omitempty"`
}

type Caid struct {
	Caid    string `json:"caid,omitempty"`
	Version string `json:"version,omitempty"`
}

type DeviceGeo struct {
	Lat float64 `json:"lat,omitempty"` //PS获取的维度信息，不传可能影响广告填充
	Lon float64 `json:"lon,omitempty"` //GPS获取的经度信息，不传可能影响广告填充
	Ts  int64   `json:"ts,omitempty"`  //10位时间戳 精确到s
}

type Network struct {
	Ip       string `json:"ip,omitempty"`       //是 客户端ip地址
	Ipv6     string `json:"ipv6,omitempty"`     //否 ipv6版本，与ip一起必须存在一个有效值
	Net_type int    `json:"net_type"`           //客户端网络类型： 0.未知 1.wifi 2.2G3.3G 4.4G 5.5G
	Net_op   string `json:"net_op,omitempty"`   //是 运营商信息： 46000(中国移动)46001(中国联调) 46003(中国电信)
	Imsi     string `json:"imsi,omitempty"`     //是 手机imsi
	Ssid     string `json:"ssid,omitempty"`     //否 无线网ssid名称 移动端必填，无线网ssid 名称，如获取不到可传空（影响填充）
	Wifi_mac string `json:"wifi_mac,omitempty"` // wifi 路由器MAC地址 移动端必填， WIFI路由器MAC地址，如获取不到 可传空(影响填充）；例如：wifi mac地 址 20:a6:cd:7e:e3:60
}

type JunYiMengRes struct {
	Code int           `json:"code"`           //	0-正常，-1-存在报错信息（报错详情见msg）
	Msg  string        `json:"msg,omitempty"`  //	具体错误信息
	Data *ResponseData `json:"data,omitempty"` //	广告内容
	Ts   int           `json:"ts,omitempty"`   //	当前响应时间戳（秒）
	Cost int           `json:"cost,omitempty"` //	响应耗时（毫秒）

}

type ResponseData struct {
	Seid    string           `json:"seid,omitempty"`    // 是  系统唯一标识id
	Req_id  string           `json:"req_id,omitempty"`  // 否	请求时唯一标识id（请求携带该传参则返回
	Ad_info []ResponseAdInfo `json:"ad_info,omitempty"` // 否 	广告素材
	Price   int              `json:"price,omitempty"`   // 价格，单位：分/CPM
}

type ResponseAdInfo struct {
	Ad_id                  int64              `json:"ad_id,omitempty"`                  //	是 	广告ID
	Action_type            int                `json:"action_type,omitempty"`            //  是	点击广告后的操作类型:1:  打开页面2：下载3：GDT下载4：DEEPLINK WEB5 :  DEEPLINK DOWNLOAD6：DEEPLINK APP
	Ad_type                int                `json:"ad_type,omitempty"`                //  是	1:普通广告 2:激励视频 3:资讯信息流
	Title                  string             `json:"title,omitempty"`                  //	是	标题
	Desc                   string             `json:"desc,omitempty"`                   //	否	描述
	Image_urls             []string           `json:"image_urls,omitempty"`             //	否	图片素材地址
	Icon_url               string             `json:"icon_url,omitempty"`               //	否	图标地址
	Logo_url               string             `json:"logo_url,omitempty"`               //	否	广告logo地址
	Click_url              string             `json:"click_url,omitempty"`              //	否	落地页地址
	Click_url_with_header  *TrackingHeaderUrl `json:"click_url_with_header,omitempty"`  //	否	落地页地址，当该项有值时，优 先 使用进行跳转
	Download_url           string             `json:"download_url,omitempty"`           //	否	下载地址，当action_type=2，3时该地址存在
	Deeplink_url           string             `json:"deeplink_url,omitempty"`           //	否	deeplink地址
	Universalurl           string             `json:"universalurl,omitempty"`           // unlk > deek > cl
	Deeplink_success_urls  []string           `json:"deeplink_success_urls,omitempty"`  //	否	dp吊起成功监控地址
	Deeplink_fail_urls     []string           `json:"deeplink_fail_urls,omitempty"`     //	否	dp吊起失败监控地址
	Show_urls              []string           `json:"show_urls,omitempty"`              //	是	曝光监控地址
	Click_urls             []string           `json:"click_urls,omitempty"`             //	是	点击监控地址
	Download_start_urls    []string           `json:"download_start_urls,omitempty"`    //	否	开始下载监控地址
	Download_end_urls      []string           `json:"download_end_urls,omitempty"`      //	否	下载完成监控地址
	Install_start_urls     []string           `json:"install_start_urls,omitempty"`     //	否	开始安装监控地址
	Install_end_urls       []string           `json:"install_end_urls,omitempty"`       //	否	安装完成监控地址
	Action_urls            []string           `json:"action_urls,omitempty"`            //	否	安装后激活打开App后上报
	Click_area_report_urls []string           `json:"click_area_report_urls,omitempty"` //	否	点击坐标打点上报，上报方法见 4.2, 部分预算可能返回
	App_info               *ResponseAppInfo   `json:"app_info,omitempty"`               //	否	app信息
	Video                  *ResponseVideo     `json:"video,omitempty"`                  //	否	视频预算
	Trackers_with_header   []TranckingHeader  `json:"trackers_with_header,omitempty"`   // 上报URL列表，必须由客户端上报， 上报时设置我方返回请求头格式（当 // 有值时，必须使用该项进行上报）

}

type TrackingHeaderUrl struct {
	Url     string         `json:"url,omitempty"`     //上报链接
	Headers []ReportHeader `json:"headers,omitempty"` //当进行监测上报时，需要使用此字段 循环里面的 key 之外，不要设置其他 的 head的 key 值，禁 止设置 cookie。 注意：如该字段为空，则无需特别设 置head，保持原有处理逻辑即可
}

type ResponseAppInfo struct {
	Name        string  `json:"name,omitempty"`        //	否 APP 名称
	Package     string  `json:"package,omitempty"`     //	否 APP 包名
	Icon_url    string  `json:"icon_url,omitempty"`    //	否 APP icon图标
	Score       float64 `json:"score,omitempty"`       //	否 APP 评分
	Version     string  `json:"version,omitempty"`     //	否 APP 版本
	Corporate   string  `json:"corporate,omitempty"`   //	否 APP 开发者名称
	Size        int64   `json:"size,omitempty"`        //	否 APP 大小
	Permissions string  `json:"permissions,omitempty"` //	否 APP 权限
	Privacy     string  `json:"privacy,omitempty"`     //	否 APP 隐私政策

}

type ResponseVideo struct {
	Video_duration  int         `json:"video_duration,omitempty"`  //是 视频时长（秒）
	Video_url       string      `json:"video_url,omitempty"`       //是 视频地址
	Play_start_urls []string    `json:"play_start_urls,omitempty"` //否 视频播放开始监控地址
	Play_end_urls   []string    `json:"play_end_urls,omitempty"`   //否 视频播放完成监控地址
	Play_close_urls []string    `json:"play_close_urls,omitempty"` //否 视频被关闭监控地址
	Play_skip_urls  []string    `json:"play_skip_urls,omitempty"`  //否 视频被跳过监控地址
	Play_tracers    *PlayTracer `json:"play_tracers,omitempty"`    //否 视频播放过程监控
	Has_end         bool        `json:"has_end,omitempty"`         //否 是否存在视频后贴
	Video_end       *VideoEnd   `json:"video_end,omitempty"`       //否 视频后贴
	Has_mid         bool        `json:"has_mid,omitempty"`         //否 是否存在视频中贴
	Video_mid       *VideoMid   `json:"video_mid,omitempty"`       //否 视频中贴
}

type PlayTracer struct {
	T    int      `json:"t,omitempty"`    //单位秒，播放到第t秒的时候发送监控 地址
	Urls []string `json:"urls,omitempty"` //监控地址
}

type VideoEnd struct {
	End_html                string   `json:"end_html,omitempty"`                // 如果该字段内容存在，播放完成后使用webview打开该html页面
	End_html_close_monitors []string `json:"end_html_close_monitors,omitempty"` //否 当用户关闭end_html落地页之后发送监控
	End_html_imp_monitors   []string `json:"end_html_imp_monitors,omitempty"`   //否 当展示end_html之后发送监控
	End_title               string   `json:"end_title,omitempty"`               //后贴内容标题
	End_desc                string   `json:"end_desc,omitempty"`                //后贴内容描述
	End_icon_url            string   `json:"end_icon_url,omitempty"`            //后贴logo地址
	End_cover_url           string   `json:"end_cover_url,omitempty"`           //后贴推广图地址
	End_action_text         string   `json:"end_action_text,omitempty"`         //后贴点击动作文案
	End_click_url           string   `json:"end_click_url,omitempty"`           //后贴点击落地页
	End_rating              float32  `json:"end_rating,omitempty"`              //后贴内容评分
	End_rating_count        int      `json:"end_rating_count,omitempty"`        //后贴内容评论数
}

type VideoMid struct {
	Mid_title       string `json:"mid_title,omitempty"`       //视频中贴标题
	Mid_desc        string `json:"mid_desc,omitempty"`        //视频中贴描述
	Mid_action_text string `json:"mid_action_text,omitempty"` //点击按钮文案
	Mid_logo        string `json:"mid_logo,omitempty"`        //广告logo
}

type TranckingHeader struct {
	Event       int                 `json:"event,omitempty"`       // 事件类型:1曝光 2点击 3吊起成功 4吊起失败 5开始下载 6下载完成 7开始安装 8安装完成 9视频开始播放 10视频播放结束 11视频关闭 12 视频跳过播放 13 视频播放1/4 14视频播放1/2 15视频播放3/4 16 视频加载成功 17 视频加载失败
	Header_urls []TrackingHeaderUrl `json:"header_urls,omitempty"` //上报Url对象
}

type ReportHeader struct {
	Key   string `json:"key,omitempty"`   // key
	Value string `json:"value,omitempty"` // value
}
