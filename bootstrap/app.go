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

	//初始化一些全局引擎
	InitConfig()      //初始化配置
	InitDB()          //初始化DB
	InitRedisClient() //redis客户端
	//InitRedisClusterClient() //redis集群客户端
	InitTracer()      //全链路追踪
	InitNsqProducer() //初始化nsq

	//引擎启动
	engine := gin.Default()

	//全局中间件注入
	engine.Use(Middleware.Cors())
	engine.Use(Middleware.Tracing())
	engine.Use(Middleware.Translations())
	engine.Use(Middleware.AccessLog())

	//注入路由
	engine = Route.RegisterRoutes(engine)

	//设置受信任代理,如果不设置默认信任所有代理，不安全
	engine.SetTrustedProxies([]string{configs.Proxy.TrustProxy})

	//todo::后期理解到位grpc实现的http接口路由注入到gin里面（貌似可以直接一个goroutine注入到里面启动）

	//启动gin http引擎 websocket和http服务是同时存在的 直接由主http管理启动或者开启一个后台管理的goroutine启动管理路由之类的
	engine.Run(configs.Proxy.Port)
}
