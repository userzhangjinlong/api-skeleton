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
	//WithInsecure 禁用传输安全性tls认证
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, addr, opts...)
}
