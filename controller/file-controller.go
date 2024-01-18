package controller

import (
	"common-web-framework/common"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	service service.FileService
}

func (f FileController) UploadImageFile(ctx *gin.Context) {
	var result = f.service.UploadImageFile(ctx)

	ctx.JSON(200, common.Success(result))
}

func (f FileController) UploadAvatarFile(ctx *gin.Context) {
	var result = f.service.UploadAvatarFile(ctx)

	ctx.JSON(200, common.Success(result))
}

func (f FileController) UploadOtherFile(ctx *gin.Context) {
	var result = f.service.UploadFile(ctx)

	ctx.JSON(200, common.Success(result))
}
func NewFileController(service service.FileService) FileController {
	return FileController{service: service}
}
