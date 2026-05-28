package impl

import (
	"sync"
)

var GlobalCacheResponseMap *CacheResponseMap

type CacheResponseMap struct {
	SspSlotInfo   int64
	CacheRespMaps map[int64][]BidResponse
	mu            sync.RWMutex
}

func NewCacheResponseMap() *CacheResponseMap {
	return &CacheResponseMap{
		CacheRespMaps: make(map[int64][]BidResponse),
	}
}

// AddCache 追加写入
//func (c *CacheResponseMap) AddCache(slotID int64, data []BidResponse) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	c.CacheRespMaps[slotID] = append(c.CacheRespMaps[slotID], data...)
//}

func (c *CacheResponseMap) AddCache(slotID int64, data []BidResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()

	list := append(c.CacheRespMaps[slotID], data...)

	// 超过阈值：删除最早的1000条
	if len(list) > 3000 {
		list = list[1000:]

		// 可选：重新分配底层数组，避免内存一直占用
		newList := make([]BidResponse, len(list))
		copy(newList, list)
		list = newList
	}

	c.CacheRespMaps[slotID] = list
}

func (c *CacheResponseMap) GetOne(slotID int64) (BidResponse, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	data, ok := c.CacheRespMaps[slotID]
	if !ok || len(data) == 0 {
		return BidResponse{}, false
	}

	// 取第一个
	res := data[0]

	// 剩余的放回去（相当于 pop）
	if len(data) == 1 {
		delete(c.CacheRespMaps, slotID)
	} else {
		c.CacheRespMaps[slotID] = data[1:]
	}

	return res, true
}

//  根据 slotId 获取长度
func (c *CacheResponseMap) DspSlotInfoLen(slotID int64) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.CacheRespMaps[slotID])
}
