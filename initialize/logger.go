package initialize

import (
	"fmt"
	utils2 "github.com/uuk020/gin-web-template/cmd/app/utils"
	"go.uber.org/zap"
)

// Logger 初始化日志
func Logger()  {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", utils2.Settings.LogsAddress, utils2.GetNowFormatTodayTime()),
		"stdout",
	}
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	utils2.Lg = logg
}
