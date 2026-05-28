package client

import (
	"context"
	"github.com/cxb116/DSP/global"
	"github.com/redis/go-redis/v9"
	"log"
)

// redis 集群
func RedisClusterClientConnect() (*redis.ClusterClient, error) {

	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        global.EngineConfig.Redis.Cluster.Addrs,
		Password:     global.EngineConfig.Redis.Cluster.Password,
		PoolSize:     global.EngineConfig.Redis.Cluster.PoolSize,
		DialTimeout:  global.EngineConfig.Redis.Cluster.DialTimeout,
		ReadTimeout:  global.EngineConfig.Redis.Cluster.ReadTimeout,
		WriteTimeout: global.EngineConfig.Redis.Cluster.WriteTimeout,
	})

	if status := redisClient.Ping(context.Background()); status.Err() != nil {
		return nil, status.Err()
	}
	log.Println("redis connect success")
	return redisClient, nil
}

func NewClusterClientRedis() *redis.ClusterClient {
	client, err := RedisClusterClientConnect()

	if err != nil {
		log.Printf("redis cluster connect fail: %v", err)
		return nil
	}
	log.Println("redis connect success")
	return client
}

// 单体
func RedisClientConnect() (*redis.Client, error) {
	if len(global.EngineConfig.Redis.Cluster.Addrs) == 0 {
		return nil, redis.Nil
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         global.EngineConfig.Redis.Cluster.Addrs[0],
		Password:     global.EngineConfig.Redis.Cluster.Password,
		PoolSize:     global.EngineConfig.Redis.Cluster.PoolSize,
		DialTimeout:  global.EngineConfig.Redis.Cluster.DialTimeout,
		ReadTimeout:  global.EngineConfig.Redis.Cluster.ReadTimeout,
		WriteTimeout: global.EngineConfig.Redis.Cluster.WriteTimeout,
	})
	if status := redisClient.Ping(context.Background()); status.Err() != nil {
		return nil, status.Err()
	}
	log.Println("redis connect success")
	return redisClient, nil

}

func NewRedisClient() *redis.Client {
	client, err := RedisClientConnect()

	if err != nil {
		log.Printf("redis connect fail: %v", err)
		return nil
	}
	log.Println("redis connect success")
	return client
}
