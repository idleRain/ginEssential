package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"idleRain.com/ginEssential/model"
)

var DB *gorm.DB

// InitDB 初始化 DB
func InitDB() {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	username := "root"
	database := "ginessential"
	passwrod := "123456"
	charset := "utf8"
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		passwrod,
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
