package main

import (
	"github.com/gin-gonic/gin"
	"idleRain.com/ginEssential/common"
)

func main() {
	common.InitDB()
	db := common.GetDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	// 运行 gin
	panic(r.Run())
}
