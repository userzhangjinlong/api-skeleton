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
	fmt.Println("我接通啦")
	var imRes = make([]*ImMsgRpc.ImMsg, 0)
	imRes = append(imRes, &ImMsgRpc.ImMsg{
		Id:         1,
		ToUserId:   2,
		FormUserId: 3,
		Content:    "23e234",
		MsgImg:     "4355234",
		SendTime:   123,
		CreateTime: 234,
		UpdateTime: 67768,
	})
	return &ImMsgRpc.GetMsgResponse{ImMsg: imRes, PageNum: 1, PageSize: 10}, nil
}
