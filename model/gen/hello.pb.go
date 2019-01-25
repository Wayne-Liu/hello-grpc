
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type ResponseHello struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHello) Reset()         { *m = ResponseHello{} }
func (m *ResponseHello) String() string { return proto.CompactTextString(m) }
func (*ResponseHello) ProtoMessage()    {}
func (*ResponseHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_ab818c118407ae49, []int{0}
}
func (m *ResponseHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHello.Unmarshal(m, b)
}
func (m *ResponseHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHello.Marshal(b, m, deterministic)
}
func (dst *ResponseHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHello.Merge(dst, src)
}
func (m *ResponseHello) XXX_Size() int {
	return xxx_messageInfo_ResponseHello.Size(m)
}
func (m *ResponseHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHello.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHello proto.InternalMessageInfo

func (m *ResponseHello) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseHello)(nil), "model.ResponseHello")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseHello, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseHello, error) {
	out := new(ResponseHello)
	err := c.cc.Invoke(ctx, "/model.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *empty.Empty) (*ResponseHello, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_ab818c118407ae49) }

var fileDescriptor_hello_ab818c118407ae49 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0x92,
	0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x26, 0x95, 0xa6, 0xe9, 0xbb, 0xe6, 0x16,
	0x94, 0x54, 0x42, 0xd4, 0x28, 0x69, 0x73, 0xf1, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x7a, 0x80, 0xb4, 0x0a, 0x49, 0x71, 0x71, 0x14, 0x41, 0x05, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xe0, 0x7c, 0x23, 0x2f, 0x2e, 0x1e, 0xb0, 0xa2, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54,
	0x21, 0x2b, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0x88, 0x3e, 0x31, 0x3d, 0x88, 0x35, 0x7a, 0x30, 0x6b,
	0xf4, 0xc0, 0xd6, 0x48, 0x89, 0xe8, 0x81, 0x5d, 0xa1, 0x87, 0x62, 0x8b, 0x12, 0x43, 0x12, 0x1b,
	0x58, 0x9d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x99, 0x51, 0xee, 0xb2, 0x00, 0x00, 0x00,
}