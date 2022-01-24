package bootstrap

import (
	"api-skeleton/config"
	Route "api-skeleton/routes"
)

var configs = config.InitConfig

type Server struct {
}

func (s *Server) Start() {
	//注入路由
	engine := Route.RegisterRoutes()
	//启动引擎
	engine.Run(configs.Proxy.Port)
}
