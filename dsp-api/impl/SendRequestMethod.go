package impl

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cxb116/DSP/internal/logger"
)

func SendOnPbRequest(msg []byte, url string, httpClient *http.Client) ([]byte, bool, bool) {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		logger.ErrorLog.Error().Msgf("SendRequestOnPb NewRequest Err: %v, url: %s", err, url)
		return nil, false, false
	}

	request.Header.Set("Content-Type", "application/protobuf")

	start := time.Now()

	resp, err := httpClient.Do(request)
	cost := time.Since(start)

	if err != nil {
		logger.ErrorLog.Error().Msgf(
			"SendRequestOnPb POST Err: %v, cost: %v, url: %s",
			err, cost, url,
		)
		return nil, false, false
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLog.Error().Msgf("SendRequestOnPb Read Body Err: %v, url: %s", err, url)
		return nil, false, false
	}

	//  增加：HTTP状态码判断（非常重要）
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBytes, false, false
	}

	//  慢请求标记（>1秒）
	if cost > time.Second {
		return respBytes, true, true
	}

	return respBytes, true, false
}

//	func SendOnJSONRequest(msg string, url string, httpClient *http.Client) ([]byte, bool) {
//		return SendOnJSONRequestWithContext(context.Background(), msg, url, httpClient)
//	}
//
//	body, 是否成功 true, 是否超时 false
func SendOnJSONRequestWithContext(
	ctx context.Context,
	msg []byte,
	url string,
	httpClient *http.Client,
) ([]byte, bool, bool) {
	//req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(msg))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, false, false
	}

	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	resp, err := httpClient.Do(req)
	cost := time.Since(start)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return nil, false, true
		}
		logger.ErrorLog.Warn().Msgf("SendOnJSONRequestWithContext POST err: %v, cost: %v, url: %s", err, cost, url)
		return nil, false, true
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLog.Warn().Msgf("SendOnJSONRequestWithContext read body err: %v, cost: %v, url: %s", err, cost, url)
		return nil, false, false
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.ErrorLog.Warn().Msgf("SendOnJSONRequestWithContext status err: status=%d, cost: %v, url: %s", resp.StatusCode, cost, url)
		return body, false, false
	}

	// 可选：慢请求标记（但不是超时）
	if cost > time.Second {
		return body, true, true
	}

	return body, true, false
}

//// SendOnJSONRequestWithContext 支持 context 的 JSON 请求，可以检测超时
//// 返回值: (数据, 是否成功, 是否超时)
//func SendOnJSONRequestWithContext(ctx context.Context, msg []byte, url string, httpClient *http.Client) ([]byte, bool, bool) {
//	// 1. 构造请求
//	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(msg))
//	if err != nil {
//		logger.ErrorLog.Error().Msgf("SendJSONRequest NewRequest Err: %v, url: %s", err, url)
//		return nil, false, false
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	// 2. 记录请求时间
//	startTime := time.Now()
//	// 3. 发送请求
//	resp, err := httpClient.Do(req)
//	since := time.Since(startTime).Seconds()
//	if err != nil {
//		// 3.4 其他错误（网络错误、连接失败等）
//		logger.ErrorLog.Error().Msgf("SendJSONRequest POST Err: %v, cost: %v, url: %s, body: %s", err, time.Since(startTime), url, truncateString(string(msg), 200))
//		return nil, false, false
//	}
//	defer resp.Body.Close()
//
//	// 4. 读取响应
//	respBytes, err := io.ReadAll(resp.Body)
//	if err != nil {
//		logger.ErrorLog.Error().Msgf("服务器器:Read Body Err: %v, url: %s", err, url)
//		return nil, false, false
//	}
//	if since > 1.0 {
//		return respBytes, true, true
//	}
//
//	// 5. 成功返回
//	return respBytes, true, false
//}

func SendOnGZIPJSONRequestWithContext(ctx context.Context, msg []byte, url string, httpClient *http.Client) ([]byte, bool, bool) {

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(msg))
	if err != nil {
		logger.ErrorLog.Error().Msgf("NewRequest Err: %v, url: %s", err, url)
		return nil, false, false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Encoding", "gzip")

	startTime := time.Now()

	resp, err := httpClient.Do(req)
	cost := time.Since(startTime)

	// ---------- 1. 错误处理 ----------
	if err != nil {
		// 判断是否是超时
		if ctx.Err() == context.DeadlineExceeded {
			return nil, false, true
		}

		logger.ErrorLog.Error().Msgf("POST Err: %v, cost: %v, url: %s", err, cost, url)
		return nil, false, false
	}
	defer resp.Body.Close()

	// ---------- 2. HTTP 状态码 ----------
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logger.ErrorLog.Error().Msgf("HTTP Err: %d, cost: %v, body: %s", resp.StatusCode, cost, truncateString(string(body), 200))
		return nil, false, false
	}

	// ---------- 3. gzip 处理 ----------
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			logger.ErrorLog.Error().Msgf("gzip NewReader Err: %v", err)
			return nil, false, false
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	respBytes, err := io.ReadAll(reader)
	if err != nil {
		logger.ErrorLog.Error().Msgf("Read Body Err: %v", err)
		return nil, false, false
	}

	// ---------- 4. 延迟标记 ----------
	isSlow := cost > 1000*time.Millisecond
	return respBytes, true, isSlow
}

// -----------------------------
// 2. Content-Encoding: gzip
// -----------------------------
func SendOnJSONCONTENTGZIPRequest(msg string, url string, httpClient *http.Client) ([]byte, bool) {
	// ---------- 1. gzip 压缩 ----------
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	if _, err := gzipWriter.Write([]byte(msg)); err != nil {
		logger.ErrorLog.Error().Msgf("gzip Write Err: %v, url: %s", err, url)
		return nil, false
	}

	if err := gzipWriter.Close(); err != nil {
		logger.ErrorLog.Error().Msgf("gzip Close Err: %v, url: %s", err, url)
		return nil, false
	}

	// ---------- 2. 构建请求 ----------
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		logger.ErrorLog.Error().Msgf("NewRequest Err: %v, url: %s", err, url)
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip") // ⭐必须
	req.Header.Set("Accept-Encoding", "gzip")

	// ---------- 3. 发送请求 ----------
	startTime := time.Now()
	resp, err := httpClient.Do(req)
	if err != nil {
		logger.ErrorLog.Error().Msgf("POST Err: %v, cost: %v, body len: %d",
			err, time.Since(startTime), len(msg))
		return nil, false
	}
	defer resp.Body.Close()

	// ---------- 4. 检查状态码 ----------
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logger.ErrorLog.Error().Msgf("HTTP Status Err: %d, body: %s", resp.StatusCode, string(body))
		return nil, false
	}

	// ---------- 5. 处理 gzip 响应 ----------
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			logger.ErrorLog.Error().Msgf("gzip NewReader Err: %v, url: %s", err, url)
			return nil, false
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	respBytes, err := io.ReadAll(reader)
	if err != nil {
		logger.ErrorLog.Error().Msgf("Read Body Err: %v, url: %s", err, url)
		return nil, false
	}

	return respBytes, true
}

// truncateString 截断字符串用于日志输出
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...(truncated)"
}

func SendOnJSONACCEPTGZIPRequest(msg string, url string, httpClient *http.Client) ([]byte, bool, bool) {
	// 1. 压缩请求体
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write([]byte(msg))
	if err != nil {
		return nil, false, false
	}
	gzipWriter.Close()

	// 2. 构建请求
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return nil, false, false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip") //  请求体是 gzip
	req.Header.Set("Accept-Encoding", "gzip")  //  允许响应 gzip

	startTime := time.Now()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, false, false
	}
	defer resp.Body.Close()

	// 3. 状态码检查（必须）
	if resp.StatusCode != http.StatusOK {
		return nil, false, false
	}

	// 4. 不要手动解 gzip（默认 client 会自动解）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, false
	}

	cost := time.Since(startTime)

	return body, true, cost > time.Second
}

func SendOnPbSRequest(msg []byte, url string, httpClient *http.Client, headers ...string) (bool, []byte) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		logger.ErrorLog.Error().Msgf("SendOnPbSRequest POST Err1: %v, url: %s", err, url)
		return false, nil
	}
	request.Header.Set("Content-Type", "application/x-protobuf")

	for _, head := range headers {
		tmpArr := strings.SplitN(head, ":", 2)
		if len(tmpArr) != 2 {
			logger.ErrorLog.Warn().Msgf("SendOnPbSRequest POST headers is wrong: %v", headers)
			return false, nil
		}
		key := strings.TrimSpace(tmpArr[0])
		value := strings.TrimSpace(tmpArr[1])
		request.Header.Set(key, value)
	}

	t := time.Now()
	resp, err := httpClient.Do(request) //发送请求
	if err != nil {
		logger.ErrorLog.Error().Msgf("SendOnPbSRequest POST Err2: %v, cost: %v, body: %v", err, time.Since(t), msg)
		return false, nil
	}
	defer resp.Body.Close()

	data, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		logger.ErrorLog.Error().Msgf("SendOnPbSRequest POST Err5: %v, url: %s", err1, url)
		return false, nil
	}
	return true, data
}
