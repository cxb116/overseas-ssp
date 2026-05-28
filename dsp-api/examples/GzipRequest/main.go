package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func gzipBody(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	gz.Close()
	return buf.Bytes(), nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		return io.ReadAll(gr)
	}

	return io.ReadAll(resp.Body)
}
func TestDSPRequest(
	url string,
	bidId string,
	gzipRequest bool,
	acceptGzipResp bool,
) error {

	// 构造请求 JSON
	reqJSON := []byte(fmt.Sprintf(`{"bidId":"%s"}`, bidId))

	var body io.Reader

	if gzipRequest {
		zb, err := gzipBody(reqJSON)
		if err != nil {
			return err
		}
		body = bytes.NewReader(zb)
	} else {
		body = bytes.NewReader(reqJSON)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	if gzipRequest {
		req.Header.Set("Content-Encoding", "gzip")
	}
	if acceptGzipResp {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := readResponseBody(resp)
	if err != nil {
		return err
	}

	fmt.Println("====== DSP RESPONSE ======")
	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))
	fmt.Println("Body:", string(respBody))
	fmt.Println("==========================")

	return nil
}
func main() {
	url := "http://127.0.0.1:8888/dsp"

	fmt.Println("\n1 普通请求 + gzip 响应")
	TestDSPRequest(url, "bid-001", false, true)

	fmt.Println("\n2 gzip 请求 + gzip 响应")
	TestDSPRequest(url, "bid-002", true, true)

	fmt.Println("\n3 普通请求 + 普通响应")
	TestDSPRequest(url, "bid-003", false, false)
}
