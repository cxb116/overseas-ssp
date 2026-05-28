package impl

import "sync"

var BidResponsePool = sync.Pool{
	New: func() any {
		return &BidResponse{}
	},
}

// 获取BidResponse
func GetBidResponse() *BidResponse {
	res := BidResponsePool.Get().(*BidResponse)
	res.Reset()
	return res
}

// 归还BidResponse
func PutBidResponse(res *BidResponse) {
	if res != nil {
		BidResponsePool.Put(res)
	}
}

// 清空响应体，防止内存泄漏
func (this *BidResponse) Reset() {
	this.Res = 0
	this.ResV = ""
	this.Message = ""
	this.Ad.Adw = 0
	this.Ad.Adh = 0
	this.Ad.Img = ""
	this.Ad.Img2 = ""
	this.Ad.Img3 = ""
	this.Ad.Guidepic = ""
	this.Ad.Title = ""
	this.Ad.Text = ""
	this.Ad.Click = this.Ad.Click[:0]
	this.Ad.Imp = this.Ad.Imp[:0]
	this.Ad.Act = 0
	this.Ad.Lpg = ""
	this.Ad.Dplink = ""
	this.Ad.Ulk = ""
	this.Ad.Qak = ""
	this.Ad.Logo = ""
	this.Ad.Icon = ""
	this.Ad.Html = ""
	this.Ad.Adext.Pkg = ""
	this.Ad.Adext.Nurl = this.Ad.Adext.Nurl[:0]
	this.Ad.Adext.Lurl = this.Ad.Adext.Lurl[:0]
	this.Ad.Adext.Carurl = this.Ad.Adext.Carurl[:0]
	this.Ad.Adext.Downbegin = this.Ad.Adext.Downbegin[:0]
	this.Ad.Adext.Downsucc = this.Ad.Adext.Downsucc[:0]
	this.Ad.Adext.Installbegin = this.Ad.Adext.Installbegin[:0]
	this.Ad.Adext.Installsucc = this.Ad.Adext.Installsucc[:0]
	this.Ad.Adext.Appactive = this.Ad.Adext.Appactive[:0]
	this.Ad.Adext.Ktbegin = this.Ad.Adext.Ktbegin[:0]
	this.Ad.Adext.Kt = this.Ad.Adext.Kt[:0]
	this.Ad.Adext.Ktfail = this.Ad.Adext.Ktfail[:0]
	this.Ad.Adext.Price = 0
	this.Ad.Adext.Currency = ""

	this.Ad.Videoext.Vurl = ""
	this.Ad.Videoext.Duration = 0
	this.Ad.Videoext.Keep = 0
	this.Ad.Videoext.Vhtml = ""
	this.Ad.Videoext.Lpic = ""
	this.Ad.Videoext.Endcard_img = ""
	this.Ad.Videoext.Endcard_html = ""
	this.Ad.Videoext.Video_first_urls = this.Ad.Videoext.Video_first_urls[:0]
	this.Ad.Videoext.Video_mid_urls = this.Ad.Videoext.Video_mid_urls[:0]
	this.Ad.Videoext.Video_third_urls = this.Ad.Videoext.Video_third_urls[:0]
	this.Ad.Videoext.Video_begin_urls = this.Ad.Videoext.Video_begin_urls[:0]
	this.Ad.Videoext.Video_end_urls = this.Ad.Videoext.Video_end_urls[:0]
	this.Ad.Videoext.Video_mute_urls = this.Ad.Videoext.Video_mute_urls[:0]
	this.Ad.Videoext.Video_unmute_urls = this.Ad.Videoext.Video_unmute_urls[:0]
	this.Ad.Videoext.Video_skip_urls = this.Ad.Videoext.Video_skip_urls[:0]
	this.Ad.Videoext.Video_close_urls = this.Ad.Videoext.Video_close_urls[:0]
	this.Ad.Videoext.Video_pause_urls = this.Ad.Videoext.Video_pause_urls[:0]
	this.Ad.Videoext.Video_resume_urls = this.Ad.Videoext.Video_resume_urls[:0]
	this.Ad.Videoext.Video_replay_urls = this.Ad.Videoext.Video_replay_urls[:0]
	this.Ad.Videoext.Video_fullscreen_urls = this.Ad.Videoext.Video_fullscreen_urls[:0]
	this.Ad.Videoext.Video_exit_fullscreen_urls = this.Ad.Videoext.Video_exit_fullscreen_urls[:0]

}

type BidResponse struct {
	Res     int    `json:"res,omitempty"`
	ResV    string `json:"res_v,omitempty"`
	Message string `json:"message,omitempty"`
	Ad      AdResp `json:"ad,omitempty"`
}

type AdResp struct {
	Adw      int            `json:"adw,omitempty"`
	Adh      int            `json:"adh,omitempty"`
	Img      string         `json:"img,omitempty"`
	Img2     string         `json:"img2,omitempty"`     //广告图副图
	Img3     string         `json:"img3,omitempty"`     //广告图副图
	Guidepic string         `json:"guidepic,omitempty"` //引导图
	Title    string         `json:"title,omitempty"`    //广告标题
	Text     string         `json:"text,omitempty"`     //广告文案
	Click    []string       `json:"click,omitempty"`    //点击上报地址
	Imp      []string       `json:"imp,omitempty"`      //展示上报地址
	Act      int            `json:"act,omitempty"`      //交互方式，1=跳转落地页，2=下载，3=跳转落地页下载（比如广点通），4=DeepLink
	Lpg      string         `json:"lpg,omitempty"`      //广告跳转落地页链接，支持重定向
	Dplink   string         `json:"dplink,omitempty"`   //deeplink链接，如果不为空，则点击的时候要优先处理，唤醒失败再调用lpg落地页链接
	Ulk      string         `json:"ulk,omitempty"`      //iOSuniversallink通用链接
	Qak      string         `json:"qak,omitempty"`      // 安卓快应用链接
	Logo     string         `json:"logo,omitempty"`     //logo图
	Icon     string         `json:"icon,omitempty"`     //icon图
	Html     string         `json:"html,omitempty"`     //动态广告代码
	Adext    ObjectAdExt    `json:"adext,omitempty"`    //广告扩展数据对象
	Videoext ObjectVideoExt `json:"videoext,omitempty"` //视频扩展数据
}

type ObjectAdExt struct {
	Pkg          string   `json:"pkg,omitempty"`          //包名
	Nurl         []string `json:"nurl,omitempty"`         //竞价成功上报地址
	Lurl         []string `json:"lurl,omitempty"`         //竞价失败上报地址
	Carurl       []string `json:"carurl,omitempty"`       //点击上报，POST，请参考【附录一】
	Downbegin    []string `json:"downbegin,omitempty"`    //开始下载上报地址
	Downsucc     []string `json:"downsucc,omitempty"`     //下载成功上报地址
	Installbegin []string `json:"installbegin,omitempty"` //开始安装上报地址
	Installsucc  []string `json:"installsucc,omitempty"`  //安装成功上报地址
	Appactive    []string `json:"appactive,omitempty"`    //下载类广告，下载且安装完成之后,打开应用时上报
	Ktbegin      []string `json:"ktbegin,omitempty"`      //deeplink开始唤醒上报地址
	Kt           []string `json:"kt,omitempty"`           //deeplink唤醒成功上报地址
	Ktfail       []string `json:"ktfail,omitempty"`       //deeplink唤醒失败上报地址
	Price        float64  `json:"price,omitempty"`        //double出价，单位：分/CPM，RTB结算方式时有值
	Currency     string   `json:"currency,omitempty"`     //币种，人民币=CNY，美元=USD，默认CNY
}

type ObjectVideoExt struct {
	Vurl                       string   `json:"vurl,omitempty"`                       //视频地址
	Duration                   int      `json:"duration,omitempty"`                   //播放时长(单位秒)
	Keep                       int      `json:"keep,omitempty"`                       //持续播放时长后可跳过(单位秒)
	Vhtml                      string   `json:"vhtml,omitempty"`                      //视频封面HTML代码
	Lpic                       string   `json:"lpic,omitempty"`                       //视频封面图
	Endcard_img                string   `json:"endcard_img,omitempty"`                //视频播放完后需要展示的图片地址
	Endcard_html               string   `json:"endcard_html,omitempty"`               // 视频播放完后需要展示的HTML代码
	Video_first_urls           []string `json:"video_first_urls,omitempty"`           // 视频播放25%时上报地址
	Video_mid_urls             []string `json:"video_mid_urls,omitempty"`             // 视频播放50%时上报地址
	Video_third_urls           []string `json:"video_third_urls,omitempty"`           // 视频播放75%时上报地址
	Video_begin_urls           []string `json:"video_begin_urls,omitempty"`           // 视频开始播放上报地址
	Video_end_urls             []string `json:"video_end_urls,omitempty"`             // 视频播放完成上报地址
	Video_mute_urls            []string `json:"video_mute_urls,omitempty"`            //  视频开启静音上报地址
	Video_unmute_urls          []string `json:"video_unmute_urls,omitempty"`          //视频关闭静音上报地址
	Video_skip_urls            []string `json:"video_skip_urls,omitempty"`            //跳过视频上报地址
	Video_close_urls           []string `json:"video_close_urls,omitempty"`           // 关闭视频上报地址
	Video_pause_urls           []string `json:"video_pause_urls,omitempty"`           // 视频暂停时上报地址
	Video_resume_urls          []string `json:"video_resume_urls,omitempty"`          // 视频暂停再播放时上报地址
	Video_replay_urls          []string `json:"video_replay_urls,omitempty"`          // 视频重播时上报地址
	Video_fullscreen_urls      []string `json:"video_fullscreen_urls,omitempty"`      // 视频全屏时上报地址
	Video_exit_fullscreen_urls []string `json:"video_exit_fullscreen_urls,omitempty"` //视频退出全屏时上报地址
}
