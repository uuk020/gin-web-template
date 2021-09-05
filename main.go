package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/global"
	"github.com/uuk020/gin-web-template/initialize"
	"net/http"
)

func main() {
	initialize.InitConfig()
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%d", global.Settings.Port))
}
