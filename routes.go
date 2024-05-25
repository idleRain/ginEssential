package main

import (
	"github.com/gin-gonic/gin"
	"idleRain.com/ginEssential/controller"
	"idleRain.com/ginEssential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
