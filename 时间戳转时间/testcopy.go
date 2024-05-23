package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	rawTimestamps := `
`

	// 去除前后的空白字符并按行分割字符串
	timestampStrings := strings.Fields(strings.TrimSpace(rawTimestamps))

	for _, timestampStr := range timestampStrings {
		// 将字符串转换为整型时间戳
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			fmt.Println("Error converting string to timestamp:", err)
			continue
		}

		// 将时间戳转换为time.Time类型
		t := time.Unix(timestamp, 0)

		// 格式化时间
		formattedTime := t.Format("2006-01-02 15:04:05")

		fmt.Println(formattedTime)
	}
}
