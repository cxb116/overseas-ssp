package global

import (
	"github.com/IBM/sarama"

	"github.com/cxb116/DSP/internal/config"
	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

var (
	EngineViper  *viper.Viper
	EngineConfig config.Server // 引擎配置
	EngineDB     *gorm.DB

	EngineRedis *redis.Client
	EngineKafka *sarama.KafkaVersion
	EngineETCD  *clientv3.Client
)
