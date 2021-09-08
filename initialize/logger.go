package initialize

import (
	"fmt"
	"github.com/uuk020/gin-web-template/utils"
	"go.uber.org/zap"
)

func InitLogger()  {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", utils.Settings.LogsAddress, utils.GetNowFormatTodayTime()),
		"stdout",
	}
	logg, _ := cfg.Build()
	zap.ReplaceGlobals(logg)
	utils.Lg = logg
}
