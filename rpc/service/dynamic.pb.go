// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dynamic.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BaseDynamicRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseDynamicRequest) Reset()         { *m = BaseDynamicRequest{} }
func (m *BaseDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*BaseDynamicRequest) ProtoMessage()    {}
func (*BaseDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0de932ef16148a5, []int{0}
}

func (m *BaseDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseDynamicRequest.Unmarshal(m, b)
}
func (m *BaseDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseDynamicRequest.Marshal(b, m, deterministic)
}
func (m *BaseDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseDynamicRequest.Merge(m, src)
}
func (m *BaseDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_BaseDynamicRequest.Size(m)
}
func (m *BaseDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseDynamicRequest proto.InternalMessageInfo

func (m *BaseDynamicRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type DynamicRequest struct {
	BaseCommentRequest   *BaseDynamicRequest `protobuf:"bytes,1,opt,name=baseCommentRequest,proto3" json:"baseCommentRequest,omitempty"`
	Offect               int64               `protobuf:"varint,2,opt,name=Offect,proto3" json:"Offect,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DynamicRequest) Reset()         { *m = DynamicRequest{} }
func (m *DynamicRequest) String() string { return proto.CompactTextString(m) }
func (*DynamicRequest) ProtoMessage()    {}
func (*DynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0de932ef16148a5, []int{1}
}

func (m *DynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicRequest.Unmarshal(m, b)
}
func (m *DynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicRequest.Marshal(b, m, deterministic)
}
func (m *DynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicRequest.Merge(m, src)
}
func (m *DynamicRequest) XXX_Size() int {
	return xxx_messageInfo_DynamicRequest.Size(m)
}
func (m *DynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicRequest proto.InternalMessageInfo

func (m *DynamicRequest) GetBaseCommentRequest() *BaseDynamicRequest {
	if m != nil {
		return m.BaseCommentRequest
	}
	return nil
}

func (m *DynamicRequest) GetOffect() int64 {
	if m != nil {
		return m.Offect
	}
	return 0
}

type AllDynamicRequest struct {
	BaseCommentRequest   *BaseDynamicRequest `protobuf:"bytes,1,opt,name=baseCommentRequest,proto3" json:"baseCommentRequest,omitempty"`
	Time                 int64               `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AllDynamicRequest) Reset()         { *m = AllDynamicRequest{} }
func (m *AllDynamicRequest) String() string { return proto.CompactTextString(m) }
func (*AllDynamicRequest) ProtoMessage()    {}
func (*AllDynamicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0de932ef16148a5, []int{2}
}

func (m *AllDynamicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllDynamicRequest.Unmarshal(m, b)
}
func (m *AllDynamicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllDynamicRequest.Marshal(b, m, deterministic)
}
func (m *AllDynamicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllDynamicRequest.Merge(m, src)
}
func (m *AllDynamicRequest) XXX_Size() int {
	return xxx_messageInfo_AllDynamicRequest.Size(m)
}
func (m *AllDynamicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllDynamicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllDynamicRequest proto.InternalMessageInfo

func (m *AllDynamicRequest) GetBaseCommentRequest() *BaseDynamicRequest {
	if m != nil {
		return m.BaseCommentRequest
	}
	return nil
}

func (m *AllDynamicRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type DynamicResponse struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Card                 string   `protobuf:"bytes,3,opt,name=Card,proto3" json:"Card,omitempty"`
	RID                  int64    `protobuf:"varint,4,opt,name=RID,proto3" json:"RID,omitempty"`
	Offect               int64    `protobuf:"varint,5,opt,name=Offect,proto3" json:"Offect,omitempty"`
	Type                 int32    `protobuf:"varint,6,opt,name=Type,proto3" json:"Type,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
	Time                 int64    `protobuf:"varint,8,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynamicResponse) Reset()         { *m = DynamicResponse{} }
func (m *DynamicResponse) String() string { return proto.CompactTextString(m) }
func (*DynamicResponse) ProtoMessage()    {}
func (*DynamicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0de932ef16148a5, []int{3}
}

func (m *DynamicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicResponse.Unmarshal(m, b)
}
func (m *DynamicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicResponse.Marshal(b, m, deterministic)
}
func (m *DynamicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicResponse.Merge(m, src)
}
func (m *DynamicResponse) XXX_Size() int {
	return xxx_messageInfo_DynamicResponse.Size(m)
}
func (m *DynamicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicResponse proto.InternalMessageInfo

func (m *DynamicResponse) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *DynamicResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DynamicResponse) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *DynamicResponse) GetRID() int64 {
	if m != nil {
		return m.RID
	}
	return 0
}

func (m *DynamicResponse) GetOffect() int64 {
	if m != nil {
		return m.Offect
	}
	return 0
}

func (m *DynamicResponse) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *DynamicResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DynamicResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*BaseDynamicRequest)(nil), "grpc.BaseDynamicRequest")
	proto.RegisterType((*DynamicRequest)(nil), "grpc.DynamicRequest")
	proto.RegisterType((*AllDynamicRequest)(nil), "grpc.AllDynamicRequest")
	proto.RegisterType((*DynamicResponse)(nil), "grpc.DynamicResponse")
}

func init() {
	proto.RegisterFile("dynamic.proto", fileDescriptor_b0de932ef16148a5)
}

var fileDescriptor_b0de932ef16148a5 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xbb, 0x4e, 0xc3, 0x40,
	0x10, 0xc4, 0x38, 0xb1, 0x61, 0x79, 0xaf, 0x78, 0x9c, 0xa8, 0x22, 0x17, 0x28, 0x95, 0x41, 0x41,
	0xa2, 0xa2, 0x49, 0x6c, 0x29, 0xb8, 0x01, 0xe9, 0x04, 0x0d, 0x9d, 0xe3, 0x6c, 0x90, 0x25, 0xbf,
	0xe2, 0x3b, 0x90, 0x52, 0xf1, 0x59, 0xfc, 0x1e, 0xba, 0x3b, 0x13, 0x42, 0x8c, 0xe8, 0xe8, 0xe6,
	0xc6, 0x3b, 0xb3, 0xb3, 0x23, 0xc3, 0xde, 0x74, 0x51, 0xc4, 0x79, 0x9a, 0xf8, 0x55, 0x5d, 0xca,
	0x12, 0x3b, 0x2f, 0x75, 0x95, 0x78, 0x17, 0x80, 0xa3, 0x58, 0x50, 0x68, 0x3e, 0x71, 0x9a, 0xbf,
	0x92, 0x90, 0x78, 0x08, 0xf6, 0x53, 0x14, 0x32, 0xab, 0x67, 0xf5, 0x6d, 0xae, 0xa0, 0x57, 0xc3,
	0xfe, 0xda, 0xcc, 0x1d, 0xe0, 0x24, 0x16, 0x14, 0x94, 0x79, 0x4e, 0x85, 0x6c, 0x58, 0x2d, 0xd9,
	0x19, 0x30, 0x5f, 0x99, 0xfb, 0x6d, 0x67, 0xfe, 0x8b, 0x06, 0x4f, 0xc1, 0x79, 0x98, 0xcd, 0x28,
	0x91, 0x6c, 0x53, 0x2f, 0x6c, 0x5e, 0xde, 0x1c, 0x8e, 0x86, 0x59, 0xf6, 0x6f, 0x6b, 0x11, 0x3a,
	0x8f, 0x69, 0x4e, 0xcd, 0x52, 0x8d, 0xbd, 0x0f, 0x0b, 0x0e, 0x96, 0x52, 0x51, 0x95, 0x85, 0xa0,
	0x76, 0x19, 0xc8, 0xc0, 0x0d, 0xca, 0x42, 0x52, 0x61, 0x12, 0x6f, 0xf3, 0xaf, 0xa7, 0xf2, 0x0c,
	0xe2, 0x7a, 0xca, 0x6c, 0x4d, 0x6b, 0xac, 0xf4, 0x3c, 0x0a, 0x59, 0xc7, 0xe8, 0x79, 0x14, 0xae,
	0x1c, 0xdc, 0x5d, 0x3d, 0x58, 0x27, 0x5a, 0x54, 0xc4, 0x9c, 0x9e, 0xd5, 0xef, 0x72, 0x8d, 0x15,
	0x77, 0x1f, 0xe7, 0xc4, 0x5c, 0xe3, 0xa8, 0xf0, 0x32, 0xf9, 0xd6, 0x77, 0xf2, 0xc1, 0x3b, 0xb8,
	0x4d, 0x70, 0xbc, 0x01, 0x7b, 0x4c, 0x12, 0x8f, 0x4d, 0x1b, 0x3f, 0x9b, 0x38, 0x3f, 0x59, 0x63,
	0xcd, 0x91, 0xde, 0xc6, 0x95, 0x85, 0xb7, 0xe0, 0x8c, 0x49, 0x0e, 0xb3, 0x0c, 0xcf, 0xcc, 0x50,
	0xab, 0xfd, 0x3f, 0xd4, 0xa3, 0xdd, 0x67, 0xf0, 0xfd, 0x4b, 0x41, 0xf5, 0x5b, 0x9a, 0xd0, 0xc4,
	0xd1, 0x3f, 0xd9, 0xf5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x50, 0x62, 0x33, 0x75, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DynamicClient is the client API for Dynamic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DynamicClient interface {
	Get(ctx context.Context, in *DynamicRequest, opts ...grpc.CallOption) (Dynamic_GetClient, error)
	GetAll(ctx context.Context, in *AllDynamicRequest, opts ...grpc.CallOption) (Dynamic_GetAllClient, error)
}

type dynamicClient struct {
	cc grpc.ClientConnInterface
}

func NewDynamicClient(cc grpc.ClientConnInterface) DynamicClient {
	return &dynamicClient{cc}
}

func (c *dynamicClient) Get(ctx context.Context, in *DynamicRequest, opts ...grpc.CallOption) (Dynamic_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dynamic_serviceDesc.Streams[0], "/grpc.Dynamic/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &dynamicGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dynamic_GetClient interface {
	Recv() (*DynamicResponse, error)
	grpc.ClientStream
}

type dynamicGetClient struct {
	grpc.ClientStream
}

func (x *dynamicGetClient) Recv() (*DynamicResponse, error) {
	m := new(DynamicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dynamicClient) GetAll(ctx context.Context, in *AllDynamicRequest, opts ...grpc.CallOption) (Dynamic_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dynamic_serviceDesc.Streams[1], "/grpc.Dynamic/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &dynamicGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dynamic_GetAllClient interface {
	Recv() (*DynamicResponse, error)
	grpc.ClientStream
}

type dynamicGetAllClient struct {
	grpc.ClientStream
}

func (x *dynamicGetAllClient) Recv() (*DynamicResponse, error) {
	m := new(DynamicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DynamicServer is the server API for Dynamic service.
type DynamicServer interface {
	Get(*DynamicRequest, Dynamic_GetServer) error
	GetAll(*AllDynamicRequest, Dynamic_GetAllServer) error
}

// UnimplementedDynamicServer can be embedded to have forward compatible implementations.
type UnimplementedDynamicServer struct {
}

func (*UnimplementedDynamicServer) Get(req *DynamicRequest, srv Dynamic_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDynamicServer) GetAll(req *AllDynamicRequest, srv Dynamic_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterDynamicServer(s *grpc.Server, srv DynamicServer) {
	s.RegisterService(&_Dynamic_serviceDesc, srv)
}

func _Dynamic_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DynamicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DynamicServer).Get(m, &dynamicGetServer{stream})
}

type Dynamic_GetServer interface {
	Send(*DynamicResponse) error
	grpc.ServerStream
}

type dynamicGetServer struct {
	grpc.ServerStream
}

func (x *dynamicGetServer) Send(m *DynamicResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dynamic_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllDynamicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DynamicServer).GetAll(m, &dynamicGetAllServer{stream})
}

type Dynamic_GetAllServer interface {
	Send(*DynamicResponse) error
	grpc.ServerStream
}

type dynamicGetAllServer struct {
	grpc.ServerStream
}

func (x *dynamicGetAllServer) Send(m *DynamicResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Dynamic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Dynamic",
	HandlerType: (*DynamicServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Dynamic_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAll",
			Handler:       _Dynamic_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dynamic.proto",
}
