package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string
	//Db         string
	DBAddress  string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("配置文件读取失败")
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")

}
func LoadDatabase(file *ini.File) {
	//Db = file.Section("database").Key("Db").MustString("mysql")
	DBAddress = file.Section("database").Key("DBAddress").MustString("localhost")
	DbHost = file.Section("database").Key("DbHost").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
