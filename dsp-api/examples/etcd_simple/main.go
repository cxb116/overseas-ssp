package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("🚀 ETCD 简单示例")
	fmt.Println("=================")
	fmt.Println()

	// 1. 连接 etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.72.67.186:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("❌ 连接失败: %v\n", err)
		return
	}
	defer cli.Close()

	fmt.Println("✅ 连接成功")

	// 2. 启动监听（在后台运行）
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		fmt.Println("\n🔍 开始监听: /dsp/key/")
		fmt.Println("=================")

		watchChan := cli.Watch(ctx, "/dsp/key/", clientv3.WithPrefix())
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				fmt.Printf("\n📨 收到事件:\n")
				fmt.Printf("  Key:    %s\n", string(event.Kv.Key))
				fmt.Printf("  Action: %s\n", event.Type)
				fmt.Printf("  Value:  %s\n", string(event.Kv.Value))
			}
		}
	}()

	// 等待监听启动
	time.Sleep(1 * time.Second)

	// 3. 演示操作
	fmt.Println("\n📝 演示操作:")
	fmt.Println("=================")

	// 添加数据
	fmt.Println("\n1️⃣ 添加数据...")
	cli.Put(ctx, "/dsp/key/test/add/1", `{"id":1,"name":"测试1"}`)
	time.Sleep(500 * time.Millisecond)

	// 修改数据
	fmt.Println("2️⃣ 修改数据...")
	cli.Put(ctx, "/dsp/key/test/add/1", `{"id":1,"name":"测试1-已修改"}`)
	time.Sleep(500 * time.Millisecond)

	// 删除数据
	fmt.Println("3️⃣ 删除数据...")
	cli.Delete(ctx, "/dsp/key/test/add/1")
	time.Sleep(500 * time.Millisecond)

	// 4. 查询数据
	fmt.Println("\n📊 查询所有数据:")
	fmt.Println("=================")
	resp, _ := cli.Get(ctx, "/dsp/key/", clientv3.WithPrefix())
	fmt.Printf("共找到 %d 条数据:\n", resp.Count)
	for _, kv := range resp.Kvs {
		fmt.Printf("  - %s = %s\n", string(kv.Key), string(kv.Value))
	}

	// 5. 持续监听
	fmt.Println("\n⏳ 持续监听中... (按 Ctrl+C 退出)")
	fmt.Println("=================")

	// 手动添加测试数据
	fmt.Println("\n💡 你可以在另一个终端手动添加数据测试:")
	fmt.Println("   etcdctl put /dsp/key/test/add/999 '{\"id\":999,\"name\":\"手动测试\"}'")
	fmt.Println()

	select {} // 持续运行
}
