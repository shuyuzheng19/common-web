package config

import (
	"common-web-framework/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// LoggerConfig 日志配置
type LoggerConfig struct {
	Dev         bool   `yaml:"dev" json:"dev"`
	Encoding    string `yaml:"encoding" json:"encoding"`
	OutputPaths string `yaml:"outputPaths" json:"outputPaths"`
	ErrorPaths  string `yaml:"errorPaths" json:"errorPaths"`
	Level       string `yaml:"level" json:"level"`
}

var LOGGER *zap.Logger

func LoadLogger() {

	config := zap.NewProductionConfig()

	var loggerConfig = CONFIG.Logger

	config.Encoding = loggerConfig.Encoding

	config.Development = loggerConfig.Dev

	config.OutputPaths = strings.Split(loggerConfig.OutputPaths, ",")

	config.ErrorOutputPaths = strings.Split(loggerConfig.ErrorPaths, ",")

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var logger, err = config.Build()

	helper.ErrorPanicAndMessage(err, "加载日志配置失败")

	LOGGER = logger

	defer LOGGER.Sync()
}
