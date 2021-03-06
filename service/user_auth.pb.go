// Code generated by protoc-gen-go.
// source: user_auth.proto
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

type UserAuthRequest struct {
	GithubId             string `protobuf:"bytes,1,opt,name=github_id,json=githubId" json:"github_id,omitempty"`
	GithubName           string `protobuf:"bytes,2,opt,name=github_name,json=githubName" json:"github_name,omitempty"`
	GithubProfileIconUrl string `protobuf:"bytes,4,opt,name=github_profile_icon_url,json=githubProfileIconUrl" json:"github_profile_icon_url,omitempty"`
}

func (m *UserAuthRequest) Reset()                    { *m = UserAuthRequest{} }
func (m *UserAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*UserAuthRequest) ProtoMessage()               {}
func (*UserAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *UserAuthRequest) GetGithubId() string {
	if m != nil {
		return m.GithubId
	}
	return ""
}

func (m *UserAuthRequest) GetGithubName() string {
	if m != nil {
		return m.GithubName
	}
	return ""
}

func (m *UserAuthRequest) GetGithubProfileIconUrl() string {
	if m != nil {
		return m.GithubProfileIconUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*UserAuthRequest)(nil), "service.UserAuthRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthUser service

type AuthUserClient interface {
	Execute(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type authUserClient struct {
	cc *grpc.ClientConn
}

func NewAuthUserClient(cc *grpc.ClientConn) AuthUserClient {
	return &authUserClient{cc}
}

func (c *authUserClient) Execute(ctx context.Context, in *UserAuthRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/service.AuthUser/execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthUser service

type AuthUserServer interface {
	Execute(context.Context, *UserAuthRequest) (*UserResponse, error)
}

func RegisterAuthUserServer(s *grpc.Server, srv AuthUserServer) {
	s.RegisterService(&_AuthUser_serviceDesc, srv)
}

func _AuthUser_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthUserServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AuthUser/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthUserServer).Execute(ctx, req.(*UserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.AuthUser",
	HandlerType: (*AuthUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "execute",
			Handler:    _AuthUser_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_auth.proto",
}

func init() { proto.RegisterFile("user_auth.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4e, 0x86, 0x30,
	0x14, 0x46, 0xc5, 0x18, 0x81, 0xeb, 0x40, 0x52, 0x35, 0x36, 0x38, 0x68, 0x98, 0x9c, 0x18, 0x34,
	0x6e, 0x2e, 0x6e, 0xb2, 0x18, 0x43, 0xc2, 0xdc, 0x40, 0xb9, 0x4a, 0x13, 0x68, 0xf1, 0xb6, 0x35,
	0x3e, 0x81, 0xcf, 0x6d, 0xa0, 0x5d, 0xfe, 0x7f, 0xfd, 0xce, 0x19, 0xce, 0x07, 0x85, 0xb7, 0x48,
	0xa2, 0xf7, 0x6e, 0xaa, 0x57, 0x32, 0xce, 0xb0, 0xd4, 0x22, 0xfd, 0x28, 0x89, 0xe5, 0x25, 0xa1,
	0x5d, 0x8d, 0xb6, 0x28, 0x36, 0x25, 0xd0, 0xea, 0x2f, 0x81, 0xa2, 0xb3, 0x48, 0xaf, 0xde, 0x4d,
	0x2d, 0x7e, 0x7b, 0xb4, 0x8e, 0xdd, 0x42, 0xfe, 0xa5, 0xdc, 0xe4, 0x07, 0xa1, 0x46, 0x9e, 0xdc,
	0x27, 0x0f, 0x79, 0x9b, 0x85, 0xa1, 0x19, 0xd9, 0x1d, 0x5c, 0x44, 0xa8, 0xfb, 0x05, 0xf9, 0xe9,
	0x8e, 0x21, 0x4c, 0xef, 0xfd, 0x82, 0xec, 0x19, 0x6e, 0xa2, 0xb0, 0x92, 0xf9, 0x54, 0x33, 0x0a,
	0x25, 0x8d, 0x16, 0x9e, 0x66, 0x7e, 0xb6, 0xcb, 0x57, 0x01, 0x7f, 0x04, 0xda, 0x48, 0xa3, 0x3b,
	0x9a, 0x1f, 0xdf, 0x20, 0xdb, 0x1a, 0xb6, 0x16, 0xf6, 0x02, 0x29, 0xfe, 0xa2, 0xf4, 0x0e, 0x19,
	0xaf, 0x63, 0x7e, 0x7d, 0x54, 0x59, 0x5e, 0x1f, 0x90, 0x36, 0x7e, 0xab, 0x4e, 0x86, 0xf3, 0xfd,
	0xd9, 0xd3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xcb, 0x4c, 0xf4, 0x0a, 0x01, 0x00, 0x00,
}
