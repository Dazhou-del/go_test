package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/gintest/routers"
)

func main() {
	//r := routers.SetupRouter()
	//if err := r.Run("localhost:8080"); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

	r.Run("localhost:8080")
}
