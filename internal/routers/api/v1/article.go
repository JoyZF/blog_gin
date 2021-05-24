package v1

import (
	"github.com/JoyZF/blog_gin/pkg/app"
	"github.com/JoyZF/blog_gin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {

}

func NewArticle() Tag {
	return Tag{}
}

func (a Article) Get(c *gin.Context)  {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

func (a Article) List(c *gin.Context)  {}

func (a Article) Create(c *gin.Context)  {}

func (a Article) Update(c *gin.Context)  {}

func (a Article) Delete(c *gin.Context)  {}