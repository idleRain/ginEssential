package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"idleRain.com/ginEssential/model"
)

var DB *gorm.DB

// InitDB 初始化 DB
func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	database := viper.GetString("datasource.database")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)

	if err != nil {
		panic("InitDB, err:" + err.Error())
	}
	// 自动创建数据表
	db.AutoMigrate(&model.User{})

	DB = db
}

// GetDB 获取 DB 实例
func GetDB() *gorm.DB {
	return DB
}
