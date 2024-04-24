package routers

import "github.com/gin-gonic/gin"

func cc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func LoadBlog(r *gin.Engine) {
	r.GET("/cc", cc)

}
