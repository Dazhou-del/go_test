package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	//利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "沙河小王子")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
