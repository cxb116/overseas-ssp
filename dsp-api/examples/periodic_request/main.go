package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// BidRequest 广告请求结构
type BidRequest struct {
	ID     string `json:"id"`
	Imp    Imp    `json:"imp"`
	App    App    `json:"app"`
	Device Device `json:"device"`
	User   User   `json:"user"`
	Ver    string `json:"ver"`
	HTTPS  int    `json:"https"`
}

type Imp struct {
	ID       int `json:"id"`
	W        int `json:"w"`
	H        int `json:"h"`
	AdType   int `json:"ad_type"`
	Position int `json:"position"`
	BidFloor int `json:"bid_floor"`
}

type App struct {
	Name            string `json:"name"`
	Pkg             string `json:"pkg"`
	Ver             string `json:"ver"`
	VerCode         int    `json:"ver_code"`
	URL             string `json:"url"`
	Appstorever     string `json:"appstorever"`
	AppstoreverName string `json:"appstorever_name"`
}

type Device struct {
	IP              string  `json:"ip"`
	UA              string  `json:"ua"`
	OS              int     `json:"os"`
	OSV             string  `json:"osv"`
	DeviceType      int     `json:"device_type"`
	Geo             Geo     `json:"geo"`
	Network         Network `json:"network"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	ModelCode       string  `json:"model_code"`
	Orientation     int     `json:"orientation"`
	W               int     `json:"w"`
	H               int     `json:"h"`
	DPI             int     `json:"dpi"`
	PPI             int     `json:"ppi"`
	ScreenSize      float64 `json:"screen_size"`
	OAID            string  `json:"oaid"`
	DeviceName      string  `json:"device_name"`
	Language        string  `json:"language"`
	Country         string  `json:"country"`
	BatteryStatus   int     `json:"battery_status"`
	BatteryPower    int     `json:"battery_power"`
	CpuNum          int     `json:"cpu_num"`
	TimeZone        string  `json:"time_zone"`
	Lmt             int     `json:"lmt"`
	AppStoreVersion string  `json:"app_store_version"`
}

type Geo struct {
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
	CType int     `json:"ctype"`
}

type Network struct {
	NetType int    `json:"net_type"`
	Carrier int    `json:"carrier"`
	Imsi    string `json:"imsi"`
	Mac     string `json:"mac"`
	MacMD5  string `json:"mac_md5"`
	Ssid    string `json:"ssid"`
	WifiMac string `json:"wifi_mac"`
}

type User struct {
	UID    string `json:"uid"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

// 创建示例请求 模拟下游
func createSampleRequest() *BidRequest {
	return &BidRequest{
		ID: "req-20260305-0001",
		Imp: Imp{
			ID:       25,
			W:        375,
			H:        300,
			AdType:   1,
			Position: 4,
			BidFloor: 1,
		},
		App: App{
			Name:            "NewsPro",
			Pkg:             "com.newspro.app",
			Ver:             "3.2.1",
			VerCode:         321,
			URL:             "https://apps.apple.com/app/id1234567890",
			Appstorever:     "3.2.1",
			AppstoreverName: "3.2.1",
		},
		Device: Device{
			IP:         "127.0.0.1",
			UA:         "Mozilla/5.0 (Linux; Android 10; PBCM10 Build/QKQ1.191224.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.92 Mobile Safari/537.36",
			OS:         1,
			OSV:        "17.3",
			DeviceType: 0,
			Geo: Geo{
				Lat:   16.2,
				Lon:   106.816666,
				CType: 1,
			},
			Network: Network{
				NetType: 1,
				Carrier: 0,
				Imsi:    "",
				Mac:     "",
				MacMD5:  "",
				Ssid:    "MyHomeWifi",
				WifiMac: "",
			},
			Brand:           "vivo",
			Model:           "vivo",
			ModelCode:       "A2890",
			Orientation:     1,
			W:               1179,
			H:               2556,
			DPI:             3,
			PPI:             460,
			ScreenSize:      6.1,
			OAID:            "FF5689EFE4BE4030B311F2B6AF17B2BCc004228554d3a5f134f8aa3e9842392d",
			DeviceName:      "",
			Language:        "zh",
			Country:         "CN",
			BatteryStatus:   1,
			BatteryPower:    80,
			CpuNum:          6,
			TimeZone:        "Asia/Jakarta",
			Lmt:             0,
			AppStoreVersion: "17.3",
		},
		User: User{
			UID:    "user_100001",
			Age:    28,
			Gender: 1,
		},
		Ver:   "1.0",
		HTTPS: 1,
	}
}

// 发送请求
func sendRequest(url string, req *BidRequest) error {
	// 序列化为 JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %v", err)
	}

	// 打印请求内容
	fmt.Printf("\n========== 发送请求 ==========\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Request Body:\n%s\n", string(jsonData))

	// 创建 HTTP 请求
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %v", err)
	}

	// 打印响应
	fmt.Printf("\n========== 响应结果 ==========\n")
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body:\n%s\n", string(body))
	fmt.Printf("==============================\n")

	return nil
}

func main() {
	// 配置
	serverURL := "http://127.0.0.1:9990/dsp" // 修改为你的实际服务地址
	interval := 5 * time.Second              // 发送间隔：5秒
	requestCount := 0                        // 请求计数器

	fmt.Printf("========================================\n")
	fmt.Printf("     定时请求发送工具\n")
	fmt.Printf("========================================\n")
	fmt.Printf("目标地址: %s\n", serverURL)
	fmt.Printf("发送间隔: %v\n", interval)
	fmt.Printf("按 Ctrl+C 停止程序\n")
	fmt.Printf("========================================\n\n")

	// 创建定时器
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// 立即发送第一次请求
	req := createSampleRequest()
	if err := sendRequest(serverURL, req); err != nil {
		log.Printf("❌ 请求失败: %v\n", err)
	} else {
		requestCount++
		log.Printf("✅ 第 %d 次请求发送成功\n", requestCount)
	}

	// 定时发送请求
	for {
		select {
		case <-ticker.C:
			// 每次可以修改请求ID，使其唯一
			req.ID = fmt.Sprintf("req-%d-%04d", time.Now().Unix(), requestCount%10000)

			// 发送请求
			if err := sendRequest(serverURL, req); err != nil {
				log.Printf("❌ 请求失败: %v\n", err)
			} else {
				requestCount++
				log.Printf("✅ 第 %d 次请求发送成功\n", requestCount)
			}
		}
	}
}
