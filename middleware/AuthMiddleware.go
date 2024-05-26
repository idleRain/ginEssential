package middleware

import (
	"github.com/gin-gonic/gin"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/model"
	"net/http"
	"strings"
)

// AuthMiddleware 中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 authorization header
		tokenString := context.GetHeader("Authorization")
		// 验证 token
		// 是否为空，获取是否为 Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		// 验证 token 失效或者无效
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		// 验证通过
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)
		// 用户不存在
		if userId == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		// 用户存在，将用户信息写入上下文
		context.Set("user", user)
		context.Next()
	}
}
