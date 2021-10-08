package utils

import (
	"fmt"
	"time"
)

// GetNowFormatTodayTime 获取今日时间
func GetNowFormatTodayTime() string {
	location, _ := time.LoadLocation(Settings.Timezone)
	now := time.Now().In(location)
	return fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()), now.Day())
}

func AddRoute(group string, route *Route) {
	if len(RouterGroup) == 0 {
		RouterGroup = make(map[string][]*Route)
	}
	RouterGroup[group] = append(RouterGroup[group], route)
}
