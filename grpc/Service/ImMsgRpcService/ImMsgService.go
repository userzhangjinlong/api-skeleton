package ImMsgRpcService

import (
	ImMsgRpc "api-skeleton/grpc/Proto/imMsg"
	"context"
	"errors"
	"fmt"
)

type ImMsgServiceServer struct {
}

func (i *ImMsgServiceServer) CreateMsg(ctx context.Context, req *ImMsgRpc.CreateMsgRequest) (res *ImMsgRpc.CreateMsgResponse, err error) {
	return nil, errors.New("test")
}

func (i *ImMsgServiceServer) GetMsg(ctx context.Context, req *ImMsgRpc.GetMsgRequest) (res *ImMsgRpc.GetMsgResponse, err error) {
	fmt.Printf("查看走到这里的ctx是什么：%v", ctx)
	return nil, errors.New("test")
}
