// Code generated by protoc-gen-go. DO NOT EDIT.
// source: info.proto

package info

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type System struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	LastSeen             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=lastSeen,proto3" json:"lastSeen,omitempty"`
	Hostname             string               `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Address              string               `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Manufacturer         string               `protobuf:"bytes,7,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Product              string               `protobuf:"bytes,8,opt,name=product,proto3" json:"product,omitempty"`
	ProductVersion       string               `protobuf:"bytes,9,opt,name=productVersion,proto3" json:"productVersion,omitempty"`
	SerialNumber         string               `protobuf:"bytes,10,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	BiosVendor           string               `protobuf:"bytes,11,opt,name=biosVendor,proto3" json:"biosVendor,omitempty"`
	BiosDate             string               `protobuf:"bytes,12,opt,name=biosDate,proto3" json:"biosDate,omitempty"`
	BiosVersion          string               `protobuf:"bytes,13,opt,name=biosVersion,proto3" json:"biosVersion,omitempty"`
	BootRomVersion       string               `protobuf:"bytes,14,opt,name=bootRomVersion,proto3" json:"bootRomVersion,omitempty"`
	SmcVersion           string               `protobuf:"bytes,15,opt,name=smcVersion,proto3" json:"smcVersion,omitempty"`
	MemoryB              uint64               `protobuf:"varint,16,opt,name=memoryB,proto3" json:"memoryB,omitempty"`
	MemoryGb             uint32               `protobuf:"varint,17,opt,name=memoryGb,proto3" json:"memoryGb,omitempty"`
	CpuCoresPerSocket    uint32               `protobuf:"varint,18,opt,name=cpuCoresPerSocket,proto3" json:"cpuCoresPerSocket,omitempty"`
	CpuPhysicalCores     uint32               `protobuf:"varint,19,opt,name=cpuPhysicalCores,proto3" json:"cpuPhysicalCores,omitempty"`
	CpuLogicalCores      uint32               `protobuf:"varint,20,opt,name=cpuLogicalCores,proto3" json:"cpuLogicalCores,omitempty"`
	CpuSockets           uint32               `protobuf:"varint,21,opt,name=cpuSockets,proto3" json:"cpuSockets,omitempty"`
	CpuThreadsPerCore    uint32               `protobuf:"varint,22,opt,name=cpuThreadsPerCore,proto3" json:"cpuThreadsPerCore,omitempty"`
	CpuModel             string               `protobuf:"bytes,23,opt,name=cpuModel,proto3" json:"cpuModel,omitempty"`
	CpuFlags             string               `protobuf:"bytes,24,opt,name=cpuFlags,proto3" json:"cpuFlags,omitempty"`
	Os                   string               `protobuf:"bytes,25,opt,name=os,proto3" json:"os,omitempty"`
	OsVersion            string               `protobuf:"bytes,26,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	OsBuild              string               `protobuf:"bytes,27,opt,name=osBuild,proto3" json:"osBuild,omitempty"`
	OsDescription        string               `protobuf:"bytes,28,opt,name=osDescription,proto3" json:"osDescription,omitempty"`
	Kernel               string               `protobuf:"bytes,29,opt,name=kernel,proto3" json:"kernel,omitempty"`
	KernelVersion        string               `protobuf:"bytes,30,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	KernelRelease        string               `protobuf:"bytes,31,opt,name=kernelRelease,proto3" json:"kernelRelease,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *System) Reset()         { *m = System{} }
func (m *System) String() string { return proto.CompactTextString(m) }
func (*System) ProtoMessage()    {}
func (*System) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{1}
}

func (m *System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_System.Unmarshal(m, b)
}
func (m *System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_System.Marshal(b, m, deterministic)
}
func (m *System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_System.Merge(m, src)
}
func (m *System) XXX_Size() int {
	return xxx_messageInfo_System.Size(m)
}
func (m *System) XXX_DiscardUnknown() {
	xxx_messageInfo_System.DiscardUnknown(m)
}

var xxx_messageInfo_System proto.InternalMessageInfo

func (m *System) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *System) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *System) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *System) GetLastSeen() *timestamp.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *System) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *System) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *System) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *System) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *System) GetProductVersion() string {
	if m != nil {
		return m.ProductVersion
	}
	return ""
}

func (m *System) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *System) GetBiosVendor() string {
	if m != nil {
		return m.BiosVendor
	}
	return ""
}

func (m *System) GetBiosDate() string {
	if m != nil {
		return m.BiosDate
	}
	return ""
}

func (m *System) GetBiosVersion() string {
	if m != nil {
		return m.BiosVersion
	}
	return ""
}

func (m *System) GetBootRomVersion() string {
	if m != nil {
		return m.BootRomVersion
	}
	return ""
}

func (m *System) GetSmcVersion() string {
	if m != nil {
		return m.SmcVersion
	}
	return ""
}

func (m *System) GetMemoryB() uint64 {
	if m != nil {
		return m.MemoryB
	}
	return 0
}

func (m *System) GetMemoryGb() uint32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *System) GetCpuCoresPerSocket() uint32 {
	if m != nil {
		return m.CpuCoresPerSocket
	}
	return 0
}

func (m *System) GetCpuPhysicalCores() uint32 {
	if m != nil {
		return m.CpuPhysicalCores
	}
	return 0
}

func (m *System) GetCpuLogicalCores() uint32 {
	if m != nil {
		return m.CpuLogicalCores
	}
	return 0
}

func (m *System) GetCpuSockets() uint32 {
	if m != nil {
		return m.CpuSockets
	}
	return 0
}

func (m *System) GetCpuThreadsPerCore() uint32 {
	if m != nil {
		return m.CpuThreadsPerCore
	}
	return 0
}

func (m *System) GetCpuModel() string {
	if m != nil {
		return m.CpuModel
	}
	return ""
}

func (m *System) GetCpuFlags() string {
	if m != nil {
		return m.CpuFlags
	}
	return ""
}

func (m *System) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *System) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *System) GetOsBuild() string {
	if m != nil {
		return m.OsBuild
	}
	return ""
}

func (m *System) GetOsDescription() string {
	if m != nil {
		return m.OsDescription
	}
	return ""
}

func (m *System) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *System) GetKernelVersion() string {
	if m != nil {
		return m.KernelVersion
	}
	return ""
}

func (m *System) GetKernelRelease() string {
	if m != nil {
		return m.KernelRelease
	}
	return ""
}

type KeepAliveRequest struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeepAliveRequest) Reset()         { *m = KeepAliveRequest{} }
func (m *KeepAliveRequest) String() string { return proto.CompactTextString(m) }
func (*KeepAliveRequest) ProtoMessage()    {}
func (*KeepAliveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{2}
}

func (m *KeepAliveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAliveRequest.Unmarshal(m, b)
}
func (m *KeepAliveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAliveRequest.Marshal(b, m, deterministic)
}
func (m *KeepAliveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAliveRequest.Merge(m, src)
}
func (m *KeepAliveRequest) XXX_Size() int {
	return xxx_messageInfo_KeepAliveRequest.Size(m)
}
func (m *KeepAliveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAliveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAliveRequest proto.InternalMessageInfo

func (m *KeepAliveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *KeepAliveRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SystemList struct {
	Systems              []*System `protobuf:"bytes,1,rep,name=systems,proto3" json:"systems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SystemList) Reset()         { *m = SystemList{} }
func (m *SystemList) String() string { return proto.CompactTextString(m) }
func (*SystemList) ProtoMessage()    {}
func (*SystemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{3}
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

func (m *SystemList) GetSystems() []*System {
	if m != nil {
		return m.Systems
	}
	return nil
}

type SingleRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleRequest) Reset()         { *m = SingleRequest{} }
func (m *SingleRequest) String() string { return proto.CompactTextString(m) }
func (*SingleRequest) ProtoMessage()    {}
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{4}
}

func (m *SingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleRequest.Unmarshal(m, b)
}
func (m *SingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleRequest.Marshal(b, m, deterministic)
}
func (m *SingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleRequest.Merge(m, src)
}
func (m *SingleRequest) XXX_Size() int {
	return xxx_messageInfo_SingleRequest.Size(m)
}
func (m *SingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleRequest proto.InternalMessageInfo

func (m *SingleRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type ListRequest struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	Cursor               string   `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ListRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "info.Empty")
	proto.RegisterType((*System)(nil), "info.System")
	proto.RegisterType((*KeepAliveRequest)(nil), "info.KeepAliveRequest")
	proto.RegisterType((*SystemList)(nil), "info.SystemList")
	proto.RegisterType((*SingleRequest)(nil), "info.SingleRequest")
	proto.RegisterType((*ListRequest)(nil), "info.ListRequest")
}

func init() { proto.RegisterFile("info.proto", fileDescriptor_f140d5b28dddb141) }

var fileDescriptor_f140d5b28dddb141 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x6f, 0xd3, 0x48,
	0x14, 0x5e, 0x37, 0x69, 0x2e, 0x27, 0x4d, 0x9b, 0xce, 0xee, 0x76, 0x67, 0x43, 0x69, 0x23, 0x83,
	0x2a, 0x0b, 0xa1, 0x54, 0x0a, 0x15, 0xf0, 0x4a, 0x29, 0x54, 0x88, 0x82, 0x2a, 0xa7, 0xea, 0x33,
	0x8e, 0x7d, 0x92, 0x5a, 0xb5, 0x3d, 0x66, 0x2e, 0x48, 0xf9, 0x67, 0x3c, 0xf2, 0xd3, 0xd0, 0xcc,
	0xd8, 0x89, 0xdd, 0x22, 0xca, 0xdb, 0x7c, 0x97, 0x33, 0xf3, 0x9d, 0xe3, 0xe4, 0x00, 0xc4, 0xd9,
	0x9c, 0x8d, 0x73, 0xce, 0x24, 0x23, 0x4d, 0x7d, 0x1e, 0x1e, 0x2e, 0x18, 0x5b, 0x24, 0x78, 0x6c,
	0xb8, 0x99, 0x9a, 0x1f, 0xcb, 0x38, 0x45, 0x21, 0x83, 0x34, 0xb7, 0x36, 0xb7, 0x0d, 0x9b, 0xef,
	0xd2, 0x5c, 0x2e, 0xdd, 0xef, 0x1d, 0x68, 0x4d, 0x97, 0x42, 0x62, 0x4a, 0x08, 0x34, 0x95, 0x8a,
	0x23, 0xea, 0x8c, 0x1c, 0xaf, 0xeb, 0x9b, 0x33, 0x39, 0x81, 0x76, 0xc8, 0x31, 0x90, 0x18, 0xd1,
	0x8d, 0x91, 0xe3, 0xf5, 0x26, 0xc3, 0xb1, 0xbd, 0x7a, 0x5c, 0x5e, 0x3d, 0xbe, 0x2a, 0xaf, 0xf6,
	0x4b, 0xab, 0xae, 0x52, 0x79, 0x64, 0xaa, 0x1a, 0x0f, 0x57, 0x15, 0x56, 0xf2, 0x12, 0x3a, 0x49,
	0x20, 0xe4, 0x14, 0x31, 0xa3, 0xcd, 0x07, 0xcb, 0x56, 0x5e, 0x32, 0x84, 0xce, 0x0d, 0x13, 0x32,
	0x0b, 0x52, 0xa4, 0x9b, 0x26, 0xfb, 0x0a, 0x13, 0x0a, 0xed, 0x20, 0x8a, 0x38, 0x0a, 0x41, 0x5b,
	0x46, 0x2a, 0x21, 0x71, 0x61, 0x2b, 0x0d, 0x32, 0x35, 0x0f, 0x42, 0xa9, 0x38, 0x72, 0xda, 0x36,
	0x72, 0x8d, 0xd3, 0xd5, 0x39, 0x67, 0x91, 0x0a, 0x25, 0xed, 0xd8, 0xea, 0x02, 0x92, 0x23, 0xd8,
	0x2e, 0x8e, 0xd7, 0xc8, 0x45, 0xcc, 0x32, 0xda, 0x35, 0x86, 0x3b, 0xac, 0x7e, 0x45, 0x20, 0x8f,
	0x83, 0xe4, 0xb3, 0x4a, 0x67, 0xc8, 0x29, 0xd8, 0x57, 0xaa, 0x1c, 0x39, 0x00, 0x98, 0xc5, 0x4c,
	0x5c, 0x63, 0x16, 0x31, 0x4e, 0x7b, 0xc6, 0x51, 0x61, 0x74, 0x7f, 0x1a, 0x9d, 0x05, 0x12, 0xe9,
	0x96, 0xed, 0xaf, 0xc4, 0x64, 0x04, 0x3d, 0xeb, 0xb4, 0x21, 0xfa, 0x46, 0xae, 0x52, 0x3a, 0xe9,
	0x8c, 0x31, 0xe9, 0xb3, 0xb4, 0x34, 0x6d, 0xdb, 0xa4, 0x75, 0x56, 0xa7, 0x10, 0x69, 0x58, 0x7a,
	0x76, 0x6c, 0x8a, 0x35, 0xa3, 0x67, 0x91, 0x62, 0xca, 0xf8, 0xf2, 0x94, 0x0e, 0x46, 0x8e, 0xd7,
	0xf4, 0x4b, 0xa8, 0xf3, 0xd9, 0xe3, 0xf9, 0x8c, 0xee, 0x8e, 0x1c, 0xaf, 0xef, 0xaf, 0x30, 0x79,
	0x0e, 0xbb, 0x61, 0xae, 0xde, 0x32, 0x8e, 0xe2, 0x12, 0xf9, 0x94, 0x85, 0xb7, 0x28, 0x29, 0x31,
	0xa6, 0xfb, 0x02, 0x79, 0x06, 0x83, 0x30, 0x57, 0x97, 0x37, 0x4b, 0x11, 0x87, 0x41, 0x62, 0x44,
	0xfa, 0xb7, 0x31, 0xdf, 0xe3, 0x89, 0x07, 0x3b, 0x61, 0xae, 0x2e, 0xd8, 0x62, 0x6d, 0xfd, 0xc7,
	0x58, 0xef, 0xd2, 0xba, 0xb3, 0x30, 0x57, 0xf6, 0x09, 0x41, 0xff, 0x35, 0xa6, 0x0a, 0x53, 0x64,
	0xbc, 0xba, 0xe1, 0x18, 0x44, 0x3a, 0x8c, 0xae, 0xa2, 0x7b, 0xab, 0x8c, 0x75, 0x41, 0x77, 0x1b,
	0xe6, 0xea, 0x13, 0x8b, 0x30, 0xa1, 0xff, 0xd9, 0xaf, 0x51, 0xe2, 0x42, 0x7b, 0x9f, 0x04, 0x0b,
	0x41, 0xe9, 0x4a, 0x33, 0x98, 0x6c, 0xc3, 0x06, 0x13, 0xf4, 0x7f, 0xc3, 0x6e, 0x30, 0x41, 0xf6,
	0xa1, 0xbb, 0xfe, 0x6e, 0x43, 0x43, 0xaf, 0x09, 0x3d, 0x6d, 0x26, 0x4e, 0x55, 0x9c, 0x44, 0xf4,
	0x91, 0xfd, 0xe5, 0x15, 0x90, 0x3c, 0x85, 0x3e, 0x13, 0x67, 0x28, 0x42, 0x1e, 0xe7, 0x52, 0xd7,
	0xee, 0x1b, 0xbd, 0x4e, 0x92, 0x3d, 0x68, 0xdd, 0x22, 0xcf, 0x30, 0xa1, 0x8f, 0x8d, 0x5c, 0x20,
	0x5d, 0x6d, 0x4f, 0xe5, 0xcb, 0x07, 0xb6, 0xba, 0x46, 0xae, 0x5d, 0x3e, 0x26, 0x18, 0x08, 0xa4,
	0x87, 0x55, 0x57, 0x41, 0xba, 0x5f, 0x60, 0xf0, 0x11, 0x31, 0x7f, 0x93, 0xc4, 0xdf, 0xd0, 0xc7,
	0xaf, 0x0a, 0x85, 0xfc, 0xe5, 0x0e, 0x79, 0x0d, 0xdd, 0xd5, 0xfa, 0xf9, 0x83, 0x2d, 0xb2, 0x36,
	0xbb, 0x27, 0x00, 0x76, 0x37, 0x5d, 0xc4, 0x42, 0xff, 0xe7, 0xda, 0xc2, 0x20, 0x41, 0x9d, 0x51,
	0xc3, 0xeb, 0x4d, 0xb6, 0xc6, 0x66, 0xf1, 0x59, 0x8b, 0x5f, 0x8a, 0xee, 0x13, 0xe8, 0x4f, 0xe3,
	0x6c, 0x91, 0xfc, 0x2e, 0x94, 0xfb, 0x0a, 0x7a, 0xfa, 0xd2, 0xd2, 0x32, 0x80, 0x46, 0xa6, 0x52,
	0xe3, 0x68, 0xf8, 0xfa, 0xa8, 0x27, 0x18, 0x2a, 0x2e, 0x18, 0x37, 0x91, 0xbb, 0x7e, 0x81, 0x26,
	0x3f, 0x1c, 0x68, 0x7e, 0xc8, 0xe6, 0x8c, 0x78, 0xd0, 0x3d, 0x47, 0x59, 0xec, 0xce, 0x9e, 0x8d,
	0x62, 0x76, 0xea, 0xb0, 0x96, 0xcb, 0xfd, 0x8b, 0x1c, 0x41, 0xc7, 0xc7, 0x45, 0x2c, 0x24, 0x72,
	0x52, 0xd3, 0xea, 0x4e, 0x32, 0x81, 0xee, 0x6a, 0xa0, 0x64, 0xcf, 0x4a, 0x77, 0x27, 0x3c, 0xac,
	0xbe, 0xe4, 0x39, 0x64, 0x62, 0xfb, 0xb0, 0x37, 0x08, 0xb2, 0x6b, 0xd5, 0x4a, 0x6b, 0xc3, 0x41,
	0xf5, 0x0d, 0x2d, 0xcc, 0x5a, 0x66, 0xea, 0x2f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xad,
	0x68, 0xe4, 0x38, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoClient is the client API for Info service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoClient interface {
	GetSystem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*System, error)
	Register(ctx context.Context, in *System, opts ...grpc.CallOption) (*System, error)
	KeepAlive(ctx context.Context, opts ...grpc.CallOption) (Info_KeepAliveClient, error)
	ListSystems(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*SystemList, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) GetSystem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*System, error) {
	out := new(System)
	err := c.cc.Invoke(ctx, "/info.Info/GetSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) Register(ctx context.Context, in *System, opts ...grpc.CallOption) (*System, error) {
	out := new(System)
	err := c.cc.Invoke(ctx, "/info.Info/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoClient) KeepAlive(ctx context.Context, opts ...grpc.CallOption) (Info_KeepAliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Info_serviceDesc.Streams[0], "/info.Info/KeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &infoKeepAliveClient{stream}
	return x, nil
}

type Info_KeepAliveClient interface {
	Send(*KeepAliveRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type infoKeepAliveClient struct {
	grpc.ClientStream
}

func (x *infoKeepAliveClient) Send(m *KeepAliveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *infoKeepAliveClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *infoClient) ListSystems(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*SystemList, error) {
	out := new(SystemList)
	err := c.cc.Invoke(ctx, "/info.Info/ListSystems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServer is the server API for Info service.
type InfoServer interface {
	GetSystem(context.Context, *Empty) (*System, error)
	Register(context.Context, *System) (*System, error)
	KeepAlive(Info_KeepAliveServer) error
	ListSystems(context.Context, *ListRequest) (*SystemList, error)
}

// UnimplementedInfoServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServer struct {
}

func (*UnimplementedInfoServer) GetSystem(ctx context.Context, req *Empty) (*System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (*UnimplementedInfoServer) Register(ctx context.Context, req *System) (*System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedInfoServer) KeepAlive(srv Info_KeepAliveServer) error {
	return status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedInfoServer) ListSystems(ctx context.Context, req *ListRequest) (*SystemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystems not implemented")
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/GetSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).GetSystem(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(System)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Register(ctx, req.(*System))
	}
	return interceptor(ctx, in, info, handler)
}

func _Info_KeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InfoServer).KeepAlive(&infoKeepAliveServer{stream})
}

type Info_KeepAliveServer interface {
	SendAndClose(*Empty) error
	Recv() (*KeepAliveRequest, error)
	grpc.ServerStream
}

type infoKeepAliveServer struct {
	grpc.ServerStream
}

func (x *infoKeepAliveServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *infoKeepAliveServer) Recv() (*KeepAliveRequest, error) {
	m := new(KeepAliveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Info_ListSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).ListSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.Info/ListSystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).ListSystems(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystem",
			Handler:    _Info_GetSystem_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Info_Register_Handler,
		},
		{
			MethodName: "ListSystems",
			Handler:    _Info_ListSystems_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "KeepAlive",
			Handler:       _Info_KeepAlive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "info.proto",
}
