package main

import (
	"api-skeleton/config"
	"fmt"
)

func main() {
	configClass := config.InitConfig{}
	conf := configClass.GetConfig()
	fmt.Println(conf)
}
