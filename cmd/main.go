package main

import (
	"github.com/uuk020/gin-web-template/cmd/app/middlewares"
	"github.com/uuk020/gin-web-template/cmd/app/router"
	"github.com/uuk020/gin-web-template/cmd/app/utils"
	"github.com/uuk020/gin-web-template/initialize"
)

func main() {
	// 初始化配置
	initialize.Config()
	// 初始化日志
	initialize.Logger()
	// 初始化数据库相关
	initialize.InitMysql()
	initialize.InitRedis()
	// 初始化路由
	router.App()
	// 启动 gin 服务
	utils.StartGinServ(true, middlewares.HttpRequestLogger(), middlewares.ServerDowntimeLogger(true))
}
