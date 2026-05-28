//go:build ignore

package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("       测试 etcd 监听")
	fmt.Println("========================================")
	fmt.Println()

	// 连接 etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.72.67.186:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("❌ 连接 etcd 失败: %v\n", err)
		return
	}
	defer cli.Close()

	fmt.Printf("✅ 成功连接 etcd\n")
	fmt.Println()

	// 测试监听
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	prefix := "/dsp/config/"
	watchChan := cli.Watch(ctx, prefix, clientv3.WithPrefix())

	fmt.Printf("🔍 开始监听前缀: %s\n", prefix)
	fmt.Println("========================================")
	fmt.Println("⏳ 等待事件... (请现在去后台删除一个 dspLaunch)")
	fmt.Println("========================================")
	fmt.Println()

	// 等待事件
	timeout := time.NewTimer(60 * time.Second)
	defer timeout.Stop()

	for {
		select {
		case watchResp := <-watchChan:
			if watchResp.Err() != nil {
				fmt.Printf("❌ Watch 错误: %v\n", watchResp.Err())
				continue
			}

			fmt.Printf("✅ 收到 %d 个事件\n", len(watchResp.Events))
			for i, event := range watchResp.Events {
				fmt.Printf("\n【事件 %d】\n", i+1)
				fmt.Printf("Key:      %s\n", string(event.Kv.Key))
				fmt.Printf("Type:     %s\n", event.Type)
				fmt.Printf("Value:    %s\n", string(event.Kv.Value))
				fmt.Printf("Version:  %d\n", event.Kv.Version)
			}

			if len(watchResp.Events) > 0 {
				fmt.Println("\n========================================")
				fmt.Println("✅ 监听正常工作！")
				fmt.Println("========================================")
				return
			}

		case <-timeout.C:
			fmt.Println("❌ 60 秒内没有收到任何事件")
			fmt.Println("可能的原因：")
			fmt.Println("1. 后台没有发送删除消息")
			fmt.Println("2. etcd 前缀配置不匹配")
			fmt.Println("3. 网络问题")
			return
		}
	}
}
