package bootstrap

import (
	"api-skeleton/config"
	"github.com/gin-gonic/gin"
)

var configs = config.InitConfig

type Server struct {
}

func (s *Server) Start() {
	gin.Default().Run(configs.Proxy.Port)
}
