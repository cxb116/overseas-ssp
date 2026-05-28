package main

import (
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/httpEngine"
	"github.com/cxb116/DSP/internal/client"
	"github.com/cxb116/DSP/internal/logger"
	_ "github.com/go-sql-driver/mysql"
)

// 缓存

func main() {

	global.EngineViper = client.NewClientViper() // 加载配置文件
	if global.EngineViper == nil {
		return
	}

	// 立即初始化日志系统（必须在其他组件初始化之前）
	logger.InitDefaultLogger()
	// 初始化存储物料的内存
	//GlobalCacheResponseMap = impl.NewCacheResponseMap()
	global.EngineDB = client.NewClientMysql()    // 链接mysql
	global.EngineRedis = client.NewRedisClient() // 链接redis
	global.EngineETCD = client.NewEtcdClient()   // 链接etcd
	if global.EngineDB == nil || global.EngineRedis == nil || global.EngineETCD == nil {
		logger.ErrorLog.Error().Msg("startup dependency initialization failed")
		return
	}

	if err := client.InitKafkaAsyncProducer(global.EngineConfig.Kafka.Brokers); err != nil {
		logger.ErrorLog.Error().Msgf("failed to initialize kafka producer: %v", err)
		return
	}

	// 初始化 Kafka topics（在启动时预先创建，避免每次发送时检查）
	if err := client.InitTopics(); err != nil {
		logger.ErrorLog.Error().Msgf("failed to initialize kafka topics: %v", err)
		return
	}

	// 启动 Kafka 统计监控（每 30 秒打印一次统计信息）

	// 启动 HTTP 服务器（阻塞）
	//httpEngine.ServerEngine()

	httpEngine.GnetEngine() // workerPool + fanout
}
