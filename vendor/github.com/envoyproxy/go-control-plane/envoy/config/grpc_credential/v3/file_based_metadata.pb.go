// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/grpc_credential/v3/file_based_metadata.proto

package envoy_config_grpc_credential_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type FileBasedMetadataConfig struct {
	SecretData           *v3.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	HeaderKey            string         `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	HeaderPrefix         string         `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_413c1287d6760a42, []int{0}
}

func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBasedMetadataConfig.Unmarshal(m, b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
}
func (m *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(m, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return xxx_messageInfo_FileBasedMetadataConfig.Size(m)
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *v3.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v3.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v3/file_based_metadata.proto", fileDescriptor_413c1287d6760a42)
}

var fileDescriptor_413c1287d6760a42 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x26, 0x11, 0x0a, 0x9d, 0x2a, 0x48, 0x16, 0x5a, 0x0a, 0xb5, 0x51, 0x37, 0xdd, 0x38, 0x81,
	0x66, 0xa5, 0x88, 0x8b, 0x54, 0x04, 0x11, 0xa1, 0xd6, 0x03, 0x84, 0xd7, 0xe4, 0xb5, 0x1d, 0x8c,
	0x33, 0x61, 0x66, 0x1a, 0x9a, 0x9d, 0x4b, 0xc1, 0x1b, 0xb8, 0xf6, 0x0e, 0x8a, 0x7b, 0xc1, 0xad,
	0x07, 0xf1, 0x0e, 0x32, 0x99, 0xba, 0xa8, 0x45, 0xdd, 0x7e, 0x7f, 0xbc, 0xef, 0x7b, 0xe4, 0x10,
	0x79, 0x21, 0xca, 0x20, 0x11, 0x7c, 0xcc, 0x26, 0xc1, 0x44, 0xe6, 0x49, 0x9c, 0x48, 0x4c, 0x91,
	0x6b, 0x06, 0x59, 0x50, 0x84, 0xc1, 0x98, 0x65, 0x18, 0x8f, 0x40, 0x61, 0x1a, 0xdf, 0xa2, 0x86,
	0x14, 0x34, 0xd0, 0x5c, 0x0a, 0x2d, 0xbc, 0x4e, 0x65, 0xa5, 0xd6, 0x4a, 0x7f, 0x58, 0x69, 0x11,
	0xb6, 0x3a, 0x4b, 0xd9, 0x89, 0x90, 0x68, 0x02, 0x4d, 0x96, 0x4d, 0x68, 0xf9, 0xb3, 0x34, 0x87,
	0x00, 0x38, 0x17, 0x1a, 0x34, 0x13, 0x5c, 0x05, 0x0a, 0xb9, 0x62, 0x9a, 0x15, 0xdf, 0x8a, 0xf6,
	0xaa, 0x42, 0x83, 0x9e, 0xa9, 0x05, 0xbd, 0xbb, 0x42, 0x17, 0x28, 0x15, 0x13, 0x9c, 0xf1, 0x89,
	0x95, 0xec, 0x7d, 0x3a, 0x64, 0xfb, 0x8c, 0x65, 0x18, 0x99, 0x0a, 0x97, 0x8b, 0x06, 0xfd, 0xea,
	0x24, 0xef, 0x9c, 0x34, 0x14, 0x26, 0x12, 0x75, 0x6c, 0xc0, 0xa6, 0xe3, 0x3b, 0xdd, 0x46, 0xcf,
	0xa7, 0x4b, 0xbd, 0xcc, 0xd9, 0xb4, 0x08, 0xe9, 0x29, 0x68, 0xb8, 0x16, 0x33, 0x99, 0x60, 0x54,
	0x7b, 0x79, 0x7e, 0x78, 0x72, 0x9d, 0x21, 0xb1, 0x66, 0xc3, 0x78, 0x6d, 0x42, 0xa6, 0x08, 0x29,
	0xca, 0xf8, 0x06, 0xcb, 0xa6, 0xeb, 0x3b, 0xdd, 0xfa, 0xb0, 0x6e, 0x91, 0x0b, 0x2c, 0xbd, 0x7d,
	0xb2, 0xb1, 0xa0, 0x73, 0x89, 0x63, 0x36, 0x6f, 0xae, 0x55, 0x8a, 0x75, 0x0b, 0x0e, 0x2a, 0xec,
	0xa8, 0xff, 0xf8, 0x76, 0xbf, 0x73, 0x42, 0x8e, 0xff, 0xde, 0xb5, 0x07, 0x59, 0x3e, 0x05, 0xfa,
	0x4b, 0xa7, 0xe8, 0xea, 0xf5, 0xee, 0xfd, 0xa3, 0xe6, 0x6e, 0xba, 0xe4, 0x80, 0x09, 0x5b, 0x25,
	0x97, 0x62, 0x5e, 0xd2, 0x7f, 0xbe, 0x15, 0x6d, 0xad, 0x24, 0x0e, 0xcc, 0x80, 0x03, 0x67, 0x54,
	0xab, 0x96, 0x0c, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x3e, 0x16, 0xed, 0x2c, 0x02, 0x00,
	0x00,
}
