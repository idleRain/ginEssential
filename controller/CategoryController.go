package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"idleRain.com/ginEssential/common"
	"idleRain.com/ginEssential/model"
	"idleRain.com/ginEssential/response"
	"strconv"
)

type ICategoryController interface {
	IRestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() CategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

// Create 创建文章分类
func (c CategoryController) Create(ctx *gin.Context) {
	requestCategory := model.Category{}
	ctx.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "分类名不能为空！")
		return
	}

	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

// Update 更新文章分类
func (c CategoryController) Update(ctx *gin.Context) {
	requestCategory := model.Category{}
	ctx.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "分类名不能为空！")
		return
	}
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory := model.Category{}
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	// 更新分类
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, nil, "更新成功！")
}

// Show 获取文章分类
func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category := model.Category{}
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

// Delete 删除文章分类
func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "删除失败，请重试！")
		return
	}
	response.Success(ctx, nil, "删除成功！")
}
