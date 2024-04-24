package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "ces")
		address := c.Query("address")
		//输出结果json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run("localhost:8080")
}
