package middleware

import (
	"github.com/JoyZF/blog_gin/pkg/app"
	"github.com/JoyZF/blog_gin/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, b := c.GetQuery("token");b{
			token = s
		}else{
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		}else{
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnAuth
				default:
					ecode = errcode.UnAuth
				}
			}
		}


		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()


	}
}