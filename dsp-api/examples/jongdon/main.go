package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type BidResponse struct {
	Id     string `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Seat   Seat   `json:"seat,omitempty"`
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
	ImgUrls         string   `json:"img_urls,omitempty"`
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
		Id:     "req_20260109_0001",
		Status: 0,
		Seat: Seat{
			SlotId:       "slot_10002",
			Price:        120,
			CreativeType: 2,
			Title:        "京东买卖",
			Desc:         "-叮当猫",
			IconUrl:      "https://cdn.example.com/icon.png",
			ImgUrls:      "https://cdn.example.com/img1.png,https://cdn.example.com/img2.png",
			W:            720,
			H:            1280,
			TargetUrl:    "https://www.example.com/landing",
			Deeplink:     "yymobile://Entrance/Redirect?bizType=12\u0026from_channel=cg_gp_uc_8521\u0026jump_type=live\u0026creativeid=416542570\u0026adid=1521123325\u0026projectid=124625436\u0026accountid=211546759",
			Source:       "dsp_mock",
			Nurls: []string{
				"https://track.example.com/win?price=${AUCTION_PRICE}",
				"https://wiretap.adxtop.cn/scene/win?id=1458708887246393398\u0026rid=b7a6a95a-f0be-443e-a18f-d6c7f851c558\u0026date=20260108\u0026dateh=tES2bH3GX4Pk5pVCCRrr9A%3D%3D\u0026planId=61823\u0026sspId=200077\u0026dspId=500071\u0026appId=10603\u0026slotId=31241\u0026stubId=70587\u0026px=MLIo2wCQ8IKin1eMCPhNGQ%3D%3D\u0026ts=1767853700\u0026sn=jzss\u0026dn=huichuan\u0026pkg=com.fy.bbqx.huawei\u0026data=bU6%2BKc9jfJIXxKyN3O%2FDIDbAEFuhVLZRDLUPO9amIdWOYXt3lZb%2FCgjSzxq6oopQqY7yntJQK7vLzupyX5HLSksT23407KKOhJCgPIGCN%2FXMLfYodwZDppf3xy3qyEgm1LssDSZPCD2QJnNH2SFGrfcxCkFVjMu1wCyKFTwmqbBhmzrZ0IbTaxHIKf6cy5Yv\u0026price=__PRICE__",
			},
			Ims: []string{
				"https://track.example.com/impression",
				"https://wiretap.adxtop.cn/scene2/show?id=1458708887246393398\u0026rid=b7a6a95a-f0be-443e-a18f-d6c7f851c558\u0026date=20260108\u0026dateh=tES2bH3GX4Pk5pVCCRrr9A%3D%3D\u0026planId=61823\u0026sspId=200077\u0026dspId=500071\u0026appId=10603\u0026slotId=31241\u0026stubId=70587\u0026px=MLIo2wCQ8IKin1eMCPhNGQ%3D%3D\u0026ts=1767853700\u0026sn=jzss\u0026dn=huichuan\u0026pkg=com.fy.bbqx.huawei\u0026type=__TYPE__\u0026behavior=__BEHAVIOR__\u0026playedDuration=__P_DURATION__\u0026playedRate=__P_RATE__\u0026price=__PRICE__",
				"https://huichuan.sm.cn/v?btinfo=CPTH2i8:\u0026data=ADgAAD8GAAAjVwoN7sM5lHiKQseu4qQ5h3not1Bimu-MxlQH6uhD0Nq9Jy6E9cx9sB2LlkCCbur6QAvUcRdzn17TY15t9H6h7Oemy_TaEXjxciE6Pt_OTtjjYceibCTT-51guHF79aw_56dziCdhPIcHfVITNUoJC6S6ncVAdaZXJYh597PvrQhskb_DYiu2SKSZcnYGQI2Qq3iFEAb5NynKSjaJV-h6yR0db6YX-mPoZoXy-qZh5jnBAze27qRLwKQVTji0kShUbxJFBeOaAMMsCzsKmCKAVoDC4q_57nVxNdFPKKoehviZmri1sT9yV_HsN91PSyfjjkzOa6wD1vsAjLdoxUZA7Kcu862PLzFXFo6z-8cHza1_T-yXcpn-4IIaANd4-2trFNzungL5RTbGP4xUJ3SLohRp0r5prCn8TRCq5uviQVAE712JdNJZstgFKv0yRBUo1YzRbEt5rs-Ev1GLn5ZfeQ-hywGz3BPW4HcqfVFY9tt5FD57mQ9ZRZg2FIVlxuOpTs4wgLu2a_q_EsqQpc16683jJ5vHuwWYeA7MtqKnpGg2ZaYrp3n-0vfm0yMicYny2p-lgMuYHLxEjt3pOLTkAgrOWvk4FiBKqz61F4JWnSBLRj6m1mq0PxOI7ce8-sv14Dv3pfhn6rXkxQjSQzptdvQGEfj-6zjDHbuv41TN_4EgOnk3IsAXuJMkxWBg6uVeGGUE6ZYpuFNtklc_XxSyRfURGhpDqnWKwsvXB3NFy2jyYmmCwkGHPPv9wguxii2K2lnQBKWDZwYBgcH958B7o3KrmlFh89AvJvauzQ-dccfYsSYeremX0HZzJbqRtNSM320BsfTExyI7R1Bsnl9s8z-ADqnMMLY_SdrjEdwLKJw6ldAmjFu2Ios9gYC5kFni-TZjIXXl1iXgkrMKw-pJo8Ej350LIxeAYTD28pO9N3KUTd6bCRtPr14gefre9wa-yoKaEu-pGHbp8X3V57TXbec8cX5cevWDExMFypfKNuSRjFIZnxB6rqmrLOLYE-VPTNuSm2TgimeShPPj15ZqpgtqnircuUy-u7fKcNgIJmyQtIg7XRcUGBvEhViC9IV_kc8V5AtR-JJlxrX2MAiP2JujC5yBCj5EecRW9Tjp21jhI7FCywlbs8-vAXlumrOv6C6eQFQJ6hv32RdLxUxDlvnYCsEBj--NKgWIVZ9CmM_XziW0H9XTdUY1CKSWNs6K9Otn3hGYFPl5WLVxwcNwNhFWMpyBjC2linJMdzpaPyqamQEUgq60YDmGl8NE-51AEsinddXCf-Jfd3fBPiudbVJnv5LK4y29Fi3kJ-1dahVlGwHkh_bVkCBEHVJRH8cNswxhNVVpVf97EU_jR9a_8Fwod2Z_5-B-2iZ8pCCFKRDJqcMvZYQ21t79wwUBIWigzDJi0x2TRU9VbWEON6PVBdAPcwkyS1Uyhr3647VyxBkXQ-8ElK0SeBUkGj3HmpfGfGMGOpQ-uZTYK1MUBySsL1KR4FqeNzKCV-K3uWvqQimXQVzCNmUoruX3e02eWbALIQ6VcecI2jlL3vAqByOwu3Y4MQCWokTt_ekfQ7focu0l7SVZcttbTBF0UZcOCJMY4PLaAaRfy8Fts4s0BaYYFakReR9_e_VZgWj1jT6NVP_3gZvdV7gFp4MCBjsPmflbwZJByRZn-kZhPuYqjCsdzpOOkO1NHE0_ibFShQvD57-_QBHQk0n5O6FNJ3A-swx3HONs4sveTwzRMYIR8az9Kr_pm9IfbsuBGQgSdT4gcgRZylgh-Dm0S-ADH59FjUJuqEG01tXQJxrBq9dp_xLjXK3FiW1dCMiN5-wxAd7yOILwbmUb4rc3q4IaxV8o19EU_VTaFQBkF4QhzZH-oxnM-r9dkyThf9rkL_RtQnDKlo0DySqwn9RQDQfZvLocVbFKGTLTdPy1Hqi4mtnGlsrZy3QnlEEmz6cfykCwHq3eV-RUv5eHe1ifv6dDOj_O13-8erDwiLgkOOV6nT2O68cVDYTQrC7PQh303AnvN3qA4YdPWxXTgGEk7Vqfi8wPVAdgjUF14TBD5PgXOrtZlBzRwEwaFLS8NXPC9ZixnGQ9euqnvocmJMH0YFemS_aO1WuvZGRPIowykkThbZOI3nI:\u0026stm=__TS_S__",
			},
			Clks: []string{
				"https://track.example.com/click",
				"https://c.sm.cn/c?btinfo=CPTH2i8:\u0026data=AJgAAD8GAAAjVwoN7sM5lHiKQseu4qQ5h3not1Bimu-MxlQH6uhD0Nq9Jy6E9cx9sB2LlkCCbur6QAvUcRdzn17TY15t9H6h7Oemy_TaEXjxciE6Pt_OTtjjYceibCTT-51guHF79aw_56dziCdhPIcHfVITNUoJC6S6ncVAdaZXJYh597PvrQhskb_DYiu2SKSZcnYGQI2Qq3iFEAb5NynKSjaJV-h6yR0db6YX-mPoZoXy-qZh5jnBAze27qRLwKQVTji0kShUbxJFBeOaAMMsCzsKmCKAVoDC4q_57nVxNdFPKKoehviZmri1sT9yV_HsN91PSyfjjkzOa6wD1vsAjLdoxUZA7Kcu862PLzFXFo6z-8cHza1_T-yXcpn-4IIaANd4-2trFNzungL5RTbGP4xUJ3SLohRp0r5prCn8TRCq5uviQVAE712JdNJZstgFKv0yRBUo1YzRbEt5rs-Ev1GLn5ZfeQ-hywGz3BPW4HcqfVFY9tt5FD57mQ9ZRZg2FIVlxuOpTs4wgLu2a_q_EsqQpc16683jJ5vHuwWYeA7MtqKnpGg2ZaYrp3n-0vfm0yMicYny2p-lgMuYHLxEjt3pOLTkAgrOWvk4FiBKqz61F4JWnSBLRj6m1mq0PxOI7ce8-sv14Dv3pfhn6rXkxQjSQzptdvQGEfj-6zjDHbuv41TN_4EgOnk3IsAXuJMkxWBg6uVeGGUE6ZYpuFNtklc_XxSyRfURGhpDqnWKwsvXB3NFy2jyYmmCwkGHPPv9wguxii2K2lnQBKWDZwYBgcH958B7o3KrmlFh89AvJvauzQ-dccfYsSYeremX0HZzJbqRtNSM320BsfTExyI7R1Bsnl9s8z-ADqnMMLY_SdrjEdwLKJw6ldAmjFu2Ios9gYC5kFni-TZjIXXl1iXgkrMKw-pJo8Ej350LIxeAYTD28pO9N3KUTd6bCRtPr14gefre9wa-yoKaEu-pGHbp8X3V57TXbec8cX5cevWDExMFypfKNuSRjFIZnxB6rqmrLOLYE-VPTNuSm2TgimeShPPj15ZqpgtqnircuUy-u7fKcNgIJmyQtIg7XRcUGBvEhViC9IV_kc8V5AtR-JJlxrX2MAiP2JujC5yBCj5EecRW9Tjp21jhI7FCywlbs8-vAXlumrOv6C6eQFQJ6hv32RdLxUxDlvnYCsEBj--NKgWIVZ9C0zGAqXjgb3ucl1V6Xp-r4z4cn0AOWCif7PjgPEL22gwr7bIHAeb910bj9KaUlxw957fzDLhIum4bW-OUgsMASnrc5T2B56HHRxmsZi1--zpk3RIuqD-uCd3wISoBtOybQKCtg905XoOSVsQD4ZHmMQG4nNov5Ulk0M9BO4pUfUnAGVwAmxZjX-hWYJFhHLXWkKC7mQa4X-J9lRxBECWwdtHpiXlrAczWpAwRJpSFbtm5Zz0IcvNr1eXj73lMCTkiIJcNR8VUWJIxvprRlmBQaNWfycRvEiHXX_qdXRT6GWkQEaz-sF9dMpZs3LUokbYlhJPSkQ1GSAxs7GfzW0XlQ_iJTXno_4p0UZFbevB4QujaJshQ9CLXxQhVA-mbtvSumxNDDE3GfbD-BIY1Tvj1pIEHtYF6tRtcdMI7kEaB-1RrTGJrnTM4UBmt1Nm6-NQA4RfaZvKgm4ew1CGekmPEsW8rsIArUNmlGP9QRDPePuxQZjCP0U-f8uhrd2HthhcL4xMt4ZnnqIbcD0Y2wDzXtAuW7-nm3bLqAzrjcNTJ9uATH9ZN28H296t_0nnvtDYlyjzg2NrhzuXA3fL2ykeJa9ale0Erh2lr_vsmLzO3fDgC1ZiUBnOu0IeIdzbhDv1zzlb75uq-Dr8aRoPwwLGSw5HIp35tmQ7PcLUllzKpeqk58ThzEXfKa_YwUXnc0LcriGgm6L2w8e2MV66qQpGU7Sp_Bcik-XYyJDvKEsGMd_eWgwQyNH10ysRv5neheCjCXe98DwqcGBFxbWL-__avTNqUrG52Mclzf1Z4lq_CZJPk4wSvsyjr3-gPbdwPgSCJ_g4XVYcjf_0cgH6VkF1G16ICd9qwfxOzdC0WjoPpspyh4jOFKsTBZAs8EQOR_f2aL-wfY_px1pCzDAR3sjXVS7jc-eYec3ojcGKNqTW28Zo:\u0026stm=__TS_S__",
				"https://wiretap.adxtop.cn/scene/click?id=1458708887246393398\u0026rid=b7a6a95a-f0be-443e-a18f-d6c7f851c558\u0026date=20260108\u0026dateh=tES2bH3GX4Pk5pVCCRrr9A%3D%3D\u0026planId=61823\u0026sspId=200077\u0026dspId=500071\u0026appId=10603\u0026slotId=31241\u0026stubId=70587\u0026ts=1767853700\u0026sn=jzss\u0026dn=huichuan\u0026pkg=com.fy.bbqx.huawei\u0026width=__WIDTH__\u0026height=__HEIGHT__\u0026down_x=__DOWN_X__\u0026down_y=__DOWN_Y__\u0026up_x=__UP_X__\u0026up_y=__UP_Y__\u0026click_element=__CLICK_ELEMENT__\u0026sid=9781589960222840312\u0026creative_id=416542570",
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//time.Sleep(650 * time.Millisecond)

	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/bid", bidHandler)

	addr := ":9010"
	log.Println("mock dsp server start at", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
