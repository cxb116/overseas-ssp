package impl

import (
	"net"
	"net/http"
	"time"
)

// ClientPool HTTP客户端池 - 优化超时配置用于请求上游DSP
var ClientPool = &http.Client{
	Timeout: 1000 * time.Millisecond,
	Transport: &http.Transport{
		DisableCompression: false, // 支持 gzip

		// ===== 连接池配置 =====
		MaxIdleConns:        1000,             // 最大空闲连接数（提升）
		MaxIdleConnsPerHost: 500,              // 每个host的最大空闲连接数（提升）
		MaxConnsPerHost:     2000,             // 每个host的最大连接数（提升）
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时时间

		// ===== 连接超时配置 =====
		DialContext: (&net.Dialer{
			Timeout:   1 * time.Second,  // 连接超时：3秒（快速失败）
			KeepAlive: 30 * time.Second, // keep-alive时间
			// 双栈网络优化
			DualStack: true,
		}).DialContext,

		// 强制启用 keep-alive，复用连接
		//DisableKeepAlives: false,

		// ===== 请求超时配置（按阶段） =====
		// 响应头超时：从发送请求到收到响应头的最大时间
		ResponseHeaderTimeout: 2 * time.Second, // 2秒没收到响应头就超时

		// 期望 100 Continue 超时
		ExpectContinueTimeout: 1 * time.Second,

		// TLS握手超时（如果是HTTPS）
		TLSHandshakeTimeout: 3 * time.Second,
	},
}
