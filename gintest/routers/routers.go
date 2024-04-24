package routers

import "github.com/gin-gonic/gin"

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", helloWorld)
	return r
}
