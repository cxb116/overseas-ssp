package client

import (
	"flag"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/config"
	"github.com/spf13/viper"
	"log"
	"time"
)

func NewClientViper() *viper.Viper {

	configPath := getConfigPath()

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("fatal error config file: %v", err)
		return nil
	}

	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//	if err := v.Unmarshal(&global.EngineConfig); err != nil {
	//		fmt.Println(err)
	//	}
	//})

	if err := v.Unmarshal(&global.EngineConfig); err != nil {
		log.Printf("config unmarshal failed: %v", err)
		return nil
	}

	global.EngineConfig.RequestLink = config.NormalizeDuration(global.EngineConfig.RequestLink, 1300*time.Millisecond)
	global.EngineConfig.DspRequest = config.NormalizeDuration(global.EngineConfig.DspRequest, time.Second)
	global.EngineConfig.Kafka.Frequency = config.NormalizeDuration(global.EngineConfig.Kafka.Frequency, 500*time.Millisecond)

	return v
}

func getConfigPath() (config string) {
	flag.StringVar(&config, "c", "config.json", "choose config file.")
	flag.Parse() // 解析命令行参数

	return config
}
