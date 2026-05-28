package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomPick(pkgs []string, n int) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})
	if n > len(pkgs) {
		n = len(pkgs)
	}
	return pkgs[:n]
}

func main() {
	packages := []string{"com.xhey.xcamera", "com.xt.retouch", "com.ziroom.ziroomcustomer", "com.tencent.news", "com.taobao.trip", "com.sankuai.meituan", "com.shizhuang.duapp", "com.chaoxing.mobile", "com.jd.jrapp", "com.achievo.vipshop", "com.duowan.mobile", "com.msxf.ayh", "com.tencent.qqlive", "com.qiyi.video", "com.mt.mtxx.mtxx", "com.kuaishou.nebula", "com.lingan.seeyou", "com.UCMobile", "com.tencent.jkchess", "com.taobao.litetao", "com.kmxs.reader", "com.meitu.meiyancamera", "com.baidu.tieba", "com.baidu.searchbox", "com.immomo.momo", "com.netease.cloudmusic", "com.taobao.idlefish", "com.phoenix.read", "com.quark.browser", "com.tmall.wireless", "com.clover.daysmatter", "com.wuba", "com.tmri.app.main", "com.tianyancha.skyeye", "com.jingdong.app.mall", "com.kwai.videoeditor", "com.qiyi.video.lite", "com.duowan.kiwi", "com.Qunar", "com.netease.party", "com.cubic.autohome", "com.dragon.read", "com.tencent.tmgp.sgame", "com.babycloud.hanju", "com.taobao.taobao"}

	pick := RandomPick(packages, 3)

	for _, pkg := range pick {
		fmt.Println(pkg)
	}

}
