package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/model"
	"log"
	"net/http"
)

func Register(context *gin.Context) {
	db := common.GetDB()

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
	newUSer := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUSer)
	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

// 校验手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
