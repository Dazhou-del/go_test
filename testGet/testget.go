package main

import (
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
		host = "https://dev6-admin.rrzuji.com/super/warehouse/dict/spu-relation-by-page-type"
	case "test":
		host = "https://test1-admin.rrzuji.com/super/warehouse/dict/spu-relation-by-page-type"
	}

	fullURL := fmt.Sprintf("%s?page_type=%d", host, 1)
	fmt.Println(fullURL)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	//cookieValue := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU0ODI5NTcsImlhdCI6MTcxNTM5NjU1Nywic2NlbmUiOjEsInNlcnZlcl9pZCI6LTEsInVzZXJfaWQiOjY2NjY3MDc5fQ.bFPfMq-3uAMRtzeu3yTDMlZE9OUIuCcQ2k0gYtDE0Ts; Data-Token=af0f7a06589ee168; RRZUJI_TEST_5=e69cd61f70a40db0a0e1b3e50956550726ab5e629726879b9e45cbc1cfb18442a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_5%22%3Bi%3A1%3Bs%3A51%3A%22%5B66667079%2C%22Tpsp3b3mouvMXkGKJ-d-Hg-Oh_ew9xSz%22%2C85110%5D%22%3B%7D; PHPSESSID=k5r598uiom1hf3sq6e0doepkjt"
	cookieValue := "PHPSESSID=8naeueq4i81g531d51nbkeiemj; RRZUJI_TEST_7=9898c9e3d213e7a61688add6347473f3b9062469e1dc2f56f92d538ce3a37391a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_7%22%3Bi%3A1%3Bs%3A44%3A%22%5B4%2C%22PJAP1kTQBlA2tltlg3XCu1M-wdRhKRRW%22%2C84762%5D%22%3B%7D; _csrf=0763a9281d6c230b8fbc0e7e8a63d7402d8b691272d5c0dc7b0d0a720eb705eca%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22puh1_8GtpfrbWvkLCCg81weuHWSQ2L8y%22%3B%7D"
	fmt.Println(cookieValue)
	req.Header.Add("Cookie", "RRZUJI_TEST_5=e69cd61f70a40db0a0e1b3e50956550726ab5e629726879b9e45cbc1cfb18442a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_5%22%3Bi%3A1%3Bs%3A51%3A%22%5B66667079%2C%22Tpsp3b3mouvMXkGKJ-d-Hg-Oh_ew9xSz%22%2C85110%5D%22%3B%7D; RRZUJI_TEST_6=1dbfacaa442415f7d6f39d3375bd2e23147e186fd7e7e5eaa0ffa0963958de71a%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_6%22%3Bi%3A1%3Bs%3A51%3A%22%5B66667079%2C%22Tpsp3b3mouvMXkGKJ-d-Hg-Oh_ew9xSz%22%2C84405%5D%22%3B%7D; _csrf=ccc2195ee1da9dd646e16cb930e9fb28ebc904285c53e43964eec6bed4f6a9e9a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22_2bfBomSDCdPXIUqv7mtVVdAlRFAm-Pr%22%3B%7D; Go-Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU0NTYyNzEsImlhdCI6MTcxNTQxMzA3MSwic2NlbmUiOjEsInNlcnZlcl9pZCI6LTEsInVzZXJfaWQiOjYwMDA0MzIzfQ.NqW3KDCLC5KKjLEh6Qdhe_HMntaNyFWfFQRM2XbvJAI; RRZUJI_TEST_1=8002692b44cb40c99a2b58c656b56f5cc60c22c9105a61434a56ae60ce9fc6bea%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22RRZUJI_TEST_1%22%3Bi%3A1%3Bs%3A19%3A%22%5B60004323%2C%22%22%2C87390%5D%22%3B%7D; Data-Token=f85d938269131863; PHPSESSID=b74qp5r5h8gsvmraobjgch77ls")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
	//err = json.Unmarshal(body, resp)
	fmt.Println(string(body))
	fmt.Println(fullURL)
}
