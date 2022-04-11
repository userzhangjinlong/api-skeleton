package main

import (
	"api-skeleton/app/Global"
	"api-skeleton/bootstrap"
	"api-skeleton/grpc/GrpcRoutes"
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

/**
Grpc 服务端
*/
func main() {
	bootstrap.InitConfig()
	bootstrap.InitDB()

	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer()

	//dir, _ := os.Getwd()
	//severPemPath := path.Dir(dir+"/grpc/Keys/") + "/server.pem"
	//severKeyPath := path.Dir(dir+"/grpc/Keys/") + "/server.key"

	httpMux.Handle("/", gatewayMux)
	fmt.Println("启动grpc和http同端口双流量服务")
	err := http.ListenAndServe(":"+Global.Configs.Grpc.Port, grpcHandlerFunc(grpcS, httpMux))
	//监听tls
	//err := http.ListenAndServeTLS(":"+Global.Configs.Grpc.Port, severPemPath, severKeyPath, grpcHandlerFunc(grpcS, httpMux))
	if err != nil {
		panic(fmt.Sprintf("启动grpc http同端口双流量服务异常：%s", err))
	}
}

//runGrpcServer 启动注册grpc服务
func runGrpcServer() *grpc.Server {
	//Grpc 新增Tls配置
	//tls与ssl 一样需要配置读取证书路径，本地签发无效当前方式可以链接
	//dir, _ := os.Getwd()
	//severPemPath := path.Dir(dir+"/grpc/Keys/") + "/server.pem"
	//severKeyPath := path.Dir(dir+"/grpc/Keys/") + "/server.key"
	//cred, err := credentials.
	//	NewServerTLSFromFile(severPemPath, severKeyPath)
	//if err != nil {
	//	panic(fmt.Sprintf("配置Tls启动rpc 服务端异常：%s", err))
	//}

	//grpcServer := grpc.NewServer(grpc.Creds(cred))
	grpcServer := grpc.NewServer()
	GrpcRoutes.RegisterGrpcServer(grpcServer)
	reflection.Register(grpcServer)

	return grpcServer
}

//runHttpServer 注册http1.0客户端服务
func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return serveMux
}

//runGrpcGatewayServer 注册grpc基于http1.0网关服务
func runGrpcGatewayServer() *runtime.ServeMux {
	ctx := context.Background()
	endpoint := fmt.Sprintf("%s:%s", Global.Configs.Grpc.Host, Global.Configs.Grpc.Port)
	gwmux := runtime.NewServeMux()

	//dir, _ := os.Getwd()
	//severPemPath := path.Dir(dir+"/grpc/Keys/") + "/server.pem"
	////serverName 生产pem文件时定义的CommonName
	//cred, err := credentials.
	//	NewClientTLSFromFile(severPemPath, "ApiSkeleton")
	//if err != nil {
	//	panic(fmt.Sprintf("配置Tls启动rpc 客户端异常：%s", err))
	//}

	dopts := []grpc.DialOption{
		grpc.WithInsecure(), //忽略tls认证
		//grpc.WithTransportCredentials(cred),//开启tls认证
	}
	//绑定grpc http路由
	GrpcRoutes.RegisterGrpcClient(ctx, gwmux, endpoint, dopts)

	return gwmux
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
