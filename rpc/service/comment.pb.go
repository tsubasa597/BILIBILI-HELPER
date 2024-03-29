// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

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

type BaseCommentRequest struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	RID                  int64    `protobuf:"varint,2,opt,name=RID,proto3" json:"RID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseCommentRequest) Reset()         { *m = BaseCommentRequest{} }
func (m *BaseCommentRequest) String() string { return proto.CompactTextString(m) }
func (*BaseCommentRequest) ProtoMessage()    {}
func (*BaseCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}

func (m *BaseCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseCommentRequest.Unmarshal(m, b)
}
func (m *BaseCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseCommentRequest.Marshal(b, m, deterministic)
}
func (m *BaseCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseCommentRequest.Merge(m, src)
}
func (m *BaseCommentRequest) XXX_Size() int {
	return xxx_messageInfo_BaseCommentRequest.Size(m)
}
func (m *BaseCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseCommentRequest proto.InternalMessageInfo

func (m *BaseCommentRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BaseCommentRequest) GetRID() int64 {
	if m != nil {
		return m.RID
	}
	return 0
}

type CommentRequest struct {
	BaseCommentRequest   *BaseCommentRequest `protobuf:"bytes,1,opt,name=baseCommentRequest,proto3" json:"baseCommentRequest,omitempty"`
	PageSum              int32               `protobuf:"varint,2,opt,name=PageSum,proto3" json:"PageSum,omitempty"`
	PageNum              int32               `protobuf:"varint,3,opt,name=PageNum,proto3" json:"PageNum,omitempty"`
	Sort                 int32               `protobuf:"varint,4,opt,name=Sort,proto3" json:"Sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CommentRequest) Reset()         { *m = CommentRequest{} }
func (m *CommentRequest) String() string { return proto.CompactTextString(m) }
func (*CommentRequest) ProtoMessage()    {}
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}

func (m *CommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentRequest.Unmarshal(m, b)
}
func (m *CommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentRequest.Marshal(b, m, deterministic)
}
func (m *CommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentRequest.Merge(m, src)
}
func (m *CommentRequest) XXX_Size() int {
	return xxx_messageInfo_CommentRequest.Size(m)
}
func (m *CommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentRequest proto.InternalMessageInfo

func (m *CommentRequest) GetBaseCommentRequest() *BaseCommentRequest {
	if m != nil {
		return m.BaseCommentRequest
	}
	return nil
}

func (m *CommentRequest) GetPageSum() int32 {
	if m != nil {
		return m.PageSum
	}
	return 0
}

func (m *CommentRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *CommentRequest) GetSort() int32 {
	if m != nil {
		return m.Sort
	}
	return 0
}

type AllCommentRequest struct {
	BaseCommentRequest   *BaseCommentRequest `protobuf:"bytes,1,opt,name=baseCommentRequest,proto3" json:"baseCommentRequest,omitempty"`
	Time                 int64               `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AllCommentRequest) Reset()         { *m = AllCommentRequest{} }
func (m *AllCommentRequest) String() string { return proto.CompactTextString(m) }
func (*AllCommentRequest) ProtoMessage()    {}
func (*AllCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{2}
}

func (m *AllCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllCommentRequest.Unmarshal(m, b)
}
func (m *AllCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllCommentRequest.Marshal(b, m, deterministic)
}
func (m *AllCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllCommentRequest.Merge(m, src)
}
func (m *AllCommentRequest) XXX_Size() int {
	return xxx_messageInfo_AllCommentRequest.Size(m)
}
func (m *AllCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllCommentRequest proto.InternalMessageInfo

func (m *AllCommentRequest) GetBaseCommentRequest() *BaseCommentRequest {
	if m != nil {
		return m.BaseCommentRequest
	}
	return nil
}

func (m *AllCommentRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type CommentResponse struct {
	DynamicUID           int64    `protobuf:"varint,1,opt,name=DynamicUID,proto3" json:"DynamicUID,omitempty"`
	RID                  int64    `protobuf:"varint,2,opt,name=RID,proto3" json:"RID,omitempty"`
	UID                  int64    `protobuf:"varint,3,opt,name=UID,proto3" json:"UID,omitempty"`
	Rpid                 int64    `protobuf:"varint,4,opt,name=Rpid,proto3" json:"Rpid,omitempty"`
	LikeNum              int32    `protobuf:"varint,5,opt,name=LikeNum,proto3" json:"LikeNum,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	Time                 int64    `protobuf:"varint,7,opt,name=Time,proto3" json:"Time,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentResponse) Reset()         { *m = CommentResponse{} }
func (m *CommentResponse) String() string { return proto.CompactTextString(m) }
func (*CommentResponse) ProtoMessage()    {}
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{3}
}

func (m *CommentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentResponse.Unmarshal(m, b)
}
func (m *CommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentResponse.Marshal(b, m, deterministic)
}
func (m *CommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentResponse.Merge(m, src)
}
func (m *CommentResponse) XXX_Size() int {
	return xxx_messageInfo_CommentResponse.Size(m)
}
func (m *CommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentResponse proto.InternalMessageInfo

func (m *CommentResponse) GetDynamicUID() int64 {
	if m != nil {
		return m.DynamicUID
	}
	return 0
}

func (m *CommentResponse) GetRID() int64 {
	if m != nil {
		return m.RID
	}
	return 0
}

func (m *CommentResponse) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *CommentResponse) GetRpid() int64 {
	if m != nil {
		return m.Rpid
	}
	return 0
}

func (m *CommentResponse) GetLikeNum() int32 {
	if m != nil {
		return m.LikeNum
	}
	return 0
}

func (m *CommentResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CommentResponse) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *CommentResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseCommentRequest)(nil), "grpc.BaseCommentRequest")
	proto.RegisterType((*CommentRequest)(nil), "grpc.CommentRequest")
	proto.RegisterType((*AllCommentRequest)(nil), "grpc.AllCommentRequest")
	proto.RegisterType((*CommentResponse)(nil), "grpc.CommentResponse")
}

func init() {
	proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828)
}

var fileDescriptor_749aee09ea917828 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xfd, 0xed, 0x6f, 0xd3, 0x56, 0xc7, 0xff, 0x83, 0xe2, 0xe2, 0x41, 0x4a, 0x4e, 0x3d, 0x45,
	0xa9, 0xe0, 0x41, 0xbc, 0xf4, 0x0f, 0xd4, 0x82, 0x14, 0xd9, 0xea, 0xc5, 0x5b, 0x1a, 0x87, 0x12,
	0xcc, 0x26, 0x69, 0xb2, 0x11, 0x7a, 0xf2, 0xbb, 0xf8, 0x5d, 0xfc, 0x5e, 0xb2, 0x9b, 0xfe, 0x91,
	0x46, 0xf0, 0xe4, 0xed, 0xed, 0xbc, 0x37, 0x33, 0x6f, 0x1e, 0x0b, 0x7b, 0x41, 0xa2, 0x14, 0xc5,
	0xda, 0x4b, 0xb3, 0x44, 0x27, 0xe8, 0x4c, 0xb3, 0x34, 0x70, 0x6f, 0x00, 0xbb, 0x7e, 0x4e, 0xbd,
	0x92, 0x92, 0x34, 0x2b, 0x28, 0xd7, 0x88, 0xe0, 0x3c, 0xce, 0x53, 0x12, 0xac, 0xc9, 0x5a, 0x35,
	0x69, 0x31, 0x1e, 0x02, 0x97, 0xc3, 0xbe, 0xf8, 0xdf, 0x64, 0x2d, 0x2e, 0x0d, 0x74, 0x3f, 0x18,
	0xec, 0x6f, 0x34, 0xde, 0x01, 0x4e, 0x2a, 0xe3, 0xec, 0x98, 0x9d, 0xb6, 0xf0, 0xcc, 0x46, 0xaf,
	0xba, 0x4e, 0xfe, 0xd0, 0x83, 0x02, 0x1a, 0x0f, 0xfe, 0x94, 0xc6, 0x85, 0xb2, 0x2b, 0x6b, 0x72,
	0xf9, 0x5c, 0x32, 0xa3, 0x42, 0x09, 0xbe, 0x66, 0x46, 0x85, 0x32, 0xb6, 0xc7, 0x49, 0xa6, 0x85,
	0x53, 0xda, 0x36, 0xd8, 0x9d, 0xc1, 0x51, 0x27, 0x8a, 0xfe, 0xcc, 0xa6, 0x49, 0x2a, 0x54, 0xb4,
	0x88, 0xc5, 0x62, 0xf7, 0x93, 0xc1, 0xc1, 0x4a, 0x96, 0xa7, 0x49, 0x9c, 0x13, 0x9e, 0x03, 0xf4,
	0xe7, 0xb1, 0xaf, 0xc2, 0xe0, 0x69, 0xd8, 0xb7, 0x9b, 0xb8, 0xfc, 0x56, 0xa9, 0xa6, 0x6b, 0x2a,
	0x46, 0xca, 0xcb, 0x8a, 0xd1, 0x20, 0x38, 0x32, 0x0d, 0x5f, 0xec, 0x79, 0x5c, 0x5a, 0x6c, 0xc2,
	0xb8, 0x0f, 0x5f, 0x6d, 0x18, 0xb5, 0x32, 0x8c, 0xc5, 0xd3, 0x30, 0xbd, 0x24, 0xd6, 0x14, 0x6b,
	0x51, 0x6f, 0xb2, 0xd6, 0xb6, 0x5c, 0x3e, 0x57, 0x9e, 0x1b, 0x6b, 0xcf, 0xa6, 0x36, 0xf2, 0x15,
	0x89, 0x2d, 0x2b, 0xb5, 0xb8, 0xfd, 0x6e, 0x26, 0xd8, 0x33, 0xf0, 0x16, 0xea, 0x03, 0xd2, 0x9d,
	0x28, 0xc2, 0xd3, 0x32, 0x9e, 0x4a, 0xa6, 0x67, 0x27, 0x25, 0xb1, 0x71, 0xb8, 0xfb, 0xef, 0x92,
	0xe1, 0x35, 0xf0, 0x01, 0x69, 0x3c, 0xde, 0x50, 0xfc, 0xd6, 0xd7, 0xdd, 0x7d, 0x06, 0xcf, 0xbb,
	0xc8, 0x29, 0x7b, 0x0b, 0x03, 0x9a, 0xd4, 0xed, 0xbf, 0xbd, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x07, 0x0d, 0xd8, 0xc6, 0xc8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentClient interface {
	GetAll(ctx context.Context, in *AllCommentRequest, opts ...grpc.CallOption) (Comment_GetAllClient, error)
	Get(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (Comment_GetClient, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) GetAll(ctx context.Context, in *AllCommentRequest, opts ...grpc.CallOption) (Comment_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Comment_serviceDesc.Streams[0], "/grpc.Comment/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &commentGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Comment_GetAllClient interface {
	Recv() (*CommentResponse, error)
	grpc.ClientStream
}

type commentGetAllClient struct {
	grpc.ClientStream
}

func (x *commentGetAllClient) Recv() (*CommentResponse, error) {
	m := new(CommentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commentClient) Get(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (Comment_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Comment_serviceDesc.Streams[1], "/grpc.Comment/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &commentGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Comment_GetClient interface {
	Recv() (*CommentResponse, error)
	grpc.ClientStream
}

type commentGetClient struct {
	grpc.ClientStream
}

func (x *commentGetClient) Recv() (*CommentResponse, error) {
	m := new(CommentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommentServer is the server API for Comment service.
type CommentServer interface {
	GetAll(*AllCommentRequest, Comment_GetAllServer) error
	Get(*CommentRequest, Comment_GetServer) error
}

// UnimplementedCommentServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (*UnimplementedCommentServer) GetAll(req *AllCommentRequest, srv Comment_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCommentServer) Get(req *CommentRequest, srv Comment_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCommentServer(s *grpc.Server, srv CommentServer) {
	s.RegisterService(&_Comment_serviceDesc, srv)
}

func _Comment_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllCommentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommentServer).GetAll(m, &commentGetAllServer{stream})
}

type Comment_GetAllServer interface {
	Send(*CommentResponse) error
	grpc.ServerStream
}

type commentGetAllServer struct {
	grpc.ServerStream
}

func (x *commentGetAllServer) Send(m *CommentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Comment_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommentServer).Get(m, &commentGetServer{stream})
}

type Comment_GetServer interface {
	Send(*CommentResponse) error
	grpc.ServerStream
}

type commentGetServer struct {
	grpc.ServerStream
}

func (x *commentGetServer) Send(m *CommentResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Comment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _Comment_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _Comment_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "comment.proto",
}
