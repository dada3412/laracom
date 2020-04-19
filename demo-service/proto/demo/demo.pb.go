// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/demo/demo.proto

package demo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DemoRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoRequest) Reset()         { *m = DemoRequest{} }
func (m *DemoRequest) String() string { return proto.CompactTextString(m) }
func (*DemoRequest) ProtoMessage()    {}
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b73fb6bfa52be9, []int{0}
}

func (m *DemoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoRequest.Unmarshal(m, b)
}
func (m *DemoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoRequest.Marshal(b, m, deterministic)
}
func (m *DemoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoRequest.Merge(m, src)
}
func (m *DemoRequest) XXX_Size() int {
	return xxx_messageInfo_DemoRequest.Size(m)
}
func (m *DemoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DemoRequest proto.InternalMessageInfo

func (m *DemoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DemoResponse struct {
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoResponse) Reset()         { *m = DemoResponse{} }
func (m *DemoResponse) String() string { return proto.CompactTextString(m) }
func (*DemoResponse) ProtoMessage()    {}
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b73fb6bfa52be9, []int{1}
}

func (m *DemoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoResponse.Unmarshal(m, b)
}
func (m *DemoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoResponse.Marshal(b, m, deterministic)
}
func (m *DemoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoResponse.Merge(m, src)
}
func (m *DemoResponse) XXX_Size() int {
	return xxx_messageInfo_DemoResponse.Size(m)
}
func (m *DemoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DemoResponse proto.InternalMessageInfo

func (m *DemoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*DemoRequest)(nil), "demo.DemoRequest")
	proto.RegisterType((*DemoResponse)(nil), "demo.DemoResponse")
}

func init() { proto.RegisterFile("proto/demo/demo.proto", fileDescriptor_08b73fb6bfa52be9) }

var fileDescriptor_08b73fb6bfa52be9 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0xcd, 0x85, 0x10, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0xa4, 0xc8,
	0xc5, 0xed, 0x92, 0x9a, 0x9b, 0x1f, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5,
	0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0x29, 0x71,
	0xf1, 0x40, 0x94, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0xd4, 0x94, 0xa4, 0x56, 0x94, 0x48,
	0x30, 0x41, 0xd4, 0x80, 0xd8, 0x46, 0x4e, 0x10, 0x63, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x8c, 0xb9, 0x38, 0x82, 0x13, 0x2b, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0x04, 0xf5, 0xc0,
	0x96, 0x22, 0xd9, 0x22, 0x25, 0x84, 0x2c, 0x04, 0x31, 0x55, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x2e,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x51, 0x0f, 0xe4, 0xb0, 0x00, 0x00, 0x00,
}
