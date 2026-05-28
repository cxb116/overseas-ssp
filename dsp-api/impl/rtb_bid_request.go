package impl

//
//// ==================== OpenRTB 2.6 核心对象 ====================
//
//// BidRequest 表示一个 OpenRTB 2.6 的竞标请求对象
//// 这是主要的顶级请求对象，包含一次竞价的所有信息
//type BidRequest struct {
//	// 交易所为这次竞价请求分配的唯一ID
//	Id string `json:"id,omitempty"`
//
//	// Imp 数组，包含一个或多个展示机会
//	// 每次请求必须至少包含一个 Imp 对象
//	Imp []Imp `json:"imp,omitempty"`
//
//	// Site 对象，提供有关出版商网站的详细信息
//	// 适用于 Web 环境的广告位
//	Site *Site `json:"site,omitempty"`
//
//	// App 对象，提供有关出版商应用程序的详细信息
//	// 适用于非浏览器环境的广告位
//	App *App `json:"app,omitempty"`
//
//	// Device 对象，包含用户设备的相关信息
//	Device *Device `json:"device,omitempty"`
//
//	// User 对象，包含用户的相关信息
//	User *User `json:"user,omitempty"`
//
//	// 测试模式标志：0=实时模式（可计费），1=测试模式（不可计费）
//	Test int32 `json:"test,omitempty"`
//
//	// 拍卖类型：1=第一价格，2=第二价格
//	// 大于 500 的值可用于交易所特定的拍卖类型
//	At int32 `json:"at,omitempty"`
//
//	// 交易所允许接收出价的最大时间（毫秒）
//	// 包括网络延迟，以避免超时
//	Tmax int32 `json:"tmax,omitempty"`
//
//	// 允许参与此次竞价的白名单买家席位列表
//	// 席位 ID 由交易所定义
//	Wset []string `json:"wset,omitempty"`
//
//	// 黑名单买家席位，不允许参与竞价的买家席位列表
//	Bseat []string `json:"bseat,omitempty"`
//
//	// 标志：1=所有 Imp 必须有出价才能被视为有效，0=部分出价也可接受
//	Allimps int32 `json:"allimps,omitempty"`
//
//	// 接受的货币 ISO-4217 三字母代码数组，如 ["USD", "EUR"]
//	Cur []string `json:"cur,omitempty"`
//
//	// 允许创意使用的语言代码数组，使用 ISO-639-1-alpha-2
//	Wlang []string `json:"wlang,omitempty"`
//
//	// 禁止的广告类别列表（IAB 内容分类）
//	Bcat []string `json:"bcat,omitempty"`
//
//	// 禁止的广告主域名列表，如 ["buyer1.com", "buyer2.org"]
//	Badv []string `json:"badv,omitempty"`
//
//	// 禁止的应用程序包名或站点域名列表
//	Bapp []string `json:"bapp,omitempty"`
//
//	// Source 对象，包含有关请求来源的信息
//	Source *Source `json:"source,omitempty"`
//
//	// Regs 对象，包含适用于请求的法律法规信息
//	Regs *Regs `json:"regs,omitempty"`
//
//	// Ext 扩展字段，用于容纳 OpenRTB 规范中未定义的自定义数据
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Imp (Impression) 表示一个广告展示机会
//// 单次请求可包含多个 Imp 对象
//type Imp struct {
//	// 展示机会的唯一ID，用于 BidResponse 中的 seatbid.bid.impid 字段关联
//	Id string `json:"id"`
//
//	// Metric 数组，包含出价方在竞价前可能需要的性能指标数据
//	Metric []Metric `json:"metric,omitempty"`
//
//	// Banner 对象，用于横幅展示广告
//	Banner *Banner `json:"banner,omitempty"`
//
//	// Video 对象，用于视频展示广告
//	Video *Video `json:"video,omitempty"`
//
//	// Audio 对象，用于音频展示广告
//	Audio *Audio `json:"audio,omitempty"`
//
//	// Native 对象，用于原生展示广告
//	Native *Native `json:"native,omitempty"`
//
//	// Pmp (Private Marketplace) 对象，用于私有交易
//	Pmp *Pmp `json:"pmp,omitempty"`
//
//	// 广告管理系统的名称
//	Displaymanager string `json:"displaymanager,omitempty"`
//
//	// 广告管理系统的版本
//	Displaymanagerver string `json:"displaymanagerver,omitempty"`
//
//	// 标志：1=插屏广告（广告全屏展示），0=非插屏广告
//	Instl int32 `json:"instl,omitempty"`
//
//	// 广告位标识符，由广告服务器或 SSP 使用
//	Tagid string `json:"tagid,omitempty"`
//
//	// 此 Imp 的最低出价金额（非负数，单位为分）
//	Bidfloor float64 `json:"bidfloor,omitempty"`
//
//	// bidfloor 使用的货币 ISO-4217 代码，默认为 USD
//	Bidfloorcur string `json:"bidfloorcur,omitempty"`
//
//	// 标志：用于指示点击后打开浏览器的偏好
//	Clickbrowser int32 `json:"clickbrowser,omitempty"`
//
//	// 标志：1=Imp 需要通过 HTTPS 协议返回创意，0=HTTP/HTTPS 均可
//	Secure int32 `json:"secure,omitempty"`
//
//	// iframe 破坏者列表，如果 Imp 在 iframe 中，这些域名可以突破 iframe
//	Iframebuster []string `json:"iframebuster,omitempty"`
//
//	// 响应式网页设计断点数组
//	Rwdd []string `json:"rwdd,omitempty"`
//
//	// 标志：1=服务器端广告插入（SSAI），0=非 SSAI
//	Ssai int32 `json:"ssai,omitempty"`
//
//	// 过期时间，缓存系统可使用此 Imp 的最长时间（秒）
//	Exp int32 `json:"exp,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Metric 包含出价方在竞价前可能需要的性能指标数据
//// 可用于优化出价决策
//type Metric struct {
//	// 指标类型，如 "view_rate", "click_rate", "completion_rate" 等
//	Type string `json:"type,omitempty"`
//
//	// 指标的数值
//	Value float64 `json:"value,omitempty"`
//
//	// 指标数据的来源方（如 IAB, Exchange 等）
//	Vendor string `json:"vendor,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Banner 表示横幅展示广告的规格
//type Banner struct {
//	// Format 数组，指定支持的广告尺寸
//	// 如果指定，则 W 和 H 字段应被忽略
//	Format []Format `json:"format,omitempty"`
//
//	// 广告宽度，单位为像素
//	W int64 `json:"w,omitempty"`
//
//	// 广告高度，单位为像素
//	H int64 `json:"h,omitempty"`
//
//	// 允许的横幅广告类型数组，基于 IAB 定义
//	Btype []int32 `json:"btype,omitempty"`
//
//	// 禁止的创意属性数组，基于 IAB 定义
//	Battr []int32 `json:"battr,omitempty"`
//
//	// 广告位置：0=未知，1=首屏上方，2=首屏下方，3=首屏，4=折叠下方
//	Pos int32 `json:"pos,omitempty"`
//
//	// 支持的 MIME 类型数组，如 ["image/jpeg", "image/gif"]
//	Mimes []string `json:"mimes,omitempty"`
//
//	// 标志：1=位于主框架（顶级窗口），0=位于 iframe
//	Topframe int32 `json:"json:"topframe,omitempty"`
//
//	// 允许的展开方向数组：1=左，2=右，3=上，4=下
//	Expdir []int32 `json:"expdir,omitempty"`
//
//	// 支持的 API 框架数组：1=VPAID 1.0，2=VPAID 2.0，3=MRAID-1，4=ORMMA，5=MRAID-2，6=MRAID-3
//	Api []int32 `json:"api,omitempty"`
//
//	// 视频创意的 VAST 广告版本
//	Vcm int32 `json:"vcm,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Format 指定支持的广告尺寸
//type Format struct {
//	// 宽度，单位为像素
//	W int64 `json:"w,omitempty"`
//
//	// 高度，单位为像素
//	H int64 `json:"h,omitempty"`
//
//	// 相对宽度的比例
//	Wr float64 `json:"wr,omitempty"`
//
//	// 相对高度的比例
//	Hr float64 `json:"hr,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Video 表示视频展示广告的规格
//type Video struct {
//	// 支持的 MIME 类型数组，如 ["video/mp4", "video/x-flv"]
//	Mimes []string `json:"mimes,omitempty"`
//
//	// 最小广告时长，单位为秒
//	Minduration int32 `json:"minduration,omitempty"`
//
//	// 最大广告时长，单位为秒
//	Maxduration int32 `json:"maxduration,omitempty"`
//
//	// 广告开始前的延迟时间：0=前贴片，-1=中贴片，-2=后贴片
//	Startdelay int32 `json:"startdelay,omitempty"`
//
//	// 视频广告连播（Pod）中广告的最大数量
//	MaxSeq int32 `json:"max_seq,omitempty"`
//
//	// 视频广告连播的预计总时长（秒）
//	Poddur int32 `json:"poddur,omitempty"`
//
//	// 支持的协议数组：1=VAST 1.0，2=VAST 2.0，3=VAST 3.0，4=VAST 1.0 Wrapper，5=VAST 2.0 Wrapper，6=VAST 3.0 Wrapper，7=VAST 4.0，8=VAST 4.0 Wrapper
//	Protocols []int32 `json:"protocols,omitempty"`
//
//	// 视频播放器宽度，单位为像素
//	W int64 `json:"w,omitempty"`
//
//	// 视频播放器高度，单位为像素
//	H int64 `json:"h,omitempty"`
//
//	// 视频广告连播的标识符
//	Podid string `json:"podid,omitempty"`
//
//	// 视频广告连播中的序列号，从 1 开始
//	Podseq int32 `json:"podseq,omitempty"`
//
//	// 请求的视频广告时长数组（秒）
//	Rqddurs []int32 `json:"rqddurs,omitempty"`
//
//	// 广告位置：1=前贴片，2=中贴片，3=后贴片，4=独立
//	Placement int32 `json:"placement,omitempty"`
//
//	// 广告线性：1=线性，2=非线性
//	Linearity int32 `json:"linearity,omitempty"`
//
//	// 是否允许跳过：0=不允许跳过，1=允许跳过
//	Skip int32 `json:"skip,omitempty"`
//
//	// 允许跳过前必须观看的最短时间（秒）
//	Skipmin int32 `json:"skipmin,omitempty"`
//
//	// 显示跳过按钮前必须观看的最短时间（秒）
//	Skipafter int32 `json:"skipafter,omitempty"`
//
//	// 如果视频是连播的一部分，指定广告的序列号
//	Sequence int32 `json:"sequence,omitempty"`
//
//	// 视频广告连播中的插槽位置
//	Slotinpod int32 `json:"slotinpod,omitempty"`
//
//	// 每秒最低 CPM（用于连播中的广告插播）
//	Mincpmpersec float32 `json:"mincpmpersec,omitempty"`
//
//	// 禁止的创意属性数组
//	Battr []int32 `json:"battr,omitempty"`
//
//	// 允许的最大广告扩展时长（秒）
//	Maxextended int32 `json:"maxextended,omitempty"`
//
//	// 最低比特率（Kbps）
//	Minbitrate int32 `json:"minbitrate,omitempty"`
//
//	// 最高比特率（Kbps）
//	Maxbitrate int32 `json:"maxbitrate,omitempty"`
//
//	// 是否允许信箱模式：0=不允许，1=允许
//	Boxingallowed int32 `json:"boxingallowed,omitempty"`
//
//	// 播放方式数组：1=点击播放，2=自动播放（有声），3=自动播放（静音）
//	Playbackmethod []int32 `json:"playbackmethod,omitempty"`
//
//	// 播放结束方式数组：1=滚动播放，2=点击播放
//	Playbackend []int32 `json:"playbackend,omitempty"`
//
//	// 投递方式数组：1=流媒体，2=渐进式下载
//	Delivery []int32 `json:"delivery,omitempty"`
//
//	// 广告位置：0=未知，1=前贴片，2=中贴片，3=后贴片
//	Pos int32 `json:"pos,omitempty"`
//
//	// 伴随广告数组
//	Companionad []Banner `json:"companionad,omitempty"`
//
//	// 支持的 API 框架数组
//	Api []int32 `json:"api,omitempty"`
//
//	// 伴随广告类型数组
//	Companiontype []int32 `json:"companiontype,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Audio 表示音频展示广告的规格
//type Audio struct {
//	// 支持的 MIME 类型数组
//	Mimes []string `json:"mimes,omitempty"`
//
//	// 最小广告时长，单位为秒
//	Minduration int32 `json:"minduration,omitempty"`
//
//	// 最大广告时长，单位为秒
//	Maxduration int32 `json:"maxduration,omitempty"`
//
//	// 音频广告连播的预计总时长（秒）
//	Poddur int32 `json:"poddur,omitempty"`
//
//	// 支持的协议数组
//	Protocols []int32 `json:"protocols,omitempty"`
//
//	// 广告开始前的延迟时间
//	Startdelay int32 `json:"startdelay,omitempty"`
//
//	// 请求的音频广告时长数组（秒）
//	Rqddurs []int32 `json:"rqddurs,omitempty"`
//
//	// 音频广告连播的标识符
//	Podid string `json:"podid,omitempty"`
//
//	// 音频广告连播中的序列号
//	Podseq int32 `json:"podseq,omitempty"`
//
//	// 序列号
//	Sequence int32 `json:"sequence,omitempty"`
//
//	// 插槽位置
//	Slotinpod int32 `json:"slotinpod,omitempty"`
//
//	// 每秒最低 CPM
//	Mincpmpersec float32 `json:"mincpmpersec,omitempty"`
//
//	// 禁止的创意属性数组
//	Battr []int32 `json:"battr,omitempty"`
//
//	// 允许的最大广告扩展时长（秒）
//	Maxextended int32 `json:"maxextended,omitempty"`
//
//	// 最低比特率（Kbps）
//	Minbitrate int32 `json:"minbitrate,omitempty"`
//
//	// 最高比特率（Kbps）
//	Maxbitrate int32 `json:"maxbitrate,omitempty"`
//
//	// 投递方式数组
//	Delivery []int32 `json:"delivery,omitempty"`
//
//	// 伴随广告数组
//	Companionad []Banner `json:"companionad,omitempty"`
//
//	// 支持的 API 框架数组
//	Api []int32 `json:"api,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Native 表示原生展示广告的规格
//// 原生广告的格式由请求方定义
//type Native struct {
//	// 原生广告请求对象（JSON 字符串），包含 Request 字段或 RequestNative 字段
//	// 建议使用 RequestNative 字段
//	Request string `json:"request,omitempty"`
//
//	// Version 字段，用于指定 Native 请求的版本
//	Ver string `json:"ver,omitempty"`
//
//	// 支持的原生广告 API 数组：1=动态内容，2=静态内容
//	Api []int32 `json:"api,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Pmp (Private Marketplace) 表示私有交易市场
//type Pmp struct {
//	// 私有交易 ID，唯一标识此次私有交易
//	PrivateAuction int32 `json:"private_auction,omitempty"`
//
//	// 私有交易数组，包含一个或多个交易对象
//	Deals []Deal `json:"deals,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Deal 表示一个私有交易
//type Deal struct {
//	// 交易的唯一标识符
//	Id string `json:"id"`
//
//	// 最低出价金额
//	Bidfloor float64 `json:"bidfloor,omitempty"`
//
//	// bidfloor 的货币代码
//	Bidfloorcur string `json:"bidfloorcur,omitempty"`
//
//	// 指定允许参与此交易的买家席位列表
//	At int32 `json:"at,omitempty"`
//
//	// 允许的买家席位 ID 数组
//	Wseat []string `json:"wseat,omitempty"`
//
//	// 交易类型：1=首选交易，2=竞价交易，3=私有交易
//	DealType string `json:"dealtype,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Site 表示网站的相关信息
//type Site struct {
//	// 由交易所分配的网站 ID
//	Id string `json:"id,omitempty"`
//
//	// 网站名称
//	Name string `json:"name,omitempty"`
//
//	// 网站域名，如 "example.com"
//	Domain string `json:"domain,omitempty"`
//
//	// 网站分类数组（IAB 内容分类）
//	Cat []string `json:"cat,omitempty"`
//
//	// 网站分类数组（IAB 内容分类 v2）
//	Sectioncat []string `json:"sectioncat,omitempty"`
//
//	// 页面分类数组（IAB 内容分类 v2）
//	Pagecat []string `json:"pagecat,omitempty"`
//
//	// 页面 URL
//	Page string `json:"page,omitempty"`
//
//	// 引用页 URL
//	Ref string `json:"ref,omitempty"`
//
//	// 搜索关键词
//	Search string `json:"search,omitempty"`
//
//	// 移动应用：0=否，1=是
//	Mobile int32 `json:"mobile,omitempty"`
//
//	// 预定义受众细分
//	Publisher *Publisher `json:"publisher,omitempty"`
//
//	// 内容对象
//	Content *Content `json:"content,omitempty"`
//
//	// 关键词数组
//	Keywords string `json:"keywords,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Publisher 表示出版商信息
//type Publisher struct {
//	// 出版商 ID
//	Id string `json:"id,omitempty"`
//
//	// 出版商名称
//	Name string `json:"name,omitempty"`
//
//	// 出版商域名
//	Domain string `json:"domain,omitempty"`
//
//	// 出版商分类数组
//	Cat []string `json:"cat,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Content 表示内容信息
//type Content struct {
//	// 内容 ID
//	Id string `json:"id,omitempty"`
//
//	// 内容片段（如文章）的标题
//	Title string `json:"title,omitempty"`
//
//	// 内容分类数组（IAB 内容分类）
//	Cat []string `json:"cat,omitempty"`
//
//	// 内容片段的评分
//	Contentrating string `json:"contentrating,omitempty"`
//
//	// 指示内容是否适合一般受众
//	Userrating string `json:"userrating,omitempty"`
//
//	// 上下文类型：1=视频，2=游戏，3=音乐，4=应用，5=文本，6=其他，7=未知
//	Context int32 `json:"context,omitempty"`
//
//	// 内容播放进度：0=已开始播放，1=未开始播放
//	Embedcontext int32 `json:"embedcontext,omitempty"`
//
//	// 媒体对象（如果是视频或音频内容）
//	Media *Media `json:"media,omitempty"`
//
//	// 关键词
//	Keywords string `json:"keywords,omitempty"`
//
//	// 延迟直播秒数
//	Livestream int32 `json:"livestream,omitempty"`
//
//	// 内容来源资源
//	SourceRelationship int32 `json:"sourceRelationship,omitempty"`
//
//	// 内容的长度
//	Len int32 `json:"len,omitempty"`
//
//	// 语言代码
//	Qagmediarating int32 `json:"qagmediarating,omitempty"`
//
//	// 标签
//	Tags []string `json:"tags,omitempty"`
//
//	// 生产者对象
//	Producer *Producer `json:"producer,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Media 表示媒体内容信息
//type Media struct {
//	// 媒体类型（MIME 类型）
//	Mimetype string `json:"mimetype,omitempty"`
//
//	// 媒体文件的 MIME 类型
//	AudioType string `json:"audioType,omitempty"`
//
//	// 媒体文件使用的编解码器
//	CreativeApi int32 `json:"creativeApi,omitempty"`
//
//	// 视频时长（秒）
//	Duration int32 `json:"duration,omitempty"`
//
//	// 视频质量
//	Videoq string `json:"videoq,omitempty"`
//
//	// 媒体高度
//	Height int32 `json:"height,omitempty"`
//
//	// 媒体宽度
//	Width int32 `json:"width,omitempty"`
//
//	// 缩略图 URL
//	Thumbnail string `json:"thumbnail,omitempty"`
//
//	// 媒体 bitrate
//	Bitrate int32 `json:"bitrate,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Producer 表示内容生产者信息
//type Producer struct {
//	// 生产者 ID
//	Id string `json:"id,omitempty"`
//
//	// 生产者名称
//	Name string `json:"name,omitempty"`
//
//	// 生产者分类数组
//	Cat []string `json:"cat,omitempty"`
//
//	// 生产者域名
//	Domain string `json:"domain,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// App 表示应用程序的相关信息
//type App struct {
//	// 应用 ID（交易所分配或第三方 ID）
//	Id string `json:"id,omitempty"`
//
//	// 应用名称
//	Name string `json:"name,omitempty"`
//
//	// 应用包名（如 Android 的 com.example.app）
//	Bundle string `json:"bundle,omitempty"`
//
//	// 应用域名
//	Domain string `json:"domain,omitempty"`
//
//	// 应用商店 URL
//	Storeurl string `json:"storeurl,omitempty"`
//
//	// 应用分类数组（IAB 内容分类）
//	Cat []string `json:"cat,omitempty"`
//
//	// 版本号
//	Ver string `json:"ver,omitempty"`
//
//	// 隐私政策 URL
//	Privacypolicy int32 `json:"privacypolicy,omitempty"`
//
//	// 是否已付费：0=免费，1=付费
//	Paid int32 `json:"paid,omitempty"`
//
//	// 出版商对象
//	Publisher *Publisher `json:"publisher,omitempty"`
//
//	// 内容对象
//	Content *Content `json:"content,omitempty"`
//
//	// 关键词
//	Keywords string `json:"keywords,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Device 表示用户设备的相关信息
//type Device struct {
//	// 设备唯一标识符（建议使用哈希值）
//	Ua string `json:"ua,omitempty"`
//
//	// 地理位置对象
//	Geo *Geo `json:"geo,omitempty"`
//
//	// 代理检测：0=未知，1=检测到代理，2=检测到 VPN
//	Dnt int32 `json:"dnt,omitempty"`
//
//	// 限制跟踪：0=未知，1=启用限制跟踪
//	Lmt int32 `json:"lmt,omitempty"`
//
//	// IP 地址（IPv4）
//	Ip string `json:"ip,omitempty"`
//
//	// IPv6 地址
//	Ipv6 string `json:"ipv6,omitempty"`
//
//	// 设备类型：1=手机，2=平板，3=个人电脑，4=联网电视，5=机顶盒，6=其他
//	Devicetype int32 `json:"devicetype,omitempty"`
//
//	// 运营商 ID（MCC-MNC）
//	Make string `json:"make,omitempty"`
//
//	// 设备制造商
//	Model string `json:"model,omitempty"`
//
//	// 设备型号
//	Os string `json:"os,omitempty"`
//
//	// 操作系统
//	Osv string `json:"osv,omitempty"`
//
//	// 操作系统版本
//	Hwv string `json:"hwv,omitempty"`
//
//	// 硬件版本
//	H int64 `json:"h,omitempty"`
//
//	// 屏幕高度，单位为像素
//	W int64 `json:"w,omitempty"`
//
//	// 屏幕宽度，单位为像素
//	Ppi int32 `json:"ppi,omitempty"`
//
//	// 屏幕像素密度（PPI）
//	Pxratio float64 `json:"pxratio,omitempty"`
//
//	// 屏幕像素比例
//	Flashver string `json:"flashver,omitempty"`
//
//	// Flash 版本
//	Language string `json:"language,omitempty"`
//
//	// 设备语言代码（ISO-639-1-alpha-2）
//	Carrier string `json:"carrier,omitempty"`
//
//	// 运营商名称
//	Mccmnc string `json:"mccmnc,omitempty"`
//
//	// 移动国家代码 + 移动网络代码（MCC-MNC）
//	Connectiontype int32 `json:"connectiontype,omitempty"`
//
//	// 网络连接类型：0=未知，1=以太网，2=WiFi，3=蜂窝网络（2G），4=蜂窝网络（3G），5=蜂窝网络（4G），6=蜂窝网络（5G），7=其他
//	IFA string `json:"ifa,omitempty"`
//
//	// 广告标识符（iOS 的 IDFA 或 Android 的 GAID）
//	IfaType int32 `json:"ifa_type,omitempty"`
//
//	// IFA 类型：0=未知，1=IDFA，2=GAID
//	DIDSHA1 string `json:"didsha1,omitempty"`
//
//	// 设备 ID 的 SHA-1 哈希值
//	DIDMD5 string `json:"didmd5,omitempty"`
//
//	// 设备 ID 的 MD5 哈希值
//	DPIDSHA1 string `json:"dpidsha1,omitempty"`
//
//	// 平台设备 ID（如 OpenUDID）的 SHA-1 哈希值
//	DPIDMD5 string `json:"dpidmd5,omitempty"`
//
//	// 平台设备 ID（如 OpenUDID）的 MD5 哈希值
//	MACSHA1 string `json:"macsha1,omitempty"`
//
//	// MAC 地址的 SHA-1 哈希值
//	MACMD5 string `json:"macmd5,omitempty"`
//
//	// MAC 地址的 MD5 哈希值
//	Ext interface{} `json:"ext,omitempty"`
//
//	// 扩展字段
//	IMEI string `json:"imei,omitempty"`
//
//	// 国际移动设备识别码（建议使用哈希值）
//	ANDROIDID string `json:"androidid,omitempty"`
//
//	// Android ID（建议使用哈希值）
//}
//
//// Geo 表示地理位置信息
//type Geo struct {
//	// 纬度，范围 -90.0 到 +90.0
//	Lat float64 `json:"lat,omitempty"`
//
//	// 经度，范围 -180.0 到 +180.0
//	Lon float64 `json:"lon,omitempty"`
//
//	// 地理位置类型：1=GPS/位置服务，2=IP 地址，3=用户提供
//	Type int32 `json:"type,omitempty"`
//
//	// 精度：0=未知，1=精确，2=近似，3=区域，4=国家，5=全球
//	Accuracy int32 `json:"accuracy,omitempty"`
//
//	// 上次位置更新的时间（Unix 时间戳，毫秒）
//	Lastfix int64 `json:"lastfix,omitempty"`
//
//	// GPS 服务：0=未知，1=启用，2=禁用
//	Ipservice int32 `json:"ipservice,omitempty"`
//
//	// 确定地理位置的提供商
//	Country string `json:"country,omitempty"`
//
//	// 国家代码（ISO-3166-1-alpha-3）
//	Region string `json:"region,omitempty"`
//
//	// 区域代码（ISO-3166-2）
//	RegionFips104 string `json:"regionfips104,omitempty"`
//
//	// 区域 FIPS 104 代码
//	Metro string `json:"metro,omitempty"`
//
//	// 大都市区域代码（Google Metro ID）
//	City string `json:"city,omitempty"`
//
//	// 城市名称
//	Zip string `json:"zip,omitempty"`
//
//	// 邮政编码
//	Utcoffset int32 `json:"utcoffset,omitempty"`
//
//	// UTC 偏移量（分钟）
//	Ext interface{} `json:"ext,omitempty"`
//
//	// 扩展字段
//}
//
//// User 表示用户的相关信息
//type User struct {
//	// 用户 ID（建议使用哈希值）
//	Id string `json:"id,omitempty"`
//
//	// 买家用户 ID（如需定向用户）
//	Buyeruid string `json:"buyeruid,omitempty"`
//
//	// 出生年份，格式 YYYY
//	Yob int32 `json:"yob,omitempty"`
//
//	// 性别：M=男，F=女，O=其他
//	Gender string `json:"gender,omitempty"`
//
//	// 关键词
//	Keywords string `json:"keywords,omitempty"`
//
//	// 自定义数据（由交易所定义）
//	Customdata string `json:"customdata,omitempty"`
//
//	// Geo 对象，用户注册时提供的地理位置
//	Geo *Geo `json:"geo,omitempty"`
//
//	// 数据提供者对象数组
//	Data []Data `json:"data,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Data 表示第三方数据提供者的信息
//type Data struct {
//	// 数据提供者 ID
//	Id string `json:"id,omitempty"`
//
//	// 数据提供者名称
//	Name string `json:"name,omitempty"`
//
//	// 数据片段数组
//	Segment []Segment `json:"segment,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Segment 表示数据片段
//type Segment struct {
//	// 片段 ID
//	Id string `json:"id,omitempty"`
//
//	// 片段名称
//	Name string `json:"name,omitempty"`
//
//	// 片段值
//	Value string `json:"value,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Source 表示请求来源信息
//type Source struct {
//	// 交易所标识：0=交易所自身，1=买方，2=卖家
//	Fd int32 `json:"fd,omitempty"`
//
//	// 交易 ID，用于追踪请求路径
//	Tid string `json:"tid,omitempty"`
//
//	// 付费链，用于追踪需求方平台的所有权
//	Pchain string `json:"pchain,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
//
//// Regs 表示法律法规相关信息
//type Regs struct {
//	// COPPA（儿童在线隐私保护法）：0=否，1=是
//	Coppa int32 `json:"coppa,omitempty"`
//
//	// GDPR（通用数据保护条例）：0=不适用，1=适用
//	Gdpr int32 `json:"gdpr,omitempty"`
//
//	// 美国隐私标志（US Privacy String）
//	UsPrivacy string `json:"us_privacy,omitempty"`
//
//	// 扩展字段
//	Ext interface{} `json:"ext,omitempty"`
//}
