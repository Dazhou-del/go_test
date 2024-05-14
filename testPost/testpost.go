package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	var host = "https://admin.rrzu.com"
	model := "dev"
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Failed to create cookie jar:", err)
		return
	}
	client := &http.Client{
		Jar: jar,
	}

	switch model {
	case "pro":
		host = "https://admin.rrzu.com/super/warehouse/dict/spu-relation-by-page-type"
	case "dev":
		host = "https://dev6-admin.rrzuji.com/super/warehouse/sku/get-sku-id-by-attr"
	case "test":
		host = "https://test1-admin.rrzuji.com/super/warehouse/dict/spu-relation-by-page-type"
	}
	// JSON 数据
	jsonData := `
	{
		"spu_info": [
			{
				"category_id": 1,
				"brand_id": 1,
				"model_id": 1,
				"attr": []
			},
			{
				"category_id": 1,
				"brand_id": 1,
				"model_id": 1,
				"attr": []
			}
		]
	}
	`
	// 解析 JSON 数据到 SPUInfo 结构体
	var spuInfo SPUInfo
	err = json.Unmarshal([]byte(jsonData), &spuInfo)
	if err != nil {
		fmt.Println("Failed to marshal data:", err)
	}
	fmt.Printf("spuInfo: %+v\n", spuInfo)
	//reqBody, err := json.Marshal(spuInfo)
	reqBodys, _ := json.Marshal(jsonData)
	fmt.Printf("reqBody: %+v\n", bytes.NewBuffer(reqBodys))
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(reqBodys))
	if err != nil {
		//logx.WithContext(ctx).Errorf("dingtalk send request failed [%v]", err)
		//return nil, err
		log.Fatal(err)
	}
	req.Header.Add("Cookie", "RRZUJI_TEST_5=e69cd61f70a40db0a0e1b3e50956550726ab5e629726879b9e45cbc1cfb18442a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_5%22%3Bi%3A1%3Bs%3A51%3A%22%5B66667079%2C%22Tpsp3b3mouvMXkGKJ-d-Hg-Oh_ew9xSz%22%2C85110%5D%22%3B%7D; RRZUJI_TEST_6=1dbfacaa442415f7d6f39d3375bd2e23147e186fd7e7e5eaa0ffa0963958de71a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_6%22%3Bi%3A1%3Bs%3A51%3A%22%5B66667079%2C%22Tpsp3b3mouvMXkGKJ-d-Hg-Oh_ew9xSz%22%2C84405%5D%22%3B%7D; _csrf=ccc2195ee1da9dd646e16cb930e9fb28ebc904285c53e43964eec6bed4f6a9e9a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22_2bfBomSDCdPXIUqv7mtVVdAlRFAm-Pr%22%3B%7D; Go-Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU0NTYyNzEsImlhdCI6MTcxNTQxMzA3MSwic2NlbmUiOjEsInNlcnZlcl9pZCI6LTEsInVzZXJfaWQiOjYwMDA0MzIzfQ.NqW3KDCLC5KKjLEh6Qdhe_HMntaNyFWfFQRM2XbvJAI; RRZUJI_TEST_1=8002692b44cb40c99a2b58c656b56f5cc60c22c9105a61434a56ae60ce9fc6bea%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_1%22%3Bi%3A1%3Bs%3A19%3A%22%5B60004323%2C%22%22%2C87390%5D%22%3B%7D; Data-Token=f85d938269131863; PHPSESSID=b74qp5r5h8gsvmraobjgch77ls")
	response, err := client.Do(req)
	if err != nil {
		//logx.WithContext(ctx).Errorf("dingtalk send request failed [%v]", err)
		//return nil, err
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		//logx.WithContext(ctx).Errorf("dingtalk ioutil ReadAll failed [%v]", err)
		//
		//return nil, err
		log.Fatal(err)
	}
	resp := phpResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		//logx.WithContext(ctx).Errorf("dingtalk ioutil ReadAll failed [%v]", err)
		//
		//return nil, err
		log.Fatal(err)
	}
	fmt.Println("resp:", resp)
}

// 定义产品信息结构体
type ProductInfo struct {
	category_id int      `json:"category_id"`
	brand_id    int      `json:"brand_id"`
	model_id    int      `json:"model_id"`
	attr        []string `json:"attr"`
}

type phpResp struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Error   string        `json:"error"`
	Data    []interface{} `json:"data"`
}

// 定义包含SPU信息的结构体
type SPUInfo struct {
	SPUInfo []ProductInfo `json:"spu_info"`
}
