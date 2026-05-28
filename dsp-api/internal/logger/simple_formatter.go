package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

// SimpleFormatter 简单格式化器，输出格式: 20:20:15:消息内容
type SimpleFormatter struct {
	w io.Writer
}

func (f *SimpleFormatter) Write(p []byte) (n int, err error) {
	// 尝试解析 zerolog 的 JSON 输出
	var data map[string]interface{}
	if err := json.Unmarshal(p, &data); err != nil {
		// JSON解析失败，尝试手动提取
		_, err = f.extractAndWrite(p)
		if err != nil {
			return 0, err
		}
		return len(p), nil
	}

	var buf bytes.Buffer

	// 提取时间并格式化
	if t, ok := data["time"]; ok {
		var timeStr string
		switch v := t.(type) {
		case float64: // Unix 时间戳
			seconds := int64(v)
			timeStr = time.Unix(seconds, 0).Format("15:04:05")
		case string:
			// 已经是格式化的时间，提取时间部分
			// 可能是 "2006-01-02 15:04:05" 或 "15:04:05" 格式
			if len(v) >= 19 {
				// 假设是 "2006-01-02 15:04:05" 格式，提取时间部分
				timeStr = v[11:19]  // "15:04:05"
			} else if len(v) >= 8 {
				timeStr = v
			}
		}
		if len(timeStr) > 0 {
			buf.WriteString(timeStr)
			buf.WriteString(":")  // 用冒号分隔
		}
	}

	// 提取消息
	if m, ok := data["msg"]; ok {
		if msg, ok := m.(string); ok {
			buf.WriteString(msg)
		}
	}

	// 提取错误
	if e, ok := data["err"]; ok {
		if errStr := fmt.Sprintf("%v", e); len(errStr) > 0 {
			if buf.Len() > 0 {
				buf.WriteString(" err:")
			}
			buf.WriteString(errStr)
		}
	}

	// 提取堆栈跟踪
	if stack, ok := data["stack"]; ok {
		if stackStr, ok := stack.(string); ok && len(stackStr) > 0 {
			buf.WriteString("\n")
			buf.WriteString(stackStr)
		}
	}

	// 添加换行
	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	_, err = f.w.Write(buf.Bytes())
	if err != nil {
		return 0, err
	}
	return len(p), nil  // 返回原始输入长度，不是格式化后的长度
}

// extractAndWrite 当JSON解析失败时，手动提取时间和消息
func (f *SimpleFormatter) extractAndWrite(p []byte) (n int, err error) {
	var buf bytes.Buffer
	str := string(p)

	// 提取时间字段
	timeRe := regexp.MustCompile(`"time":"?([^,"}]+)"?`)
	if matches := timeRe.FindStringSubmatch(str); len(matches) > 1 {
		timeValue := matches[1]
		// 处理不同时间格式
		if strings.Contains(timeValue, " ") && len(timeValue) >= 19 {
			// "2006-01-02 15:04:05" 格式
			buf.WriteString(timeValue[11:19])
		} else if len(timeValue) >= 8 {
			// "15:04:05" 格式或 Unix 时间戳
			if strings.Contains(timeValue, ":") {
				buf.WriteString(timeValue)
			} else {
				// 可能是 Unix 时间戳
				if timestamp, err := time.Parse(time.RFC3339, timeValue); err == nil {
					buf.WriteString(timestamp.Format("15:04:05"))
				} else {
					buf.WriteString(timeValue)
				}
			}
		}
		buf.WriteString(":")
	}

	// 提取消息字段
	msgRe := regexp.MustCompile(`"msg":"?((?:[^"\\]|\\.)*)"`)
	if matches := msgRe.FindStringSubmatch(str); len(matches) > 1 {
		// 处理转义字符
		msg := matches[1]
		msg = strings.ReplaceAll(msg, "\\n", "\n")
		msg = strings.ReplaceAll(msg, "\\t", "\t")
		msg = strings.ReplaceAll(msg, "\\\"", "\"")
		msg = strings.ReplaceAll(msg, "\\\\", "\\")
		buf.WriteString(msg)
	}

	// 添加换行
	if buf.Len() > 0 {
		buf.WriteString("\n")
	}

	_, err = f.w.Write(buf.Bytes())
	if err != nil {
		return 0, err
	}
	return len(p), nil  // 返回原始输入长度
}

// NewSimpleFormatter 创建简单格式化器
func NewSimpleFormatter(w io.Writer) io.Writer {
	return &SimpleFormatter{w: w}
}
