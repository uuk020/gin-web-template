package utils

import "github.com/gin-gonic/gin"

// Route 路由定义结构体
type Route struct {
	uRL         string
	method      string
	handlerFunc gin.HandlerFunc
	filePath    string
}

// GetURL 获取 uRL 字段
func (r *Route) GetURL() string {
	return r.uRL
}

// GetMethod 获取 method 字段
func (r *Route) GetMethod() string {
	return r.method
}

// GetHandlerFunc 获取 handlerFunc 字段
func (r *Route) GetHandlerFunc() gin.HandlerFunc  {
	return r.handlerFunc
}

// GetFilePath 获取 filePath 字段
func (r *Route) GetFilePath() string {
	return r.filePath
}

// NewRouter 构造路由结构体
func NewRouter(url, method string, handlerFunc gin.HandlerFunc) *Route {
	return &Route{
		uRL:         url,
		method:      method,
		handlerFunc: handlerFunc,
	}
}

// NewFileRouter 构造文件路由结构体
func NewFileRouter(url, method, filePath string) *Route {
	return &Route{
		uRL:         url,
		method:      method,
		filePath:    filePath,
	}
}
