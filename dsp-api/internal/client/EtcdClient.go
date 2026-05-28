package client

import (
	"context"
	"github.com/cxb116/DSP/global"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func EtcdClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{global.EngineConfig.Etcd},
		DialTimeout: 5 * time.Second,
	})
}

func NewEtcdClient() *clientv3.Client {
	EClient, err := EtcdClient()
	if err != nil {
		log.Printf("etcd client init fail: %v", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := EClient.Status(ctx, global.EngineConfig.Etcd); err != nil {
		log.Printf("etcd status fail: %v", err)
		_ = EClient.Close()
		return nil
	}

	log.Println("ETCD client init success")
	return EClient
}
