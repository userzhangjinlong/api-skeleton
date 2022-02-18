package bootstrap

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Middleware"
	"api-skeleton/app/Util"
	"api-skeleton/config"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	Route "api-skeleton/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

var configs = config.InitConfig

type Server struct {
}

func (s *Server) Start() {
	//debug环境设置
	//gin.SetMode(gin.ReleaseMode)

	//初始化一些全局引擎
	initRedisClient() //redis客户端
	initTracer()      //全链路追踪

	//引擎启动
	engine := gin.Default()

	//全局中间件注入
	engine.Use(Middleware.Cors())
	engine.Use(Middleware.Tracing())
	engine.Use(Middleware.Translations())
	engine.Use(Middleware.AccessLog())
	engine.Use(Middleware.Auth())

	//注入路由
	engine = Route.RegisterRoutes(engine)

	//设置受信任代理,如果不设置默认信任所有代理，不安全
	engine.SetTrustedProxies([]string{configs.Proxy.TrustProxy})

	//启动引擎
	engine.Run(configs.Proxy.Port)
}

func initRedisClient() {
	//初始化设置全局变量
	redisPool, _ := ConnectPoolFactory.GetRedis()
	Global.RedisClient = redisPool.Get()
}

func initTracer() {
	jaegerTracer, _, err := Util.NewJaegerTracer(
		configs.Trace.Servicename,
		fmt.Sprintf("%s%s", configs.Trace.Agenthost, configs.Trace.Port),
	)

	if err != nil {
		//todo::异常日志记录链接tracer链路追综异常
		return
	}
	Global.Tracer = jaegerTracer
	return
}
