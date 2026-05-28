package main

import (
	"fmt"
	"strings"
)

func baiduMacroReplace(url string) string {

	replacer := strings.NewReplacer(
		"__1__", "&1&",
		"__2__", "&2&",
		"__3__", "&3&",
	)
	return replacer.Replace(url)
}

var url = baiduMacroReplace("http://www.baidu.com?__1__$__2__$__3__")

func main() {

	fmt.Println(url)
}
