package config

import "time"

type Kafka struct {
	Brokers       []string      `json:"brokers" yaml:"brokers" mapstructure:"brokers"`
	Frequency     time.Duration `json:"frequency" yaml:"frequency" mapstructure:"frequency"`
	Messages      int           `json:"messages" yaml:"messages" mapstructure:"messages"`
	ReturnSuccess bool          `json:"returnSuccess" yaml:"returnSuccess" mapstructure:"returnSuccess"`
	TopicReqData  string        `json:"topicReqData" yaml:"topicReqData" mapstructure:"topicReqData"` // 请求的topic
	TopicResData  string        `json:"topicResData" yaml:"topicResData" mapstructure:"topicResData"` // 响应的topic
}
