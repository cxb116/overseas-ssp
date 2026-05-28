package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type BidResponse struct {
	Id   string `json:"id,omitempty"`
	Code int    `json:"code,omitempty"`
	Seat []Seat `json:"seat,omitempty"`
}

type Seat struct {
	SlotId          string   `json:"slotId,omitempty"`
	Price           int64    `json:"price,omitempty"`
	CreativeType    int      `json:"creative_type,omitempty"`
	Html            string   `json:"html,omitempty"`
	InteractionType int      `json:"interaction_type,omitempty"`
	Title           string   `json:"title,omitempty"`
	Desc            string   `json:"desc,omitempty"`
	IconUrl         string   `json:"icon_url,omitempty"`
	ImgUrls         []string `json:"img_urls,omitempty"`
	W               int32    `json:"w,omitempty"`
	H               int32    `json:"h,omitempty"`
	QuickAppUrl     string   `json:"quick_app_url,omitempty"`
	TargetUrl       string   `json:"target_url,omitempty"`
	Deeplink        string   `json:"deeplink,omitempty"`
	MarketUrl       string   `json:"market_url,omitempty"`
	UniverseUrl     string   `json:"universe_url,omitempty"`
	Source          string   `json:"source,omitempty"`
	Nurls           []string `json:"nurls,omitempty"`
	Lurls           []string `json:"lurls,omitempty"`
	Ims             []string `json:"ims,omitempty"`
	Clks            []string `json:"clks,omitempty"`
	DpTryUrls       []string `json:"dp_try_urls,omitempty"`
	DpSucUrls       []string `json:"dp_suc_urls,omitempty"`
	DpFailUrls      []string `json:"dp_fail_urls,omitempty"`
	WxAppId         string   `json:"wx_app_id,omitempty"`
	WxId            string   `json:"wx_id,omitempty"`
	WxPath          string   `json:"wx_path,omitempty"`
	App             AppRes   `json:"app,omitempty"`
	Video           VideoRes `json:"video,omitempty"`
}

type AppRes struct {
	DownloadUrl         string   `json:"download_url,omitempty"`
	Name                string   `json:"name,omitempty"`
	Bundle              string   `json:"bundle,omitempty"`
	Version             string   `json:"version,omitempty"`
	Developer           string   `json:"developer,omitempty"`
	PrivacyUrl          string   `json:"privacy_url,omitempty"`
	PermissionUrl       string   `json:"permission_url,omitempty"`
	FunctionDevUrl      string   `json:"function_dev_url,omitempty"`
	StartDownloadUrls   []string `json:"start_download_urls,omitempty"`
	SuccessDownloadUrls []string `json:"success_download_urls,omitempty"`
	InstallStartUrls    []string `json:"install_start_urls,omitempty"`
	InstallSuccessUrls  []string `json:"install_success_urls,omitempty"`
	ActiveSuccessUrls   []string `json:"active_success_urls,omitempty"`
}

type VideoRes struct {
	Url            string   `json:"url,omitempty"`
	Duration       int32    `json:"duration,omitempty"`
	Size           int32    `json:"sizem,omitempty"`
	W              int32    `json:"w,omitempty"`
	H              int32    `json:"h,omitempty"`
	CoverUrl       string   `json:"cover_url,omitempty"`
	StartUrls      []string `json:"start_urls,omitempty"`
	FourthUrls     []string `json:"fourth_urls,omitempty"`
	HalfUrls       []string `json:"half_urls,omitempty"`
	ThdUrls        []string `json:"thd_urls,omitempty"`
	EndUrls        []string `json:"end_urls,omitempty"`
	SkipUrls       []string `json:"skip_urls,omitempty"`
	CompleteUrls   []string `json:"complete_urls,omitempty"`
	CloseUrls      []string `json:"close_urls,omitempty"`
	SuspendUrls    []string `json:"suspend_urls,omitempty"`
	MuteUrls       []string `json:"mute_urls,omitempty"`
	CancelMuteUrls []string `json:"cancel_mute_urls,omitempty"`
	FullUrls       []string `json:"full_urls,omitempty"`
	CancelFullUrls []string `json:"cancel_full_urls,omitempty"`
}

func bidHandler(w http.ResponseWriter, r *http.Request) {
	// 记录请求（可选）
	log.Printf("receive bid request: %s %s", r.Method, r.RemoteAddr)

	resp := BidResponse{
		Id:   "req_20260109_0001",
		Code: 10,
		Seat: []Seat{
			{
				SlotId:       "1",
				Price:        100,
				CreativeType: 2,
				Title:        "高效清理手机垃圾",
				Desc:         "一键清理，释放手机空间",
				IconUrl:      "https://cdn.example.com/icon.png",
				ImgUrls:      []string{"https://cdn.example.com/img1.png,https://cdn.example.com/img2.png"},
				W:            720,
				H:            1280,
				TargetUrl:    "https://www.example.com/landing",
				Deeplink:     "yymobile://Entrance/Redirect?bizType=12\u0026from_channel=cg_gp_uc_8521\u0026jump_type=live\u0026creativeid=416542570\u0026adid=1521123325\u0026projectid=124625436\u0026accountid=211546759",
				Source:       "dsp_mock",
				Nurls:        []string{},
				Ims: []string{
					"http://adx-bid.gladdigit.com/ad-core-master/api/feedback/pv\\u0026winprice=200&click_time=__CLICK_TIME_END_S__&click_time_start=__CLICK_TIME_START__",
				},
				Clks: []string{
					"http://adx-bid.gladdigit.com/ad-core-master/api/feedback/pv\\u0026winprice=200&click_time=__CLICK_TIME_END_S__&click_time_start=__CLICK_TIME_START__",
				},
				App: AppRes{
					DownloadUrl: "https://download.example.com/app.apk",
					Name:        "极速清理大师",
					Bundle:      "com.example.cleaner",
					Version:     "3.2.1",
					Developer:   "Example Tech",
				},
				Video: VideoRes{
					Url:      "https://cdn.example.com/video.mp4",
					Duration: 30,
					W:        720,
					H:        1280,
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/bid", bidHandler)

	addr := ":9009"
	log.Println("mock dsp server start at", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
