package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 处理浏览器跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:12345") // 允许访问的域名
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")                     // 允许访问的请求方式
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")                     // 允许携带的请求头
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")                       // 缓存时限
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(http.StatusOK)
		} else {
			context.Next()
		}
	}
}
