// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imMsg.proto

package ImMsgRpc // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ImMsg struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	FormUserId           uint64   `protobuf:"varint,2,opt,name=FormUserId" json:"FormUserId,omitempty"`
	ToUserId             uint64   `protobuf:"varint,3,opt,name=ToUserId" json:"ToUserId,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
	MsgImg               string   `protobuf:"bytes,5,opt,name=MsgImg" json:"MsgImg,omitempty"`
	SendTime             uint64   `protobuf:"varint,6,opt,name=SendTime" json:"SendTime,omitempty"`
	CreateTime           uint64   `protobuf:"varint,7,opt,name=CreateTime" json:"CreateTime,omitempty"`
	UpdateTime           uint64   `protobuf:"varint,8,opt,name=UpdateTime" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImMsg) Reset()         { *m = ImMsg{} }
func (m *ImMsg) String() string { return proto.CompactTextString(m) }
func (*ImMsg) ProtoMessage()    {}
func (*ImMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_imMsg_9789ddd040263930, []int{0}
}
func (m *ImMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImMsg.Unmarshal(m, b)
}
func (m *ImMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImMsg.Marshal(b, m, deterministic)
}
func (dst *ImMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImMsg.Merge(dst, src)
}
func (m *ImMsg) XXX_Size() int {
	return xxx_messageInfo_ImMsg.Size(m)
}
func (m *ImMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ImMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ImMsg proto.InternalMessageInfo

func (m *ImMsg) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ImMsg) GetFormUserId() uint64 {
	if m != nil {
		return m.FormUserId
	}
	return 0
}

func (m *ImMsg) GetToUserId() uint64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *ImMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ImMsg) GetMsgImg() string {
	if m != nil {
		return m.MsgImg
	}
	return ""
}

func (m *ImMsg) GetSendTime() uint64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *ImMsg) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *ImMsg) GetUpdateTime() uint64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type CreateMsgRequest struct {
	ImMsg                *ImMsg   `protobuf:"bytes,1,opt,name=imMsg" json:"imMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMsgRequest) Reset()         { *m = CreateMsgRequest{} }
func (m *CreateMsgRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMsgRequest) ProtoMessage()    {}
func (*CreateMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_imMsg_9789ddd040263930, []int{1}
}
func (m *CreateMsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMsgRequest.Unmarshal(m, b)
}
func (m *CreateMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMsgRequest.Marshal(b, m, deterministic)
}
func (dst *CreateMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMsgRequest.Merge(dst, src)
}
func (m *CreateMsgRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMsgRequest.Size(m)
}
func (m *CreateMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMsgRequest proto.InternalMessageInfo

func (m *CreateMsgRequest) GetImMsg() *ImMsg {
	if m != nil {
		return m.ImMsg
	}
	return nil
}

type CreateMsgResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMsgResponse) Reset()         { *m = CreateMsgResponse{} }
func (m *CreateMsgResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMsgResponse) ProtoMessage()    {}
func (*CreateMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_imMsg_9789ddd040263930, []int{2}
}
func (m *CreateMsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMsgResponse.Unmarshal(m, b)
}
func (m *CreateMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMsgResponse.Marshal(b, m, deterministic)
}
func (dst *CreateMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMsgResponse.Merge(dst, src)
}
func (m *CreateMsgResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMsgResponse.Size(m)
}
func (m *CreateMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMsgResponse proto.InternalMessageInfo

func (m *CreateMsgResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetMsgRequest struct {
	FormUserId           uint64   `protobuf:"varint,1,opt,name=formUserId" json:"formUserId,omitempty"`
	ToUserId             uint64   `protobuf:"varint,2,opt,name=toUserId" json:"toUserId,omitempty"`
	PageSize             uint64   `protobuf:"varint,3,opt,name=pageSize" json:"pageSize,omitempty"`
	PageNum              uint64   `protobuf:"varint,4,opt,name=pageNum" json:"pageNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMsgRequest) Reset()         { *m = GetMsgRequest{} }
func (m *GetMsgRequest) String() string { return proto.CompactTextString(m) }
func (*GetMsgRequest) ProtoMessage()    {}
func (*GetMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_imMsg_9789ddd040263930, []int{3}
}
func (m *GetMsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMsgRequest.Unmarshal(m, b)
}
func (m *GetMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMsgRequest.Marshal(b, m, deterministic)
}
func (dst *GetMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMsgRequest.Merge(dst, src)
}
func (m *GetMsgRequest) XXX_Size() int {
	return xxx_messageInfo_GetMsgRequest.Size(m)
}
func (m *GetMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMsgRequest proto.InternalMessageInfo

func (m *GetMsgRequest) GetFormUserId() uint64 {
	if m != nil {
		return m.FormUserId
	}
	return 0
}

func (m *GetMsgRequest) GetToUserId() uint64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *GetMsgRequest) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetMsgRequest) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetMsgResponse struct {
	PageSize             uint64   `protobuf:"varint,1,opt,name=pageSize" json:"pageSize,omitempty"`
	PageNum              uint64   `protobuf:"varint,2,opt,name=pageNum" json:"pageNum,omitempty"`
	ImMsg                []*ImMsg `protobuf:"bytes,3,rep,name=imMsg" json:"imMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMsgResponse) Reset()         { *m = GetMsgResponse{} }
func (m *GetMsgResponse) String() string { return proto.CompactTextString(m) }
func (*GetMsgResponse) ProtoMessage()    {}
func (*GetMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_imMsg_9789ddd040263930, []int{4}
}
func (m *GetMsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMsgResponse.Unmarshal(m, b)
}
func (m *GetMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMsgResponse.Marshal(b, m, deterministic)
}
func (dst *GetMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMsgResponse.Merge(dst, src)
}
func (m *GetMsgResponse) XXX_Size() int {
	return xxx_messageInfo_GetMsgResponse.Size(m)
}
func (m *GetMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMsgResponse proto.InternalMessageInfo

func (m *GetMsgResponse) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetMsgResponse) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *GetMsgResponse) GetImMsg() []*ImMsg {
	if m != nil {
		return m.ImMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*ImMsg)(nil), "ImMsg.ImMsg")
	proto.RegisterType((*CreateMsgRequest)(nil), "ImMsg.CreateMsgRequest")
	proto.RegisterType((*CreateMsgResponse)(nil), "ImMsg.CreateMsgResponse")
	proto.RegisterType((*GetMsgRequest)(nil), "ImMsg.GetMsgRequest")
	proto.RegisterType((*GetMsgResponse)(nil), "ImMsg.GetMsgResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ImMsgService service

type ImMsgServiceClient interface {
	CreateMsg(ctx context.Context, in *CreateMsgRequest, opts ...grpc.CallOption) (*CreateMsgResponse, error)
	GetMsg(ctx context.Context, in *GetMsgRequest, opts ...grpc.CallOption) (*GetMsgResponse, error)
}

type imMsgServiceClient struct {
	cc *grpc.ClientConn
}

func NewImMsgServiceClient(cc *grpc.ClientConn) ImMsgServiceClient {
	return &imMsgServiceClient{cc}
}

func (c *imMsgServiceClient) CreateMsg(ctx context.Context, in *CreateMsgRequest, opts ...grpc.CallOption) (*CreateMsgResponse, error) {
	out := new(CreateMsgResponse)
	err := grpc.Invoke(ctx, "/ImMsg.ImMsgService/CreateMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imMsgServiceClient) GetMsg(ctx context.Context, in *GetMsgRequest, opts ...grpc.CallOption) (*GetMsgResponse, error) {
	out := new(GetMsgResponse)
	err := grpc.Invoke(ctx, "/ImMsg.ImMsgService/GetMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ImMsgService service

type ImMsgServiceServer interface {
	CreateMsg(context.Context, *CreateMsgRequest) (*CreateMsgResponse, error)
	GetMsg(context.Context, *GetMsgRequest) (*GetMsgResponse, error)
}

func RegisterImMsgServiceServer(s *grpc.Server, srv ImMsgServiceServer) {
	s.RegisterService(&_ImMsgService_serviceDesc, srv)
}

func _ImMsgService_CreateMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImMsgServiceServer).CreateMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImMsg.ImMsgService/CreateMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImMsgServiceServer).CreateMsg(ctx, req.(*CreateMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImMsgService_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImMsgServiceServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImMsg.ImMsgService/GetMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImMsgServiceServer).GetMsg(ctx, req.(*GetMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImMsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ImMsg.ImMsgService",
	HandlerType: (*ImMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMsg",
			Handler:    _ImMsgService_CreateMsg_Handler,
		},
		{
			MethodName: "GetMsg",
			Handler:    _ImMsgService_GetMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imMsg.proto",
}

func init() { proto.RegisterFile("imMsg.proto", fileDescriptor_imMsg_9789ddd040263930) }

var fileDescriptor_imMsg_9789ddd040263930 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0xea, 0x30,
	0x14, 0xc5, 0xe5, 0x00, 0x01, 0x2e, 0x7f, 0xf4, 0x9e, 0xf5, 0xfe, 0x58, 0x0c, 0x08, 0x65, 0x62,
	0x29, 0x95, 0xa8, 0xda, 0xa5, 0x52, 0x87, 0x22, 0xb5, 0xca, 0x40, 0x87, 0x00, 0x4b, 0x37, 0x9a,
	0xdc, 0x46, 0xa9, 0x94, 0x38, 0x8d, 0x4d, 0x87, 0x6e, 0x1d, 0xfa, 0x41, 0xfb, 0x4d, 0x2a, 0xdb,
	0x31, 0x09, 0x54, 0x2c, 0xc8, 0xe7, 0x1c, 0xdf, 0xcb, 0xf1, 0x4f, 0x81, 0x5e, 0x92, 0x2e, 0x45,
	0x3c, 0xcb, 0x0b, 0x2e, 0x39, 0x6d, 0xf9, 0x4a, 0x78, 0x5f, 0x04, 0xcc, 0x89, 0x0e, 0xc1, 0xf1,
	0x23, 0x46, 0x26, 0x64, 0xda, 0x0c, 0x1c, 0x3f, 0xa2, 0x63, 0x80, 0x3b, 0x5e, 0xa4, 0x1b, 0x81,
	0x85, 0x1f, 0x31, 0x47, 0xfb, 0x35, 0x87, 0x8e, 0xa0, 0xb3, 0xe6, 0x65, 0xda, 0xd0, 0xe9, 0x5e,
	0x53, 0x06, 0xed, 0x05, 0xcf, 0x24, 0x66, 0x92, 0x35, 0x27, 0x64, 0xda, 0x0d, 0xac, 0xa4, 0xff,
	0xc0, 0x5d, 0x8a, 0xd8, 0x4f, 0x63, 0xd6, 0xd2, 0x41, 0xa9, 0xd4, 0xb6, 0x15, 0x66, 0xd1, 0x3a,
	0x49, 0x91, 0xb9, 0x66, 0x9b, 0xd5, 0xaa, 0xc9, 0xa2, 0xc0, 0xad, 0x44, 0x9d, 0xb6, 0x4d, 0x93,
	0xca, 0x51, 0xf9, 0x26, 0x8f, 0x6c, 0xde, 0x31, 0x79, 0xe5, 0x78, 0x57, 0xf0, 0xcb, 0xdc, 0x5e,
	0x8a, 0x38, 0xc0, 0xd7, 0x1d, 0x0a, 0x49, 0x3d, 0x68, 0x69, 0x1a, 0xfa, 0xc1, 0xbd, 0x79, 0x7f,
	0xa6, 0x21, 0x98, 0xdf, 0xc0, 0x44, 0xde, 0x19, 0xfc, 0xae, 0xcd, 0x89, 0x9c, 0x67, 0x02, 0xd5,
	0xd3, 0xc4, 0x2e, 0x0c, 0x51, 0x08, 0x3d, 0xda, 0x09, 0xac, 0xf4, 0x3e, 0x08, 0x0c, 0xee, 0x51,
	0xd6, 0xfe, 0x64, 0x0c, 0xf0, 0x5c, 0x21, 0x34, 0x68, 0x6b, 0x8e, 0x7a, 0xb4, 0xe4, 0x07, 0x80,
	0xf7, 0x5a, 0x65, 0xf9, 0x36, 0xc6, 0x55, 0xf2, 0x8e, 0x16, 0xaf, 0xd5, 0xaa, 0x83, 0x3a, 0x3f,
	0xec, 0x52, 0x8d, 0xb7, 0x19, 0x58, 0xe9, 0xbd, 0xc0, 0xd0, 0x56, 0x28, 0xfb, 0xd6, 0xf7, 0x90,
	0xd3, 0x7b, 0x9c, 0x83, 0x3d, 0x15, 0x9e, 0xc6, 0xa4, 0x71, 0x02, 0xcf, 0xfc, 0x93, 0x40, 0x5f,
	0x1b, 0x2b, 0x2c, 0xde, 0x92, 0x10, 0xe9, 0x0d, 0x74, 0xf7, 0xbc, 0xe8, 0xff, 0x72, 0xe4, 0x98,
	0xfc, 0x88, 0xfd, 0x0c, 0xca, 0xaa, 0x97, 0xe0, 0x9a, 0xf2, 0xf4, 0x4f, 0x79, 0xe7, 0x00, 0xe7,
	0xe8, 0xef, 0x91, 0x6b, 0xc6, 0x6e, 0x07, 0x8f, 0xbd, 0xd9, 0xf9, 0xb5, 0xa9, 0x96, 0x87, 0x4f,
	0xae, 0xfe, 0xbe, 0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xe3, 0x56, 0x8d, 0xee, 0x02,
	0x00, 0x00,
}
