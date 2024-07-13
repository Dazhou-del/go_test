package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	content := `{
  "user": {
    "name": "dj",
    "age": 18,
    "address": {
      "provice": "shanghai",
      "district": "xuhui"
    },
    "hobbies":["chess", "programming", "game"]
  }
}`
	jsonq := gojsonq.New().FromString(content)
	address := jsonq.Find("user.address")
	fmt.Println(address)
	jsonq.Reset()
}
