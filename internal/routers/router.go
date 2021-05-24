package routers

import (
	_ "github.com/JoyZF/blog_gin/docs"
	"github.com/JoyZF/blog_gin/global"
	"github.com/JoyZF/blog_gin/internal/middleware"
	"github.com/JoyZF/blog_gin/internal/routers/api"
	v1 "github.com/JoyZF/blog_gin/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine  {
	r := gin.New()
	r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/upload/file", Upload{}.UploadFile)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth",api.GetAuth)

	tag := v1.Tag{}
	article := v1.Article{}
	apiv1 := r.Group("/api/v1",middleware.JWT())
	{
		apiv1.POST("/tags",tag.Create)
		apiv1.DELETE("/tags/:id",tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update)
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.List)


		apiv1.POST("/articles",article.Create)
		apiv1.DELETE("/articles/:id",article.Delete)
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles",article.Get)
	}
	return r
}