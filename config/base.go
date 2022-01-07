package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config interface {
	GetConfig() map[string]interface{}
}

type InitConfig struct {
	conf map[string]interface{}
}

func init() {
	dir, _ := os.Getwd()
	file, err := ioutil.ReadFile(dir + "/env.yaml")
	if err != nil {
		log.Panic(err)
	}
	var initConfig InitConfig
	initConfig.conf = make(map[string]interface{})
	err = yaml.Unmarshal(file, initConfig.conf)

	if err != nil {
		return
	}
}
