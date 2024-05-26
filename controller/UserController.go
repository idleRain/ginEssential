package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/dto"
	"idleRain.com/ginEssential/model"
	"idleRain.com/ginEssential/response"
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
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空！")
		return
	}
	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号长度必须为11位！")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能少于6位！")
		return
	}
	if isTelephoneExist(db, telephone) {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已存在！")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "注册失败！")
		return
	}

	// 校验通过，开始新建用户
	newUSer := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	db.Create(&newUSer)
	response.Success(context, nil, "注册成功")
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
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号长度必须为11位！")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能少于6位！")
		return
	}
	// 判断用户是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(context, http.StatusBadRequest, 400, nil, "用户不存在！")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误！")
		return
	}
	// 验证通过，发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "服务器异常！")
		log.Printf("token 生成失败：%v", err)
		return
	}
	// 登录成功
	response.Success(context, gin.H{"token": token}, "登陆成功")
}

// Info 返回用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "成功")
}
