package utils

import (
	"fmt"
	"hash/crc32"
)

// ims:oaid:dspslotId:sspSlotId:ts
func ImsHSetCheKey(deviceId, dspSlotId string, sspSlotId, ts int64) string {
	h := crc32.ChecksumIEEE([]byte(deviceId))
	key := fmt.Sprintf("ims:%d:%s:%d:%d", h, dspSlotId, sspSlotId, ts/1000)
	return key
}

func ClsHSetCheKey(deviceId, dspSlotId string, sspSlotId, ts int64) string {
	h := crc32.ChecksumIEEE([]byte(deviceId))
	key := fmt.Sprintf("cls:%d:%s:%d:%d", h, dspSlotId, sspSlotId, ts/1000)
	return key
}

func DownHSetCheKey(deviceId, dspSlotId string, sspSlotId, ts int64) string {
	h := crc32.ChecksumIEEE([]byte(deviceId))
	key := fmt.Sprintf("down:%d:%s:%d:%d", h, dspSlotId, sspSlotId, ts/1000)
	return key
}

func InsHSetCheKey(deviceId, dspSlotId string, sspSlotId, ts int64) string {
	h := crc32.ChecksumIEEE([]byte(deviceId))
	key := fmt.Sprintf("ins:%d:%s:%d:%d", h, dspSlotId, sspSlotId, ts/1000)
	return key
}

func ActHSetCheKey(deviceId, dspSlotId string, sspSlotId, ts int64) string {
	h := crc32.ChecksumIEEE([]byte(deviceId))
	key := fmt.Sprintf("act:%d:%s:%d:%d", h, dspSlotId, sspSlotId, ts/1000)
	return key
}
