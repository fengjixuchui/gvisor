// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/sentry/unimpl/unimplemented_syscall.proto

package gvisor

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	registers_go_proto "gvisor.dev/gvisor/pkg/sentry/arch/registers_go_proto"
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

type UnimplementedSyscall struct {
	Tid                  int32                         `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Registers            *registers_go_proto.Registers `protobuf:"bytes,2,opt,name=registers,proto3" json:"registers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UnimplementedSyscall) Reset()         { *m = UnimplementedSyscall{} }
func (m *UnimplementedSyscall) String() string { return proto.CompactTextString(m) }
func (*UnimplementedSyscall) ProtoMessage()    {}
func (*UnimplementedSyscall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddc2fcd2bea3c75d, []int{0}
}

func (m *UnimplementedSyscall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnimplementedSyscall.Unmarshal(m, b)
}
func (m *UnimplementedSyscall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnimplementedSyscall.Marshal(b, m, deterministic)
}
func (m *UnimplementedSyscall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnimplementedSyscall.Merge(m, src)
}
func (m *UnimplementedSyscall) XXX_Size() int {
	return xxx_messageInfo_UnimplementedSyscall.Size(m)
}
func (m *UnimplementedSyscall) XXX_DiscardUnknown() {
	xxx_messageInfo_UnimplementedSyscall.DiscardUnknown(m)
}

var xxx_messageInfo_UnimplementedSyscall proto.InternalMessageInfo

func (m *UnimplementedSyscall) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *UnimplementedSyscall) GetRegisters() *registers_go_proto.Registers {
	if m != nil {
		return m.Registers
	}
	return nil
}

func init() {
	proto.RegisterType((*UnimplementedSyscall)(nil), "gvisor.UnimplementedSyscall")
}

func init() {
	proto.RegisterFile("pkg/sentry/unimpl/unimplemented_syscall.proto", fileDescriptor_ddc2fcd2bea3c75d)
}

var fileDescriptor_ddc2fcd2bea3c75d = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0xc8, 0x4e, 0xd7,
	0x2f, 0x4e, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2f, 0xcd, 0xcb, 0xcc, 0x2d, 0xc8, 0x81, 0x52, 0xa9,
	0xb9, 0xa9, 0x79, 0x25, 0xa9, 0x29, 0xf1, 0xc5, 0x95, 0xc5, 0xc9, 0x89, 0x39, 0x39, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0xe9, 0x65, 0x99, 0xc5, 0xf9, 0x45, 0x52, 0xf2, 0x48, 0xda,
	0x12, 0x8b, 0x92, 0x33, 0xf4, 0x8b, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x8a, 0x21, 0x0a,
	0x95, 0x22, 0xb9, 0x44, 0x42, 0x91, 0xcd, 0x09, 0x86, 0x18, 0x23, 0x24, 0xc0, 0xc5, 0x5c, 0x92,
	0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x0a, 0xe9, 0x73, 0x71, 0xc2, 0x35,
	0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xea, 0x41, 0xac, 0xd1, 0x0b, 0x82, 0x49, 0x04,
	0x21, 0xd4, 0x24, 0xb1, 0x81, 0x6d, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51, 0x4a, 0x47,
	0x79, 0xbb, 0x00, 0x00, 0x00,
}
