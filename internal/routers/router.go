package routers

import (
	v1 "github.com/JoyZF/blog_gin/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	tag := v1.Tag{}
	article := v1.Article{}
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/:id",tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update)
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.Get)


		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/:id",article.Delete)
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles",article.Get)
	}
	return r
}