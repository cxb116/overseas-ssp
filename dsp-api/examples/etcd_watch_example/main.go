package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// DspCompany 预算公司结构
type DspCompany struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	DspCode  string `json:"dsp_code"`
	URL      string `json:"url"`
	Method   string `json:"method"`
	Timeout  int    `json:"timeout"`
}

const (
	ETCD_ENDPOINTS = "117.72.67.186:2379"
	DSP_PREFIX     = "/dsp/key"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("       ETCD 监听示例")
	fmt.Println("========================================")
	fmt.Println()

	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{ETCD_ENDPOINTS},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("❌ 连接 etcd 失败: %v\n", err)
		return
	}
	defer cli.Close()

	fmt.Printf("✅ 成功连接 etcd: %s\n", ETCD_ENDPOINTS)
	fmt.Println()

	// 创建上下文，可以控制程序退出
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动监听 goroutine
	fmt.Println("========================================")
	fmt.Println("       启动 etcd 监听")
	fmt.Println("========================================")
	fmt.Println()

	go watchEtcdKeys(ctx, cli, DSP_PREFIX+"/")

	// 等待监听启动
	time.Sleep(1 * time.Second)

	// 演示添加、修改、删除数据
	demonstrateEtcdOperations(cli)

	// 让程序运行一段时间，观察监听输出
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("⏳ 持续监听中... (按 Ctrl+C 退出)")
	fmt.Println("========================================")

	// 持续运行，直到用户中断
	select {}
}

// watchEtcdKeys 监听 etcd key 变化
func watchEtcdKeys(ctx context.Context, cli *clientv3.Client, prefix string) {
	fmt.Printf("🔍 开始监听前缀: %s\n", prefix)
	fmt.Println()

	// 创建 watch 监听器
	watchChan := cli.Watch(ctx, prefix, clientv3.WithPrefix())

	for watchResp := range watchChan {
		if watchResp.Err() != nil {
			fmt.Printf("❌ Watch 错误: %v\n", watchResp.Err())
			continue
		}

		for _, event := range watchResp.Events {
			printEtcdEvent(event)
		}
	}
}

// printEtcdEvent 打印 etcd 事件
func printEtcdEvent(event *clientv3.Event) {
	fmt.Println("========================================")
	fmt.Println("📨 收到 etcd 事件")
	fmt.Println("========================================")

	// 打印基本信息
	fmt.Printf("Key:     %s\n", string(event.Kv.Key))
	fmt.Printf("Action:  %s\n", event.Type)
	fmt.Printf("Version: %d\n", event.Kv.Version)

	// 根据操作类型打印不同信息
	switch event.Type {
	case clientv3.EventTypePut:
		if event.IsModify() {
			fmt.Println("Type:    修改 (MODIFY)")
		} else {
			fmt.Println("Type:    添加 (ADD)")
		}
		fmt.Printf("Value:   %s\n", string(event.Kv.Value))

		// 尝试解析为 JSON
		var data map[string]interface{}
		if err := json.Unmarshal(event.Kv.Value, &data); err == nil {
			fmt.Println("解析 JSON 成功:")
			prettyJSON, _ := json.MarshalIndent(data, "  ", "  ")
			fmt.Printf("  %s\n", string(prettyJSON))
		}

	case clientv3.EventTypeDelete:
		fmt.Println("Type:    删除 (DELETE)")
		fmt.Printf("旧值:    %s\n", string(event.Kv.Value))
	}

	fmt.Println("========================================")
	fmt.Println()
}

// demonstrateEtcdOperations 演示 etcd 操作
func demonstrateEtcdOperations(cli *clientv3.Client) {
	fmt.Println("========================================")
	fmt.Println("       演示 ETCD 操作")
	fmt.Println("========================================")
	fmt.Println()

	// 1. 添加预算公司
	fmt.Println("【操作 1】添加预算公司")
	fmt.Println("--------------------")

	company1 := DspCompany{
		ID:      1001,
		Name:    "测试公司1001",
		DspCode: "TEST1001",
		URL:     "https://test1001.com",
		Method:  "POST",
		Timeout: 500,
	}

	key1 := fmt.Sprintf("%s/company/add/%d", DSP_PREFIX, company1.ID)
	value1, _ := json.Marshal(company1)

	fmt.Printf("Key:   %s\n", key1)
	fmt.Printf("Value: %s\n", string(value1))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cli.Put(ctx, key1, string(value1))
	if err != nil {
		fmt.Printf("❌ 添加失败: %v\n", err)
	} else {
		fmt.Println("✅ 添加成功")
	}
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 2. 添加预算广告位
	fmt.Println("【操作 2】添加预算广告位")
	fmt.Println("--------------------")

	slot := map[string]interface{}{
		"id":               1001,
		"dsp_company_id":   1001,
		"dsp_slot_code":    "TEST_SLOT_1001",
		"dsp_app_key":      "test_key_1001",
		"dsp_app_secret":   "test_secret_1001",
		"price_encrypt_key": "encrypt_1001",
		"dsp_pay_type":     2,
		"floor_price":      100,
	}

	key2 := fmt.Sprintf("%s/dsp/add/%d", DSP_PREFIX, 1001)
	value2, _ := json.Marshal(slot)

	fmt.Printf("Key:   %s\n", key2)
	fmt.Printf("Value: %s\n", string(value2))

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	_, err = cli.Put(ctx2, key2, string(value2))
	if err != nil {
		fmt.Printf("❌ 添加失败: %v\n", err)
	} else {
		fmt.Println("✅ 添加成功")
	}
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 3. 修改预算公司
	fmt.Println("【操作 3】修改预算公司")
	fmt.Println("--------------------")

	company1.Name = "修改后的公司1001"
	company1.Timeout = 800
	value3, _ := json.Marshal(company1)

	fmt.Printf("Key:   %s\n", key1)
	fmt.Printf("Value: %s\n", string(value3))

	ctx3, cancel3 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel3()

	_, err = cli.Put(ctx3, key1, string(value3))
	if err != nil {
		fmt.Printf("❌ 修改失败: %v\n", err)
	} else {
		fmt.Println("✅ 修改成功")
	}
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 4. 查询所有数据
	fmt.Println("【操作 4】查询所有 DSP 数据")
	fmt.Println("--------------------")

	ctx4, cancel4 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel4()

	resp, err := cli.Get(ctx4, DSP_PREFIX+"/", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("❌ 查询失败: %v\n", err)
	} else {
		fmt.Printf("找到 %d 条数据:\n", resp.Count)
		for _, kv := range resp.Kvs {
			fmt.Printf("  - %s\n", string(kv.Key))
		}
	}
	fmt.Println()

	time.Sleep(1 * time.Second)

	// 5. 删除预算公司
	fmt.Println("【操作 5】删除预算公司")
	fmt.Println("--------------------")

	fmt.Printf("Key:   %s\n", key1)

	ctx5, cancel5 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel5()

	_, err = cli.Delete(ctx5, key1)
	if err != nil {
		fmt.Printf("❌ 删除失败: %v\n", err)
	} else {
		fmt.Println("✅ 删除成功")
	}
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("       演示完成")
	fmt.Println("========================================")
	fmt.Println()
}
