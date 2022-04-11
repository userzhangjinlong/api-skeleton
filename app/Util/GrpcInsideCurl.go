package Util

import (
	"api-skeleton/app/Global"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

//GrpcClientConn inside内部调用获取grpc 客户端链接
//todo::单列的去实现创建客户端，防止无止境的开启客户端造成memory溢出 gorutinue等问题
func GrpcClientConn() (*grpc.ClientConn, error) {
	ctx := context.Background()
	clientConn, err := getClientCoon(ctx,
		fmt.Sprintf("%s:%s", Global.Configs.Grpc.Host, Global.Configs.Grpc.Port),
		nil)
	if err != nil {
		return nil, err
	}

	return clientConn, nil
}

//getClientCoon 获取地址封装
func getClientCoon(ctx context.Context, addr string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	//dir, _ := os.Getwd()
	//severPemPath := path.Dir(dir+"/grpc/Keys/") + "/server.pem"
	////serverName 生产pem文件是定义的服务名称
	//cred, err := credentials.
	//	NewClientTLSFromFile(severPemPath, addr)
	//if err != nil {
	//	panic(fmt.Sprintf("配置Tls启动rpc 客户端异常：%s", err))
	//}
	//WithInsecure 禁用传输安全性tls认证
	opts = append(opts, grpc.WithInsecure())
	//opts = append(opts, grpc.WithTransportCredentials(cred))
	return grpc.DialContext(ctx, addr, opts...)
}
