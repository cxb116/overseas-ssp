package config

type DbBase struct {
	Host         string `json:"host" yaml:"host" mapstructure:"host"`
	Port         string `json:"port" yaml:"port" mapstructure:"port"`
	Username     string `json:"username" yaml:"username" mapstructure:"username"`
	Password     string `json:"password" yaml:"password" mapstructure:"password"`
	DBname       string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
}
