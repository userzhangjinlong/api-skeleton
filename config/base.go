package config

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var InitConfig = new(System)

func init() {
	dir, _ := os.Getwd()
	file, err := ioutil.ReadFile(dir + "/env.yaml")
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal(file, InitConfig)

	if err != nil {
		return
	}

	//init logrus日志配置组件
	logrus.SetFormatter(&logrus.JSONFormatter{})
	filePath := dir + InitConfig.Log.Path
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 天轮转一个新文件，保留最近 15 天半个月的日志文件，多余的自动清理掉。
	writer, _ := rotatelogs.New(
		filePath+"%Y%m%d%H%M"+".log",
		rotatelogs.WithLinkName(filePath+"/logs"),
		rotatelogs.WithMaxAge(time.Hour*24*15),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	logrus.SetOutput(writer)

}
