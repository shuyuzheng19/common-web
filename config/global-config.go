package config

import (
	"common-web-framework/helper"
	"gopkg.in/yaml.v3"
	"os"
)

// CONFIG 全局配置
var CONFIG *GlobalConfig

// ServerConfig 服务器配置
type ServerConfig struct {
	Addr         string `yaml:"addr" json:"addr"`
	ApiPrefix    string `yaml:"apiPrefix" json:"apiPrefix"`
	ReadTimeOut  int    `yaml:"readTimeOut" json:"readTimeOut"`
	WriteTimeOut int    `yaml:"writeTimeOut" json:"writeTimeOut"`
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	//IP数据库路径
	IpDbPath string `yaml:"ipDbPath" json:"ipDbPath"`
	//server配置
	Server ServerConfig `yaml:"server" json:"server"`
	//db配置
	Db DbConfig `yaml:"db" json:"db"`
	//邮箱配置
	Email EmailConfig `yaml:"email" json:"email"`
	//日志配置
	Logger LoggerConfig `yaml:"logger" json:"logger"`
	//redis配置
	Redis RedisConfig `yaml:"redis" json:"redis"`
	//上传文件配置
	Upload UploadConfig `yaml:"upload" json:"upload"`
}

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig() {
	var file, err = os.ReadFile("application.yml")

	helper.ErrorPanic(err)

	yaml.Unmarshal(file, &CONFIG)
}
