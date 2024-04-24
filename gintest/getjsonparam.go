package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 定义map或结构体
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})
	r.Run("localhost:8080")
}
