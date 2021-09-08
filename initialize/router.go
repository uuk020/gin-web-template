package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/app/middlewares"
	"github.com/uuk020/gin-web-template/router"
	"github.com/uuk020/gin-web-template/utils"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 注册路由定义
	router.App()
	// 启动 gin 服务
	r := gin.Default()
	// 解析路由定义
	for group, routes := range utils.RouterGroup {
		groupRouter := r.Group(group)
		for _, route := range routes {
			if route.GetFilePath() != "" {
				switch route.GetMethod() {
				case "Static":
					groupRouter.Static(route.GetURL(), route.GetFilePath())
				case "StaticFS":
					groupRouter.StaticFS(route.GetURL(), http.Dir(route.GetFilePath()))
				default:
					groupRouter.StaticFile(route.GetURL(), route.GetFilePath())
				}
			} else {
				if route.GetMethod() == "Any" {
					groupRouter.Any(route.GetURL(), route.GetHandlerFunc())
				} else {
					groupRouter.Handle(route.GetMethod(), route.GetURL(), route.GetHandlerFunc())
				}
			}
		}
	}
	// 加载全局中间件
	r.Use(middlewares.HttpRequestLogger(), middlewares.ServerDowntimeLogger(true))
	return r
}