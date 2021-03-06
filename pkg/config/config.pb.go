// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import config "github.com/mjpitz/rds/pkg/config"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CloneOverride struct {
	RepositoryRoot       *wrappers.StringValue `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	Depth                *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CloneOverride) Reset()         { *m = CloneOverride{} }
func (m *CloneOverride) String() string { return proto.CompactTextString(m) }
func (*CloneOverride) ProtoMessage()    {}
func (*CloneOverride) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_3c000abeacc11847, []int{0}
}
func (m *CloneOverride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneOverride.Unmarshal(m, b)
}
func (m *CloneOverride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneOverride.Marshal(b, m, deterministic)
}
func (dst *CloneOverride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneOverride.Merge(dst, src)
}
func (m *CloneOverride) XXX_Size() int {
	return xxx_messageInfo_CloneOverride.Size(m)
}
func (m *CloneOverride) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneOverride.DiscardUnknown(m)
}

var xxx_messageInfo_CloneOverride proto.InternalMessageInfo

func (m *CloneOverride) GetRepositoryRoot() *wrappers.StringValue {
	if m != nil {
		return m.RepositoryRoot
	}
	return nil
}

func (m *CloneOverride) GetDepth() *wrappers.Int32Value {
	if m != nil {
		return m.Depth
	}
	return nil
}

type CloneConfiguration struct {
	RepositoryRoot       *wrappers.StringValue     `protobuf:"bytes,1,opt,name=repository_root,json=repositoryRoot,proto3" json:"repository_root,omitempty"`
	Depth                *wrappers.Int32Value      `protobuf:"bytes,2,opt,name=depth,proto3" json:"depth,omitempty"`
	Overrides            map[string]*CloneOverride `protobuf:"bytes,10,rep,name=overrides,proto3" json:"overrides,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CloneConfiguration) Reset()         { *m = CloneConfiguration{} }
func (m *CloneConfiguration) String() string { return proto.CompactTextString(m) }
func (*CloneConfiguration) ProtoMessage()    {}
func (*CloneConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_3c000abeacc11847, []int{1}
}
func (m *CloneConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloneConfiguration.Unmarshal(m, b)
}
func (m *CloneConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloneConfiguration.Marshal(b, m, deterministic)
}
func (dst *CloneConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloneConfiguration.Merge(dst, src)
}
func (m *CloneConfiguration) XXX_Size() int {
	return xxx_messageInfo_CloneConfiguration.Size(m)
}
func (m *CloneConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_CloneConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_CloneConfiguration proto.InternalMessageInfo

func (m *CloneConfiguration) GetRepositoryRoot() *wrappers.StringValue {
	if m != nil {
		return m.RepositoryRoot
	}
	return nil
}

func (m *CloneConfiguration) GetDepth() *wrappers.Int32Value {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *CloneConfiguration) GetOverrides() map[string]*CloneOverride {
	if m != nil {
		return m.Overrides
	}
	return nil
}

type Configuration struct {
	Mount                string              `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount,omitempty"`
	Accounts             []*config.Account   `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Clone                *CloneConfiguration `protobuf:"bytes,3,opt,name=clone,proto3" json:"clone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_3c000abeacc11847, []int{2}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetMount() string {
	if m != nil {
		return m.Mount
	}
	return ""
}

func (m *Configuration) GetAccounts() []*config.Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Configuration) GetClone() *CloneConfiguration {
	if m != nil {
		return m.Clone
	}
	return nil
}

func init() {
	proto.RegisterType((*CloneOverride)(nil), "mjpitz.gitfs.config.CloneOverride")
	proto.RegisterType((*CloneConfiguration)(nil), "mjpitz.gitfs.config.CloneConfiguration")
	proto.RegisterMapType((map[string]*CloneOverride)(nil), "mjpitz.gitfs.config.CloneConfiguration.OverridesEntry")
	proto.RegisterType((*Configuration)(nil), "mjpitz.gitfs.config.Configuration")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_3c000abeacc11847) }

var fileDescriptor_config_3c000abeacc11847 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x51, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0x56, 0x52, 0xa5, 0x6a, 0x5d, 0x28, 0xc8, 0x30, 0x44, 0x81, 0xa1, 0xca, 0x42, 0x17, 0x6c,
	0xd1, 0x4a, 0x55, 0x41, 0x62, 0x80, 0xaa, 0x03, 0x13, 0x52, 0x40, 0x0c, 0x2c, 0x90, 0x26, 0xae,
	0x6b, 0x9a, 0xfa, 0x58, 0x8e, 0x53, 0x14, 0xde, 0x80, 0x77, 0xe0, 0x79, 0x78, 0xae, 0xab, 0xd8,
	0x4d, 0x7f, 0xd4, 0x7b, 0xa5, 0xbb, 0xdd, 0xed, 0x58, 0xfe, 0xfe, 0xce, 0x77, 0xd0, 0xa3, 0x0c,
	0xe4, 0x5a, 0x70, 0xa2, 0x34, 0x18, 0xc0, 0xcf, 0x76, 0xbf, 0x94, 0x30, 0x7f, 0x08, 0x17, 0x66,
	0x5d, 0x12, 0xf7, 0x15, 0xbd, 0xe5, 0xc2, 0x6c, 0xaa, 0x15, 0xc9, 0x60, 0x47, 0x39, 0x14, 0xa9,
	0xe4, 0xd4, 0xa2, 0x57, 0xd5, 0x9a, 0x2a, 0x53, 0x2b, 0x56, 0xd2, 0xdf, 0x3a, 0x55, 0x8a, 0xe9,
	0xd3, 0xe0, 0xf4, 0xa2, 0xd7, 0x67, 0x54, 0x27, 0x4d, 0x75, 0x5e, 0x52, 0xb5, 0xe5, 0xd4, 0x89,
	0xd3, 0x73, 0xfb, 0xf8, 0xaf, 0x87, 0x1e, 0x2f, 0x0a, 0x90, 0xec, 0xf3, 0x9e, 0x69, 0x2d, 0x72,
	0x86, 0x97, 0xe8, 0x89, 0x66, 0x0a, 0x4a, 0x61, 0x40, 0xd7, 0x3f, 0x34, 0x80, 0x09, 0xbd, 0x91,
	0x37, 0x1e, 0x4c, 0x5e, 0x12, 0x0e, 0xc0, 0x0b, 0x46, 0xda, 0x28, 0xe4, 0x8b, 0xd1, 0x42, 0xf2,
	0x6f, 0x69, 0x51, 0xb1, 0x64, 0x78, 0x22, 0x25, 0x00, 0x06, 0xbf, 0x41, 0x41, 0xce, 0x94, 0xd9,
	0x84, 0xbe, 0x25, 0xbf, 0xb8, 0x22, 0x7f, 0x92, 0x66, 0x3a, 0x71, 0x5c, 0x87, 0x8c, 0xff, 0xfb,
	0x08, 0xdb, 0x2c, 0x0b, 0x9b, 0xb0, 0xd2, 0xa9, 0x11, 0x20, 0x1f, 0x2e, 0x10, 0xfe, 0x8a, 0xfa,
	0x70, 0xa8, 0xa5, 0x0c, 0xd1, 0xa8, 0x33, 0x1e, 0x4c, 0x66, 0xe4, 0x96, 0x7b, 0x91, 0xeb, 0xd4,
	0xa4, 0xed, 0xb3, 0x5c, 0x4a, 0xa3, 0xeb, 0xe4, 0x24, 0x14, 0xfd, 0x44, 0xc3, 0xcb, 0x4f, 0xfc,
	0x14, 0x75, 0xb6, 0xac, 0xb6, 0x5b, 0xf5, 0x93, 0x66, 0xc4, 0x73, 0x14, 0xec, 0x9b, 0x24, 0x87,
	0xb0, 0xf1, 0xdd, 0xae, 0xad, 0x54, 0xe2, 0x08, 0xef, 0xfc, 0xb9, 0x17, 0xff, 0x6b, 0x8e, 0x7a,
	0xd1, 0xe1, 0x73, 0x14, 0xec, 0xa0, 0x92, 0xe6, 0xe0, 0xe1, 0x1e, 0x78, 0x86, 0x7a, 0x69, 0x96,
	0x35, 0x63, 0x19, 0xfa, 0x76, 0xbd, 0xa8, 0x35, 0xd2, 0xf9, 0xd1, 0xe6, 0x83, 0x83, 0x24, 0x47,
	0x2c, 0x7e, 0x8f, 0x82, 0xac, 0xf1, 0x0e, 0x3b, 0x36, 0xdd, 0xab, 0x7b, 0x76, 0x92, 0x38, 0xd6,
	0xc7, 0xde, 0xf7, 0xae, 0xc3, 0xac, 0xba, 0xb6, 0xfc, 0xe9, 0x4d, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfe, 0x04, 0xd3, 0x66, 0x13, 0x03, 0x00, 0x00,
}
