# 定时请求发送工具

这是一个定时发送广告请求的示例程序，用于测试 DSP 广告系统。

## 功能特性

- ✅ 定时发送广告请求（默认 5 秒间隔）
- ✅ 完整的 BidRequest JSON 结构
- ✅ 自动生成唯一请求 ID
- ✅ 详细的请求/响应日志
- ✅ 可配置的目标地址和时间间隔
- ✅ Ctrl+C 优雅退出

## 使用方法

### 1. 修改配置

编辑 `main.go` 文件中的配置：

```go
serverURL := "http://localhost:8080/api/bid" // 修改为你的实际服务地址
interval := 5 * time.Second                 // 发送间隔：5秒
```

### 2. 运行程序

```bash
# 进入目录
cd examples/periodic_request

# 运行程序
go run main.go
```

### 3. 停止程序

按 `Ctrl+C` 停止程序

## 输出示例

```
========================================
     定时请求发送工具
========================================
目标地址: http://localhost:8080/api/bid
发送间隔: 5s
按 Ctrl+C 停止程序
========================================

========== 发送请求 ==========
URL: http://localhost:8080/api/bid
Request Body:
{"id":"req-20260305-0001","imp":{...},"app":{...},...}

========== 响应结果 ==========
Status Code: 200
Response Body:
{"code":0,"msg":"success","data":{...}}
==============================

✅ 第 1 次请求发送成功
```

## 自定义请求

你可以修改 `createSampleRequest()` 函数来自定义请求内容：

```go
func createSampleRequest() *BidRequest {
    return &BidRequest{
        ID: "req-20260305-0001",
        Imp: Imp{
            ID:       19,        // 修改广告位ID
            W:        375,       // 修改宽度
            H:        300,       // 修改高度
            AdType:   1,         // 修改广告类型
            Position: 4,         // 修改位置
            BidFloor: 1,         // 修改底价
        },
        // ... 其他字段
    }
}
```

## 请求参数说明

| 字段 | 说明 | 示例值 |
|------|------|--------|
| id | 请求ID | "req-20260305-0001" |
| imp.id | 广告位ID | 19 |
| imp.w | 广告位宽度 | 375 |
| imp.h | 广告位高度 | 300 |
| imp.ad_type | 广告类型 | 1: 开屏, 2: 横幅, 3: 插屏 |
| imp.bid_floor | 底价 | 1 (分) |
| device.os | 操作系统 | 1: Android, 2: iOS |
| device.oaid | OAID | 设备标识 |
| device.network.net_type | 网络类型 | 1: WiFi, 2: 2G, 3: 3G, 4: 4G, 5: 5G |

## 常见问题

### Q: 如何修改发送间隔？
A: 修改 `interval` 变量：
```go
interval := 10 * time.Second  // 改为 10 秒
```

### Q: 如何只发送一次请求？
A: 注释掉定时器循环：
```go
// ticker := time.NewTicker(interval)
// for range ticker.C { ... }
```

### Q: 如何模拟多个不同的请求？
A: 创建多个请求模板：
```go
requests := []*BidRequest{
    createRequest1(),
    createRequest2(),
    createRequest3(),
}
for _, req := range requests {
    sendRequest(serverURL, req)
}
```

## 高级用法

### 带上下文的超时控制

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, "POST", url, body)
```

### 并发发送

```go
for i := 0; i < 10; i++ {
    go func() {
        sendRequest(serverURL, req)
    }()
}
```

### 随机请求间隔

```go
interval := time.Duration(rand.Intn(10)+1) * time.Second // 1-10秒随机
```

## 技术栈

- Go 1.21+
- HTTP Client
- JSON encoding/json
- Ticker 定时器

## 许可证

MIT License
