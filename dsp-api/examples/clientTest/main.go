package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	url         = "http://127.0.0.1:9999/dsp"
	concurrency = 50   // 并发数
	timeoutMs   = 2000 // 请求超时(ms)
)

var jsonData = []byte(`{
	"id": "req-20260107-0001",
	"imp": {"id":1,"w":1080,"h":1920,"ad_type":2,"position":1,"bid_floor":500},
	"app": {"app_id":"app_987654321","name":"示例资讯App","bundle":"com.example.news","ver":"5.3.1"},
	"device": {
		"ip":"123.125.68.45",
		"ua":"Mozilla/5.0",
		"os":1,
		"osv":"12",
		"model":"Mi 10",
		"geo":{"lat":39.9042,"lon":116.4074,"ctype":1},
		"device_type":0,
		"network":{"net_type":5,"carrier":1,"imsi":"460001","mac":"02:00:00:00:00:00"},
		"androidId":"9774d56d682e549c",
		"oaid":"7f3b2c1d-1234-5678-9abc-abcdef123456"
	},
	"ver": "1",
	"support_dp": 1,
	"support_wx": 1,
	"user": {"uid":"user_123456789","age":28,"gender":1},
	"https": true
}`)

func main() {
	client := &fasthttp.Client{
		MaxConnsPerHost: concurrency * 2,
		ReadTimeout:     time.Millisecond * timeoutMs,
		WriteTimeout:    time.Millisecond * timeoutMs,
	}

	var (
		success int64
		failed  int64
		sem     = make(chan struct{}, concurrency)
		costs   = make([]int64, 0)
		lock    sync.Mutex
	)

	// 实时打印的定时器
	go func() {
		for {
			// 每秒打印一次当前状态
			time.Sleep(time.Second)
			printStats(success, failed)
		}
	}()

	// 无限请求
	for {
		sem <- struct{}{}
		go func() {
			req := fasthttp.AcquireRequest()
			resp := fasthttp.AcquireResponse()

			req.SetRequestURI(url)
			req.Header.SetMethod("POST")
			req.Header.SetContentType("application/json")
			req.SetBody(jsonData)

			begin := time.Now()
			err := client.Do(req, resp)
			cost := time.Since(begin).Milliseconds()

			// 请求处理结果
			if err != nil || resp.StatusCode() != 200 {
				atomic.AddInt64(&failed, 1)
			} else {
				atomic.AddInt64(&success, 1)
			}

			lock.Lock()
			costs = append(costs, cost)
			lock.Unlock()

			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)

			<-sem // 释放一个并发
		}()
	}
}

func printStats(success, failed int64) {
	total := success + failed
	qps := float64(total) / time.Since(time.Now()).Seconds()

	// 简单排序算 P95 / P99
	// sortInt64(costs) 可以不使用，直接打印了，因为我们不再处理总请求

	fmt.Println("====== DSP 压测实时结果 ======")
	fmt.Printf("总请求数: %d\n", total)
	fmt.Printf("成功数: %d\n", success)
	fmt.Printf("失败数: %d\n", failed)
	fmt.Printf("QPS: %.2f\n", qps)

	// 如果你要计算P95/P99等，可以根据需要进行排序
}

func sortInt64(a []int64) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}
