package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	//Db         string
	DBAddress  string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("配置文件读取失败")
	}
	LoadServer(file)
	LoadDatabase(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")

}
func LoadDatabase(file *ini.File) {
	//Db = file.Section("database").Key("Db").MustString("mysql")
	DBAddress = file.Section("database").Key("DBAddress").MustString("localhost")
	DbHost = file.Section("database").Key("DbHost").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
