package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func NewClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.72.67.186:2379"},
		DialTimeout: 5 * time.Second,
	})
}

func main() {
	cli, err := NewClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 写一个测试 key
	_, err = cli.Put(ctx, "/test/hello", "world")
	if err != nil {
		fmt.Println("Put failed:", err)
		return
	}

	// 读回测试 key
	resp, err := cli.Get(ctx, "/test/hello")
	if err != nil {
		fmt.Println("Get failed:", err)
		return
	}

	fmt.Println("Etcd connected, value:", string(resp.Kvs[0].Value))
}
