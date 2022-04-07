package GrpcRoutes

import (
	ImMsgRpc "api-skeleton/grpc/Proto/imMsg"
	UserRpc "api-skeleton/grpc/Proto/user"
	"api-skeleton/grpc/Service/ImMsgRpcService"
	"api-skeleton/grpc/Service/UserRpcService"
	"google.golang.org/grpc"
)

//RegisterGrpcServer 注册所有的rpc http2.0服务层路由接口内网rpc远程过程调用接口
func RegisterGrpcServer(s *grpc.Server) {
	//注册rpc服务端路由
	UserRpc.RegisterUserServiceServer(s, new(UserRpcService.UserServiceServer))
	ImMsgRpc.RegisterImMsgServiceServer(s, new(ImMsgRpcService.ImMsgServiceServer))
}
