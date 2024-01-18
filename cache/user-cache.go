package cache

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

type UserCache struct {
	redis *redis.Client
}

func (u UserCache) SetToken(id int, token string) error {
	return u.redis.Set(fmt.Sprintf(common.UserTokenKey+"%d", id), token, common.TokenExpire).Err()
}

func (u UserCache) GetToken(id int) string {
	return u.redis.Get(fmt.Sprintf(common.UserTokenKey+"%d", id)).Val()
}

func (u UserCache) SetUser(user *models.User) error {
	var buff, _ = json.Marshal(&user)
	return u.redis.Set(fmt.Sprintf(common.UserInfoKey+"%d", user.Id), string(buff), common.UserInfoKeyExpire).Err()
}

func (u UserCache) GetUser(id int) (user *models.User) {
	var result, _ = u.redis.Get(fmt.Sprintf(common.UserTokenKey+"%d", id)).Bytes()
	json.Unmarshal(result, &user)
	return user
}

func (u UserCache) SetEmailCode(code, email string) error {
	return u.redis.Set(fmt.Sprintf(common.EmailCodeKey+"%s", email), code, common.EmailCodeKeyExpire).Err()
}

func (u UserCache) GetEmailCode(email string) string {
	return u.redis.Get(fmt.Sprintf(common.EmailCodeKey+"%s", email)).Val()
}

func NewUserCache(r *redis.Client) UserCache {
	return UserCache{redis: r}
}
