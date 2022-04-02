package UserRpcService

import (
	"api-skeleton/app/Global"
	UserRpc "api-skeleton/grpc/Proto/user"
	"context"
	"errors"
	"fmt"
)

type UserServiceServer struct {
}

//CreateUser Grpc 内部服务调用创建用户
func (u *UserServiceServer) CreateUser(ctx context.Context, req *UserRpc.CreateUserRequest) (res *UserRpc.CreateUserResponse, err error) {
	fmt.Println(&req.User)
	result := Global.DB.Create(&req.User)
	if result.Error != nil {
		return nil, result.Error
	}
	res.Id = req.User.Id
	return res, nil
}

func (u *UserServiceServer) UpdateUser(ctx context.Context, req *UserRpc.UpdateUserRequest) (res *UserRpc.UpdateUserResponse, err error) {
	return nil, errors.New("ceshi")
}

//GetUser rpc实现服务层逻辑
func (u *UserServiceServer) GetUser(ctx context.Context, req *UserRpc.GetUserRequest) (res *UserRpc.GetUserResponse, err error) {
	var resultUser UserRpc.User
	result := Global.DB.Table("user").Where("username = ? and password = ?", req.Username, req.Password).Find(&resultUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &UserRpc.GetUserResponse{
		User: &resultUser,
	}, nil
}

func (u *UserServiceServer) DelUser(ctx context.Context, req *UserRpc.DelUserRequest) (res *UserRpc.DelUserResponse, err error) {
	return nil, errors.New("t12")
}
