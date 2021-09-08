package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/utils"
	"net/http"
)
// App 路由定义文件
func App() {
	utils.AddRoute("/", utils.NewRouter("/hello", "GET", func(context *gin.Context) {
		context.String(http.StatusOK, "hello World")
	}))
}
