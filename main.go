package main

import (
	"fmt"
	"github.com/uuk020/gin-web-template/initialize"
	"github.com/uuk020/gin-web-template/utils"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化路由并且启动 gin 服务
	r := initialize.InitRouter()
	err := r.Run(fmt.Sprintf(":%d", utils.Settings.Port))
	if err != nil {
		panic("gin 服务运行失败: " + err.Error())
	}
}
