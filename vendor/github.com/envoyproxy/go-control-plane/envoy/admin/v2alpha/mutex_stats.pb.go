// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/mutex_stats.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type MutexStats struct {
	NumContentions       uint64   `protobuf:"varint,1,opt,name=num_contentions,json=numContentions,proto3" json:"num_contentions,omitempty"`
	CurrentWaitCycles    uint64   `protobuf:"varint,2,opt,name=current_wait_cycles,json=currentWaitCycles,proto3" json:"current_wait_cycles,omitempty"`
	LifetimeWaitCycles   uint64   `protobuf:"varint,3,opt,name=lifetime_wait_cycles,json=lifetimeWaitCycles,proto3" json:"lifetime_wait_cycles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutexStats) Reset()         { *m = MutexStats{} }
func (m *MutexStats) String() string { return proto.CompactTextString(m) }
func (*MutexStats) ProtoMessage()    {}
func (*MutexStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1145f4b9ed4752, []int{0}
}

func (m *MutexStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutexStats.Unmarshal(m, b)
}
func (m *MutexStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutexStats.Marshal(b, m, deterministic)
}
func (m *MutexStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutexStats.Merge(m, src)
}
func (m *MutexStats) XXX_Size() int {
	return xxx_messageInfo_MutexStats.Size(m)
}
func (m *MutexStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MutexStats.DiscardUnknown(m)
}

var xxx_messageInfo_MutexStats proto.InternalMessageInfo

func (m *MutexStats) GetNumContentions() uint64 {
	if m != nil {
		return m.NumContentions
	}
	return 0
}

func (m *MutexStats) GetCurrentWaitCycles() uint64 {
	if m != nil {
		return m.CurrentWaitCycles
	}
	return 0
}

func (m *MutexStats) GetLifetimeWaitCycles() uint64 {
	if m != nil {
		return m.LifetimeWaitCycles
	}
	return 0
}

func init() {
	proto.RegisterType((*MutexStats)(nil), "envoy.admin.v2alpha.MutexStats")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/mutex_stats.proto", fileDescriptor_2c1145f4b9ed4752)
}

var fileDescriptor_2c1145f4b9ed4752 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xd9, 0x28, 0x29, 0xb6, 0x30, 0xba, 0xb1, 0x08, 0x82, 0xa0, 0x82, 0x68, 0xb5, 0x2b,
	0xda, 0x5b, 0x24, 0xb5, 0x10, 0xb4, 0xb0, 0x3c, 0xc6, 0xcb, 0x8a, 0x0b, 0xb7, 0x33, 0xc7, 0xdd,
	0x6c, 0xbc, 0xeb, 0x7c, 0x0b, 0xdf, 0xc5, 0x27, 0xb0, 0xf5, 0x8d, 0x64, 0x87, 0x93, 0x43, 0x48,
	0x3b, 0xff, 0xf7, 0x0d, 0x7c, 0xfa, 0xd2, 0xe3, 0x96, 0x7a, 0x07, 0x9b, 0x18, 0xd0, 0x6d, 0x6f,
	0xa1, 0xaa, 0xdf, 0xc0, 0xc5, 0xc4, 0xbe, 0x2b, 0x5a, 0x06, 0x6e, 0x6d, 0xdd, 0x10, 0x93, 0x99,
	0x0b, 0x66, 0x05, 0xb3, 0x03, 0x76, 0x72, 0x9a, 0x36, 0x35, 0x38, 0x40, 0x24, 0x06, 0x0e, 0x84,
	0xad, 0xcb, 0x4a, 0x1a, 0x9c, 0x8b, 0x4f, 0xa5, 0xf5, 0x43, 0xfe, 0xf4, 0x94, 0x1f, 0x99, 0x2b,
	0x3d, 0xc3, 0x14, 0x8b, 0x92, 0x90, 0x3d, 0x0a, 0xbe, 0x50, 0x67, 0xea, 0x7a, 0xff, 0xf1, 0x00,
	0x53, 0x5c, 0x8d, 0x57, 0x63, 0xf5, 0xbc, 0x4c, 0x4d, 0xe3, 0x91, 0x8b, 0x77, 0x08, 0x5c, 0x94,
	0x7d, 0x59, 0xf9, 0x76, 0x31, 0x11, 0xf8, 0x68, 0x98, 0x9e, 0x21, 0xf0, 0x4a, 0x06, 0x73, 0xa3,
	0x8f, 0xab, 0xf0, 0xea, 0x39, 0x44, 0xff, 0x4f, 0xd8, 0x13, 0xc1, 0xfc, 0x6d, 0xa3, 0xb1, 0xbc,
	0xff, 0xfa, 0xf8, 0xfe, 0x99, 0x4e, 0x0e, 0x95, 0x3e, 0x0f, 0x64, 0x25, 0xad, 0x6e, 0xa8, 0xeb,
	0xed, 0x8e, 0xca, 0xe5, 0x6c, 0x6c, 0x58, 0xe7, 0xae, 0xb5, 0x7a, 0x99, 0x4a, 0xe0, 0xdd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x63, 0x31, 0x12, 0x3d, 0x01, 0x00, 0x00,
}
