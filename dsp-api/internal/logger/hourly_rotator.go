package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

// HourlyRotator 按小时轮转的日志写入器
type HourlyRotator struct {
	dir             string
	filePrefix      string
	fileExt         string
	retentionDays   int // 日志保留天数，0表示不删除

	currentFile *os.File
	currentHour int
	mu          sync.Mutex

	closed bool
}

// NewHourlyRotator 创建一个按小时轮转的日志写入器
func NewHourlyRotator(dir, filePrefix, fileExt string, retentionDays int) (*HourlyRotator, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	hr := &HourlyRotator{
		dir:           dir,
		filePrefix:    filePrefix,
		fileExt:       fileExt,
		retentionDays: retentionDays,
	}

	if err := hr.openCurrentFile(); err != nil {
		return nil, err
	}

	// 启动后台协程检查小时变化
	go hr.watchHourChange()

	// 启动后台协程定期清理旧日志
	if retentionDays > 0 {
		go hr.startCleaner()
	}

	return hr, nil
}

// openCurrentFile 打开当前小时的日志文件
func (hr *HourlyRotator) openCurrentFile() error {
	now := time.Now()
	currentHour := now.Hour()
	dateStr := now.Format("2006-01-02")
	hourStr := fmt.Sprintf("%02d", currentHour)

	filename := fmt.Sprintf("%s-%s-%s.%s", hr.filePrefix, dateStr, hourStr, hr.fileExt)
	filepath := filepath.Join(hr.dir, filename)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	hr.mu.Lock()
	if hr.currentFile != nil {
		hr.currentFile.Close()
	}
	hr.currentFile = file
	hr.currentHour = currentHour
	hr.mu.Unlock()

	return nil
}

// watchHourChange 监控小时变化并轮转日志文件
func (hr *HourlyRotator) watchHourChange() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for !hr.closed {
		<-ticker.C
		now := time.Now()
		currentHour := now.Hour()

		hr.mu.Lock()
		if hr.currentHour != currentHour {
			hr.mu.Unlock()
			hr.openCurrentFile()
		} else {
			hr.mu.Unlock()
		}
	}
}

// Write 实现 io.Writer 接口
func (hr *HourlyRotator) Write(p []byte) (n int, err error) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	if hr.closed {
		return 0, fmt.Errorf("hourly rotator is closed")
	}

	if hr.currentFile == nil {
		if err := hr.openCurrentFile(); err != nil {
			return 0, err
		}
	}

	return hr.currentFile.Write(p)
}

// Close 关闭日志写入器
func (hr *HourlyRotator) Close() error {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.closed = true

	if hr.currentFile != nil {
		return hr.currentFile.Close()
	}

	return nil
}

// Sync 刷新文件缓冲区
func (hr *HourlyRotator) Sync() error {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	if hr.currentFile == nil {
		return nil
	}

	return hr.currentFile.Sync()
}

// startCleaner 启动日志清理协程
func (hr *HourlyRotator) startCleaner() {
	// 每天凌晨1点执行清理
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	// 计算下一次凌晨1点的时间
	now := time.Now()
	nextClean := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, now.Location())
	if nextClean.Before(now) {
		nextClean = nextClean.Add(24 * time.Hour)
	}
	initialDelay := nextClean.Sub(now)
	if initialDelay > 0 {
		time.Sleep(initialDelay)
		hr.cleanOldLogs()
	}

	for !hr.closed {
		<-ticker.C
		hr.cleanOldLogs()
	}
}

// cleanOldLogs 清理过期的日志文件
func (hr *HourlyRotator) cleanOldLogs() {
	if hr.retentionDays <= 0 {
		return
	}

	// 计算截止时间
	cutoffTime := time.Now().AddDate(0, 0, -hr.retentionDays)

	// 遍历日志目录
	entries, err := os.ReadDir(hr.dir)
	if err != nil {
		return
	}

	// 匹配日志文件名的正则表达式，例如：data-2026-03-04-15.log 或 err-2026-03-04-15.log
	pattern := fmt.Sprintf(`^%s-(\d{4})-(\d{2})-(\d{2})-(\d{2})\.%s$`, hr.filePrefix, hr.fileExt)
	re := regexp.MustCompile(pattern)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		matches := re.FindStringSubmatch(filename)
		if matches == nil {
			continue
		}

		// 解析文件名中的时间
		year := matches[1]
		month := matches[2]
		day := matches[3]
		hour := matches[4]

		fileTime, err := time.Parse("2006-01-02-15", fmt.Sprintf("%s-%s-%s-%s", year, month, day, hour))
		if err != nil {
			continue
		}

		// 如果文件时间早于截止时间，删除文件
		if fileTime.Before(cutoffTime) {
			filePath := filepath.Join(hr.dir, filename)
			os.Remove(filePath)
		}
	}
}
