package utils

import (
	"github.com/uuk020/gin-web-template/configs"
	"go.uber.org/zap"
)
// 全局变量
var (
	// Settings 系统配置
	Settings configs.ServerConfig
	// Lg 日志读取器
	Lg *zap.Logger
	// RouterGroup 路由规则
	RouterGroup map[string][]*Route
)
