package main

import (
	"api-skeleton/config"
	"fmt"
)

func main() {
	fmt.Println(config.InitConfig.App.Debug)
}
