package request

import (
	"common-web-framework/models"
	"common-web-framework/utils"
)

type UserRequest struct {
	Username string `json:"username" validate:"required,min=8,max=16"`
	Password string `json:"password" validate:"required,min=8,max=16"`
	Email    string `json:"email" validate:"required,email"`
	NickName string `json:"nick_name" validate:"required,max=50,min=1"`
	Code     string `json:"code" validate:"required,min=6,max=6"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r UserRequest) ToUserDo() models.User {
	return models.User{
		Username: r.Username,
		Password: utils.EncryptPassword(r.Password),
		Email:    r.Email,
		NickName: r.NickName,
		RoleId:   1,
	}
}
