// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarding/endpoint/remote/protocol.proto

package remote

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/havoc-io/mutagen/pkg/forwarding"
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

// InitializeForwardingRequest is the initialization request sent to remote
// forwarding endpoints.
type InitializeForwardingRequest struct {
	// Field 1 is reserved for the session identifier.
	// Version is the session version.
	Version forwarding.Version `protobuf:"varint,2,opt,name=version,proto3,enum=forwarding.Version" json:"version,omitempty"`
	// Configuration is the session configuration.
	Configuration *forwarding.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Protocol is the protocol specification for the listener/dialer.
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Address it the bind address for a listener or dial address for a dialer.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Listener indicates whether this endpoint should function as a listener or
	// dialer for the associated address.
	Listener             bool     `protobuf:"varint,6,opt,name=listener,proto3" json:"listener,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeForwardingRequest) Reset()         { *m = InitializeForwardingRequest{} }
func (m *InitializeForwardingRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeForwardingRequest) ProtoMessage()    {}
func (*InitializeForwardingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb0a3c00a9f3e65, []int{0}
}

func (m *InitializeForwardingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeForwardingRequest.Unmarshal(m, b)
}
func (m *InitializeForwardingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeForwardingRequest.Marshal(b, m, deterministic)
}
func (m *InitializeForwardingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeForwardingRequest.Merge(m, src)
}
func (m *InitializeForwardingRequest) XXX_Size() int {
	return xxx_messageInfo_InitializeForwardingRequest.Size(m)
}
func (m *InitializeForwardingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeForwardingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeForwardingRequest proto.InternalMessageInfo

func (m *InitializeForwardingRequest) GetVersion() forwarding.Version {
	if m != nil {
		return m.Version
	}
	return forwarding.Version_Invalid
}

func (m *InitializeForwardingRequest) GetConfiguration() *forwarding.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *InitializeForwardingRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *InitializeForwardingRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *InitializeForwardingRequest) GetListener() bool {
	if m != nil {
		return m.Listener
	}
	return false
}

// InitializeForwardingResponse is the initialization response sent from remote
// forwarding endpoint.
type InitializeForwardingResponse struct {
	// Error is any error that occurred during initialization.
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeForwardingResponse) Reset()         { *m = InitializeForwardingResponse{} }
func (m *InitializeForwardingResponse) String() string { return proto.CompactTextString(m) }
func (*InitializeForwardingResponse) ProtoMessage()    {}
func (*InitializeForwardingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb0a3c00a9f3e65, []int{1}
}

func (m *InitializeForwardingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeForwardingResponse.Unmarshal(m, b)
}
func (m *InitializeForwardingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeForwardingResponse.Marshal(b, m, deterministic)
}
func (m *InitializeForwardingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeForwardingResponse.Merge(m, src)
}
func (m *InitializeForwardingResponse) XXX_Size() int {
	return xxx_messageInfo_InitializeForwardingResponse.Size(m)
}
func (m *InitializeForwardingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeForwardingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeForwardingResponse proto.InternalMessageInfo

func (m *InitializeForwardingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*InitializeForwardingRequest)(nil), "remote.InitializeForwardingRequest")
	proto.RegisterType((*InitializeForwardingResponse)(nil), "remote.InitializeForwardingResponse")
}

func init() {
	proto.RegisterFile("forwarding/endpoint/remote/protocol.proto", fileDescriptor_4bb0a3c00a9f3e65)
}

var fileDescriptor_4bb0a3c00a9f3e65 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x89, 0xba, 0x7f, 0x8c, 0xe8, 0x21, 0x7a, 0x88, 0x55, 0xa4, 0xec, 0xa9, 0x1e, 0xb6,
	0x81, 0xd5, 0x93, 0x08, 0x82, 0x82, 0xe0, 0xb5, 0x07, 0x0f, 0xde, 0xb2, 0xed, 0x6c, 0x37, 0xd8,
	0x66, 0xea, 0x24, 0x5d, 0xc1, 0x2f, 0xeb, 0x57, 0x11, 0x5b, 0x5b, 0xbb, 0xa0, 0xb7, 0xbc, 0xbc,
	0xdf, 0x9b, 0xe4, 0x31, 0xfc, 0x72, 0x85, 0xf4, 0xae, 0x29, 0x33, 0x36, 0x57, 0x60, 0xb3, 0x0a,
	0x8d, 0xf5, 0x8a, 0xa0, 0x44, 0x0f, 0xaa, 0x22, 0xf4, 0x98, 0x62, 0x11, 0x37, 0x07, 0x31, 0x6e,
	0xaf, 0x83, 0x8b, 0x41, 0x24, 0x45, 0xbb, 0x32, 0x79, 0x4d, 0xda, 0x1b, 0xb4, 0x2d, 0x17, 0xc8,
	0x81, 0xbf, 0x01, 0x72, 0xbd, 0x33, 0xfb, 0x64, 0xfc, 0xec, 0xc9, 0x1a, 0x6f, 0x74, 0x61, 0x3e,
	0xe0, 0xb1, 0xc7, 0x12, 0x78, 0xab, 0xc1, 0x79, 0x31, 0xe7, 0x93, 0x9f, 0x80, 0xdc, 0x09, 0x59,
	0x74, 0xb4, 0x38, 0x8e, 0x7f, 0x67, 0xc5, 0xcf, 0xad, 0x95, 0x74, 0x8c, 0xb8, 0xe3, 0x87, 0x5b,
	0xef, 0xcb, 0xdd, 0x90, 0x45, 0x07, 0x8b, 0xd3, 0x61, 0xe8, 0x61, 0x08, 0x24, 0xdb, 0xbc, 0x08,
	0xf8, 0xb4, 0xeb, 0x28, 0xf7, 0x42, 0x16, 0xed, 0x27, 0xbd, 0x16, 0x92, 0x4f, 0x74, 0x96, 0x11,
	0x38, 0x27, 0x47, 0x8d, 0xd5, 0xc9, 0xef, 0x54, 0x61, 0x9c, 0x07, 0x0b, 0x24, 0xc7, 0x21, 0x8b,
	0xa6, 0x49, 0xaf, 0x67, 0xd7, 0xfc, 0xfc, 0xef, 0x82, 0xae, 0x42, 0xeb, 0x40, 0x9c, 0xf0, 0x11,
	0x10, 0x21, 0x49, 0xd6, 0xcc, 0x6c, 0xc5, 0xfd, 0xed, 0xcb, 0x4d, 0x6e, 0xfc, 0xba, 0x5e, 0xc6,
	0x29, 0x96, 0x6a, 0xad, 0x37, 0x98, 0xce, 0x0d, 0xaa, 0xb2, 0xf6, 0x3a, 0x07, 0xab, 0xaa, 0xd7,
	0x5c, 0xfd, 0xbf, 0xa6, 0xe5, 0xb8, 0xf9, 0xf3, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24,
	0x98, 0xbd, 0xd0, 0xcb, 0x01, 0x00, 0x00,
}
