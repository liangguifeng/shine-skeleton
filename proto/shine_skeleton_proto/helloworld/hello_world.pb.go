// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/shine_skeleton_proto/helloworld/hello_world.proto

package hello_world

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloWorldRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46da450421b28425, []int{0}
}

func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

type HelloWorldResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46da450421b28425, []int{1}
}

func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "hello_world.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "hello_world.HelloWorldResponse")
}

func init() {
	proto.RegisterFile("proto/shine_skeleton_proto/helloworld/hello_world.proto", fileDescriptor_46da450421b28425)
}

var fileDescriptor_46da450421b28425 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xc8, 0xcc, 0x4b, 0x8d, 0x2f, 0xce, 0x4e, 0xcd, 0x49, 0x2d, 0xc9, 0xcf,
	0x8b, 0x87, 0x08, 0x66, 0xa4, 0xe6, 0xe4, 0xe4, 0x97, 0xe7, 0x17, 0xe5, 0xa4, 0x40, 0x98, 0xf1,
	0x60, 0xb6, 0x1e, 0x58, 0x52, 0x88, 0x1b, 0x49, 0x48, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27,
	0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf,
	0x18, 0xa2, 0x54, 0x49, 0x98, 0x4b, 0xd0, 0x03, 0xa4, 0x38, 0x1c, 0xa4, 0x36, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x49, 0x8f, 0x4b, 0x08, 0x59, 0xb0, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc6, 0x35, 0xaa, 0x46, 0x36, 0x24, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28,
	0x8d, 0x8b, 0x0b, 0x21, 0x28, 0x24, 0xa7, 0x87, 0xec, 0x4c, 0x0c, 0x2b, 0xa5, 0xe4, 0x71, 0xca,
	0x43, 0x6c, 0x57, 0x12, 0x6f, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xa0, 0x10, 0xbf, 0x7e, 0x99, 0x21,
	0xc4, 0xcb, 0xfa, 0x60, 0xb5, 0x49, 0x6c, 0x60, 0x8f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0x98, 0x0a, 0x31, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/hello_world.HelloWorldService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

// UnimplementedHelloWorldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServiceServer struct {
}

func (*UnimplementedHelloWorldServiceServer) HelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello_world.HelloWorldService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello_world.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _HelloWorldService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shine_skeleton_proto/helloworld/hello_world.proto",
}