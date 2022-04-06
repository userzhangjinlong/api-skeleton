package main

import (
	"api-skeleton/app/Global"
	"api-skeleton/bootstrap"
	ImMsgRpc "api-skeleton/grpc/Proto/imMsg"
	UserRpc "api-skeleton/grpc/Proto/user"
	"api-skeleton/grpc/Service/ImMsgRpcService"
	"api-skeleton/grpc/Service/UserRpcService"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"strings"
)

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	bootstrap.InitConfig()
	runServer(Global.Configs.HttpGrpc.Port)
}

func runServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return serveMux
}

func runGrpcGatewayServer() *runtime.ServeMux {
	ctx := context.Background()
	endpoint := fmt.Sprintf("%s:%s", Global.Configs.HttpGrpc.Host, Global.Configs.HttpGrpc.Port)
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	//绑定grpc http路由
	err := ImMsgRpc.RegisterImMsgServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
	err = UserRpc.RegisterUserServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
	if err != nil {
		panic(fmt.Sprintf("启动http网关异常:%s", err))
	}

	return gwmux
}

func runGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	//pb.RegisterTagServiceServer(s, server.NewTagServer())
	UserRpc.RegisterUserServiceServer(s, new(UserRpcService.UserServiceServer))
	ImMsgRpc.RegisterImMsgServiceServer(s, new(ImMsgRpcService.ImMsgServiceServer))
	reflection.Register(s)

	return s
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
