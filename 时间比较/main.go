package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now().Unix()
	if int64(1724766860) >= nowTime {
		fmt.Println("asdwdwww")
	}

	overdueStarTime := time.Unix(int64(1724766860), 0)
	nowDate := time.Unix(nowTime, 0)

	if int(nowDate.Sub(overdueStarTime).Hours()/24) >= 30 {
		fmt.Println("asdwdwww")
		fmt.Println(nowDate.Sub(overdueStarTime).Hours() / 24)
	}

}
