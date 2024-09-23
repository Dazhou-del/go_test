package main

import (
	"fmt"
	"regexp"
)

func main() {
	idCard := "511502199103223189"

	// 身份证号码正则表达式
	regex := `[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`

	re := regexp.MustCompile(regex)

	fmt.Println(re.MatchString(idCard))

	a := int64(22)
	b := int64(33)
	c := int64(12)

	d := b - a - c
	fmt.Println(d)
}
