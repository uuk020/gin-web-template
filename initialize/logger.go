package initialize

import (
	"fmt"
	"github.com/uuk020/gin-web-template/app/utils"
	"github.com/uuk020/gin-web-template/global"
	"go.uber.org/zap"
)

func InitLogger()  {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	global.Lg = logg
}
