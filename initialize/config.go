package initialize

import (
	"github.com/spf13/viper"
	"github.com/uuk020/gin-web-template/config"
	"github.com/uuk020/gin-web-template/utils"
	"log"
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	utils.Settings = serverConfig
	log.Printf("读取配置成功, 应用名称和日志地址: %s, %s", utils.Settings.Name, utils.Settings.LogsAddress)
}
