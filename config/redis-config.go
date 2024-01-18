package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
	MaxSize  int    `yaml:"max_size" json:"maxSize"`
	MinIdle  int    `yaml:"min_idle" json:"minIdle"`
	Timeout  int    `yaml:"timeout"`
}

var REDIS *redis.Client

func LoadRedis() {
	var redisConfig = CONFIG.Redis
	REDIS = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		DB:           redisConfig.Db,
		Password:     redisConfig.Password,
		PoolSize:     redisConfig.MaxSize,
		MinIdleConns: redisConfig.MinIdle,
		PoolTimeout:  time.Duration(redisConfig.Timeout) * time.Second,
	})
}
