package initialize

import (
	"github.com/spf13/viper"
	"github.com/uuk020/gin-web-template/configs"
	"github.com/uuk020/gin-web-template/internal/app/utils"
	"log"
)

// Config 初始化配置
func Config() {
	v := viper.New()
	v.SetConfigFile("../settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := configs.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	utils.Settings = serverConfig
	log.Printf("读取配置成功, 应用名称和日志地址: %s, %s", utils.Settings.Name, utils.Settings.LogsAddress)
}
