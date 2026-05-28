# ETCD 监听示例

这是一个完整的 etcd 监听示例，演示如何：
1. 连接 etcd
2. 监听 key 的变化（添加、修改、删除）
3. 向 etcd 写入数据
4. 接收并打印事件

## 运行方法

### Windows
```bash
cd examples\etcd_watch_example
run.bat
```

### 手动运行
```bash
cd examples\etcd_watch_example
go run main.go
```

## 程序说明

### 功能演示

1. **连接 etcd**
   - 连接到 `117.72.67.186:2379`
   - 创建客户端

2. **启动监听**
   - 监听前缀：`/dsp/key/`
   - 自动接收所有变化事件

3. **演示操作**
   - 添加预算公司 (ID=1001)
   - 添加预算广告位 (ID=1001)
   - 修改预算公司
   - 查询所有数据
   - 删除预算公司

4. **持续监听**
   - 程序会持续运行，监听新的变化
   - 按 `Ctrl+C` 退出

## 输出示例

```
========================================
       ETCD 监听示例
========================================

✅ 成功连接 etcd: 117.72.67.186:2379

========================================
       启动 etcd 监听
========================================

🔍 开始监听前缀: /dsp/key/

【操作 1】添加预算公司
--------------------
Key:   /dsp/key/company/add/1001
Value: {"id":1001,"name":"测试公司1001",...}
✅ 添加成功

========================================
📨 收到 etcd 事件
========================================
Key:     /dsp/key/company/add/1001
Action:  PUT
Type:    添加 (ADD)
Value:   {"id":1001,"name":"测试公司1001",...}
解析 JSON 成功:
  {
    "id": 1001,
    "name": "测试公司1001",
    "dsp_code": "TEST1001",
    ...
  }
========================================

【操作 2】添加预算广告位
...
```

## 代码结构

### main 函数
- 创建 etcd 客户端
- 启动监听 goroutine
- 演示各种操作
- 持续运行

### watchEtcdKeys
- 监听指定前缀的所有 key
- 使用 `clientv3.WithPrefix()` 监听子 key
- 实时接收事件

### printEtcdEvent
- 打印事件详情
- 解析 JSON 数据
- 区分 PUT/DELETE 操作

### demonstrateEtcdOperations
- 演示完整的 CRUD 操作
- 添加、修改、查询、删除

## 关键 API

### 创建客户端
```go
cli, err := clientv3.New(clientv3.Config{
    Endpoints:   []string{"117.72.67.186:2379"},
    DialTimeout: 5 * time.Second,
})
```

### 监听 Key
```go
watchChan := cli.Watch(ctx, "/dsp/key/", clientv3.WithPrefix())
for event := range watchChan {
    // 处理事件
}
```

### 添加/修改数据
```go
cli.Put(ctx, key, value)
```

### 删除数据
```go
cli.Delete(ctx, key)
```

### 查询数据
```go
cli.Get(ctx, prefix, clientv3.WithPrefix())
```

## 注意事项

1. **etcd 连接**：确保 etcd 服务可访问
2. **权限**：需要有读写权限
3. **网络**：确保网络连接正常
4. **退出**：按 `Ctrl+C` 退出程序

## 扩展使用

### 修改监听前缀
```go
const DSP_PREFIX = "/your/custom/prefix"
```

### 修改 etcd 地址
```go
const ETCD_ENDPOINTS = "your-etcd-server:2379"
```

### 添加自定义数据处理
在 `printEtcdEvent` 函数中添加你的处理逻辑。

## 相关文件

- DSP 主程序：`impl/LoopSubscribeMessage.go`
- etcd 文档：`docs/etcd_api_guide.md`
- 日志系统：`docs/logger_usage.md`

---

**作者：** DSP Team
**日期：** 2026-02-19
