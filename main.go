package main

import (
	"common-web-framework/config"
	"common-web-framework/middle"
	"common-web-framework/route"
	"common-web-framework/utils"
	"os"
)

func AutoCreateUploadPath() {
	os.MkdirAll(config.CONFIG.Upload.Path, os.ModePerm)
}

func LoadConfig() {
	config.LoadGlobalConfig()

	config.LoadDBConfig()

	config.LoadLogger()

	config.LoadRedis()

	utils.LoadIpDB()

	AutoCreateUploadPath()
}

func main() {

	LoadConfig()

	var serverConfig = config.CONFIG.Server

	route.NewRouter(serverConfig).
		AddMiddles(middle.LoggerMiddleware).
		LoadUserController().
		LoadFileController().
		RunServer()

}
