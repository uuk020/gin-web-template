package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// StartGinServ 启动 gin 服务
func StartGinServ(isDefault bool, middleware ...gin.HandlerFunc) {
	var r *gin.Engine
	if isDefault {
		r = gin.Default()
	} else {
		r = gin.New()
	}
	r.Use(middleware...)
	// 初始化路由
	parseRoute(r)
	// 优雅关停服务
	gracefulShutdown(r)
}

// parseRoute 解析路由
func parseRoute(r *gin.Engine) {
	// 解析路由定义
	for group, routes := range RouterGroup {
		groupRouter := r.Group(group)
		for _, route := range routes {
			// 是文件路由构建文件路由
			if route.GetFilePath() != "" {
				switch route.GetMethod() {
				case "Static":
					groupRouter.Static(route.GetURL(), route.GetFilePath())
				case "StaticFS":
					groupRouter.StaticFS(route.GetURL(), http.Dir(route.GetFilePath()))
				default:
					groupRouter.StaticFile(route.GetURL(), route.GetFilePath())
				}
			} else {
				if route.GetMethod() == "Any" {
					groupRouter.Any(route.GetURL(), route.ExecHandler()...)
				} else {
					groupRouter.Handle(route.GetMethod(), route.GetURL(), route.ExecHandler()...)
				}
			}
		}
	}
}

// gracefulShutdown 优雅关停服务
func gracefulShutdown(r *gin.Engine) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", Settings.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("正常关闭服务, 请按 Ctrl+C 强制退出")

	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx1); err != nil {
		log.Fatal("服务已被强制关闭: ", err)
	}
	log.Println("服务已退出")
}
