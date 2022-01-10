package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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
}
