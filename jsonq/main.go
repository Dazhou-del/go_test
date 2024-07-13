package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	//	content := `{
	//  "user": {
	//    "name": "dj",
	//    "age": 18,
	//    "address": {
	//      "provice": "shanghai",
	//      "district": "xuhui"
	//    },
	//    "hobbies":["chess", "programming", "game"]
	//  }
	//}`
	//	jsonq := gojsonq.New().FromString(content)
	//	address := jsonq.Find("user.address")
	//	fmt.Println(address)
	//	jsonq.Reset()
	//
	//	fmt.Println(jsonq.Find("user.hobbies.[0]"))

	//get := gojsonq.New().File("jsonq/data.json").From("items").Select("id", "name").Get()
	//
	//data, _ := json.MarshalIndent(get, "", " ")
	//
	//fmt.Println(string(data))

	//file := gojsonq.New().File("jsonq/data.json")
	//
	//get := file.From("items").Select("id", "name").Where("id", "=", 1).OrWhere("id", "=", 2).Get()
	//
	//fmt.Println(get)

	//file := gojsonq.New().File("jsonq/data.json")
	//
	//rs := file.From("items").Select("id", "name", "count").
	//	Where("count", ">", 1).Where("price", "<", 100).Get()
	//
	//fmt.Println(rs)

	//gq := gojsonq.New().File("jsonq/data.json").From("items")
	//
	//fmt.Println("Total Count:", gq.Sum("count"))
	//fmt.Println("Min Price:", gq.Min("price"))
	//fmt.Println("Max Price:", gq.Max("price"))
	//fmt.Println("Avg Price:", gq.Avg("price"))

	gq := gojsonq.New().File("jsonq/data.json").From("items")

	fmt.Println(gq.From("items").GroupBy("price").Get())
	gq.Reset()
	fmt.Println(gq.From("items").SortBy("price", "desc").Get())
}
