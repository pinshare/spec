// Code generated by protoc-gen-go.
// source: pin_delete.proto
// DO NOT EDIT!

package service

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

type DeleteRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeleteRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PinDeleteResponse struct {
}

func (m *PinDeleteResponse) Reset()                    { *m = PinDeleteResponse{} }
func (m *PinDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*PinDeleteResponse) ProtoMessage()               {}
func (*PinDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*DeleteRequest)(nil), "service.DeleteRequest")
	proto.RegisterType((*PinDeleteResponse)(nil), "service.PinDeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeletePin service

type DeletePinClient interface {
	Execute(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*PinDeleteResponse, error)
}

type deletePinClient struct {
	cc *grpc.ClientConn
}

func NewDeletePinClient(cc *grpc.ClientConn) DeletePinClient {
	return &deletePinClient{cc}
}

func (c *deletePinClient) Execute(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*PinDeleteResponse, error) {
	out := new(PinDeleteResponse)
	err := grpc.Invoke(ctx, "/service.DeletePin/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeletePin service

type DeletePinServer interface {
	Execute(context.Context, *DeleteRequest) (*PinDeleteResponse, error)
}

func RegisterDeletePinServer(s *grpc.Server, srv DeletePinServer) {
	s.RegisterService(&_DeletePin_serviceDesc, srv)
}

func _DeletePin_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeletePinServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.DeletePin/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeletePinServer).Execute(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeletePin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.DeletePin",
	HandlerType: (*DeletePinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _DeletePin_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pin_delete.proto",
}

func init() { proto.RegisterFile("pin_delete.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc8, 0xcc, 0x8b,
	0x4f, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x92, 0xe7, 0xe2, 0x75, 0x01, 0x4b, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06,
	0x31, 0x65, 0xa6, 0x28, 0x09, 0x73, 0x09, 0x06, 0x64, 0xe6, 0xc1, 0xd4, 0x14, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x1a, 0xf9, 0x70, 0x71, 0x42, 0x44, 0x02, 0x32, 0xf3, 0x84, 0xec, 0xb9, 0xd8, 0x5d,
	0x2b, 0x52, 0x93, 0x4b, 0x4b, 0x52, 0x85, 0xc4, 0xf4, 0xa0, 0xe6, 0xea, 0xa1, 0x18, 0x2a, 0x25,
	0x05, 0x17, 0xc7, 0x30, 0x4b, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x67, 0xe9, 0xe2, 0x39, 0xa7, 0x00, 0x00, 0x00,
}
