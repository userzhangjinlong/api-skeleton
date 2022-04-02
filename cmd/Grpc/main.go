package main

import (
	"api-skeleton/app/Global"
	"api-skeleton/bootstrap"
	ImMsgRpc "api-skeleton/grpc/Proto/imMsg"
	UserRpc "api-skeleton/grpc/Proto/user"
	"api-skeleton/grpc/Service/ImMsgRpcService"
	"api-skeleton/grpc/Service/UserRpcService"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

/**
grpc 服务端
*/
func main() {
	bootstrap.InitConfig()
	bootstrap.InitDB()

	grpcServer := grpc.NewServer()
	//统一注册多个grpc服务
	registerRpcServer(grpcServer)

	lis, err := net.Listen("tcp", ":"+Global.Configs.Grpc.Port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	//在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(grpcServer)
	fmt.Println("启动grpc服务")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}

}

//registerRpcServer 注册所有的rpc服务
func registerRpcServer(s *grpc.Server) {
	//用户rpc服务
	UserRpc.RegisterUserServiceServer(s, new(UserRpcService.UserServiceServer))
	ImMsgRpc.RegisterImMsgServiceServer(s, new(ImMsgRpcService.ImMsgServiceServer))
}
