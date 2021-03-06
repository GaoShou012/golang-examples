// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jwt.proto

package proto

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

type JwtUser struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Id                   uint64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	LoginTime            int64    `protobuf:"varint,6,opt,name=loginTime,proto3" json:"loginTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtUser) Reset()         { *m = JwtUser{} }
func (m *JwtUser) String() string { return proto.CompactTextString(m) }
func (*JwtUser) ProtoMessage()    {}
func (*JwtUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cacc9c239fa0c7, []int{0}
}

func (m *JwtUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtUser.Unmarshal(m, b)
}
func (m *JwtUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtUser.Marshal(b, m, deterministic)
}
func (m *JwtUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtUser.Merge(m, src)
}
func (m *JwtUser) XXX_Size() int {
	return xxx_messageInfo_JwtUser.Size(m)
}
func (m *JwtUser) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtUser.DiscardUnknown(m)
}

var xxx_messageInfo_JwtUser proto.InternalMessageInfo

func (m *JwtUser) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JwtUser) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *JwtUser) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *JwtUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *JwtUser) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *JwtUser) GetLoginTime() int64 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

type JwtEncodeRequest struct {
	User                 *JwtUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtEncodeRequest) Reset()         { *m = JwtEncodeRequest{} }
func (m *JwtEncodeRequest) String() string { return proto.CompactTextString(m) }
func (*JwtEncodeRequest) ProtoMessage()    {}
func (*JwtEncodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cacc9c239fa0c7, []int{1}
}

func (m *JwtEncodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtEncodeRequest.Unmarshal(m, b)
}
func (m *JwtEncodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtEncodeRequest.Marshal(b, m, deterministic)
}
func (m *JwtEncodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtEncodeRequest.Merge(m, src)
}
func (m *JwtEncodeRequest) XXX_Size() int {
	return xxx_messageInfo_JwtEncodeRequest.Size(m)
}
func (m *JwtEncodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtEncodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JwtEncodeRequest proto.InternalMessageInfo

func (m *JwtEncodeRequest) GetUser() *JwtUser {
	if m != nil {
		return m.User
	}
	return nil
}

type JwtEncodeResponse struct {
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtEncodeResponse) Reset()         { *m = JwtEncodeResponse{} }
func (m *JwtEncodeResponse) String() string { return proto.CompactTextString(m) }
func (*JwtEncodeResponse) ProtoMessage()    {}
func (*JwtEncodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cacc9c239fa0c7, []int{2}
}

func (m *JwtEncodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtEncodeResponse.Unmarshal(m, b)
}
func (m *JwtEncodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtEncodeResponse.Marshal(b, m, deterministic)
}
func (m *JwtEncodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtEncodeResponse.Merge(m, src)
}
func (m *JwtEncodeResponse) XXX_Size() int {
	return xxx_messageInfo_JwtEncodeResponse.Size(m)
}
func (m *JwtEncodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtEncodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JwtEncodeResponse proto.InternalMessageInfo

func (m *JwtEncodeResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type JwtDecodeRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtDecodeRequest) Reset()         { *m = JwtDecodeRequest{} }
func (m *JwtDecodeRequest) String() string { return proto.CompactTextString(m) }
func (*JwtDecodeRequest) ProtoMessage()    {}
func (*JwtDecodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cacc9c239fa0c7, []int{3}
}

func (m *JwtDecodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtDecodeRequest.Unmarshal(m, b)
}
func (m *JwtDecodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtDecodeRequest.Marshal(b, m, deterministic)
}
func (m *JwtDecodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtDecodeRequest.Merge(m, src)
}
func (m *JwtDecodeRequest) XXX_Size() int {
	return xxx_messageInfo_JwtDecodeRequest.Size(m)
}
func (m *JwtDecodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtDecodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JwtDecodeRequest proto.InternalMessageInfo

func (m *JwtDecodeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type JwtDecodeResponse struct {
	User                 *JwtUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtDecodeResponse) Reset()         { *m = JwtDecodeResponse{} }
func (m *JwtDecodeResponse) String() string { return proto.CompactTextString(m) }
func (*JwtDecodeResponse) ProtoMessage()    {}
func (*JwtDecodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5cacc9c239fa0c7, []int{4}
}

func (m *JwtDecodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtDecodeResponse.Unmarshal(m, b)
}
func (m *JwtDecodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtDecodeResponse.Marshal(b, m, deterministic)
}
func (m *JwtDecodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtDecodeResponse.Merge(m, src)
}
func (m *JwtDecodeResponse) XXX_Size() int {
	return xxx_messageInfo_JwtDecodeResponse.Size(m)
}
func (m *JwtDecodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtDecodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JwtDecodeResponse proto.InternalMessageInfo

func (m *JwtDecodeResponse) GetUser() *JwtUser {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*JwtUser)(nil), "proto.JwtUser")
	proto.RegisterType((*JwtEncodeRequest)(nil), "proto.JwtEncodeRequest")
	proto.RegisterType((*JwtEncodeResponse)(nil), "proto.JwtEncodeResponse")
	proto.RegisterType((*JwtDecodeRequest)(nil), "proto.JwtDecodeRequest")
	proto.RegisterType((*JwtDecodeResponse)(nil), "proto.JwtDecodeResponse")
}

func init() {
	proto.RegisterFile("jwt.proto", fileDescriptor_b5cacc9c239fa0c7)
}

var fileDescriptor_b5cacc9c239fa0c7 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0xeb, 0xfc, 0x7b, 0xdf, 0x1c, 0x52, 0x05, 0x27, 0x24, 0xac, 0x8a, 0x21, 0xf2, 0x14,
	0x96, 0x0e, 0x45, 0x82, 0x89, 0x0d, 0x96, 0x8c, 0x11, 0x7c, 0x00, 0x68, 0x4e, 0xc8, 0x94, 0xda,
	0x21, 0x71, 0x14, 0x31, 0xf3, 0x21, 0xf8, 0xba, 0x28, 0x76, 0xd2, 0x3a, 0x42, 0x4c, 0xf6, 0xdd,
	0xf3, 0xdc, 0x3d, 0xbf, 0x83, 0xf4, 0xad, 0x37, 0xeb, 0xba, 0xd1, 0x46, 0x63, 0x6c, 0x1f, 0xf1,
	0xcd, 0xe0, 0x5f, 0xd1, 0x9b, 0xa7, 0x96, 0x1a, 0x44, 0x88, 0xcc, 0x67, 0x4d, 0x9c, 0x65, 0x2c,
	0x4f, 0x4b, 0xfb, 0x1f, 0x7a, 0x5d, 0x27, 0x2b, 0x1e, 0xb8, 0xde, 0xf0, 0xc7, 0x25, 0x04, 0xb2,
	0xe2, 0x61, 0xc6, 0xf2, 0xa8, 0x0c, 0x64, 0x85, 0x2b, 0xf8, 0xdf, 0xb5, 0xd4, 0xa8, 0xe7, 0x3d,
	0xf1, 0xc8, 0xfa, 0x0e, 0xf5, 0xa0, 0x29, 0xb9, 0xdd, 0x59, 0x2d, 0x76, 0xda, 0x54, 0xe3, 0x25,
	0xa4, 0xef, 0xfa, 0x55, 0xaa, 0x47, 0xb9, 0x27, 0x9e, 0x64, 0x2c, 0x0f, 0xcb, 0x63, 0x43, 0xdc,
	0xc0, 0x69, 0xd1, 0x9b, 0x07, 0xb5, 0xd5, 0x15, 0x95, 0xf4, 0xd1, 0x51, 0x6b, 0x50, 0x40, 0x34,
	0x6c, 0xb6, 0x84, 0x27, 0x9b, 0xa5, 0x3b, 0x65, 0x3d, 0xf2, 0x97, 0x56, 0x13, 0x57, 0x70, 0xe6,
	0xcd, 0xb5, 0xb5, 0x56, 0x2d, 0xe1, 0x39, 0xc4, 0x46, 0xef, 0x48, 0x8d, 0x77, 0xb8, 0x42, 0xe4,
	0x36, 0xe2, 0x9e, 0xfc, 0x88, 0x83, 0x93, 0xf9, 0xce, 0x5b, 0xbb, 0x74, 0x72, 0x8e, 0x4b, 0x27,
	0x9a, 0xe0, 0x6f, 0x9a, 0xcd, 0x17, 0x83, 0xb0, 0xe8, 0x0d, 0xde, 0x41, 0xe2, 0x90, 0xf0, 0xe2,
	0xe8, 0x9b, 0x1d, 0xb7, 0xe2, 0xbf, 0x05, 0x17, 0x24, 0x16, 0xc3, 0xb8, 0x0b, 0xf7, 0xc7, 0x67,
	0xe0, 0xfe, 0xf8, 0x9c, 0x53, 0x2c, 0x5e, 0x12, 0x2b, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x37, 0x77, 0x2c, 0x2d, 0x00, 0x02, 0x00, 0x00,
}
