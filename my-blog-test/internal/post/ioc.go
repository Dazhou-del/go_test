package post

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	engine := gin.Default()
	postHandler := NewPostHandler(IPostService{})
	engine.GET("/posts/:id", postHandler.GetPost)
	return engine
}
