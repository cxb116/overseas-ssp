package httpEngine

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	gz *gzip.Writer
}

var gzipWriterPool = sync.Pool{
	New: func() interface{} {
		// 使用 BestSpeed 降低 CPU，适合高吞吐接口场景。
		w, _ := gzip.NewWriterLevel(io.Discard, gzip.BestSpeed)
		return w
	},
}

func (w *gzipResponseWriter) WriteHeader(statusCode int) {
	// gzip 后长度不可预知，必须删除
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.gz.Write(b)
}

func GzipResHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		//  设置标准响应头
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")

		gz := gzipWriterPool.Get().(*gzip.Writer)
		gz.Reset(w)
		defer func() {
			_ = gz.Close()
			gzipWriterPool.Put(gz)
		}()

		// 包装 ResponseWriter
		gzw := &gzipResponseWriter{
			ResponseWriter: w,
			gz:             gz,
		}

		// 进入真实竞价逻辑
		next.ServeHTTP(gzw, r)
	})
}
