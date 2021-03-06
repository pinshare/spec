// Code generated by protoc-gen-go.
// source: response_pin.proto
// DO NOT EDIT!

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PinResponse struct {
	Id          int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Url         string   `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Phrase      string   `protobuf:"bytes,5,opt,name=phrase" json:"phrase,omitempty"`
	Timestamp   int64    `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Description string   `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
}

func (m *PinResponse) Reset()                    { *m = PinResponse{} }
func (m *PinResponse) String() string            { return proto.CompactTextString(m) }
func (*PinResponse) ProtoMessage()               {}
func (*PinResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PinResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PinResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PinResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PinResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PinResponse) GetPhrase() string {
	if m != nil {
		return m.Phrase
	}
	return ""
}

func (m *PinResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PinResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PinResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*PinResponse)(nil), "service.PinResponse")
}

func init() { proto.RegisterFile("response_pin.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xc9, 0xe6, 0x76, 0xcf, 0x9d, 0x03, 0x91, 0x41, 0x64, 0x0a, 0x8b, 0x60, 0x95, 0xca,
	0xc6, 0xa7, 0xb0, 0x93, 0xbc, 0x80, 0xac, 0x97, 0x41, 0x07, 0xee, 0x92, 0x90, 0xc9, 0xf9, 0x8e,
	0xbe, 0x95, 0x5c, 0x76, 0x41, 0xbb, 0xff, 0xfb, 0x98, 0xaf, 0x18, 0xc0, 0xca, 0x5a, 0x72, 0x52,
	0x7e, 0x2f, 0x92, 0x9e, 0x4b, 0xcd, 0x2d, 0xe3, 0x5e, 0xb9, 0x7e, 0xcb, 0x91, 0x9f, 0x7e, 0x0c,
	0x1c, 0xde, 0x24, 0x85, 0xed, 0x04, 0x6f, 0x61, 0x90, 0x48, 0xc6, 0x19, 0x3f, 0x86, 0x41, 0x22,
	0x3e, 0xc0, 0x74, 0x51, 0xae, 0xaf, 0x91, 0x06, 0x67, 0xfc, 0x1c, 0x36, 0xc2, 0x7b, 0x18, 0x9b,
	0xb4, 0x13, 0x93, 0xed, 0x7a, 0x05, 0xbc, 0x03, 0x7b, 0xa9, 0x27, 0xda, 0x75, 0x77, 0x9d, 0xd7,
	0xbe, 0x7c, 0xd5, 0x45, 0x99, 0xc6, 0xb5, 0x5f, 0x09, 0x1f, 0x61, 0x6e, 0x72, 0x66, 0x6d, 0xcb,
	0xb9, 0xd0, 0xe4, 0x8c, 0xb7, 0xe1, 0x4f, 0xa0, 0x83, 0x43, 0x64, 0x3d, 0x56, 0x29, 0x4d, 0x72,
	0xa2, 0x7d, 0x4f, 0xff, 0x2b, 0x44, 0xd8, 0xb5, 0xe5, 0x53, 0xe9, 0xc6, 0x59, 0x3f, 0x87, 0xbe,
	0x3f, 0xa6, 0xfe, 0xdb, 0xcb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x22, 0x03, 0x98, 0xf1,
	0x00, 0x00, 0x00,
}
