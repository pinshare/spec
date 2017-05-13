// Code generated by protoc-gen-go.
// source: pin_add.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	pin_add.proto
	pin_delete.proto
	pin_update.proto
	response_pin.proto
	response_user.proto
	user_auth.proto
	user_get.proto

It has these top-level messages:
	AddRequest
	DeleteRequest
	PinDeleteResponse
	UpdateRequest
	PinResponse
	UserResponse
	UserAuthRequest
	UserIdRequest
	UserTokenRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	Title       string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Url         string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Phrase      string   `protobuf:"bytes,4,opt,name=phrase" json:"phrase,omitempty"`
	Timestamp   int64    `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AddRequest) GetPhrase() string {
	if m != nil {
		return m.Phrase
	}
	return ""
}

func (m *AddRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AddRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "service.AddRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddPin service

type AddPinClient interface {
	Execute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PinResponse, error)
}

type addPinClient struct {
	cc *grpc.ClientConn
}

func NewAddPinClient(cc *grpc.ClientConn) AddPinClient {
	return &addPinClient{cc}
}

func (c *addPinClient) Execute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PinResponse, error) {
	out := new(PinResponse)
	err := grpc.Invoke(ctx, "/service.AddPin/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddPin service

type AddPinServer interface {
	Execute(context.Context, *AddRequest) (*PinResponse, error)
}

func RegisterAddPinServer(s *grpc.Server, srv AddPinServer) {
	s.RegisterService(&_AddPin_serviceDesc, srv)
}

func _AddPin_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddPinServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AddPin/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddPinServer).Execute(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddPin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.AddPin",
	HandlerType: (*AddPinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _AddPin_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pin_add.proto",
}

func init() { proto.RegisterFile("pin_add.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x69, 0x13, 0x32, 0x22, 0xc8, 0x58, 0x64, 0x29, 0x1e, 0x42, 0x4f, 0x39, 0xe5,
	0xa0, 0xe0, 0xd9, 0x1e, 0x3c, 0x78, 0x2b, 0xfb, 0x07, 0x4a, 0xcc, 0x0e, 0x3a, 0xd0, 0x6e, 0xd6,
	0x9d, 0x89, 0xf8, 0xd3, 0xfc, 0x79, 0xd2, 0x4d, 0xb0, 0xde, 0xde, 0xfb, 0xe6, 0xcd, 0xe1, 0x3d,
	0xb8, 0x0e, 0xec, 0xf7, 0x9d, 0x73, 0x6d, 0x88, 0x83, 0x0e, 0x58, 0x0a, 0xc5, 0x2f, 0xee, 0x69,
	0x8d, 0x91, 0x24, 0x0c, 0x5e, 0x68, 0x1f, 0xd8, 0x4f, 0xc7, 0xcd, 0x4f, 0x06, 0xb0, 0x75, 0xce,
	0xd2, 0xe7, 0x48, 0xa2, 0xb8, 0x82, 0xa5, 0xb2, 0x1e, 0xc8, 0x64, 0x75, 0xd6, 0x54, 0x76, 0x32,
	0x78, 0x07, 0xc5, 0x28, 0x14, 0x5f, 0x9d, 0xb9, 0x4c, 0x78, 0x76, 0x78, 0x03, 0xf9, 0x18, 0x0f,
	0x26, 0x4f, 0xf0, 0x24, 0x4f, 0xc9, 0xf0, 0x11, 0x3b, 0x21, 0xb3, 0x98, 0x92, 0x93, 0xc3, 0x7b,
	0xa8, 0x94, 0x8f, 0x24, 0xda, 0x1d, 0x83, 0x59, 0xd6, 0x59, 0x93, 0xdb, 0x33, 0xc0, 0x1a, 0xae,
	0x1c, 0x49, 0x1f, 0x39, 0x28, 0x0f, 0xde, 0x14, 0xe9, 0xf5, 0x3f, 0x42, 0x84, 0x85, 0x76, 0xef,
	0x62, 0xca, 0x3a, 0x6f, 0x2a, 0x9b, 0xf4, 0xc3, 0x33, 0x14, 0x5b, 0xe7, 0x76, 0xec, 0xf1, 0x09,
	0xca, 0x97, 0x6f, 0xea, 0x47, 0x25, 0xbc, 0x6d, 0xe7, 0xb6, 0xed, 0xb9, 0xd5, 0x7a, 0xf5, 0x07,
	0x77, 0xec, 0xed, 0x3c, 0xc2, 0xe6, 0xe2, 0xad, 0x48, 0x1b, 0x3c, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0xf4, 0xb0, 0x92, 0x31, 0x01, 0x00, 0x00,
}
