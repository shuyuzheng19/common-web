package common

import "time"

const (
	UserTokenKey       = "USER-TOKEN:"
	EmailCodeKey       = "EMAIL-CODE:"
	EmailCodeKeyExpire = time.Minute
	UserInfoKey        = "USER-INFO:"
	UserInfoKeyExpire  = time.Minute * 30
)
