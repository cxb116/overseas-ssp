package main

import (
	"fmt"
	"strings"
)

func main() {
	if strings.Contains( "com.baidu.searchbox", "\u200C") {
		fmt.Println("包含零宽字符")
	}
}
