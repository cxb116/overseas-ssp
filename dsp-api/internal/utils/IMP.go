package utils

//pkgList := []string{"tv.danmaku.bili",
//"com.ss.android.article.news",
//"com.ximalaya.ting.android",
//"com.tmall.wireless",
//"com.quark.browser",
//"com.eg.android.AlipayGphone",
//"com.taobao.taobao",
//"com.xs.fm",
//"com.xs.fm.lite",
//"com.baidu.searchbox",
//"com.baidu.searchbox.lite",
//"com.baidu.netdisk",
//"com.phoenix.read",
//"com.taobao.idlefish",
//"com.immomo.momo",
//"me.ele",
//}

func UniquePackages(arrays ...[]string) []string {
	set := make(map[string]struct{})

	for _, arr := range arrays {
		for _, pkg := range arr {
			if pkg == "" {
				continue
			}
			set[pkg] = struct{}{}
		}
	}

	result := make([]string, 0, len(set))
	for pkg := range set {
		result = append(result, pkg)
	}

	return result
}
