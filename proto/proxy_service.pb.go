// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy_service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proxy_service.proto

It has these top-level messages:
	Payload
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Payload struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto1.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto1.RegisterType((*Payload)(nil), "proto.Payload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProxyService service

type ProxyServiceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (ProxyService_StreamClient, error)
}

type proxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewProxyServiceClient(cc *grpc.ClientConn) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (ProxyService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ProxyService_serviceDesc.Streams[0], c.cc, "/proto.ProxyService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceStreamClient{stream}
	return x, nil
}

type ProxyService_StreamClient interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type proxyServiceStreamClient struct {
	grpc.ClientStream
}

func (x *proxyServiceStreamClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyServiceStreamClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ProxyService service

type ProxyServiceServer interface {
	Stream(ProxyService_StreamServer) error
}

func RegisterProxyServiceServer(s *grpc.Server, srv ProxyServiceServer) {
	s.RegisterService(&_ProxyService_serviceDesc, srv)
}

func _ProxyService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServiceServer).Stream(&proxyServiceStreamServer{stream})
}

type ProxyService_StreamServer interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type proxyServiceStreamServer struct {
	grpc.ServerStream
}

func (x *proxyServiceStreamServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyServiceStreamServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _ProxyService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy_service.proto",
}

func init() { proto1.RegisterFile("proxy_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0xaf,
	0xa8, 0x8c, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x4a, 0xb2, 0x5c, 0xec, 0x01, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x42, 0x42,
	0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x91,
	0x1d, 0x17, 0x4f, 0x00, 0x48, 0x73, 0x30, 0x44, 0xaf, 0x90, 0x1e, 0x17, 0x5b, 0x70, 0x49, 0x51,
	0x6a, 0x62, 0xae, 0x10, 0x1f, 0xc4, 0x1c, 0x3d, 0xa8, 0x6e, 0x29, 0x34, 0xbe, 0x12, 0x83, 0x06,
	0xa3, 0x01, 0x63, 0x12, 0x1b, 0x58, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x1b, 0xa8,
	0x0c, 0x83, 0x00, 0x00, 0x00,
}
