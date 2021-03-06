// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/core/cache.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// CacheEntry represents cache data for a file on disk.
type CacheEntry struct {
	// Mode stores the value of the POSIX mode bits (i.e. the st_mode member of
	// struct stat). On Windows, this value is computed using the Go os.FileMode
	// value retrieved through the os package (for which bit definitions are
	// guaranteed to be stable).
	Mode uint32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// ModificationTime is the cached modification time.
	ModificationTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=modificationTime,proto3" json:"modificationTime,omitempty"`
	// Size is the cached size.
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// FileID is the file identifier. On POSIX systems it is the inode number.
	// On Windows it is currently 0.
	FileID uint64 `protobuf:"varint,4,opt,name=fileID,proto3" json:"fileID,omitempty"`
	// Digest is the cached digest for file entries.
	Digest               []byte   `protobuf:"bytes,9,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CacheEntry) Reset()         { *m = CacheEntry{} }
func (m *CacheEntry) String() string { return proto.CompactTextString(m) }
func (*CacheEntry) ProtoMessage()    {}
func (*CacheEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56e0ad2085c3d36, []int{0}
}

func (m *CacheEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheEntry.Unmarshal(m, b)
}
func (m *CacheEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheEntry.Marshal(b, m, deterministic)
}
func (m *CacheEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheEntry.Merge(m, src)
}
func (m *CacheEntry) XXX_Size() int {
	return xxx_messageInfo_CacheEntry.Size(m)
}
func (m *CacheEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheEntry.DiscardUnknown(m)
}

var xxx_messageInfo_CacheEntry proto.InternalMessageInfo

func (m *CacheEntry) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *CacheEntry) GetModificationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ModificationTime
	}
	return nil
}

func (m *CacheEntry) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CacheEntry) GetFileID() uint64 {
	if m != nil {
		return m.FileID
	}
	return 0
}

func (m *CacheEntry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

// Cache provides a store for file metadata and digets to allow for efficient
// rescans.
type Cache struct {
	// Entries is a map from scan path to cache entry.
	Entries              map[string]*CacheEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Cache) Reset()         { *m = Cache{} }
func (m *Cache) String() string { return proto.CompactTextString(m) }
func (*Cache) ProtoMessage()    {}
func (*Cache) Descriptor() ([]byte, []int) {
	return fileDescriptor_b56e0ad2085c3d36, []int{1}
}

func (m *Cache) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cache.Unmarshal(m, b)
}
func (m *Cache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cache.Marshal(b, m, deterministic)
}
func (m *Cache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cache.Merge(m, src)
}
func (m *Cache) XXX_Size() int {
	return xxx_messageInfo_Cache.Size(m)
}
func (m *Cache) XXX_DiscardUnknown() {
	xxx_messageInfo_Cache.DiscardUnknown(m)
}

var xxx_messageInfo_Cache proto.InternalMessageInfo

func (m *Cache) GetEntries() map[string]*CacheEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheEntry)(nil), "core.CacheEntry")
	proto.RegisterType((*Cache)(nil), "core.Cache")
	proto.RegisterMapType((map[string]*CacheEntry)(nil), "core.Cache.EntriesEntry")
}

func init() { proto.RegisterFile("synchronization/core/cache.proto", fileDescriptor_b56e0ad2085c3d36) }

var fileDescriptor_b56e0ad2085c3d36 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6b, 0xc2, 0x30,
	0x14, 0xc6, 0x89, 0x56, 0x87, 0xd1, 0x81, 0xe4, 0x30, 0x82, 0x97, 0x15, 0x0f, 0xa3, 0x97, 0x25,
	0xe0, 0x60, 0xc8, 0x8e, 0xdb, 0x1c, 0x0c, 0x76, 0x2a, 0x9e, 0x76, 0xab, 0xf1, 0x19, 0x1f, 0x9a,
	0x44, 0xda, 0x74, 0xa0, 0xff, 0xc1, 0xfe, 0x97, 0xfd, 0x91, 0x23, 0xa9, 0x05, 0x61, 0xbb, 0x7d,
	0xef, 0xfb, 0xde, 0xfb, 0xb5, 0x5f, 0x4b, 0xd3, 0xea, 0x68, 0xd5, 0xb6, 0x74, 0x16, 0x4f, 0x85,
	0x47, 0x67, 0xa5, 0x72, 0x25, 0x48, 0x55, 0xa8, 0x2d, 0x88, 0x43, 0xe9, 0xbc, 0x63, 0x49, 0x70,
	0x26, 0xb7, 0xda, 0x39, 0xbd, 0x07, 0x19, 0xbd, 0x55, 0xbd, 0x91, 0x1e, 0x0d, 0x54, 0xbe, 0x30,
	0x87, 0x66, 0x6d, 0xfa, 0x43, 0x28, 0x7d, 0x09, 0x67, 0x0b, 0xeb, 0xcb, 0x23, 0x63, 0x34, 0x31,
	0x6e, 0x0d, 0x9c, 0xa4, 0x24, 0xbb, 0xce, 0xa3, 0x66, 0x6f, 0x74, 0x6c, 0xdc, 0x1a, 0x37, 0xa8,
	0xe2, 0xa3, 0x96, 0x68, 0x80, 0x77, 0x52, 0x92, 0x0d, 0x67, 0x13, 0xd1, 0xe0, 0x45, 0x8b, 0x17,
	0xcb, 0x16, 0x9f, 0xff, 0xb9, 0x09, 0xec, 0x0a, 0x4f, 0xc0, 0xbb, 0x29, 0xc9, 0x92, 0x3c, 0x6a,
	0x76, 0x43, 0xfb, 0x1b, 0xdc, 0xc3, 0xfb, 0x2b, 0x4f, 0xa2, 0x7b, 0x9e, 0x82, 0xbf, 0x46, 0x0d,
	0x95, 0xe7, 0x83, 0x94, 0x64, 0xa3, 0xfc, 0x3c, 0x4d, 0xbf, 0x09, 0xed, 0xc5, 0xd7, 0x65, 0x33,
	0x7a, 0x05, 0xd6, 0x97, 0x08, 0x15, 0x27, 0x69, 0x37, 0x1b, 0xce, 0xb8, 0x08, 0x8d, 0x45, 0x4c,
	0xc5, 0xa2, 0x89, 0x62, 0xa9, 0xbc, 0x5d, 0x9c, 0x7c, 0xd0, 0xd1, 0x65, 0xc0, 0xc6, 0xb4, 0xbb,
	0x83, 0x63, 0x2c, 0x3b, 0xc8, 0x83, 0x64, 0x77, 0xb4, 0xf7, 0x55, 0xec, 0xeb, 0xb6, 0xe0, 0xf8,
	0x82, 0xd9, 0xb0, 0x9a, 0xf8, 0xa9, 0x33, 0x27, 0xcf, 0xf3, 0xcf, 0x47, 0x8d, 0x7e, 0x5b, 0xaf,
	0x84, 0x72, 0x46, 0x9a, 0xda, 0x17, 0x1a, 0xec, 0x3d, 0xba, 0x56, 0xca, 0xc3, 0x4e, 0xcb, 0xff,
	0xfe, 0xd3, 0xaa, 0x1f, 0xbf, 0xd7, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x67, 0xab,
	0x4f, 0xc6, 0x01, 0x00, 0x00,
}
