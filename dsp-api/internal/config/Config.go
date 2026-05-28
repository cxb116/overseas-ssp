package config

import "time"

type Server struct {
	Version      string        `json:"version" yaml:"version" mapstructure:"version"`
	Port         string        `json:"port" yaml:"port" mapstructure:"port"`
	RequestLink  time.Duration `json:"requestLink" yaml:"requestLink" mapstructure:"requestLink"`
	DspRequest   time.Duration `json:"dspRequest" yaml:"dspRequest" mapstructure:"dspRequest"`
	ImsLink      string        `json:"ims_link" yaml:"ims_link" mapstructure:"ims_link"`
	ClkLink      string        `json:"clk_link" yaml:"clk_link" mapstructure:"clk_link"`
	PriceEncrypt string        `json:"priceEncrypt" yaml:"priceEncrypt" mapstructure:"priceEncrypt"`
	WorkerSize   int           `json:"workerSize" yaml:"workerSize" mapstructure:"workerSize"`
	TaskQueue    int           `json:"taskQueue" yaml:"taskQueue" mapstructure:"taskQueue"`
	Redis        Redis         `json:"redis" yaml:"redis" mapstructure:"redis"`
	Database     DbBase        `json:"database" yaml:"database" mapstructure:"database"`
	Kafka        Kafka         `json:"kafka" yaml:"kafka" mapstructure:"kafka"`
	Etcd         string        `json:"etcd" yaml:"etcd" mapstructure:"etcd"`
	EtcdPrefix   string        `json:"etcdPrefix" yaml:"etcdPrefix" mapstructure:"etcdPrefix"` // etcd key 前缀，如 /dsp/config

	ReqTotalMinute  string `json:"reqTotalMinute" yaml:"reqTotalMinute" mapstructure:"reqTotalMinute"`
	ReqDiscarMinute string `json:"reqDiscarMinute" yaml:"reqDiscarMinute" mapstructure:"reqDiscarMinute"`

	//上报链接
	Ims  string `json:"ims" yaml:"ims" mapstructure:"ims"`
	Cls  string `json:"cls" yaml:"cls" mapstructure:"cls"`
	Down string `json:"down" yaml:"down" mapstructure:"down"`
	Ins  string `json:"ins" yaml:"ins" mapstructure:"ins"`
	Act  string `json:"act" yaml:"act" mapstructure:"act"`

	Log Log `json:"log" yaml:"log" mapstructure:"log"`
}
