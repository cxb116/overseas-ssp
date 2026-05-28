package httpEngine

import (
	"context"
	_ "net/http/pprof"

	"github.com/cxb116/DSP/document"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/gnethttp"
	"github.com/cxb116/DSP/impl"
	"github.com/cxb116/DSP/internal/logger"
)

//type Engine struct {
//	EngineHttpServer *http.Server
//	EngineContext    context.Context
//}
//
//// 缓存
//
//func newEngineHttpServer() *http.Server {
//	logger.Log.Info().Msgf("Starting HTTP Server,启动端口: %s", global.EngineConfig.Port)
//	return &http.Server{
//		Addr:           global.EngineConfig.Port,
//		ReadTimeout:    1000 * time.Millisecond, // 读取请求超时
//		WriteTimeout:   2 * time.Second,         // 写入响应超时
//		IdleTimeout:    30 * time.Second,        // 空闲连接超时
//		MaxHeaderBytes: 1 << 20,                 // 最大header大小 1MB
//	}
//}
//

//func ServerEngine() {
//	engine := newEngineWithConfig()
//	engine.EngineWithEtcdInit()
//
//	// 启动 pprof
//	go func() {
//		if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
//			log.Println("[pprof] failed:", err)
//		}
//	}()
//
//	muxHttp := http.NewServeMux()
//	//muxHttp.Handle("/api/adx", RequestTimingMiddleware(GzipResHandler(http.HandlerFunc(BidRequestManager))))
//	muxHttp.Handle("/api/adx", http.HandlerFunc(BidRequestManager))
//	muxHttp.Handle("/api/adx/healthy", http.HandlerFunc(healthy))
//
//	//muxHttp.Handle("/api/adx/v1", GzipResHandler(http.HandlerFunc(BidRequestPost)))
//	muxHttp.Handle("/api/adx/v1", RequestTimingMiddleware(http.HandlerFunc(BidRequestPost)))
//	engine.EngineHttpServer.Handler = muxHttp
//	if err := engine.EngineHttpServer.ListenAndServe(); err != nil {
//		logger.Log.Info().Msgf("服务器异常 err: %v", err)
//	}
//
//}

type Engine struct {
	EngineHttpServer *gnethttp.Server
	EngineContext    context.Context
}

func newEngineWithConfig() *Engine {
	return &Engine{
		EngineHttpServer: newEngineHttpServer(),
		EngineContext:    context.Background(),
	}
}

func newEngineHttpServer() *gnethttp.Server {
	logger.Log.Info().Msgf("Starting HTTP Server,启动端口: %s", global.EngineConfig.Port)
	return &gnethttp.Server{}
}

func (engine *Engine) EngineWithEtcdInit() {

	//impl.GlobalCacheResponseMap = impl.NewCacheResponseMap()
	//impl.DspSlotInitMapsHandler = impl.NewDspSlotMaps()
	impl.DspSlotGlobalData = impl.NewDspSlotManager()
	// 初始化协程开启数据加载
	go impl.DspSlotGlobalData.RegLoopSubscribeMessage(engine.EngineContext)
	// 加载工作池初始化
	impl.InitGlobalWorkerChannelHandler(global.EngineConfig.WorkerSize, global.EngineConfig.TaskQueue)
	// 初始化对接文档
	document.DocumentManager()

}

func GnetEngine() {
	router := gnethttp.NewRouter()
	engine := newEngineWithConfig()
	engine.EngineWithEtcdInit()

	// 启动 pprof
	//go func() {
	//	if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
	//		log.Println("[pprof] failed:", err)
	//	}
	//}()

	router.POST("/api/ssp/v2", func(req *gnethttp.Request) *gnethttp.Response {
		return BidRequestOverseasSsp(req)
	})
	// 健康检查
	router.POST("/api/ssp/v2/healthy", func(req *gnethttp.Request) *gnethttp.Response {
		return healthy(req)
	})

	gnethttp.ListenAndServe(":6666", router)
}

func healthy(req *gnethttp.Request) *gnethttp.Response {
	return &gnethttp.Response{
		Code: 200,
	}
}
