package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 默认 JSON payload
var payload = map[string]interface{}{
	"adh": 1080,
	"adw": 1920,
	"app": map[string]string{
		"an":  "次元滤镜",
		"pkg": "com.eryou.ciyuanlj",
		"vc":  "1.2.3",
		"ver": "1.25",
	},
	"appid": 2002271,
	"device": map[string]interface{}{
		"memory":           10,
		"hard_disk":        111,
		"device_name":      "Android",
		"timezone":         "Asia/Shanghai",
		"app_store_vc":     "",
		"boot_mark":        "fe675732-5a63-4a78-9ac0-17df4c2e0d09",
		"brd":              "OPPO",
		"density":          3,
		"dpi":              480,
		"dso":              2,
		"dvt":              1,
		"hardware_model":   "PBCM10",
		"imei":             "869468020290123",
		"lg":               "zh-CN",
		"mac":              "58:C6:F0:BE:37:0B",
		"mdl":              "PBCM10",
		"net":              1,
		"aid":              "d6b4b18a1ea597a0",
		"oaid":             "FF5689EFE4BE4030B311F2B6AF17B2BCc004228554d3a5f134f8aa3e9842392d",
		"opt":              46000,
		"ov":               "10",
		"plt":              1,
		"ppi":              480,
		"sheight":          2132,
		"ssid":             "10465",
		"swidth":           1080,
		"ua":               "Mozilla/5.0 (Linux; Android 10; PBCM10 Build/QKQ1.191224.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.92 Mobile Safari/537.36",
		"update_mark":      "2490696.709999980",
		"vendor":           "OPPO",
		"startup_ts":       "1705351197",
		"init_ts":          "1705351197",
		"upgrade_ts":       "1705351197",
		"sys_compiling_ts": "1705351197",
		"country":          "CN",
	},
	"geo": map[string]float64{
		"lat": 30.763441,
		"lon": 120.29801,
	},
	"ip":       "39.172.26.8",
	"ishttps":  1,
	"slotid":   182,
	"bidfloor": 1,
}

func sendRequest(client *http.Client, url string, payload map[string]interface{}) (bool, float64) {
	start := time.Now()
	data, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return false, 0
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		return false, elapsed
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, elapsed
	}
	return false, elapsed
}

func main() {
	// 命令行参数
	url := flag.String("url", "http://115.190.55.62:8080/api/adx", "测试地址")
	concurrency := flag.Int("concurrency", 800, "并发数")
	totalRequests := flag.Int("requests", 10000, "总请求数")
	timeout := flag.Int("timeout", 2, "请求超时时间（秒）")
	flag.Parse()

	if *url == "" {
		fmt.Println("请提供目标接口 URL: -url https://example.com/bidrequest")
		return
	}

	var wg sync.WaitGroup
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	successCount := 0
	failCount := 0
	var totalTime float64
	var mu sync.Mutex

	fmt.Printf("开始压力测试: URL=%s, 并发=%d, 总请求=%d\n", *url, *concurrency, *totalRequests)

	sem := make(chan struct{}, *concurrency) // 控制并发

	for i := 0; i < *totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			success, elapsed := sendRequest(client, *url, payload)
			mu.Lock()
			totalTime += elapsed
			if success {
				successCount++
			} else {
				failCount++
			}
			mu.Unlock()
			<-sem
		}()
	}

	wg.Wait()

	fmt.Println("\n=== 测试结果 ===")
	fmt.Printf("总请求数: %d\n", *totalRequests)
	fmt.Printf("成功请求: %d\n", successCount)
	fmt.Printf("失败请求: %d\n", failCount)
	fmt.Printf("平均响应时间: %.3f 秒\n", totalTime/float64(*totalRequests))
}
