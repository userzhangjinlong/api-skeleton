package Route

import (
	"api-skeleton/app/Http/Middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"strings"
)

func RegisterRoutes(router *gin.Engine) *gin.Engine {
	// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
	pprof.Register(router)

	webRoute := setWebRoute()

	for group, routes := range webRoute {
		//字符串分割
		groupAndAuth := strings.Split(group, "-")

		var routeGroup *gin.RouterGroup
		if groupAndAuth[1] == "login" {
			routeGroup = router.Group(groupAndAuth[0])
			routeGroup.Use(Middleware.Auth())
		} else {
			routeGroup = router.Group(groupAndAuth[0])
		}
		for i := 0; i < len(routes); i++ {
			switch routes[i].Method {
			case MethodGet:
				routeGroup.GET(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
			case MethodPost:
				routeGroup.POST(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
			case MethodPut:
				routeGroup.PUT(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
			case MethodDelete:
				routeGroup.DELETE(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
			case MethodAny:
				routeGroup.Any(routes[i].Pattern, routes[i].Callback.(func(context *gin.Context)))
			}
		}
	}

	return router
}
