package utils

import (
	"github.com/uuk020/gin-web-template/config"
	"go.uber.org/zap"
)

var (
	// Settings 系统配置
	Settings config.ServerConfig
	// Lg 日志读取器
	Lg *zap.Logger
	// RouterGroup 路由规则
	RouterGroup map[string][]*Route
)
