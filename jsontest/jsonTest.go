package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetSkuIdByAttrReqData 定义结构体
type GetSkuIdByAttrReqData struct {
	CategoryID int      `json:"category_id"`
	BrandID    int      `json:"brand_id"`
	ModelID    int      `json:"model_id"`
	Attr       []string `json:"attr"`
}

// GetSkuIdByAttrReq 定义结构体
type GetSkuIdByAttrReq struct {
	SPUInfo []GetSkuIdByAttrReqData `json:"spu_info"`
}

// toLowerCaseJSON 将结构体转换为小写JSON字符串的辅助函数
func toLowerCaseJSON(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	lowerCaseMap := make(map[string]interface{})
	for k, val := range m {
		lowerCaseMap[strings.ToLower(k)] = val
	}

	return json.Marshal(lowerCaseMap)
}

func main() {
	req := GetSkuIdByAttrReq{
		SPUInfo: []GetSkuIdByAttrReqData{
			{
				CategoryID: 1,
				BrandID:    2,
				ModelID:    3,
				Attr:       []string{"attr1", "attr2"},
			},
		},
	}

	jsonStr, err := toLowerCaseJSON(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonStr))
}
