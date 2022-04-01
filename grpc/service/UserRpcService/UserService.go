package UserRpcService

import (
	"api-skeleton/app/Global"
	UserRpc "api-skeleton/grpc/Proto/user"
	"context"
	"errors"
)

type UserServiceServer struct {
}

//todo:: 实现rpc服务层
func (u *UserServiceServer) CreateUser(ctx context.Context, request *UserRpc.CreateUserRequest) (res *UserRpc.CreateUserResponse, err error) {
	Global.DB.DB()
	return nil, errors.New("ceshi")
}

func (u *UserServiceServer) UpdateUser(ctx context.Context, req *UserRpc.UpdateUserRequest) (res *UserRpc.UpdateUserResponse, err error) {
	return nil, errors.New("ceshi")
}

func (u *UserServiceServer) GetUser(ctx context.Context, req *UserRpc.GetUserRequest) (res *UserRpc.GetUserResponse, err error) {
	return nil, errors.New("te")
}

func (u *UserServiceServer) DelUser(ctx context.Context, req *UserRpc.DelUserRequest) (res *UserRpc.DelUserResponse, err error) {
	return nil, errors.New("t12")
}
