package GrpcRoutes

import (
	ImMsgRpc "api-skeleton/grpc/Proto/imMsg"
	UserRpc "api-skeleton/grpc/Proto/user"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

//RegisterGrpcClient 统一管理注册http客户端路由
func RegisterGrpcClient(ctx context.Context, gwmux *runtime.ServeMux, addr string, opts []grpc.DialOption) {
	err := ImMsgRpc.RegisterImMsgServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	err = UserRpc.RegisterUserServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		panic(fmt.Sprintf("注册http客户端路由服务异常:%s", err))
	}
}
