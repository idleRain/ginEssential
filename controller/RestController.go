package controller

import "github.com/gin-gonic/gin"

type IRestController interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Show(context *gin.Context)
	Delete(context *gin.Context)
}
