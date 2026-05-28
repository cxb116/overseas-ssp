package impl

import (
	"context"
	"encoding/json"
	"math/rand/v2"
	"strconv"
	"strings"
	"sync"

	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

//var DspSlotInitMapsHandler *LoopDspEventMaps

var DspSlotGlobalData *LoopDspEventMaps

//func NewDspSlotMaps() *LoopDspEventMaps {
//	return &LoopDspEventMaps{
//		SspSlotInfoMaps: make(map[int64]*DspEventMaps),
//	}
//}

// 接收全部数据
func NewDspSlotManager() *LoopDspEventMaps {
	return &LoopDspEventMaps{
		DspSlotInfoMaps: make(map[int64]*DspSlotInfo),
		DspCompanyMaps:  make(map[int64]*DspCompany),
		DspLaunchMaps:   make(map[int64]*DspLaunch),
		SspSlotInfoMaps: make(map[int64]*SspSlotInfo),
	}
}

//type LoopDspEventMaps struct {
//	SspSlotInfoMaps map[int64]*DspEventMaps
//	LoopMutex       sync.Mutex
//}

type DspEventMaps struct {
	SspSlotInfo     *SspSlotInfo
	DspSlotInfoMaps map[int64]*DspSlotInfo // int64 是只记得 dspSlotId
	DspMutex        sync.Mutex
}

type LoopDspEventMaps struct {
	DspSlotInfoMaps map[int64]*DspSlotInfo
	DspCompanyMaps  map[int64]*DspCompany
	DspLaunchMaps   map[int64]*DspLaunch
	SspSlotInfoMaps map[int64]*SspSlotInfo
	RwMutex         sync.RWMutex
}

func (this *LoopDspEventMaps) RegLoopSubscribeMessage(ctx context.Context) {

	this.AllDbData()
	go this.DspSlotInfoWatchChan(ctx)
}

func (this *LoopDspEventMaps) AllDbData() {
	if global.EngineDB == nil {
		logger.Log.Error().Msg("数据库初始化 internal error")
		return
	}

	var dspSlotInfos []DspSlotInfo
	if err := global.EngineDB.Table("dsp_slot_info").
		Select("dsp_slot_info.*,pro.name AS product_name").Joins("LEFT JOIN dsp_product as pro ON dsp_slot_info.product_id = pro.id").Find(&dspSlotInfos).Error; err != nil {
		logger.Log.Error().Msgf("log %v", err)
	} else {
		DspSlotGlobalData.RwMutex.Lock()
		for _, dspslot := range dspSlotInfos {
			DspSlotGlobalData.DspSlotInfoMaps[dspslot.Id] = &dspslot
		}
		DspSlotGlobalData.RwMutex.Unlock()

		logger.Log.Info().Msgf(" DspSlotGlobalData DspSlotInfo 的长度 %v", len(DspSlotGlobalData.DspSlotInfoMaps))
	}

	var dspCompany []DspCompany
	if err := global.EngineDB.Find(&dspCompany).Error; err != nil {
		logger.Log.Error().Msgf("预算公司查询异常 %v", err)
	} else {
		DspSlotGlobalData.RwMutex.Lock()
		for _, company := range dspCompany {
			DspSlotGlobalData.DspCompanyMaps[company.Id] = &company
		}
		DspSlotGlobalData.RwMutex.Unlock()

		logger.Log.Info().Msgf("预算公司数据长度 %v", len(DspSlotGlobalData.DspCompanyMaps))
	}

	var dspLaunches []DspLaunch
	if err := global.EngineDB.Find(&dspLaunches).Error; err != nil {
		logger.Log.Error().Msgf("媒体预算关联数据异常: %v", err)
	} else {
		DspSlotGlobalData.RwMutex.Lock()
		for _, launch := range dspLaunches {
			DspSlotGlobalData.DspLaunchMaps[launch.Id] = &launch
		}
		DspSlotGlobalData.RwMutex.Unlock()

		logger.Log.Info().Msgf("媒体预算关联数据长度 %v", len(DspSlotGlobalData.DspLaunchMaps))
	}

	var sspSlotsInfo []SspSlotInfo
	err := global.EngineDB.
		Table("ssp_slot_info").
		Select("ssp_slot_info.id, ssp_slot_info.ad_type_id, ssp_slot_info.ssp_pay_type, ssp_slot_info.ssp_deal_ratio, ssp_slot_info.height, ssp_slot_info.width, ssp_slot_info.interaction_type, ssp_app.os_type as os_type, ssp_app.id as app_id").
		Joins("LEFT JOIN ssp_app ON ssp_slot_info.app_id = ssp_app.id").
		Where("ssp_slot_info.enable=1").
		Find(&sspSlotsInfo).Error

	if err != nil {
		logger.Log.Error().Msgf("媒体数据查询异常 %v", err)
	} else {
		DspSlotGlobalData.RwMutex.Lock()
		for _, sspslot := range sspSlotsInfo {
			DspSlotGlobalData.SspSlotInfoMaps[sspslot.Id] = &sspslot
		}
		DspSlotGlobalData.RwMutex.Unlock()
	}

}

// 根据SspSlotId 获取多个DspLaunchs
func GetDspLaunchesBySspSlotId(sspSlotId int64) []*DspLaunch {
	DspSlotGlobalData.RwMutex.RLock()
	defer DspSlotGlobalData.RwMutex.RUnlock()
	var dspLaunches []*DspLaunch
	for _, launch := range DspSlotGlobalData.DspLaunchMaps {
		if launch.SspSlotId == sspSlotId {
			dspLaunches = append(dspLaunches, launch)
		}
	}
	return dspLaunches
}

// 根据 dspSlotId 获取 DspSlotInfo 信息
func GetDspSlotInfo(dspSlotId int64) *DspSlotInfo {
	DspSlotGlobalData.RwMutex.RLock()
	defer DspSlotGlobalData.RwMutex.RUnlock()

	if DspSlotGlobalData.DspSlotInfoMaps[dspSlotId] == nil {
		return nil
	}
	return DspSlotGlobalData.DspSlotInfoMaps[dspSlotId]

}

// 根据 公司id 获取 DspCompany 信息
func GetDspCompany(companyId int64) *DspCompany {
	DspSlotGlobalData.RwMutex.RLock()
	defer DspSlotGlobalData.RwMutex.RUnlock()

	if DspSlotGlobalData.DspCompanyMaps[companyId] == nil {
		return nil
	}
	return DspSlotGlobalData.DspCompanyMaps[companyId]
}

func GetSspSlotInfo(sspSlotId int64) *SspSlotInfo {
	DspSlotGlobalData.RwMutex.RLock()
	defer DspSlotGlobalData.RwMutex.RUnlock()

	if DspSlotGlobalData.SspSlotInfoMaps[sspSlotId] == nil {
		return nil
	}
	return DspSlotGlobalData.SspSlotInfoMaps[sspSlotId]

}

func GetDspLaunch(dspLaunchId int64) *DspLaunch {
	DspSlotGlobalData.RwMutex.RLock()
	defer DspSlotGlobalData.RwMutex.RUnlock()

	if DspSlotGlobalData.DspLaunchMaps[dspLaunchId] == nil {
		return nil
	}

	return DspSlotGlobalData.DspLaunchMaps[dspLaunchId]
}

// 跟新内存数据DspSlotInfo
func AddOrUpdateDspSlotInfo(dspSlotInfo *DspSlotInfo) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	DspSlotGlobalData.DspSlotInfoMaps[dspSlotInfo.Id] = dspSlotInfo
}

func DeleteDspSlotInfo(dspSlotId int64) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	if DspSlotGlobalData.DspSlotInfoMaps[dspSlotId] != nil {
		delete(DspSlotGlobalData.DspSlotInfoMaps, dspSlotId)
	}

}

func AddOrUpdateDspCompany(dspCompany *DspCompany) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	DspSlotGlobalData.DspCompanyMaps[dspCompany.Id] = dspCompany

}

func DeleteDspCompany(companyId int64) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	if DspSlotGlobalData.DspCompanyMaps[companyId] != nil {
		delete(DspSlotGlobalData.DspCompanyMaps, companyId)
	}

}

func AddOrUpdateDspLaunch(dspLaunch *DspLaunch) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	DspSlotGlobalData.DspLaunchMaps[dspLaunch.Id] = dspLaunch
}

func DeleteDspLaunch(launchId int64) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	if DspSlotGlobalData.DspLaunchMaps[launchId] != nil {
		delete(DspSlotGlobalData.DspLaunchMaps, launchId)
	}
}

func AddOrUpdateSspSlotInfo(sspSlotInfo *SspSlotInfo) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	DspSlotGlobalData.SspSlotInfoMaps[sspSlotInfo.Id] = sspSlotInfo
}

func DeleteSspSlotInfo(sspSlotId int64) {
	DspSlotGlobalData.RwMutex.Lock()
	defer DspSlotGlobalData.RwMutex.Unlock()

	if DspSlotGlobalData.SspSlotInfoMaps[sspSlotId] != nil {
		delete(DspSlotGlobalData.SspSlotInfoMaps, sspSlotId)
	}

}

func (this *LoopDspEventMaps) DspSlotInfoWatchChan(ctx context.Context) {
	logger.Log.Info().Msgf("log")

	if global.EngineETCD == nil {
		logger.Log.Error().Msg("ETCD client not initialized")
		return
	}

	prefix := global.EngineConfig.EtcdPrefix
	if prefix == "" {
		prefix = "/dsp/config"
		logger.Log.Warn().Msgf("log")
	}

	logger.Log.Info().Msgf("log")
	logger.Log.Info().Msgf("log %v", global.EngineConfig.Etcd)
	logger.Log.Info().Msgf("log %v", prefix)

	watchPrefix := prefix + "/"
	logger.Log.Info().Msgf("log %v", watchPrefix)

	watchChan := global.EngineETCD.Watch(ctx, watchPrefix, clientv3.WithPrefix())

	logger.Log.Info().Msgf("log")
	logger.Log.Info().Msgf("========================================")

	eventCount := 0

	for {
		select {
		case <-ctx.Done():
			logger.Log.Warn().Msg("warning")
			return

		case watchResp, ok := <-watchChan:
			if !ok {
				logger.Log.Error().Msg("internal error")
				return
			}
			if watchResp.Err() != nil {
				logger.Log.Error().Msg("etcd watch response error")
				continue
			}
			eventCount += len(watchResp.Events)
			for _, event := range watchResp.Events {
				this.handleEtcdEvent(event, prefix)
			}
		}
	}
}

func (this *LoopDspEventMaps) handleEtcdEvent(event *clientv3.Event, prefix string) {
	key := string(event.Kv.Key)

	relKey := strings.TrimPrefix(key, prefix)
	relKey = strings.Trim(relKey, "/")

	parts := strings.Split(relKey, "/")
	if len(parts) != 2 {
		logger.Log.Warn().Msgf("log %v %v %v",
			key, relKey, len(parts))
		return
	}

	dataType := parts[0] // dsp, company, launch, sspslot
	idStr := parts[1]    // id

	logger.Log.Info().Msgf("log %v %v", event.Type, key)
	logger.Log.Info().Msgf("log %v %v", dataType, idStr)

	if event.Type == clientv3.EventTypeDelete {
		this.handleEtcdDelete(dataType, idStr)
		return
	}

	if event.Type == clientv3.EventTypePut {
		this.handleEtcdPut(dataType, idStr, event)
		return
	}

	logger.Log.Warn().Msgf("log %v %v", event.Type, key)
}

func (this *LoopDspEventMaps) handleEtcdPut(dataType, idStr string, event *clientv3.Event) {
	logger.Log.Debug().Msgf("log %v", string(event.Kv.Value))

	switch dataType {
	case "dsp":
		this.handleDspSlotInfoPut(idStr, event)
	case "company":
		this.handleDspCompanyPut(idStr, event)
	case "launch":
		this.handleDspLaunchPut(idStr, event)
	case "sspslot":
		this.handleSspSlotInfoPut(idStr, event)
	default:
		logger.Log.Warn().Msgf("log %v", dataType)
	}

	logger.Log.Info().Msgf("log")

}

func (this *LoopDspEventMaps) handleEtcdDelete(dataType, idStr string) {
	id := parseInt64(idStr)

	logger.Log.Info().Msgf("log")
	logger.Log.Info().Msgf("log %v %v", dataType, id)

	switch dataType {
	case "dsp":
		logger.Log.Info().Msgf("log %v", id)
		DeleteDspSlotInfo(id)
		logger.Log.Info().Msg("done")
	case "company":
		logger.Log.Info().Msgf("log %v", id)
		DeleteDspCompany(id)
		logger.Log.Info().Msg("delete company done")
	case "launch":
		logger.Log.Info().Msgf("log %v", id)
		DeleteDspLaunch(id)
		logger.Log.Info().Msg("delete launch done")
	case "sspslot":
		logger.Log.Info().Msgf("log %v", id)
		DeleteSspSlotInfo(id)
		logger.Log.Info().Msg("done")
	default:
		logger.ErrorLog.Warn().Msgf("没有匹配到响应的修改策略 %v", dataType)
	}

	logger.Log.Info().Msgf("log")
}

func (this *LoopDspEventMaps) handleDspSlotInfoPut(idStr string, event *clientv3.Event) {
	var dspSlotInfo DspSlotInfo
	if err := json.Unmarshal(event.Kv.Value, &dspSlotInfo); err != nil {
		logger.Log.Error().Msgf("dsp slot info unmarshal error: %v, value: %s", err, string(event.Kv.Value))
		return
	}
	AddOrUpdateDspSlotInfo(&dspSlotInfo)
	logger.Log.Info().Msgf("dsp slot info updated: id=%d name=%s code=%s", dspSlotInfo.Id, dspSlotInfo.Name, dspSlotInfo.DspSlotCode)
}

func (this *LoopDspEventMaps) handleDspCompanyPut(idStr string, event *clientv3.Event) {
	var dspCompany DspCompany
	if err := json.Unmarshal(event.Kv.Value, &dspCompany); err != nil {
		logger.Log.Error().Msgf("log %v %v", err, string(event.Kv.Value))
		return
	}
	AddOrUpdateDspCompany(&dspCompany)
	logger.Log.Info().Msgf("log %v %v %v", dspCompany.Id, dspCompany.Name, dspCompany.DspCode)
}

func (this *LoopDspEventMaps) handleDspLaunchPut(idStr string, event *clientv3.Event) {
	var dspLaunch DspLaunch
	if err := json.Unmarshal(event.Kv.Value, &dspLaunch); err != nil {
		logger.Log.Error().Msgf("log %v %v", err, string(event.Kv.Value))
		return
	}
	AddOrUpdateDspLaunch(&dspLaunch)
	logger.Log.Info().Msgf("log %v %v %v %v",
		dspLaunch.Id, dspLaunch.SspSlotId, dspLaunch.DspSlotId, dspLaunch.TrafficWeight)
}

func (this *LoopDspEventMaps) handleSspSlotInfoPut(idStr string, event *clientv3.Event) {
	var sspSlotInfo SspSlotInfo
	if err := json.Unmarshal(event.Kv.Value, &sspSlotInfo); err != nil {
		logger.Log.Error().Msgf("log %v %v", err, string(event.Kv.Value))
		return
	}
	AddOrUpdateSspSlotInfo(&sspSlotInfo)
	logger.Log.Info().Msgf("log %v %v %v",
		sspSlotInfo.Id, sspSlotInfo.OsType, sspSlotInfo.AdTypeId)
}

func parseInt64(s string) int64 {
	id, _ := strconv.ParseInt(s, 10, 64)
	return id
}

type KfHandler struct {
	SspSlotId  int64
	DspSlotIds []int64
}

func (this *LoopDspEventMaps) MatchBudgetHandler(request *BidRequest) *KfHandler {
	// 使用 this 而不是全局变量，避免锁嵌套
	this.RwMutex.RLock()
	defer this.RwMutex.RUnlock()

	// 直接访问 this.DspLaunchMaps，避免重复获取锁
	dspLaunchs := make([]*DspLaunch, 0)
	for _, launch := range this.DspLaunchMaps {
		if launch.SspSlotId == request.SlotId {
			dspLaunchs = append(dspLaunchs, launch)
		}
	}

	if len(dspLaunchs) == 0 {
		return nil
	}
	// 创建一个 权重搜集器，将流量分配到这个容器中
	launchSnapshot := make([]DspLaunch, 0, len(dspLaunchs))
	// 权重值
	weightCandidates := make([]DspLaunch, 0, len(launchSnapshot))
	seen := make(map[int64]struct{})

	for _, dsplaunch := range dspLaunchs {
		if _, ok := seen[int64(dsplaunch.Indexs)]; ok {
			continue
		}
		seen[int64(dsplaunch.Indexs)] = struct{}{}
		launchSnapshot = append(launchSnapshot, *dsplaunch)
	}
	for _, weightLaunch := range launchSnapshot {
		weightCandidates = append(weightCandidates, DspLaunch{
			SspSlotId:     weightLaunch.SspSlotId,
			Indexs:        weightLaunch.Indexs,
			TrafficWeight: weightLaunch.TrafficWeight,
		})
	}

	index, sspSlotId := returnDspSlotIdWeightDropOut(weightCandidates)
	if sspSlotId == 0 {
		logger.ErrorLog.Warn().Msgf("Weight selection failed, SspSlotId=%d", request.SlotId)
		return nil
	}
	var DspSlotIds []int64
	// 直接使用已获取的 dspLaunchs，避免重复查询
	for _, dspLaunchOn := range dspLaunchs {
		if dspLaunchOn.Indexs == index {
			DspSlotIds = append(DspSlotIds, dspLaunchOn.DspSlotId)
		}
	}

	return &KfHandler{
		SspSlotId:  sspSlotId,
		DspSlotIds: DspSlotIds,
	}
}

func returnDspSlotIdWeightDropOut(dspLaunchArr []DspLaunch) (int, int64) {
	if len(dspLaunchArr) == 1 {
		return dspLaunchArr[0].Indexs, dspLaunchArr[0].SspSlotId
	}

	launch, b := weightDropOut(dspLaunchArr, func(launch DspLaunch) int64 {
		return int64(launch.TrafficWeight)
	})

	if b {
		return launch.Indexs, launch.SspSlotId
	}
	return 0, 0
}

func weightDropOut[T any](ps []T, weightFunc func(T) int64) (T, bool) {
	var zero T
	if len(ps) == 0 {
		return zero, false
	}

	var total int64
	for i := range ps {
		w := weightFunc(ps[i])
		if w > 0 {
			total += w
		}
	}

	if total <= 0 {
		return zero, false
	}

	r := rand.Int64N(total)

	for i := range ps {
		w := weightFunc(ps[i])
		if w <= 0 {
			continue
		}
		if r < w {
			return ps[i], true
		}
		r -= w
	}

	return zero, false
}
