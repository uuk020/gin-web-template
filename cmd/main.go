package main

import (
	"fmt"
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
	// 初始化路由并且启动 gin 服务
	r := initialize.Router()
	err := r.Run(fmt.Sprintf(":%d", utils.Settings.Port))
	if err != nil {
		panic("gin 服务运行失败: " + err.Error())
	}
}
