package middlewares

import (

	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/cmd/app/utils"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("authorization")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "请登录",
			})
			context.Abort()
			return
		}
		j := utils.NewJWT()
		cliams, err := j.ParseToken(token)
		if err != nil {
			if err == utils.ErrTokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"message": "已过期",
				})
				context.Abort()
				return
			}
		}
		context.Set("cliams", cliams)
		context.Set("userId", cliams.ID)
		context.Next()
	}
}
