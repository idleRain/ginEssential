package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/model"
	"log"
	"net/http"
)

// Register 用户注册
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

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "注册失败！"})
		return
	}

	// 校验通过，开始新建用户
	newUSer := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
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

// Login 用户登录
func Login(context *gin.Context) {
	db := common.GetDB()
	// 获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")

	if len(telephone) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号长度必须为11位！"})
		return
	}
	if len(password) < 6 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能少于6位！"})
		return
	}
	// 判断用户是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 验证通过，发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器异常"})
		log.Printf("token 生成失败：%v", err)
		return
	}
	// 登录成功
	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登陆成功", "data": gin.H{"token": token}})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
