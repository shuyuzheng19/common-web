package controller

import (
	"common-web-framework/common"
	"common-web-framework/request"
	"common-web-framework/service"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func (u UserController) SendEmail(ctx *gin.Context) {

	var email = ctx.Query("email")

	if email == "" {
		ctx.JSON(200, common.AutoFail(common.BadRequestCode))
	}

	u.service.SendCodeToEmail(email)

	ctx.JSON(200, common.OK())

}

func (u UserController) RegisteredUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	ctx.ShouldBindJSON(&userRequest)

	u.service.RegisteredUser(userRequest)

	ctx.JSON(200, common.OK())
}

func (u UserController) GetUserInfo(ctx *gin.Context) {
	ctx.JSON(200, common.Success(utils.GetUserInfo(ctx).ToUserResponse()))
}

func (u UserController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest

	ctx.ShouldBindJSON(&loginRequest)

	var tokenResponse = u.service.Login(loginRequest)

	ctx.JSON(200, common.Success(tokenResponse))
}

func NewUserController(service service.UserService) UserController {
	return UserController{service: service}
}
