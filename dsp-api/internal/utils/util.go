package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func MD5String(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(s string) int {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return atoi
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StringToFloat(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0)
	}
	return float
}

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}
	return i
}

func Int64ToFloat64(i int64) float64 {
	return float64(i)
}
