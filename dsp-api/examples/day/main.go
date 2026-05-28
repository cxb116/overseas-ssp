package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cxb116/DSP/internal/logger"
)

func CurrentDay() string {
	return time.Now().Format("20060102")
}

//获取当前分钟（YYYYMMDDHHmm）
func CurrentMinute() string {
	return time.Now().Format("200601021504")
}

func CurrentTenMinuteWindow() time.Time {
	now := time.Now()
	minute := (now.Minute() / 10) * 10
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		now.Hour(), minute, 0, 0,
		now.Location(),
	)
}

//func CurrentTenMinuteKey() string {
//	return CurrentTenMinuteWindow().Format("200601021504")
//}
//
//func main() {
//	//minute := CurrentMinute()
//	//
//	//day := CurrentDay()
//	//
//	//fmt.Println(minute, day)
//	//
//	//window := CurrentTenMinuteWindow
//	//
//	//fmt.Println(window)
//	//
//	//key := CurrentTenMinuteKey
//	//fmt.Println(key)
//
//	//event := TenMinuteWindowByEvent
//
//	fmt.Println(TimeNowUnix() / 1000)
//}
//
//func TimeNowUnix() int64 {
//	return time.Now().Unix()
//}
//
//func TenMinuteWindowByEvent(eventTime int64) time.Time {
//
//	fmt.Println("eventTime:", TimeNowUnix()/1000)
//	return time.Unix(TimeNowUnix()/1000, 0)
//}
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

func main() {
	key := Get10MinTime(time.Now())
	fmt.Printf("key:%d", key)
}
