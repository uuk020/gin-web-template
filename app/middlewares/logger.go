package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/uuk020/gin-web-template/utils"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func HttpRequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		query := context.Request.URL.RawQuery
		context.Next()
		cost := time.Since(start)
		if context.Writer.Status() != http.StatusOK {
			zap.L().Error(path,
				zap.Int("status", context.Writer.Status()),
				zap.String("method", context.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", context.ClientIP()),
				zap.String("user-agent", context.Request.UserAgent()),
				zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		} else {
			zap.L().Info(path,
				zap.Int("status", context.Writer.Status()),
				zap.String("method", context.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", context.ClientIP()),
				zap.String("user-agent", context.Request.UserAgent()),
				zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}

func ServerDowntimeLogger(stack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.
							ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(context.Request, false)
				if brokenPipe {
					utils.Lg.Error(context.Request.URL.Path, zap.Any("error", err), zap.String("request",
						string(httpRequest)))
					_ = context.Error(err.(error))
					context.Abort()
					return
				}
				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		context.Next()
	}
}
