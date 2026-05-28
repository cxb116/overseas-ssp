package config

type Log struct {
	LogLevel        string `json:"logLevel" yaml:"logLevel" mapstructure:"logLevel"`
	LogConsole      string `json:"logConsole" yaml:"logConsole" mapstructure:"logConsole"`
	LogDir          string `json:"logDir" yaml:"logDir" mapstructure:"logDir"`
	LogMaxSize      int    `json:"logMaxSize" yaml:"logMaxSize" mapstructure:"logMaxSize"`
	LogMaxBackup    int    `json:"logMaxBackup" yaml:"logMaxBackup" mapstructure:"logMaxBackup"`
	LogEncodingMode string `json:"logEncodingMode" yaml:"logEncodingMode" mapstructure:"logEncodingMode"`
	LogNoCaller     bool   `json:"logNoCaller" yaml:"logNoCaller" mapstructure:"logNoCaller"`
	LogTimestamp    bool   `json:"logTimestamp" yaml:"logTimestamp" mapstructure:"logTimestamp"`
	LogAsync        bool   `json:"logAsync" yaml:"logAsync" mapstructure:"logAsync"`
	ErrorStackTrace bool   `json:"errorStackTrace" yaml:"errorStackTrace" mapstructure:"errorStackTrace"`
	LogConsoleColor bool   `json:"logConsoleColor" yaml:"logConsoleColor" mapstructure:"logConsoleColor"`
	OpenCheat       bool   `json:"openCheat" yaml:"openCheat" mapstructure:"openCheat"`
	LogUtcTime      bool   `json:"logUtcTime" yaml:"logUtcTime" mapstructure:"logUtcTime"`
	// 业务日志目录
	BusinessLogDir string `json:"businessLogDir" yaml:"businessLogDir" mapstructure:"businessLogDir"`
	// 错误日志目录
	ErrorLogDir string `json:"errorLogDir" yaml:"errorLogDir" mapstructure:"errorLogDir"`
	// 日志保留天数（0表示不删除）
	LogRetentionDays int `json:"logRetentionDays" yaml:"logRetentionDays" mapstructure:"logRetentionDays"`
}
