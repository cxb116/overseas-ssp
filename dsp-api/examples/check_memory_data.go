//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MemoryStats struct {
	DspSlotInfoCount int      `json:"dsp_slot_info_count"`
	DspCompanyCount  int      `json:"dsp_company_count"`
	DspLaunchCount   int      `json:"dsp_launch_count"`
	SspSlotInfoCount int      `json:"ssp_slot_info_count"`
	DspLaunches      []Launch `json:"dsp_launches"`
}

type Launch struct {
	Id          int64  `json:"id"`
	SspSlotId   int64  `json:"ssp_slot_id"`
	DspSlotId   int64  `json:"dsp_slot_id"`
	Weight      int    `json:"weight"`
	FloorPrice  int    `json:"floor_price"`
}

func main() {
	fmt.Println("========================================")
	fmt.Println("       检查 dspEngine 内存数据")
	fmt.Println("========================================")
	fmt.Println()

	// 请求内存统计接口
	url := "http://localhost:9990/dsp/memory/stats"
	fmt.Printf("🔍 请求 URL: %s\n", url)
	fmt.Println()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ 请求失败: %v\n", err)
		fmt.Println("\n请确保：")
		fmt.Println("1. dspEngine 正在运行")
		fmt.Println("2. 端口 9990 可访问")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ 读取响应失败: %v\n", err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Printf("❌ HTTP 状态码: %d\n", resp.StatusCode)
		fmt.Printf("响应内容: %s\n", string(body))
		return
	}

	var stats MemoryStats
	if err := json.Unmarshal(body, &stats); err != nil {
		fmt.Printf("❌ 解析 JSON 失败: %v\n", err)
		fmt.Printf("原始响应: %s\n", string(body))
		return
	}

	// 打印统计信息
	fmt.Println("📊 内存数据统计")
	fmt.Println("========================================")
	fmt.Printf("DspSlotInfo:  %d 条\n", stats.DspSlotInfoCount)
	fmt.Printf("DspCompany:   %d 条\n", stats.DspCompanyCount)
	fmt.Printf("DspLaunch:    %d 条\n", stats.DspLaunchCount)
	fmt.Printf("SspSlotInfo:  %d 条\n", stats.SspSlotInfoCount)
	fmt.Println()

	// 打印所有 DspLaunch
	fmt.Println("📋 DspLaunch 详情")
	fmt.Println("========================================")
	if len(stats.DspLaunches) == 0 {
		fmt.Println("⚠️ 没有任何 DspLaunch 数据")
	} else {
		for _, launch := range stats.DspLaunches {
			fmt.Printf("ID: %d, SspSlotId: %d, DspSlotId: %d, Weight: %d, FloorPrice: %d\n",
				launch.Id, launch.SspSlotId, launch.DspSlotId, launch.Weight, launch.FloorPrice)
		}
	}

	fmt.Println("========================================")
	fmt.Println("✅ 检查完成")
	fmt.Println()

	// 查找特定 ID
	var targetID int64 = 112
	fmt.Printf("🔍 查找 ID=%d 的 DspLaunch...\n", targetID)
	found := false
	for _, launch := range stats.DspLaunches {
		if launch.Id == targetID {
			fmt.Printf("✅ 找到 ID=%d 的记录\n", targetID)
			fmt.Printf("   SspSlotId: %d, DspSlotId: %d\n", launch.SspSlotId, launch.DspSlotId)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("❌ 未找到 ID=%d 的记录（可能已被删除）\n", targetID)
	}
}
