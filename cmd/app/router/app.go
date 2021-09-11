package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/cmd/app/middlewares"
	"github.com/uuk020/gin-web-template/cmd/app/utils"
	"net/http"
)
// App 路由定义文件
func App() {
	// 示例代码
	utils.AddRoute("/", utils.NewRouter("/hello", "GET", utils.NewRouteHandle(func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})))
	utils.AddRoute("/", utils.NewRouter("/test", "GET", utils.NewRouteHandle(func(context *gin.Context) {
		context.String(http.StatusOK, "hello test")
	}, middlewares.IPAuth())))
}
