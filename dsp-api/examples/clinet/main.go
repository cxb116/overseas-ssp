package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	// 注意：你的引擎监听端口以实际启动为准（常见是 8081；如果你改成 9999 也可以）
	url      = "http://127.0.0.1:9999/dsp"
	jsonData = `{
	"id": "req-20260107-0001",
	"imp": {"id":1,"w":1080,"h":1920,"ad_type":2,"position":1,"bid_floor":500},
	"app": {"app_id":"app_987654321","name":"示例资讯App","bundle":"com.example.news","ver":"5.3.1"},
	"device": {"ip":"123.125.68.45","ua":"Mozilla/5.0","os":1,"osv":"12","model":"Mi 10","geo":{"lat":39.9042,"lon":116.4074,"ctype":1},"device_type":0,"network":{"net_type":5,"carrier":1,"imsi":"460001","mac":"02:00:00:00:00:00"},"androidId":"9774d56d682e549c","oaid":"7f3b2c1d-1234-5678-9abc-abcdef123456"},
	"ver": "1", "support_dp": 1, "support_wx": 1,
	"user": {"uid":"user_123456789","age":28,"gender":1},
	"https": true
	}`

	count         int64
	fail          int64
	baiduCount    int64 // 百度请求计数
	jingdongCount int64 // 京东请求计数
	wg            sync.WaitGroup
)

const (
	Concurrency     = 50                      // 并发线程数（可以修改为更大的值模拟并发）
	TotalRequests   = 700                     // 总请求次数
	RequestInterval = 50 * time.Millisecond   // 每次请求间隔
	RequestTimeout  = 1300 * time.Millisecond // 请求超时
)

func main() {
	startTime := time.Now()
	fmt.Printf("开始压测：并发数=%d, 总请求数=%d\n", Concurrency, TotalRequests)
	fmt.Println("========================================")

	// 启动并发协程
	for i := 0; i < Concurrency; i++ {
		wg.Add(1)
		go DoRequest(i)
	}

	// 等待所有请求完成
	wg.Wait()

	// 统计结果
	duration := time.Since(startTime)
	total := atomic.LoadInt64(&count) + atomic.LoadInt64(&fail)
	success := atomic.LoadInt64(&count)
	failed := atomic.LoadInt64(&fail)
	baidu := atomic.LoadInt64(&baiduCount)
	jingdong := atomic.LoadInt64(&jingdongCount)

	fmt.Println("========================================")
	fmt.Printf("压测完成！\n")
	fmt.Printf("总耗时: %v\n", duration)
	fmt.Printf("总请求数: %d\n", total)
	fmt.Printf("成功: %d (%.2f%%)\n", success, float64(success)/float64(total)*100)
	fmt.Printf("失败: %d (%.2f%%)\n", failed, float64(failed)/float64(total)*100)
	fmt.Printf("QPS: %.2f\n", float64(total)/duration.Seconds())
	fmt.Println("----------------------------------------")
	fmt.Printf("流量分配统计：\n")
	fmt.Printf("  百度 (9009): %d 次 (%.2f%%)\n", baidu, float64(baidu)/float64(success)*100)
	fmt.Printf("  京东 (9010): %d 次 (%.2f%%)\n", jingdong, float64(jingdong)/float64(success)*100)
	fmt.Println("========================================")
}

// DoRequest 循环发送请求
func DoRequest(id int) {
	defer wg.Done()

	client := &fasthttp.Client{
		MaxConnsPerHost: 1000,
		ReadTimeout:     RequestTimeout,
		WriteTimeout:    RequestTimeout,
	}

	// 计算每个线程需要发送的请求数
	requestsPerThread := TotalRequests / Concurrency
	if id < TotalRequests%Concurrency {
		requestsPerThread++ // 多余的请求分配给前面的线程
	}

	for i := 0; i < requestsPerThread; i++ {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()

		req.SetRequestURI(url)
		req.Header.SetMethod(fasthttp.MethodPost)
		req.Header.SetContentType("application/json")
		req.SetBody([]byte(jsonData))

		err := client.Do(req, resp)
		if err != nil {
			atomic.AddInt64(&fail, 1)
			fmt.Printf("[线程 %d] 请求 %d/%d 失败: %v\n", id, i+1, requestsPerThread, err)
		} else {
			atomic.AddInt64(&count, 1)
			total := atomic.LoadInt64(&count)

			body := resp.Body()
			statusCode := resp.StatusCode()

			// 尝试解析 JSON 判断是百度还是京东
			var result map[string]interface{}
			if err := json.Unmarshal(body, &result); err == nil {
				if statusCode == 200 {
					if seat, ok := result["seat"].(map[string]interface{}); ok {
						// 用 slotId 统计（避免 title + slotId 双计数）
						if slotId, ok := seat["slotId"].(string); ok {
							if slotId == "slot_10001" {
								atomic.AddInt64(&baiduCount, 1)
							} else if slotId == "slot_10002" {
								atomic.AddInt64(&jingdongCount, 1)
							}
						}
					}
				}
			}

			// 每10次请求打印一次进度
			if total%10 == 0 {
				fmt.Printf("[线程 %d] 进度: %d/%d (总计: %d, 失败: %d)\n",
					id, i+1, requestsPerThread, total, atomic.LoadInt64(&fail))
			}
		}

		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)

		// 控制请求速率
		if i < requestsPerThread-1 { // 最后一次请求不需要 sleep
			time.Sleep(RequestInterval)
		}
	}

	fmt.Printf("[线程 %d] 完成，共发送 %d 个请求\n", id, requestsPerThread)
}
