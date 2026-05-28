package dsp

//import (
//	"strings"
//
//	"github.com/cxb116/DSP/constant"
//	"github.com/cxb116/DSP/dspAuction"
//	"github.com/cxb116/DSP/impl"
//)
//
//func TestHuanXiDocument(dspCode int64) {
//	impl.DspRegister(dspCode, func(reqContext *impl.RequestContext) impl.DspHandler {
//		return &THuanXiDsp{
//			DspCode:        dspCode,
//			RequestContext: nil,
//			DspSlotInfo:    nil,
//		}
//	})
//}
//func (this *THuanXiDsp) ReturnSelf() impl.DspHandler {
//	return this
//}
//
//type THuanXiDsp struct {
//	DspCode        int64
//	RequestContext *impl.RequestContext
//	DspSlotInfo    *impl.DspSlotInfo
//}
//
//func (this *THuanXiDsp) SetDspRequestContext(reqCtx *impl.RequestContext, dspSlotInfo *impl.DspSlotInfo) {
//	this.RequestContext = reqCtx
//	this.DspSlotInfo = dspSlotInfo
//}
//
//func (this *THuanXiDsp) DspRequestHandler() impl.BidResponse {
//	res := impl.BidResponse{}
//	res.Ad.Act = 2
//	res.Ad.Adw = 360
//	res.Ad.Adh = 720
//	res.Ad.Img = "http://open.bigmodel.cn/finance-center/finance/overview"
//	res.Ad.Img2 = "http://open.bigmodel.cn/finance/overview"
//	res.Ad.Img3 = "http://open.bigmodel.cn/finance/overview/img.html"
//	res.Ad.Title = "质谱AI"
//	res.Ad.Text = "上质谱,用AI"
//	res.Ad.Click = []string{
//		"http://open.bigmodel/trace?data=6b9a476e26a2483150d03e7237e94fca78ffc231938c70010bc7f19245e416c3884c03f3822ae7e253ad6f7bf73e3485f01736fb340e15328f6026a4210868f8e248ecff1b55f602727940b6256f5ffaf22f60adc49bece339f5d60129116f14fcfa6bb33aee2bf3677b62a3c4c74",
//		"http://open.bigmodel/cm?info=MQIzAjE2ODM0ODIzNjICYTgwYWRhNTQtNjI1MS00YzY0LWFlNzctNWQ1YzNhNmUwOWJmAgIyMDAyMjcxAmNvbS5wa2cuZGVtbwLlupTnlKjlkI3np7ACMS4wLjJfMjAxNjAzMjgCMS4wLjICMTc1LjI1LjE4OC4zAgICODY5NDY4MDIwMjkwMTIzAgICNDYwMDMwOTExMDcyNTY1AjAyOjAwOjAwOjAwOjAwOjAwAmEwMmU3NjcyMTA0MDIyMgICMAIxAjYuMAIzAjcyMAIxMDgwAjACSFRDAiBEODE2ZAJ6aAIzAjQ2MDAzAmU0OmY0OmM2OjA2OmJlOjI3AkRPTk9QTy1BUDkCTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDYuMC4xOyBPUFBPIFI5c2sgQnVpbGQvTU1CMjlNOyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzQ2LjAuMjQ5MC43NiBNb2JpbGUgU2FmYXJpLzUzNy4zNgIwAjAC5YyX5Lqs5rW35reA5Yy6AjMwMDE3ODYwODICMTYwAjMyMAI1MAIxLDIsNCwCMAIxAjc2YTA0ZjU3LThlNTMtNGRlZi1iNDZjLTMwMGYyZTJlNWYzYwI0AjMCMAIyAmNvbS51ZWxpdmUuc2hvd3ZpZGVvLmFjdGl2aXR5Amh0dHA6Ly9kMi55b3V5aTUyMC5jb20vYS9jLzUzL1VFbGl2ZS5jb21fNTM1MDMuYXBrAgJzYW55b3UCMAIw",
//	}
//	res.Ad.Imp = []string{
//		"http://open.bigmodel/trace?data=98a6b727a779ba0e83f646074bb58eaa26acd150279ca851725667cd9302af7a8bf822500714377759fe4ccad9fbf4d667c0993ea3ca3bf650ca62b8b017a33b2f396920f832e971c3b7b832078e02d37780c6de656c146fd5b2ec26a4e0d73ba2d69d7d24625ca4af717e027d6896ebebc3547cdd2ef0adda41521db46418564aec91c8300a2580c5dd667e9d8c50921beb92827bd41d7b978541270728d2750f5d08bd0b8c06ed2b41a86bec2f06d930578d7da6ee8f6d0510fb3353eac8362b715c2d8d2bd4447838dae2067a4c3504e395b9c0691db0824781190cf0d7d9b5f3e702dae5b9dfa8923322641b42fac3220855d6a22d5f49494c1fdaf77968e4a439409fa8984b30897b5b2b5566e4b82fc2a26bd6beab73251a99a3923c1c301d9c96120355e51bbdf3f878f8021f83be22aac63d92db",
//		"http://open.bigmodel/vm?info=MQIyAjE2ODM0ODIzNjICYTgwYWRhNTQtNjI1MS00YzY0LWFlNzctNWQ1YzNhNmUwOWJmAgIyMDAyMjcxAmNvbS5wa2cuZGVtbwLlupTnlKjlkI3np7ACMS4wLjJfMjAxNjAzMjgCMS4wLjICMTc1LjI1LjE4OC4zAgICODY5NDY4MDIwMjkwMTIzAgICNDYwMDMwOTExMDcyNTY1AjAyOjAwOjAwOjAwOjAwOjAwAmEwMmU3NjcyMTA0MDIyMgICMAIxAjYuMAIzAjcyMAIxMDgwAjACSFRDAiBEODE2ZAJ6aAIzAjQ2MDAzAmU0OmY0OmM2OjA2OmJlOjI3AkRPTk9QTy1BUDkCTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDYuMC4xOyBPUFBPIFI5c2sgQnVpbGQvTU1CMjlNOyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzQ2LjAuMjQ5MC43NiBNb2JpbGUgU2FmYXJpLzUzNy4zNgIwAjAC5YyX5Lqs5rW35reA5Yy6AjMwMDE3ODYwODICMTYwAjMyMAI1MAIxLDIsNCwCMAIxAjc2YTA0ZjU3LThlNTMtNGRlZi1iNDZjLTMwMGYyZTJlNWYzYwI0AjMCMAIyAmNvbS51ZWxpdmUuc2hvd3ZpZGVvLmFjdGl2aXR5Amh0dHA6Ly9kMi55b3V5aTUyMC5jb20vYS9jLzUzL1VFbGl2ZS5jb21fNTM1MDMuYXBrAgJzYW55b3UCMAIw",
//	}
//	res.Ad.Lpg = "http://open.bigmodel.cn/DownLoad-center/finance/"
//	res.Ad.Adext.Price = 1
//	res.Ad.Lpg = "http://open.bigmodel/EgZjaHJvbWUqEAgAEAAYgwEY4wIYsQMYgAQyEAgAEAAYgwEY4wIYsQMYgAQyEwgBEC4YgwEYxwEYsQMY0QMYgAQyDQgCEAAYgwEYsQMYgAQyDQgDEAAYgwEYsQMYgAQyDQgEEAAYgwEYsQMYgAQyDQgFEAAYgwEYsQMYgAQyCggGEAAYsQMYgAQyCggHEAAYsQMYgAQyCggIEAAYsQMYgAQ"
//	res.Ad.Dplink = "http://open.bigmodel/EgZjaHJvbWUqEAgAEAAYgwEY4wIYsQMYgAQyEAgAEAAYgwEY4wIYsQMYgAQyEwgBEC4YgwEYxwEYsQMY0QMYgAQyDQgCEAAYgwEYsQMYgAQyDQgDEAAYgwEYsQMYgAQyDQgEEAAYgwEYsQMYgAQyDQgFEAAYgwEYsQMYgAQyCggGEAAYsQMYgAQyCggHEAAYsQMYgAQyCggIEAAYsQMYgAQ"
//	res.Ad.Adext.Pkg = "open.bigmode.cn"
//	//res.Ad.Adext.Nurl = ""
//	res.Ad.Adext.Kt = []string{
//		"http://open.bigmodel/trace?data=6b9a476e26a2483150d03e7237e94fca78ffc231938c70010bc7f19245e416c3884c03f3822ae7e253ad6f7bf73e3485f01736fb340e15328f6026a4210868f8e248ecff1b55f602727940b6256f5ffaf22f60adc49bece339f5d60129116f14fcfa6bb33aee2bf3677b62a3c4c74",
//	}
//
//	//logger.Log.Info().Msgf("TTTT 欢喜测试----------------------------")
//
//	//filterFloorPriceBool := dspAuction.FloorPriceFilter(this.RequestContext, this.DspSlotInfo, float64(res.Ad.Adext.Price))
//	//if filterFloorPriceBool == true {
//	//	res.Res = constant.REQ_CODE_DSP_NOT_FLOORPRICE
//	//	return impl.BidResponse{Res: constant.REQ_CODE_DSP_NOT_FLOORPRICE, Message: "预算出价太低"}
//	//}
//
//	price, _ := dspAuction.SetAuctionPrice(this.RequestContext, this.DspSlotInfo, float64(res.Ad.Adext.Price))
//	this.DspSlotInfo.BidPrice = price
//	this.DspSlotInfo.AuctionPrice = res.Ad.Adext.Price
//	res.Ad.Adext.Price = price
//	res.Res = constant.REQ_CODE_SUC // 一定要这这个否则没有上报链接
//	return res
//}
//
//func (this *THuanXiDsp) GetDspCode() int64 {
//	return this.DspCode
//}
//
//func THuanXiReplaceWithParams(url string, winPrice float64, w, h int) string {
//
//	// 宏替换
//	replacer := strings.NewReplacer()
//	return replacer.Replace(url)
//}
