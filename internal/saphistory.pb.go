// Code generated by protoc-gen-go. DO NOT EDIT.
// source: saphistory.proto

package internal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Secret struct {
	Name                 map[string][]byte `protobuf:"bytes,1,rep,name=Name,proto3" json:"Name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{0}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() map[string][]byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type D1StringList struct {
	Data1                []string `protobuf:"bytes,1,rep,name=Data1,proto3" json:"Data1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *D1StringList) Reset()         { *m = D1StringList{} }
func (m *D1StringList) String() string { return proto.CompactTextString(m) }
func (*D1StringList) ProtoMessage()    {}
func (*D1StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{1}
}

func (m *D1StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_D1StringList.Unmarshal(m, b)
}
func (m *D1StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_D1StringList.Marshal(b, m, deterministic)
}
func (m *D1StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_D1StringList.Merge(m, src)
}
func (m *D1StringList) XXX_Size() int {
	return xxx_messageInfo_D1StringList.Size(m)
}
func (m *D1StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_D1StringList.DiscardUnknown(m)
}

var xxx_messageInfo_D1StringList proto.InternalMessageInfo

func (m *D1StringList) GetData1() []string {
	if m != nil {
		return m.Data1
	}
	return nil
}

type D2StringList struct {
	Data2                []*D1StringList `protobuf:"bytes,1,rep,name=Data2,proto3" json:"Data2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *D2StringList) Reset()         { *m = D2StringList{} }
func (m *D2StringList) String() string { return proto.CompactTextString(m) }
func (*D2StringList) ProtoMessage()    {}
func (*D2StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{2}
}

func (m *D2StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_D2StringList.Unmarshal(m, b)
}
func (m *D2StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_D2StringList.Marshal(b, m, deterministic)
}
func (m *D2StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_D2StringList.Merge(m, src)
}
func (m *D2StringList) XXX_Size() int {
	return xxx_messageInfo_D2StringList.Size(m)
}
func (m *D2StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_D2StringList.DiscardUnknown(m)
}

var xxx_messageInfo_D2StringList proto.InternalMessageInfo

func (m *D2StringList) GetData2() []*D1StringList {
	if m != nil {
		return m.Data2
	}
	return nil
}

type D3StringList struct {
	Data3                []*D2StringList `protobuf:"bytes,1,rep,name=Data3,proto3" json:"Data3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *D3StringList) Reset()         { *m = D3StringList{} }
func (m *D3StringList) String() string { return proto.CompactTextString(m) }
func (*D3StringList) ProtoMessage()    {}
func (*D3StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{3}
}

func (m *D3StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_D3StringList.Unmarshal(m, b)
}
func (m *D3StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_D3StringList.Marshal(b, m, deterministic)
}
func (m *D3StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_D3StringList.Merge(m, src)
}
func (m *D3StringList) XXX_Size() int {
	return xxx_messageInfo_D3StringList.Size(m)
}
func (m *D3StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_D3StringList.DiscardUnknown(m)
}

var xxx_messageInfo_D3StringList proto.InternalMessageInfo

func (m *D3StringList) GetData3() []*D2StringList {
	if m != nil {
		return m.Data3
	}
	return nil
}

type TransactionData struct {
	Ts                   string        `protobuf:"bytes,1,opt,name=Ts,proto3" json:"Ts,omitempty"`
	Tcodes               *D1StringList `protobuf:"bytes,4,opt,name=Tcodes,proto3" json:"Tcodes,omitempty"`
	Hdata                *D2StringList `protobuf:"bytes,5,opt,name=Hdata,proto3" json:"Hdata,omitempty"`
	Tdata                *D3StringList `protobuf:"bytes,6,opt,name=Tdata,proto3" json:"Tdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}
func (*TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{4}
}

func (m *TransactionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionData.Unmarshal(m, b)
}
func (m *TransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionData.Marshal(b, m, deterministic)
}
func (m *TransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionData.Merge(m, src)
}
func (m *TransactionData) XXX_Size() int {
	return xxx_messageInfo_TransactionData.Size(m)
}
func (m *TransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionData proto.InternalMessageInfo

func (m *TransactionData) GetTs() string {
	if m != nil {
		return m.Ts
	}
	return ""
}

func (m *TransactionData) GetTcodes() *D1StringList {
	if m != nil {
		return m.Tcodes
	}
	return nil
}

func (m *TransactionData) GetHdata() *D2StringList {
	if m != nil {
		return m.Hdata
	}
	return nil
}

func (m *TransactionData) GetTdata() *D3StringList {
	if m != nil {
		return m.Tdata
	}
	return nil
}

type SapSelection struct {
	Sid                  string   `protobuf:"bytes,2,opt,name=Sid,proto3" json:"Sid,omitempty"`
	Ts                   string   `protobuf:"bytes,3,opt,name=Ts,proto3" json:"Ts,omitempty"`
	Direction            int32    `protobuf:"varint,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SapSelection) Reset()         { *m = SapSelection{} }
func (m *SapSelection) String() string { return proto.CompactTextString(m) }
func (*SapSelection) ProtoMessage()    {}
func (*SapSelection) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{5}
}

func (m *SapSelection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SapSelection.Unmarshal(m, b)
}
func (m *SapSelection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SapSelection.Marshal(b, m, deterministic)
}
func (m *SapSelection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SapSelection.Merge(m, src)
}
func (m *SapSelection) XXX_Size() int {
	return xxx_messageInfo_SapSelection.Size(m)
}
func (m *SapSelection) XXX_DiscardUnknown() {
	xxx_messageInfo_SapSelection.DiscardUnknown(m)
}

var xxx_messageInfo_SapSelection proto.InternalMessageInfo

func (m *SapSelection) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *SapSelection) GetTs() string {
	if m != nil {
		return m.Ts
	}
	return ""
}

func (m *SapSelection) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type SystemInfo struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=Sid,proto3" json:"Sid,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=Client,proto3" json:"Client,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Sysnr                string   `protobuf:"bytes,4,opt,name=Sysnr,proto3" json:"Sysnr,omitempty"`
	Hostname             string   `protobuf:"bytes,5,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	Username             string   `protobuf:"bytes,6,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,7,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{6}
}

func (m *SystemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemInfo.Unmarshal(m, b)
}
func (m *SystemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemInfo.Marshal(b, m, deterministic)
}
func (m *SystemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemInfo.Merge(m, src)
}
func (m *SystemInfo) XXX_Size() int {
	return xxx_messageInfo_SystemInfo.Size(m)
}
func (m *SystemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemInfo proto.InternalMessageInfo

func (m *SystemInfo) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *SystemInfo) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *SystemInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SystemInfo) GetSysnr() string {
	if m != nil {
		return m.Sysnr
	}
	return ""
}

func (m *SystemInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *SystemInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SystemInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SystemList struct {
	Systems              []*SystemInfo `protobuf:"bytes,1,rep,name=Systems,proto3" json:"Systems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SystemList) Reset()         { *m = SystemList{} }
func (m *SystemList) String() string { return proto.CompactTextString(m) }
func (*SystemList) ProtoMessage()    {}
func (*SystemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{7}
}

func (m *SystemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemList.Unmarshal(m, b)
}
func (m *SystemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemList.Marshal(b, m, deterministic)
}
func (m *SystemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemList.Merge(m, src)
}
func (m *SystemList) XXX_Size() int {
	return xxx_messageInfo_SystemList.Size(m)
}
func (m *SystemList) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemList.DiscardUnknown(m)
}

var xxx_messageInfo_SystemList proto.InternalMessageInfo

func (m *SystemList) GetSystems() []*SystemInfo {
	if m != nil {
		return m.Systems
	}
	return nil
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_164980ab29709afc, []int{8}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Secret)(nil), "private.Secret")
	proto.RegisterMapType((map[string][]byte)(nil), "private.Secret.NameEntry")
	proto.RegisterType((*D1StringList)(nil), "private.D1StringList")
	proto.RegisterType((*D2StringList)(nil), "private.D2StringList")
	proto.RegisterType((*D3StringList)(nil), "private.D3StringList")
	proto.RegisterType((*TransactionData)(nil), "private.TransactionData")
	proto.RegisterType((*SapSelection)(nil), "private.SapSelection")
	proto.RegisterType((*SystemInfo)(nil), "private.SystemInfo")
	proto.RegisterType((*SystemList)(nil), "private.SystemList")
	proto.RegisterType((*Nothing)(nil), "private.Nothing")
}

func init() { proto.RegisterFile("saphistory.proto", fileDescriptor_164980ab29709afc) }

var fileDescriptor_164980ab29709afc = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xe5, 0x76, 0x9b, 0x92, 0xd9, 0x00, 0xc5, 0x7c, 0x28, 0x54, 0x1c, 0xaa, 0x88, 0x43,
	0x25, 0xd4, 0x4a, 0x9b, 0x8a, 0x0f, 0xd1, 0x03, 0x12, 0x04, 0x51, 0x24, 0x54, 0xa1, 0xa4, 0x7b,
	0xe1, 0x66, 0x1a, 0xb3, 0x6b, 0x91, 0x75, 0x22, 0xdb, 0x5b, 0x94, 0x57, 0xe2, 0x31, 0x78, 0x2c,
	0x4e, 0xc8, 0x76, 0xbe, 0x48, 0xe1, 0xb0, 0xb7, 0xcc, 0xfc, 0x7f, 0x33, 0xf3, 0xb7, 0xc7, 0x2d,
	0x4c, 0x24, 0x29, 0x2e, 0x99, 0x54, 0xb9, 0x28, 0x97, 0x85, 0xc8, 0x55, 0x8e, 0xc7, 0x85, 0x60,
	0x07, 0xa2, 0x68, 0x50, 0x80, 0x93, 0xd0, 0xbd, 0xa0, 0x0a, 0x2f, 0xe0, 0x64, 0x4b, 0xae, 0xa8,
	0x8f, 0x66, 0xc3, 0xf9, 0x69, 0xf8, 0x78, 0x59, 0x11, 0x4b, 0x2b, 0x2f, 0xb5, 0xf6, 0x9e, 0x2b,
	0x51, 0xc6, 0x06, 0x9b, 0xbe, 0x04, 0xb7, 0x49, 0xe1, 0x09, 0x0c, 0xbf, 0xd3, 0xd2, 0x47, 0x33,
	0x34, 0x77, 0x63, 0xfd, 0x89, 0x1f, 0xc0, 0xe8, 0x40, 0xb2, 0x6b, 0xea, 0x0f, 0x66, 0x68, 0xee,
	0xc5, 0x36, 0x78, 0x3d, 0x78, 0x85, 0x82, 0xa7, 0xe0, 0x45, 0x67, 0x89, 0x12, 0x8c, 0x5f, 0x7c,
	0x62, 0x52, 0x69, 0x32, 0x22, 0x8a, 0x9c, 0x99, 0xc1, 0x6e, 0x6c, 0x83, 0x60, 0x0d, 0x5e, 0x14,
	0x76, 0xa8, 0x67, 0x96, 0x0a, 0x2b, 0x7b, 0x0f, 0x1b, 0x7b, 0xdd, 0x5e, 0xb6, 0x38, 0x34, 0xc5,
	0xab, 0xe3, 0xe2, 0xd5, 0x71, 0x71, 0xd8, 0x2f, 0x5e, 0x05, 0x3f, 0x11, 0xdc, 0xdd, 0x09, 0xc2,
	0x25, 0xd9, 0x2b, 0x96, 0x73, 0x9d, 0xc4, 0x77, 0x60, 0xb0, 0x93, 0xd5, 0xf1, 0x06, 0x3b, 0x89,
	0x17, 0xe0, 0xec, 0xf6, 0x79, 0x4a, 0xa5, 0x7f, 0x32, 0x43, 0xff, 0xb7, 0x53, 0x41, 0x7a, 0xfe,
	0x26, 0x25, 0x8a, 0xf8, 0xa3, 0x3e, 0xfd, 0xd7, 0x7c, 0xc3, 0x68, 0x78, 0x67, 0x60, 0xa7, 0x0f,
	0xaf, 0xba, 0xb0, 0x61, 0x82, 0x2d, 0x78, 0x09, 0x29, 0x12, 0x9a, 0x51, 0x63, 0x56, 0x2f, 0x22,
	0x61, 0xa9, 0xb9, 0x74, 0x37, 0xd6, 0x9f, 0x95, 0xf5, 0x61, 0x63, 0xfd, 0x09, 0xb8, 0x11, 0x13,
	0x16, 0x37, 0xee, 0x47, 0x71, 0x9b, 0x08, 0x7e, 0x21, 0x80, 0xa4, 0x94, 0x8a, 0x5e, 0x7d, 0xe4,
	0xdf, 0xf2, 0xba, 0x1d, 0x6a, 0xdb, 0x3d, 0x02, 0xe7, 0x5d, 0xc6, 0x28, 0x57, 0xd5, 0x8c, 0x2a,
	0xc2, 0x33, 0x38, 0x8d, 0xa8, 0xdc, 0x0b, 0x56, 0x98, 0xc6, 0x76, 0x5e, 0x37, 0xa5, 0xf7, 0x9c,
	0x94, 0x92, 0x0b, 0x33, 0xd4, 0x8d, 0x6d, 0x80, 0xa7, 0x70, 0x6b, 0x93, 0x4b, 0xc5, 0xf5, 0xcb,
	0x1b, 0x19, 0xa1, 0x89, 0xb5, 0x76, 0x2e, 0xa9, 0x30, 0x9a, 0x63, 0xb5, 0x3a, 0xd6, 0xda, 0x67,
	0x22, 0xe5, 0x8f, 0x5c, 0xa4, 0xfe, 0xd8, 0x6a, 0x75, 0x1c, 0xac, 0xeb, 0x33, 0x98, 0xe5, 0x2f,
	0x60, 0x6c, 0x23, 0x59, 0xad, 0xff, 0x7e, 0xfb, 0xb4, 0x9b, 0x93, 0xc6, 0x35, 0x13, 0xb8, 0x30,
	0xde, 0xe6, 0xea, 0x92, 0xf1, 0x8b, 0xf0, 0x37, 0x82, 0x7b, 0x09, 0x29, 0x36, 0xf6, 0x97, 0x93,
	0x50, 0x71, 0x60, 0x7b, 0x8a, 0xdf, 0x80, 0xf7, 0x81, 0x2a, 0x7d, 0xeb, 0x8a, 0xa8, 0x6b, 0x89,
	0xdb, 0x05, 0x75, 0x37, 0x31, 0xf5, 0x9b, 0x74, 0xff, 0x31, 0xbd, 0x80, 0xdb, 0xba, 0x41, 0xeb,
	0x70, 0xd2, 0xa0, 0xd5, 0xe4, 0x69, 0xdf, 0xa2, 0xc1, 0x9e, 0x83, 0x77, 0x5e, 0xa4, 0x44, 0x51,
	0x9b, 0xc3, 0xff, 0x3a, 0xc7, 0xf4, 0xa8, 0x97, 0x2e, 0x8b, 0x68, 0x46, 0x6f, 0x58, 0xf6, 0xd6,
	0xfb, 0x02, 0xcb, 0x35, 0xe3, 0x4a, 0xdf, 0x77, 0xf6, 0xd5, 0x31, 0x7f, 0x1b, 0xab, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0xd3, 0x41, 0xb3, 0x4a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SapHistoryServiceClient is the client API for SapHistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SapHistoryServiceClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetSapStatus(ctx context.Context, in *SapSelection, opts ...grpc.CallOption) (*TransactionData, error)
	GetSystemList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*SystemList, error)
	UpdateSystem(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Nothing, error)
	DeleteSystem(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Nothing, error)
}

type sapHistoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSapHistoryServiceClient(cc *grpc.ClientConn) SapHistoryServiceClient {
	return &sapHistoryServiceClient{cc}
}

func (c *sapHistoryServiceClient) GetSapStatus(ctx context.Context, in *SapSelection, opts ...grpc.CallOption) (*TransactionData, error) {
	out := new(TransactionData)
	err := c.cc.Invoke(ctx, "/private.SapHistoryService/GetSapStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sapHistoryServiceClient) GetSystemList(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*SystemList, error) {
	out := new(SystemList)
	err := c.cc.Invoke(ctx, "/private.SapHistoryService/GetSystemList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sapHistoryServiceClient) UpdateSystem(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/private.SapHistoryService/UpdateSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sapHistoryServiceClient) DeleteSystem(ctx context.Context, in *SystemInfo, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/private.SapHistoryService/DeleteSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SapHistoryServiceServer is the server API for SapHistoryService service.
type SapHistoryServiceServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetSapStatus(context.Context, *SapSelection) (*TransactionData, error)
	GetSystemList(context.Context, *Nothing) (*SystemList, error)
	UpdateSystem(context.Context, *SystemInfo) (*Nothing, error)
	DeleteSystem(context.Context, *SystemInfo) (*Nothing, error)
}

// UnimplementedSapHistoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSapHistoryServiceServer struct {
}

func (*UnimplementedSapHistoryServiceServer) GetSapStatus(ctx context.Context, req *SapSelection) (*TransactionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSapStatus not implemented")
}
func (*UnimplementedSapHistoryServiceServer) GetSystemList(ctx context.Context, req *Nothing) (*SystemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemList not implemented")
}
func (*UnimplementedSapHistoryServiceServer) UpdateSystem(ctx context.Context, req *SystemInfo) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSystem not implemented")
}
func (*UnimplementedSapHistoryServiceServer) DeleteSystem(ctx context.Context, req *SystemInfo) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSystem not implemented")
}

func RegisterSapHistoryServiceServer(s *grpc.Server, srv SapHistoryServiceServer) {
	s.RegisterService(&_SapHistoryService_serviceDesc, srv)
}

func _SapHistoryService_GetSapStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SapSelection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SapHistoryServiceServer).GetSapStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.SapHistoryService/GetSapStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SapHistoryServiceServer).GetSapStatus(ctx, req.(*SapSelection))
	}
	return interceptor(ctx, in, info, handler)
}

func _SapHistoryService_GetSystemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SapHistoryServiceServer).GetSystemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.SapHistoryService/GetSystemList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SapHistoryServiceServer).GetSystemList(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _SapHistoryService_UpdateSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SapHistoryServiceServer).UpdateSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.SapHistoryService/UpdateSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SapHistoryServiceServer).UpdateSystem(ctx, req.(*SystemInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SapHistoryService_DeleteSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SapHistoryServiceServer).DeleteSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.SapHistoryService/DeleteSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SapHistoryServiceServer).DeleteSystem(ctx, req.(*SystemInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _SapHistoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "private.SapHistoryService",
	HandlerType: (*SapHistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSapStatus",
			Handler:    _SapHistoryService_GetSapStatus_Handler,
		},
		{
			MethodName: "GetSystemList",
			Handler:    _SapHistoryService_GetSystemList_Handler,
		},
		{
			MethodName: "UpdateSystem",
			Handler:    _SapHistoryService_UpdateSystem_Handler,
		},
		{
			MethodName: "DeleteSystem",
			Handler:    _SapHistoryService_DeleteSystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saphistory.proto",
}
