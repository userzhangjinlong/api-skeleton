package Route

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
	pprof.Register(router)

	webRoute := setWebRoute()

	for group, routes := range webRoute {
		group := router.Group(group)
		{
			for i := 0; i < len(routes); i++ {
				switch routes[i].Method {
				case MethodGet:
					//这里后续写增加回调方法的工厂方法调用指定位置的回调方法
					if routes[i].Middleware == nil {
						group.GET(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					} else {
						group.Use(routes[i].Middleware.(gin.HandlerFunc)).
							GET(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					}
				case MethodPost:
					if routes[i].Middleware == nil {
						group.POST(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					} else {
						group.Use(routes[i].Middleware.(gin.HandlerFunc)).
							POST(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					}
				case MethodPut:
					if routes[i].Middleware == nil {
						group.PUT(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					} else {
						group.Use(routes[i].Middleware.(gin.HandlerFunc)).
							PUT(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					}
				case MethodDelete:
					if routes[i].Middleware != nil {
						group.DELETE(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					} else {
						group.Use(routes[i].Middleware.(gin.HandlerFunc)).
							DELETE(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					}
				case MethodAny:
					if routes[i].Middleware != nil {
						group.Any(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					} else {
						group.Use(routes[i].Middleware.(gin.HandlerFunc)).
							Any(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
					}
				}
			}
		}

	}

	return router
}
