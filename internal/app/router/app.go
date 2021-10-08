package router

import (
	"github.com/uuk020/gin-web-template/internal/app/controller"
	"github.com/uuk020/gin-web-template/internal/app/middlewares"
	"github.com/uuk020/gin-web-template/internal/app/utils"
)

// App 路由定义文件
func App() {
	// 示例代码
	utils.AddRoute("/", utils.NewRouter("/hello", "GET", utils.NewRouteHandle(controller.Hello)))
	utils.AddRoute("/", utils.NewRouter("/test", "GET", utils.NewRouteHandle(controller.ModuleMiddleware, middlewares.IPAuth())))
	utils.AddRoute("/", utils.NewFileRouter("/favicon.ico", "StaticFile", utils.StaticResources+"favicon.png"))
	utils.AddRoute("/", utils.NewRouter("/token", "POST", utils.NewRouteHandle(controller.Token)))
}
