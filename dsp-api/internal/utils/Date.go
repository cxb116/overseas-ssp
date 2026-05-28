package utils

import (
	"strconv"
	"time"

	"github.com/cxb116/DSP/internal/logger"
)

func Gen10MinKeyTime(t time.Time) string {
	aligned := t.Truncate(10 * time.Minute)
	return aligned.Format("200601021504")
}

// 获取前10分钟时间，比如12:35,获取的时间说就是12:40
func Get10MinTime(t time.Time) int64 {
	str := t.Add(9 * time.Minute).
		Truncate(10 * time.Minute).
		Format("200601021504")
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logger.ErrorLog.Error().Msgf("Get10MinTime strconv.ParseInt failed: %v", err)
		return 0
	}
	return i
}

func Get10MinTimeString(t time.Time) string {
	return t.Add(9 * time.Minute).
		Truncate(10 * time.Minute).
		Format("200601021504")
}

func TimeNowUnix() int64 {
	return time.Now().Unix()
}

func TineNowUnixString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func NowTime() string {
	today := time.Now().Format("2006010215")
	return today
}
