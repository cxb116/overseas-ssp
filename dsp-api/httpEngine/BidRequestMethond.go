package httpEngine

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/cxb116/DSP/impl"
	"github.com/cxb116/DSP/internal/logger"
)

//var adxV1IngressLimiter = make(chan struct{}, 3000)

const adxV1MaxBodyBytes = 1 << 20

// BidRequestPost 媒体请求入口
func BidRequestPost(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 限制请求体太大，不能超过1MB
	req.Body = http.MaxBytesReader(w, req.Body, adxV1MaxBodyBytes)
	defer req.Body.Close()

	var reader io.Reader = req.Body
	//if req.Header.Get("Content-Encoding") == "gzip" {
	//	gz, err := gzip.NewReader(req.Body)
	//	if err != nil {
	//		http.Error(w, "invalid gzip request body", http.StatusBadRequest)
	//		return
	//	}
	//	defer gz.Close()
	//	reader = gz
	//}

	timeout := 1100 * time.Millisecond
	//if global.EngineConfig.RequestLink > 0 {
	//	timeout = global.EngineConfig.RequestLink
	//}
	//now := time.Now()
	//if parentDeadline, ok := req.Context().Deadline(); ok {
	//	logger.Log.Info().Msgf(
	//		"当前的context 时间 deadline=%s remaining=%s now=%s",
	//		parentDeadline.Format(time.RFC3339Nano),
	//		time.Until(parentDeadline).String(),
	//		now.Format(time.RFC3339Nano),
	//	)
	//} else {
	//	logger.Log.Info().Msgf(
	//		"当前的context 时间 deadline now=%s",
	//		now.Format(time.RFC3339Nano),
	//	)
	//}
	//logger.Log.Info().Msgf("当前的context 时间 timeout=%s", timeout.String())

	_, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()
	//if childDeadline, ok := reqCtx.Deadline(); ok {
	//	logger.Log.Info().Msgf(
	//		"请求的超时的时间 deadline=%s remaining=%s",
	//		childDeadline.Format(time.RFC3339Nano),
	//		time.Until(childDeadline).String(),
	//	)
	//}

	bidReq := impl.BidRequest{}

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&bidReq); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		logger.Log.Info().Msgf("bidRequest json decode error: %v", err)
		return
	}

	//start := time.Now()
	//res := impl.HandleBidRequestDirect(reqCtx, &bidReq, timeout)
	//requestOut := time.Now()
	//if err := reqCtx.Err(); err != nil {
	//	logger.Log.Warn().Msgf(
	//		"[adx-v1] request 响应时间 err=%v elapsed=%s slotId=%d res=%d",
	//		err,
	//		time.Since(start).String(),
	//		bidReq.SlotId,
	//		res.Res,
	//	)
	//} else {
	//	logger.Log.Info().Msgf(
	//		"[adx-v1] request 报错是响应时间 elapsed=%s slotId=%d res=%d",
	//		time.Since(start).String(),
	//		bidReq.SlotId,
	//		res.Res,
	//	)
	//}
	//writeBidResponse(w, &res)
	//logger.Log.Info().Msgf(
	//	"http request breakdown, path=%s slotID=%d res=%d http_total=%s request_in=%s request_out=%s remote=%s",
	//	req.URL.Path,
	//	bidReq.SlotId,
	//	res.Res,
	//	requestOut.Sub(requestIn).String(),
	//	requestIn.Format(time.RFC3339Nano),
	//	requestOut.Format(time.RFC3339Nano),
	//	req.RemoteAddr,
	//)
}

func writeBidResponse(w http.ResponseWriter, res *impl.BidResponse) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.ErrorLog.Error().Msgf("json encode failed: %v", err)
	}
}
