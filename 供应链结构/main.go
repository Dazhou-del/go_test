package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type StockItem struct {
	WarehouseID   int `json:"warehouse_id"`
	Distributable int `json:"distributable"`
	StockType     int `json:"stockType,omitempty"`
	StockDetail   int `json:"stockDetail,omitempty"`
}

type Root struct {
	Data map[string][]StockItem `json:"data"`
}

func main() {
	jsonData := `
    {
        "data": {
            "34781004111": [
                {
                    "warehouse_id": 1,
                    "distributable": 1,
                    "stockType": 1,
                    "stockDetail": 2
                },
                {
                    "warehouse_id": 2,
                    "distributable": 2
                },
                {
                    "warehouse_id": 3,
                    "distributable": 4
                }
            ],
            "34781004112": [
                {
                    "warehouse_id": 1,
                    "distributable": 1,
                    "stockType": 1,
                    "stockDetail": 2
                },
                {
                    "warehouse_id": 2,
                    "distributable": 2
                },
                {
                    "warehouse_id": 3,
                    "distributable": 4
                }
            ]
        }
    }
    `

	var root Root
	err := json.Unmarshal([]byte(jsonData), &root)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// 打印解析后的数据
	fmt.Println("解析后的数据:", root.Data)

	for k, v := range root.Data {
		fmt.Println("k", k)
		fmt.Println("v", v)
	}
}
