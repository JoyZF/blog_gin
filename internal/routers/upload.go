package routers

import (
	"github.com/JoyZF/blog_gin/global"
	"github.com/JoyZF/blog_gin/internal/service"
	"github.com/JoyZF/blog_gin/pkg/app"
	"github.com/JoyZF/blog_gin/pkg/convert"
	"github.com/JoyZF/blog_gin/pkg/errcode"
	"github.com/JoyZF/blog_gin/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf( "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
	return
}