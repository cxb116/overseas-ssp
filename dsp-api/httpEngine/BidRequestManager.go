package httpEngine

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cxb116/DSP/gnethttp"
	"github.com/cxb116/DSP/impl"
	"github.com/cxb116/DSP/internal/logger"
)

// BidRequestManager 媒体请求入口
func BidRequestManager(w http.ResponseWriter, req *http.Request) {
	var reader io.Reader = req.Body
	if req.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(req.Body)
		if err != nil {
			http.Error(w, "invalid gzip request body", http.StatusBadRequest)
			return
		}
		defer gz.Close()
		reader = gz
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, "read body failed", http.StatusBadRequest)
		return
	}

	reqCtx := impl.GetRequestContext()
	defer impl.PutRequestContext(reqCtx)

	if err := json.Unmarshal(body, reqCtx.BidRequest); err != nil {
		http.Error(w, "invalid json body:", http.StatusBadRequest)
		logger.Log.Info().Msgf("bidRequest json unmarshal error: %v", err)
		return
	}

	var res *impl.BidResponse
	if !impl.GlobalWorkerChannelHandler.SendRequestToTaskQueue(reqCtx) {
		res = &impl.BidResponse{
			Res:     0,
			ResV:    "",
			Message: "",
		}
	} else {
		res = waitForWinner(*reqCtx)
	}

	if err1 := json.NewEncoder(w).Encode(res); err1 != nil {
		logger.ErrorLog.Error().Msgf("json encode failed: %v", err1)
	}
}

func waitForWinner(reqCtx impl.RequestContext) *impl.BidResponse {
	select {
	case response := <-reqCtx.ResponseChan:
		return response
	case <-reqCtx.Context.Done():
		return &impl.BidResponse{
			Res:     0,
			ResV:    "",
			Message: "",
		}
	}
}

func BidRequestOverseasSsp(req *gnethttp.Request) *gnethttp.Response {

	var bidReq impl.BidRequest
	if err := json.Unmarshal(req.Body, &bidReq); err != nil {
		logger.Log.Info().Msgf("bidRequest json decode error: %v", err)
		return gnethttp.BadRequest("invalid json body")
	}

	fmt.Printf("bidRequest json decode: %v\n", bidReq)
	res := impl.BidRequestHandleDirect(bidReq)
	// 使用 bidReq...
	return gnethttp.JSON(res)
}
