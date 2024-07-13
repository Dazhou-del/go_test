package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	morning = 1
	duck    = 2
	night   = 3
)

func main() {
	// 定义一个map来存储数据
	schedule := map[string][]int{
		"Monday":    {2, 3},
		"Tuesday":   {1, 2},
		"Wednesday": {1, 2},
		"Thursday":  {1, 2},
		"Friday":    {1, 2, 3},
		"Saturday":  {2},
		"Sunday":    {},
	}

	// 输入要查询的日期
	input := "Saturday"
	findScheduleFromJSON(input)

	// 查询并打印结果
	if days, exists := schedule[input]; exists {
		fmt.Printf("%s: %v\n", input, days)
	} else {
		log.Printf("No schedule found for %s", input)
	}
}

// 上面的代码直接使用map进行查询，如果你的输入是一个JSON字符串，可以使用下面的方式解析后查询
func findScheduleFromJSON(inputDay string) {
	jsonStr := `{"Monday":{},"Tuesday":[1,2],"Wednesday":[1,2],"Thursday":[1,2],"Friday":[1,2,3],"Saturday":[2],"Sunday":[]}`
	var schedule map[string][]int
	err := json.Unmarshal([]byte(jsonStr), &schedule)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	values := schedule[inputDay]
	if len(values) == 1 {
		if values[0] == morning {
			fmt.Println("早班吗?:", values[0])
		}
		if values[0] == duck {
			fmt.Println("晚班吗?:", values[0])
		}
	}
	if len(values) == 2 {
		if values[0] == morning && values[1] == duck {
			fmt.Println("早晚吗?:", values[0], values[1])
		}
		if values[0] == morning && values[1] == night {
			fmt.Println("早夜班?:", values[0], values[1])
		}
		if values[0] == duck && values[1] == night {
			fmt.Println("晚夜班?:", values[0], values[1])
		}
	}

	if len(values) == 3 {
		fmt.Println("这是早晚班:values:", values)
	}

}
