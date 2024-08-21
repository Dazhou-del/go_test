package main

import (
	"fmt"
)

func isChinese(r rune) bool {
	if r >= 0x4e00 && r <= 0x9fff {
		return true
	}
	return false
}

func main() {
	testRune := '⽬'
	if isChinese(testRune) {
		fmt.Printf("'%c' 是中文字符。\n", testRune)
	} else {
		fmt.Printf("'%c' 不是中文字符。\n", testRune)
	}
}
