package utils

import (
	"common-web-framework/config"
	"common-web-framework/helper"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net/http"
	"strings"
)

func GetIPAddress(request *http.Request) string {
	ipAddress := request.Header.Get("X-Forwarded-For")
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("WL-Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.RemoteAddr
	}
	return ipAddress
}

var searcher *xdb.Searcher

func LoadIpDB() {
	var err error

	searcher, err = xdb.NewWithFileOnly(config.CONFIG.IpDbPath)

	if err != nil {
		helper.ErrorPanicAndMessage(err, "加载IP数据库失败")
	}
}

func GetIpCity(ip string) string {
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "未知"
	}

	var split = strings.Split(region, "|")

	return strings.ReplaceAll(split[0]+" "+split[2]+" "+split[3], "0", "")
}
