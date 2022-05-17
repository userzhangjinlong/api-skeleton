package wsRoutes

import (
	"api-skeleton/grateway"
	"fmt"
	"strings"
	"sync"
)

//ws路由回调执行方法
type HandlerFunc func(req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp)
type MiddlewareFunc func(req *grateway.WsMsgReq) HandlerFunc
type HandlersChain []HandlerFunc
type MiddlewareHandlersChain []MiddlewareFunc
type RouteHandlers map[string]HandlerFunc

type WsGroup struct {
	basePath           string
	MiddlewareHandlers MiddlewareHandlersChain
	RouteHandlers      RouteHandlers
	engine             *WsEngine
	mutex              sync.RWMutex
}

//路由组定义已经 组方法添加定义
func (wg *WsGroup) Group(groupPath string, middlewareHandlers ...MiddlewareFunc) *WsGroup {
	return &WsGroup{
		basePath:           groupPath,
		MiddlewareHandlers: wg.combineMiddlewareHandlers(middlewareHandlers),
	}
}

//ws 路由引入入的中间件
func (wg *WsGroup) Use(middleware ...MiddlewareFunc) {
	wg.MiddlewareHandlers = append(wg.MiddlewareHandlers, middleware...)
}

//AddRoute wg路由添加
func (wg *WsGroup) AddRoute(path string, handler HandlerFunc) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	routePath := wg.getWsRoutePath(path)
	if len(wg.RouteHandlers) > 0 {
		wg.RouteHandlers[routePath] = handler
	} else {
		wg.RouteHandlers = make(map[string]HandlerFunc, 1000)
		wg.RouteHandlers[routePath] = handler
	}

}

//Exec 执行响应的回调路由
func (wg *WsGroup) Exec(path string, req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp) {
	//todo::执行正常方法之前先执行中间件
	routePath := wg.getWsRoutePath(path)
	if handler, ok := wg.RouteHandlers[routePath]; ok {
		//正常运行handler方法
		handler(req, rsp)
	} else {
		//抛异常路由不存在
		panic(fmt.Sprintf("%s:方法不存在", routePath))
	}
}

//拼接ws路由组
func (wg *WsGroup) getWsRoutePath(groupPath string) string {
	tmpPathSlice := []string{wg.basePath, groupPath}
	return strings.Join(tmpPathSlice, ".")
}

//combineHandlers 合并中间件
func (wg *WsGroup) combineMiddlewareHandlers(handlers MiddlewareHandlersChain) MiddlewareHandlersChain {
	length := len(wg.MiddlewareHandlers) + len(handlers)
	mergeMiddlewareHandlerChain := make(MiddlewareHandlersChain, length)
	copy(mergeMiddlewareHandlerChain, wg.MiddlewareHandlers)
	copy(mergeMiddlewareHandlerChain[len(wg.MiddlewareHandlers):], handlers)
	return mergeMiddlewareHandlerChain
}
