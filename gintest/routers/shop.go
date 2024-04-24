package routers

import "github.com/gin-gonic/gin"

func dd(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func LoadShop(r *gin.Engine) {
	r.GET("/dd", dd)

}
