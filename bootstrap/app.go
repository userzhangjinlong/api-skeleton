package bootstrap

import (
	"api-skeleton/app/Http/Middleware"
	"api-skeleton/config"
	Route "api-skeleton/routes"
	"github.com/gin-gonic/gin"
)

var configs = config.InitConfig

type Server struct {
}

func (s *Server) Start() {
	//debug环境设置
	//gin.SetMode(gin.ReleaseMode)

	//引擎启动
	engine := gin.Default()

	//全局中间件注入
	engine.Use(Middleware.Cors())
	engine.Use(Middleware.Translations())

	//注入路由
	engine = Route.RegisterRoutes(engine)

	//设置受信任代理,如果不设置默认信任所有代理，不安全
	engine.SetTrustedProxies([]string{configs.Proxy.TrustProxy})

	//启动引擎
	engine.Run(configs.Proxy.Port)
}
