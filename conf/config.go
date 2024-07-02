package conf

import (
	"im/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	// 服务器设置相关
	AppMode  string
	HttpPort string
	// 数据库配置相关
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// 静态路径相关
	Host       string
	AvatarPath string

	// 邮箱相关
	SmtpEmail string
	SmtpHost  string
	SmtpPort  int
	SmtpPass  string
)

// 读取配置文件
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMySql(file)
	LoadPath(file)
	LoadEmail(file)
	// mysql读路径
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	// mysql写路径
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadMySql(f *ini.File) {
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}

func LoadServer(f *ini.File) {
	AppMode = f.Section("service").Key("AppMode").String()
	HttpPort = f.Section("service").Key("HttpPort").String()
}

func LoadPath(f *ini.File) {
	Host = f.Section("path").Key("Host").String()
	AvatarPath = f.Section("path").Key("AvatarPath").String()
}

func LoadEmail(f *ini.File) {
	SmtpHost = f.Section("email").Key("SmtpHost").String()
	SmtpEmail = f.Section("email").Key("SmtpEmail").String()
	SmtpPort = f.Section("email").Key("SmtpPort").MustInt()
	SmtpPass = f.Section("email").Key("SmtpPass").String()
}
