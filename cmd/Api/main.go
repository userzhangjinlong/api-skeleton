package main

import (
	app "api-skeleton/bootstrap"
)

var App *app.Server

func main() {
	App.Start()
}
