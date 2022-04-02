package Util

import (
	"api-skeleton/app/Global"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

//GrpcClientConn inside内部调用获取grpc 客户端链接
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
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, addr, opts...)
}
