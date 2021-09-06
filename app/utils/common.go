package utils

import (
	"fmt"
	"github.com/uuk020/gin-web-template/global"
	"time"
)

func GetNowFormatTodayTime() string {
	location, _ := time.LoadLocation(global.Settings.Timezone)
	now := time.Now().In(location)
	return fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
}
