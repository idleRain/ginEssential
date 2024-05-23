package main

import (
	"github.com/gin-gonic/gin"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/controller"
)

func main() {
	common.InitDB()
	db := common.GetDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)

	// 运行 gin
	panic(r.Run())
}
