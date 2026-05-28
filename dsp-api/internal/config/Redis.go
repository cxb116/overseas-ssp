package config

import "time"

type Redis struct {
	Cluster Cluster `json:"cluster" yaml:"cluster" mapstructure:"cluster"`
}

type Cluster struct {
	Addrs        []string      `json:"addrs" yaml:"addrs" mapstructure:"addrs"`
	Password     string        `json:"password" yaml:"password" mapstructure:"password"`
	PoolSize     int           `json:"pool_size" yaml:"pool_size" mapstructure:"pool_size"`
	DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout" mapstructure:"write_timeout"`
}
