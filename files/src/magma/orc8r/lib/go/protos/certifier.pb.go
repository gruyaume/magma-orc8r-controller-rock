// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/certifier.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CertType int32

const (
	CertType_DEFAULT CertType = 0
	CertType_VPN     CertType = 1
)

var CertType_name = map[int32]string{
	0: "DEFAULT",
	1: "VPN",
}

var CertType_value = map[string]int32{
	"DEFAULT": 0,
	"VPN":     1,
}

func (x CertType) String() string {
	return proto.EnumName(CertType_name, int32(x))
}

func (CertType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_309897dc79f61bc0, []int{0}
}

type CSR struct {
	Id                   *Identity          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ValidTime            *duration.Duration `protobuf:"bytes,2,opt,name=valid_time,json=validTime,proto3" json:"valid_time,omitempty"`
	CsrDer               []byte             `protobuf:"bytes,3,opt,name=csr_der,json=csrDer,proto3" json:"csr_der,omitempty"`
	CertType             CertType           `protobuf:"varint,4,opt,name=cert_type,json=certType,proto3,enum=magma.orc8r.CertType" json:"cert_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CSR) Reset()         { *m = CSR{} }
func (m *CSR) String() string { return proto.CompactTextString(m) }
func (*CSR) ProtoMessage()    {}
func (*CSR) Descriptor() ([]byte, []int) {
	return fileDescriptor_309897dc79f61bc0, []int{0}
}

func (m *CSR) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSR.Unmarshal(m, b)
}
func (m *CSR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSR.Marshal(b, m, deterministic)
}
func (m *CSR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSR.Merge(m, src)
}
func (m *CSR) XXX_Size() int {
	return xxx_messageInfo_CSR.Size(m)
}
func (m *CSR) XXX_DiscardUnknown() {
	xxx_messageInfo_CSR.DiscardUnknown(m)
}

var xxx_messageInfo_CSR proto.InternalMessageInfo

func (m *CSR) GetId() *Identity {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CSR) GetValidTime() *duration.Duration {
	if m != nil {
		return m.ValidTime
	}
	return nil
}

func (m *CSR) GetCsrDer() []byte {
	if m != nil {
		return m.CsrDer
	}
	return nil
}

func (m *CSR) GetCertType() CertType {
	if m != nil {
		return m.CertType
	}
	return CertType_DEFAULT
}

type Certificate struct {
	Sn                   *Certificate_SN      `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	NotBefore            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NotAfter             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	CertDer              []byte               `protobuf:"bytes,4,opt,name=cert_der,json=certDer,proto3" json:"cert_der,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_309897dc79f61bc0, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetSn() *Certificate_SN {
	if m != nil {
		return m.Sn
	}
	return nil
}

func (m *Certificate) GetNotBefore() *timestamp.Timestamp {
	if m != nil {
		return m.NotBefore
	}
	return nil
}

func (m *Certificate) GetNotAfter() *timestamp.Timestamp {
	if m != nil {
		return m.NotAfter
	}
	return nil
}

func (m *Certificate) GetCertDer() []byte {
	if m != nil {
		return m.CertDer
	}
	return nil
}

type Certificate_SN struct {
	Sn                   string   `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Certificate_SN) Reset()         { *m = Certificate_SN{} }
func (m *Certificate_SN) String() string { return proto.CompactTextString(m) }
func (*Certificate_SN) ProtoMessage()    {}
func (*Certificate_SN) Descriptor() ([]byte, []int) {
	return fileDescriptor_309897dc79f61bc0, []int{1, 0}
}

func (m *Certificate_SN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate_SN.Unmarshal(m, b)
}
func (m *Certificate_SN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate_SN.Marshal(b, m, deterministic)
}
func (m *Certificate_SN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate_SN.Merge(m, src)
}
func (m *Certificate_SN) XXX_Size() int {
	return xxx_messageInfo_Certificate_SN.Size(m)
}
func (m *Certificate_SN) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate_SN.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate_SN proto.InternalMessageInfo

func (m *Certificate_SN) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

type CACert struct {
	Cert                 []byte   `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CACert) Reset()         { *m = CACert{} }
func (m *CACert) String() string { return proto.CompactTextString(m) }
func (*CACert) ProtoMessage()    {}
func (*CACert) Descriptor() ([]byte, []int) {
	return fileDescriptor_309897dc79f61bc0, []int{2}
}

func (m *CACert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CACert.Unmarshal(m, b)
}
func (m *CACert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CACert.Marshal(b, m, deterministic)
}
func (m *CACert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CACert.Merge(m, src)
}
func (m *CACert) XXX_Size() int {
	return xxx_messageInfo_CACert.Size(m)
}
func (m *CACert) XXX_DiscardUnknown() {
	xxx_messageInfo_CACert.DiscardUnknown(m)
}

var xxx_messageInfo_CACert proto.InternalMessageInfo

func (m *CACert) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func init() {
	proto.RegisterEnum("magma.orc8r.CertType", CertType_name, CertType_value)
	proto.RegisterType((*CSR)(nil), "magma.orc8r.CSR")
	proto.RegisterType((*Certificate)(nil), "magma.orc8r.Certificate")
	proto.RegisterType((*Certificate_SN)(nil), "magma.orc8r.Certificate.SN")
	proto.RegisterType((*CACert)(nil), "magma.orc8r.CACert")
}

func init() { proto.RegisterFile("orc8r/protos/certifier.proto", fileDescriptor_309897dc79f61bc0) }

var fileDescriptor_309897dc79f61bc0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x5d, 0xab, 0xd3, 0x30,
	0x18, 0xc7, 0x6d, 0x37, 0xd6, 0xf5, 0xe9, 0xe1, 0x70, 0x08, 0x8a, 0x5d, 0x77, 0xd0, 0x32, 0x10,
	0x86, 0x42, 0x0a, 0xc7, 0x0b, 0x8f, 0x97, 0x3b, 0xad, 0x82, 0x20, 0x43, 0xb2, 0xea, 0x85, 0x37,
	0xa5, 0x2f, 0x69, 0x09, 0xac, 0x4d, 0x49, 0x73, 0x84, 0x7d, 0x2e, 0x3f, 0x90, 0x5f, 0x45, 0x92,
	0xa6, 0xa2, 0x1b, 0x78, 0xd5, 0xa4, 0xcf, 0x2f, 0xfc, 0x5f, 0x78, 0xe0, 0x96, 0x8b, 0xf2, 0x5e,
	0x44, 0xbd, 0xe0, 0x92, 0x0f, 0x51, 0x49, 0x85, 0x64, 0x35, 0xa3, 0x02, 0xeb, 0x1f, 0xc8, 0x6b,
	0xf3, 0xa6, 0xcd, 0xb1, 0x66, 0x82, 0xf5, 0x3f, 0x28, 0xab, 0x68, 0x27, 0x99, 0x3c, 0x8d, 0x64,
	0xf0, 0xb2, 0xe1, 0xbc, 0x39, 0xd2, 0x71, 0x5a, 0x3c, 0xd6, 0x91, 0x64, 0x2d, 0x1d, 0x64, 0xde,
	0xf6, 0x06, 0x78, 0x71, 0x0e, 0x54, 0x8f, 0x22, 0x97, 0x8c, 0x77, 0xe3, 0x7c, 0xf3, 0xd3, 0x82,
	0x59, 0x7c, 0x20, 0xe8, 0x15, 0xd8, 0xac, 0xf2, 0xad, 0xd0, 0xda, 0x7a, 0x77, 0xcf, 0xf0, 0x5f,
	0xfa, 0xf8, 0x93, 0x51, 0x24, 0x36, 0xab, 0xd0, 0x3d, 0xc0, 0x8f, 0xfc, 0xc8, 0xaa, 0x4c, 0xe9,
	0xf8, 0xb6, 0xc6, 0x57, 0x78, 0xd4, 0xc0, 0x93, 0x06, 0x4e, 0x8c, 0x06, 0x71, 0x35, 0x9c, 0xb2,
	0x96, 0xa2, 0xe7, 0xe0, 0x94, 0x83, 0xc8, 0x2a, 0x2a, 0xfc, 0x59, 0x68, 0x6d, 0xaf, 0xc8, 0xa2,
	0x1c, 0x44, 0x42, 0x05, 0xba, 0x03, 0x57, 0xe5, 0xcf, 0xe4, 0xa9, 0xa7, 0xfe, 0x3c, 0xb4, 0xb6,
	0xd7, 0x67, 0x06, 0x62, 0x2a, 0x64, 0x7a, 0xea, 0x29, 0x59, 0x96, 0xe6, 0xb4, 0xf9, 0x65, 0x81,
	0x17, 0x8f, 0xa5, 0x95, 0xb9, 0xa4, 0xe8, 0x0d, 0xd8, 0x43, 0x67, 0xdc, 0xaf, 0x2f, 0x1e, 0x1b,
	0x0a, 0x1f, 0xf6, 0xc4, 0x1e, 0x3a, 0xf4, 0x1e, 0xa0, 0xe3, 0x32, 0x2b, 0x68, 0xcd, 0xc5, 0x94,
	0x21, 0xb8, 0xc8, 0x90, 0x4e, 0x45, 0x12, 0xb7, 0xe3, 0xf2, 0x41, 0xc3, 0xe8, 0x1d, 0xa8, 0x4b,
	0x96, 0xd7, 0xd2, 0xc4, 0xf8, 0xff, 0xcb, 0x65, 0xc7, 0xe5, 0x4e, 0xb1, 0x68, 0x05, 0xda, 0xbc,
	0x8e, 0x3f, 0xd7, 0xf1, 0x1d, 0x75, 0x4f, 0xa8, 0x08, 0x9e, 0x82, 0x7d, 0xd8, 0xa3, 0xeb, 0x3f,
	0x09, 0x5c, 0x65, 0x72, 0x73, 0x0b, 0x8b, 0x78, 0xa7, 0xcc, 0x23, 0x04, 0x73, 0x85, 0xea, 0xd9,
	0x15, 0xd1, 0xe7, 0xd7, 0x21, 0x2c, 0xa7, 0x56, 0x90, 0x07, 0x4e, 0xf2, 0xe1, 0xe3, 0xee, 0xeb,
	0xe7, 0xf4, 0xe6, 0x09, 0x72, 0x60, 0xf6, 0xed, 0xcb, 0xfe, 0xc6, 0x7a, 0x58, 0x7f, 0x5f, 0xe9,
	0x1a, 0xa2, 0x71, 0x7b, 0x8e, 0xac, 0x88, 0x1a, 0x6e, 0x96, 0xa8, 0x58, 0xe8, 0xef, 0xdb, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x4f, 0xa9, 0x0d, 0x86, 0x02, 0x00, 0x00,
}
