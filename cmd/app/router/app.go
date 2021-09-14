package router

import (
	"github.com/uuk020/gin-web-template/cmd/app/controller"
	"github.com/uuk020/gin-web-template/cmd/app/middlewares"
	"github.com/uuk020/gin-web-template/cmd/app/utils"
)
// App 路由定义文件
func App() {
	// 示例代码
	utils.AddRoute("/", utils.NewRouter("/hello", "GET", utils.NewRouteHandle(controller.Hello, middlewares.JWTAuth())))
	utils.AddRoute("/", utils.NewRouter("/test", "GET", utils.NewRouteHandle(controller.ModuleMiddleware, middlewares.IPAuth())))
	utils.AddRoute("/", utils.NewFileRouter("/favicon.ico", "StaticFile", utils.StaticResources + "favicon.png"))
	utils.AddRoute("/", utils.NewRouter("/token", "POST", utils.NewRouteHandle(controller.Token)))
}
