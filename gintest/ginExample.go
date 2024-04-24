package main

import "github.com/gin-gonic/gin"



func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//GET 请求 路径为/test
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/test", func(c *gin.Context) {
		//c.json,返回jons格式的数据
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run("localhost:8077")
}
