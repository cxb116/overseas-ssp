package httpEngine

import (
	"net/http"
	"time"

	"github.com/cxb116/DSP/internal/logger"
)

type requestTimingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *requestTimingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func RequestTimingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestIn := time.Now()
		rw := &requestTimingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		requestOut := time.Now()
		logger.Log.Info().Msgf(
			"http request timing, method=%s path=%s status=%d request_in=%s request_out=%s elapsed=%s remote=%s",
			r.Method,
			r.URL.Path,
			rw.statusCode,
			requestIn.Format(time.RFC3339Nano),
			requestOut.Format(time.RFC3339Nano),
			requestOut.Sub(requestIn).String(),
			r.RemoteAddr,
		)
	})
}
