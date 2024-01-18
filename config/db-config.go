package config

import (
	"common-web-framework/helper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

var DB *gorm.DB

// DbConfig 关系型数据库配置
type DbConfig struct {
	//数据库时区
	Timezone string `yaml:"timezone" json:"timezone"`
	//数据库厂商 如mysql postgresql等 仅限gorm支持的数据库
	Database string `yaml:"database" json:"database"`
	//数据库远程地址
	Host string `yaml:"host" json:"host"`
	//数据库端口
	Port int `yaml:"port" json:"port"`
	//数据库登录账号
	Username string `yaml:"username" json:"username"`
	//数据库密码
	Password string `yaml:"password" json:"password"`
	//数据库名称
	Dbname string `yaml:"dbname" json:"dbname"`
}

// getDSN 根据 DbConfig 返回数据库连接字符串
func getDataBaseDSN(config DbConfig) string {
	var database = strings.ToLower(config.Database)
	fmt.Println(database)
	switch database {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "postgresql":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Host, config.Port, config.Username, config.Dbname, config.Password)
	case "sqlite":
		return config.Dbname // Assuming that Dbname is the path to SQLite file
	case "sqlserver":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "oracle":
		return fmt.Sprintf("%s/%s@%s:%d/%s",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "cockroachdb":
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			config.Username, config.Password, config.Host, config.Port, config.Dbname)
	case "clickhouse":
		return fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s",
			config.Host, config.Port, config.Username, config.Password, config.Dbname)
	case "bigquery":
		return fmt.Sprintf("bigquery://%s:%s@projectID:%s/datasetID",
			config.Username, config.Password, config.Dbname)
	default:
		var err interface{} = fmt.Sprintf("未知的数据库 %s", config.Database)
		panic(err)
	}
}

func LoadDBConfig() {
	//数据库链接
	var dsn = getDataBaseDSN(CONFIG.Db)
	//假设是mysql 当然也可以使用其他的数据，只需导入相应的数据库驱动即可
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//判断是否链接成功
	helper.ErrorPanic(err)
	//如果连接成功 打印 DataBase Connection SUCCESS
	log.Println("DataBase Connection SUCCESS")
}
