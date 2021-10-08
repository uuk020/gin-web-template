package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IPAuth IP 白名单
func IPAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		ipList := []string{
			"127.0.0.2",
		}
		flag := false
		clientIP := context.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag {
			context.String(http.StatusForbidden, "%s, not in ip list", clientIP)
			context.Abort()
		}
		context.Next()
	}
}
