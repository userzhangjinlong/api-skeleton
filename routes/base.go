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
					group.GET(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
				case MethodPost:
					group.POST(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
				case MethodPut:
					group.PUT(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
				case MethodDelete:
					group.DELETE(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
				case MethodAny:
					group.Any(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
				}
			}
		}

	}

	return router
}
