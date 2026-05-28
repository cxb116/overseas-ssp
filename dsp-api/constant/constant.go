package constant

// 在redis中存储控制台用户令牌的密钥
const RedisSubscribeMessageChannel = "redis:subscribe-message-channel"

const NOT_WORKER = -1 // 没有可用workerId
//
const ETCD_KEY = "/dsp/config"

const (
	// 新增公司, 数据库新增成功以后回填的id,再去put到etcd中
	ADD_COMPANY = "/company/add/"
	// 修改公司
	UPDATE_COMPANY = "/company/update/"
	// 新增预算位
	ADD_DSP = "/dsp/add/"
	// 修改预算
	UPDATE_DSP = "/dsp/update/"
	// 新增权重表
	ADD_LAUNCH = "/launch/add/"
	// 修改权重表
	UPDATE_LAUNCH = "/launch/update/"
)

const (
	EventAllAdd = 1 // 后台首次发布全部数据
	EventAdd    = 2 // 后台添加数据
	EventRemove = 3 // 后台删除数据
	EventUpdate = 4 // 后台跟新数据

	EventCompayAdd    = 5
	EventCompayRemove = 6
	EventCompayUpdate = 7
)

// 同一个预算下的流量权重
const TrafficValue = 100 // 分流总值

// 响应状态值
const (
	REQ_CODE_SUC = 1 // 有填充
	//REQ_CODE_FAIL       = 2 // 请求失败，这个是统一失败，用来判断是否成功的标准
	REQ_CODE_SLOT_ID    = 2 // 广告位id不存在
	REQ_CODE_DEVICE_ERR = 3 // 设备信息异常
	REQ_CODE_DEVICE_SUC = 4 // 设备信息正常

	// 预算相关的状态码
	REQ_CODE_DSP_TIMEOUT        = 100 // DSP请求超时
	REQ_CODE_DSP_ERROR          = 101 // 上游DSP异常
	REQ_CODE_DSP_NIL            = 102 // 没填充
	REQ_CODE_DSP_NOT_SLOT_ID    = 103 // 媒体广告位id没有找到预算
	REQ_CODE_DSP_NOT_MATCH      = 104 // 媒体流量没有找到对应的预算dspCode
	REQ_CODE_DSP_NOT_FLOORPRICE = 105 // 底价过滤

	REQ_CODE_DSP_NOT_AD = -1
)

// 结算方式
const (
	PAY_PRICE_COMPOSITION = 1 //分成
	PAY_PRICE_RTB         = 2
)

const (
	POST_METHOD = "POST"
	GET_METHOD  = "GET"
)

const (
	REQ_ANDROID = 1
	REQ_IOS     = 2
)

const (
	KAFKA_BUDGET_CONSUMER_GROUP  = "budget-consumer-group"  // 预算请求响应分组
	KAFKA_TRAFFIC_CONSUMER_GROUP = "traffic-consumer-group" // 流量响应分组

)

// 本次是请求还是响应 req_type
const (
	SSP = 1
	DSP = 2
)

// 请求是否响应成功  status
const (
	SUCCESS = 1
	FAIL    = 2
)

// 当kafka status为2时， 这个就是消息失败原因，作为归因使用  event
const (
	DEFAULT     = 0 //
	TIMEOUT     = 1 // 超时
	NOTMATCH    = 2 // 物料不匹配
	LESSPRICE   = 3 // 低于底价
	AllTIMEOUT  = 4 // 流程超时
	ERRORDEVICE = 5 // 请求流量参数错误
)

// 交互方式
const (
	OPEN               = 1 //打开网页
	DEEPLINK           = 2 // deeplink
	DOWNLOAD           = 3 // 直接下载
	GDT                = 4 // 广点通
	WECHAT             = 5 // 小程序
	APP_STORE_DOWNLOAD = 6 // APP 商店下载
	QUICK_APP          = 7 // 快应用
)

const (
	ON  = 1
	OFF = 0
)
