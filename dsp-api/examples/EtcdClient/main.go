package main

import (
	"context"
	"fmt"
	"github.com/cxb116/DSP/global"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 1️⃣ 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{global.EngineConfig.Etcd},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("etcd connect error:", err)
	}
	defer cli.Close()

	ctx := context.Background()

	// 2️⃣ 启动监听
	go watchKey(cli, "/test/")

	// 3️⃣ 模拟发送数据
	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("/test/key%d", i)
		value := fmt.Sprintf("value-%d", i)

		_, err := cli.Put(ctx, key, value)
		if err != nil {
			log.Println("put error:", err)
			continue
		}

		fmt.Println("写入:", key, value)
		time.Sleep(2 * time.Second)
	}

	select {}
}

func watchKey(cli *clientv3.Client, prefix string) {
	rch := cli.Watch(context.Background(), prefix, clientv3.WithPrefix())

	fmt.Println("开始监听:", prefix)

	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("监听到 PUT 事件: key=%s value=%s\n",
					string(ev.Kv.Key), string(ev.Kv.Value))

			case clientv3.EventTypeDelete:
				fmt.Printf("监听到 DELETE 事件: key=%s\n",
					string(ev.Kv.Key))
			}
		}
	}
}
