package models

import (
	"common-web-framework/response"
	"gorm.io/gorm"
)

// gorm生成的用户表名
const userTableName = "users"

// User 用户表模型
type User struct {
	//自动生成创建时间 修改时间 删除时间
	gorm.Model
	//user 主键ID
	Id int `gorm:"primaryKey"`
	//用户名称 不能为null 唯一 最大长度为16个字符
	Username string `gorm:"column:username;type:varchar(16);not null;unique"`
	//用户密码 不能为null 最大长度16个字符 使用哈希算法加密存储。
	Password string `gorm:"column:password;not null"`
	//用户邮箱 唯一 不能为null 用于找回密码和接收通知。
	Email string `gorm:"column:email;not null"`
	//用户头像，可以是图片地址的URL
	Avatar string `gorm:"column:avatar;default:test.png"`
	//用户名称 不能为空且不能超过50个字符
	NickName string `gorm:"column:nick_name;not null;type:varchar(50)"`
	//用户角色ID
	RoleId int `gorm:"column:role_id"`
	//用户包含的角色
	Role Role
}

func (*User) TableName() string { return userTableName }

func (u User) ToUserResponse() response.UserResponse {
	return response.UserResponse{
		Nickname: u.NickName,
		Avatar:   u.Avatar,
		Role:     u.Role.Name,
	}
}
