package logger

import (
	"fmt"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/config"

	"io"
	"os"

	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ravinggo/zerolog"
	"github.com/ravinggo/zerolog/log"
	"github.com/ravinggo/zerolog/pkgerrors"
)

type ILogger interface {
	Trace() *Event
	Debug() *Event
	Info() *Event
	Warn() *Event
	Error() *Event
	Fatal() *Event
	Panic() *Event
	NoLevel() *Event
	Disabled() *Event
	WithLevel(Level) *Event
}

type (
	Logger  = zerolog.Logger
	Context = zerolog.Context
	Event   = zerolog.Event
	Level   = zerolog.Level
)

var (
	Log         *Logger // 系统日志
	LogSkip3    Logger  // 跳过3帧调用栈的日志
	BusinessLog *Logger // 业务日志
	ErrorLog    *Logger // 错误日志
)

func InitDefaultLogger() {
	envCnf := global.EngineConfig.Log
	// 强制使用可读时间格式
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.TimestampFunc = func() time.Time {
		return time.Now()
	}
	if envCnf.LogUtcTime {
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().UTC()
		}
	}
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "err"
	// 禁用所有字段，只保留时间和消息
	zerolog.LevelFieldName = ""
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return ""
	}
	zerolog.DisableSampling(true)

	// 配置控制台输出 Writer
	var writers []io.Writer
	if envCnf.LogConsole == "stdout" {
		writers = append(writers, os.Stdout)
	} else if envCnf.LogConsole == "stderr" {
		writers = append(writers, os.Stderr)
	}

	// 系统日志使用纯文本格式输出到控制台
	var writer io.Writer
	if len(writers) > 1 {
		writer = zerolog.MultiLevelWriter(writers...)
	} else if len(writers) == 0 {
		writer = io.Discard
	} else {
		writer = writers[0]
	}

	// 包装 SimpleFormatter 以确保控制台输出也是纯文本格式
	writer = NewSimpleFormatter(writer)

	if envCnf.LogAsync {
		writer = NewAsync(writer)
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	ctx := zerolog.New(writer).
		Level(getLoggerLevel()).With().Timestamp()
	// 系统日志也使用纯文本格式，不使用 console formatter
	if !envCnf.LogNoCaller {
		ctx = ctx.Caller()
	}
	log.Logger = ctx.Logger()
	Log = &log.Logger
	LogSkip3 = Log.With().CallerWithSkipFrameCount(3).Logger()

	// 初始化业务日志和错误日志
	initClassifiedLoggers(envCnf)
}

// initClassifiedLoggers 初始化业务日志和错误日志
func initClassifiedLoggers(envCnf config.Log) {
	// 初始化业务日志（只输出到文件，不输出到控制台）
	if envCnf.BusinessLogDir != "" {
		businessRotator, err := NewHourlyRotator(envCnf.BusinessLogDir, "data", "log", envCnf.LogRetentionDays)
		if err != nil {
			log.Warn().Msgf("Failed to create business log rotator: %v", err)
		} else {
			// 业务日志文件 Writer（带格式化）
			businessFileWriter := NewSimpleFormatter(businessRotator)

			// 只对整体 writer 应用一次 Async
			if envCnf.LogAsync {
				businessFileWriter = NewAsync(businessFileWriter)
			}

			// 业务日志：使用纯文本格式，只写文件
			businessCtx := zerolog.New(businessFileWriter).
				Level(getLoggerLevel()).With().Timestamp()
			if !envCnf.LogNoCaller {
				businessCtx = businessCtx.Caller()
			}

			logger := businessCtx.Logger()
			BusinessLog = &logger
		}
	}

	// 初始化错误日志（只输出到文件，不输出到控制台）
	if envCnf.ErrorLogDir != "" {
		errorRotator, err := NewHourlyRotator(envCnf.ErrorLogDir, "err", "log", envCnf.LogRetentionDays)
		if err != nil {
			log.Warn().Msgf("Failed to create error log rotator: %v", err)
		} else {
			// 错误日志文件 Writer（带格式化）
			errorFileWriter := NewSimpleFormatter(errorRotator)

			// 只对整体 writer 应用一次 Async
			if envCnf.LogAsync {
				errorFileWriter = NewAsync(errorFileWriter)
			}

			// 错误日志：使用纯文本格式，只写文件
			errorCtx := zerolog.New(errorFileWriter).
				Level(zerolog.ErrorLevel).With().Timestamp()
			if !envCnf.LogNoCaller {
				errorCtx = errorCtx.Caller()
			}

			logger := errorCtx.Logger()
			ErrorLog = &logger
		}
	}

	// 如果业务日志未初始化，使用系统日志
	if BusinessLog == nil {
		BusinessLog = Log
	}

	// 如果错误日志未初始化，使用系统日志
	if ErrorLog == nil {
		ErrorLog = Log
	}
}

// SetLogger set default logger
func SetLogger(l Logger) {
	log.Logger = l
	Log = &l
	LogSkip3 = Log.With().CallerWithSkipFrameCount(3).Logger()
}

func getLoggerLevel() zerolog.Level {
	switch strings.ToUpper(global.EngineConfig.Log.LogLevel) {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "PANIC":
		return zerolog.PanicLevel
	case "DISABLED":
		return zerolog.Disabled
	default:
		return zerolog.DebugLevel
	}
}

func MarshalStack(err error) interface{} {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	var sterr stackTracer
	var ok bool
	for err != nil {
		sterr, ok = err.(stackTracer)
		if ok {
			break
		}

		u, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			return nil
		}

		err = u.Unwrap()
	}
	if sterr == nil {
		return nil
	}

	st := sterr.StackTrace()
	return fmt.Sprintf("%+v", st)
}
