package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/router"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	db := common.GetDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)

	// 运行 gin
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
}

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       // 读取配置文件名
	viper.SetConfigType("yml")               // 读取配置文件类型
	viper.AddConfigPath(workDir + "/config") // 添加文件位置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
