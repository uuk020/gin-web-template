package utils

import "github.com/gin-gonic/gin"

// Route 路由定义结构体
type Route struct {
	uRL      string
	method   string
	handler  *RouteHandle
	filePath string
}

// RouteHandle 路由执行结构体
type RouteHandle struct {
	middlewares []gin.HandlerFunc
	handlerFunc gin.HandlerFunc
}

func NewRouteHandle(handlerFunc gin.HandlerFunc, moduleMiddlewares ...gin.HandlerFunc) *RouteHandle {
	return &RouteHandle{
		middlewares: moduleMiddlewares,
		handlerFunc: handlerFunc,
	}
}

// GetURL 获取 uRL 字段
func (r *Route) GetURL() string {
	return r.uRL
}

// GetMethod 获取 method 字段
func (r *Route) GetMethod() string {
	return r.method
}

// ExecHandler 执行模型中间件和 handlerFunc
func (r *Route) ExecHandler() []gin.HandlerFunc {
	temp := make([]gin.HandlerFunc, 0)
	if len(r.handler.middlewares) > 0 {
		for _, middleware := range r.handler.middlewares {
			temp = append(temp, middleware)
		}
	}
	temp = append(temp, r.handler.handlerFunc)
	return temp
}

// GetFilePath 获取 filePath 字段
func (r *Route) GetFilePath() string {
	return r.filePath
}

// NewRouter 构造路由结构体
func NewRouter(url, method string, handler *RouteHandle) *Route {
	return &Route{
		uRL:     url,
		method:  method,
		handler: handler,
	}
}

// NewFileRouter 构造文件路由结构体
func NewFileRouter(url, method, filePath string) *Route {
	return &Route{
		uRL:      url,
		method:   method,
		filePath: filePath,
	}
}
