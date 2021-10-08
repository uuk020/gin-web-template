package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/internal/app/utils"
	"net/http"
)

// Hello this is hello world
func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}

// ModuleMiddleware 模块中间件
func ModuleMiddleware(context *gin.Context) {
	context.String(http.StatusOK, "hello test")
}
// Token 创建 token
func Token(context *gin.Context)  {
	getToken, err := utils.CreateToken(1, "wythe", 1)
	if err != nil {
		context.String(http.StatusInternalServerError, "token 创建失败")
		context.Abort()
		return
	}
	context.String(http.StatusOK, getToken)
}