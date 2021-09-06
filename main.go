package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/app/middlewares"
	"github.com/uuk020/gin-web-template/global"
	"github.com/uuk020/gin-web-template/initialize"
	"net/http"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	r := gin.Default()
	r.Use(middlewares.HttpRequestLogger(), middlewares.ServerDowntimeLogger(true))
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
