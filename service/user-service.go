package service

import (
	"common-web-framework/cache"
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/models"
	"common-web-framework/repository"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/utils"
)

type UserService interface {
	// RegisteredUser 注册用户
	RegisteredUser(request request.UserRequest)
	// SendCodeToEmail 发送验证码到用户邮箱
	SendCodeToEmail(email string)
	// ValidateEmailCode 验证验证码是否正确
	ValidateEmailCode(email string, code string)
	// Login 登录
	Login(request request.LoginRequest) response.TokenResponse
	// GetUser 通过ID获取用户
	GetUser(id int) *models.User
}

var maps = make(map[string]string)

type UserServiceImpl struct {
	repository repository.UserRepository
	cache      cache.UserCache
}

func (u UserServiceImpl) GetUser(id int) *models.User {
	if user := u.cache.GetUser(id); user == nil {
		var dbUser = u.repository.FindById(id)
		u.cache.SetUser(dbUser)
		return dbUser
	} else {
		return user
	}
}

func (u UserServiceImpl) Login(request request.LoginRequest) response.TokenResponse {
	config.ValidateError(request)

	var encodingPassword = utils.EncryptPassword(request.Password)

	var user = u.repository.FindByUsernameAndPassword(request.Username, encodingPassword)

	if user == nil {
		helper.ErrorCommonF(common.AutoFail(common.LoginFail))
	}

	var token = utils.CreateAccessToken(user.Id, user.Username)

	u.cache.SetToken(user.Id, token.Token)

	return token
}

func (u UserServiceImpl) ValidateEmailCode(email string, code string) {

	var cacheCode = u.cache.GetEmailCode(email)

	if cacheCode == "" || cacheCode != code {
		helper.ErrorCommonF(common.AutoFail(common.EmailCodeValidate))
	}
}

func (u UserServiceImpl) SendCodeToEmail(email string) {

	var code = utils.RandomNumberCode()

	config.CONFIG.Email.SendEmail(email, "注册验证码", false, code)

	u.cache.SetEmailCode(code, email)
}

func (u UserServiceImpl) RegisteredUser(request request.UserRequest) {
	config.ValidateError(request)

	u.ValidateEmailCode(request.Email, request.Code)

	var user = request.ToUserDo()

	if err := u.repository.Save(user); err != nil {
		helper.ErrorCommonF(common.AutoFail(common.RegisteredCode))
	}
}

func NewUserService() UserService {
	var repository = repository.NewUserRepository(config.DB)
	var cache = cache.NewUserCache(config.REDIS)
	return UserServiceImpl{repository: repository, cache: cache}
}
