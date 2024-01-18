package route

import (
	"common-web-framework/config"
	"common-web-framework/controller"
	"common-web-framework/helper"
	"common-web-framework/middle"
	"common-web-framework/models"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Router struct {
	//服务器监听地址
	addr string
	//gin服务
	service *gin.Engine
	//gin路由分组
	apiGroup *gin.RouterGroup
	//读取超时时间
	readTimeOut time.Duration
	//写入超时时间
	writeTimeOut time.Duration
	//获取用户
	getUser func(id int) *models.User
}

// RunServer 运行Server
func (r Router) RunServer() {

	var server = &http.Server{
		Addr:         r.addr,
		Handler:      r.service,
		ReadTimeout:  r.readTimeOut,
		WriteTimeout: r.writeTimeOut,
	}

	r.service.Static("static", config.CONFIG.Upload.Path)

	var err = server.ListenAndServe()

	helper.ErrorPanic(err)

}

func (r *Router) AddMiddles(f ...gin.HandlerFunc) *Router {
	r.apiGroup.Use(f...)
	return r
}

func (r *Router) LoadFileController() *Router {
	var group = r.apiGroup.Group("file")

	var fileService = service.NewFileService()

	var fileController = controller.NewFileController(fileService)

	{
		group.POST("upload/avatar", fileController.UploadAvatarFile)
		group.POST("upload/image", middle.JwtMiddle(middle.AdminRole, r.getUser), fileController.UploadImageFile)
		group.POST("upload/file", middle.JwtMiddle(middle.AdminRole, r.getUser), fileController.UploadOtherFile)
	}

	return r
}

func (r *Router) LoadUserController() *Router {

	var group = r.apiGroup.Group("user")

	var userService = service.NewUserService()

	r.getUser = userService.GetUser

	var userController = controller.NewUserController(userService)

	{
		group.POST("registered", userController.RegisteredUser)
		group.POST("login", userController.Login)
		group.GET("send_email", userController.SendEmail)
		group.GET("get", middle.JwtMiddle(middle.UserRole, r.getUser), userController.GetUserInfo)
	}

	return r
}

func NewRouter(config config.ServerConfig) *Router {

	var service = gin.Default()

	var group = service.Group(config.ApiPrefix)

	return &Router{
		addr:         config.Addr,
		service:      service,
		apiGroup:     group,
		writeTimeOut: time.Second * time.Duration(config.WriteTimeOut),
		readTimeOut:  time.Second * time.Duration(config.ReadTimeOut),
	}
}
