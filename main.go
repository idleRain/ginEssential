package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// User 用户接口体
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(context *gin.Context) {
		// 获取参数
		name := context.PostForm("name")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")

		log.Println(name, telephone, password)

		// 校验
		if len(name) == 0 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不能为空！"})
			return
		}
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号长度必须为11位！"})
			return
		}
		if len(password) < 6 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能少于6位！"})
			return
		}
		if isTelephoneExist(db, telephone) {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在！"})
			return
		}

		// 校验通过，开始新建用户
		newUSer := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUSer)
		context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
	})

	// 运行 gin
	panic(r.Run())
}

// InitDB 初始化 DB
func InitDB() *gorm.DB {
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
	db.AutoMigrate(&User{})
	return db
}

// 校验手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
