package router

import (
	"github.com/gin-gonic/gin"
	"idleRain.com/ginEssential/controller"
	"idleRain.com/ginEssential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	authRoutes := r.Group("/api/auth")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)
	authRoutes.POST("/info", middleware.AuthMiddleware(), controller.Info)

	// 文章分类路由组
	categoryRoutes := r.Group("/api/category")
	// 创建 categoryController
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("/createCategory", categoryController.Create)
	categoryRoutes.POST("/getCategory/:id", categoryController.Show)
	categoryRoutes.POST("/updateCategory/:id", categoryController.Update)
	categoryRoutes.POST("/deleteCategory/:id", categoryController.Delete)

	return r
}
