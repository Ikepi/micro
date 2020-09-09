// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/store.proto

package store

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Field struct {
	// type of value e.g string, int, int64, bool, float64
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// the actual value
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Record struct {
	// key of the record
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value in the record
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// time.Duration (signed int64 nanoseconds)
	Expiry int64 `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// the associated metadata
	Metadata             map[string]*Field `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *Record) GetMetadata() map[string]*Field {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ReadOptions struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Prefix               bool     `protobuf:"varint,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix               bool     `protobuf:"varint,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Limit                uint64   `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{2}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

func (m *ReadOptions) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ReadOptions) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ReadOptions) GetPrefix() bool {
	if m != nil {
		return m.Prefix
	}
	return false
}

func (m *ReadOptions) GetSuffix() bool {
	if m != nil {
		return m.Suffix
	}
	return false
}

func (m *ReadOptions) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadOptions) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ReadRequest struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Options              *ReadOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReadRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type WriteOptions struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteOptions) Reset()         { *m = WriteOptions{} }
func (m *WriteOptions) String() string { return proto.CompactTextString(m) }
func (*WriteOptions) ProtoMessage()    {}
func (*WriteOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{5}
}

func (m *WriteOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteOptions.Unmarshal(m, b)
}
func (m *WriteOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteOptions.Marshal(b, m, deterministic)
}
func (m *WriteOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteOptions.Merge(m, src)
}
func (m *WriteOptions) XXX_Size() int {
	return xxx_messageInfo_WriteOptions.Size(m)
}
func (m *WriteOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteOptions.DiscardUnknown(m)
}

var xxx_messageInfo_WriteOptions proto.InternalMessageInfo

func (m *WriteOptions) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *WriteOptions) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

type WriteRequest struct {
	Record               *Record       `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Options              *WriteOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{6}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *WriteRequest) GetOptions() *WriteOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{7}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type DeleteOptions struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOptions) Reset()         { *m = DeleteOptions{} }
func (m *DeleteOptions) String() string { return proto.CompactTextString(m) }
func (*DeleteOptions) ProtoMessage()    {}
func (*DeleteOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{8}
}

func (m *DeleteOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOptions.Unmarshal(m, b)
}
func (m *DeleteOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOptions.Marshal(b, m, deterministic)
}
func (m *DeleteOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOptions.Merge(m, src)
}
func (m *DeleteOptions) XXX_Size() int {
	return xxx_messageInfo_DeleteOptions.Size(m)
}
func (m *DeleteOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOptions proto.InternalMessageInfo

func (m *DeleteOptions) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *DeleteOptions) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

type DeleteRequest struct {
	Key                  string         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Options              *DeleteOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeleteRequest) GetOptions() *DeleteOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListOptions struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Table                string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix               string   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Limit                uint64   `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOptions) Reset()         { *m = ListOptions{} }
func (m *ListOptions) String() string { return proto.CompactTextString(m) }
func (*ListOptions) ProtoMessage()    {}
func (*ListOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{11}
}

func (m *ListOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptions.Unmarshal(m, b)
}
func (m *ListOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptions.Marshal(b, m, deterministic)
}
func (m *ListOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptions.Merge(m, src)
}
func (m *ListOptions) XXX_Size() int {
	return xxx_messageInfo_ListOptions.Size(m)
}
func (m *ListOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptions proto.InternalMessageInfo

func (m *ListOptions) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ListOptions) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ListOptions) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListOptions) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *ListOptions) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOptions) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListRequest struct {
	Options              *ListOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{12}
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

func (m *ListRequest) GetOptions() *ListOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListResponse struct {
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{13}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DatabasesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabasesRequest) Reset()         { *m = DatabasesRequest{} }
func (m *DatabasesRequest) String() string { return proto.CompactTextString(m) }
func (*DatabasesRequest) ProtoMessage()    {}
func (*DatabasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{14}
}

func (m *DatabasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasesRequest.Unmarshal(m, b)
}
func (m *DatabasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasesRequest.Marshal(b, m, deterministic)
}
func (m *DatabasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasesRequest.Merge(m, src)
}
func (m *DatabasesRequest) XXX_Size() int {
	return xxx_messageInfo_DatabasesRequest.Size(m)
}
func (m *DatabasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasesRequest proto.InternalMessageInfo

type DatabasesResponse struct {
	Databases            []string `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabasesResponse) Reset()         { *m = DatabasesResponse{} }
func (m *DatabasesResponse) String() string { return proto.CompactTextString(m) }
func (*DatabasesResponse) ProtoMessage()    {}
func (*DatabasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{15}
}

func (m *DatabasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasesResponse.Unmarshal(m, b)
}
func (m *DatabasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasesResponse.Marshal(b, m, deterministic)
}
func (m *DatabasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasesResponse.Merge(m, src)
}
func (m *DatabasesResponse) XXX_Size() int {
	return xxx_messageInfo_DatabasesResponse.Size(m)
}
func (m *DatabasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasesResponse proto.InternalMessageInfo

func (m *DatabasesResponse) GetDatabases() []string {
	if m != nil {
		return m.Databases
	}
	return nil
}

type TablesRequest struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesRequest) Reset()         { *m = TablesRequest{} }
func (m *TablesRequest) String() string { return proto.CompactTextString(m) }
func (*TablesRequest) ProtoMessage()    {}
func (*TablesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{16}
}

func (m *TablesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesRequest.Unmarshal(m, b)
}
func (m *TablesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesRequest.Marshal(b, m, deterministic)
}
func (m *TablesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesRequest.Merge(m, src)
}
func (m *TablesRequest) XXX_Size() int {
	return xxx_messageInfo_TablesRequest.Size(m)
}
func (m *TablesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TablesRequest proto.InternalMessageInfo

func (m *TablesRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type TablesResponse struct {
	Tables               []string `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TablesResponse) Reset()         { *m = TablesResponse{} }
func (m *TablesResponse) String() string { return proto.CompactTextString(m) }
func (*TablesResponse) ProtoMessage()    {}
func (*TablesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8549980b097f750b, []int{17}
}

func (m *TablesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TablesResponse.Unmarshal(m, b)
}
func (m *TablesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TablesResponse.Marshal(b, m, deterministic)
}
func (m *TablesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TablesResponse.Merge(m, src)
}
func (m *TablesResponse) XXX_Size() int {
	return xxx_messageInfo_TablesResponse.Size(m)
}
func (m *TablesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TablesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TablesResponse proto.InternalMessageInfo

func (m *TablesResponse) GetTables() []string {
	if m != nil {
		return m.Tables
	}
	return nil
}

func init() {
	proto.RegisterType((*Field)(nil), "store.Field")
	proto.RegisterType((*Record)(nil), "store.Record")
	proto.RegisterMapType((map[string]*Field)(nil), "store.Record.MetadataEntry")
	proto.RegisterType((*ReadOptions)(nil), "store.ReadOptions")
	proto.RegisterType((*ReadRequest)(nil), "store.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "store.ReadResponse")
	proto.RegisterType((*WriteOptions)(nil), "store.WriteOptions")
	proto.RegisterType((*WriteRequest)(nil), "store.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "store.WriteResponse")
	proto.RegisterType((*DeleteOptions)(nil), "store.DeleteOptions")
	proto.RegisterType((*DeleteRequest)(nil), "store.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "store.DeleteResponse")
	proto.RegisterType((*ListOptions)(nil), "store.ListOptions")
	proto.RegisterType((*ListRequest)(nil), "store.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "store.ListResponse")
	proto.RegisterType((*DatabasesRequest)(nil), "store.DatabasesRequest")
	proto.RegisterType((*DatabasesResponse)(nil), "store.DatabasesResponse")
	proto.RegisterType((*TablesRequest)(nil), "store.TablesRequest")
	proto.RegisterType((*TablesResponse)(nil), "store.TablesResponse")
}

func init() { proto.RegisterFile("store/store.proto", fileDescriptor_8549980b097f750b) }

var fileDescriptor_8549980b097f750b = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x13, 0xdb, 0x8d, 0x27, 0x49, 0x49, 0xb7, 0xa1, 0x58, 0x81, 0x43, 0xb4, 0x12, 0xc2,
	0x52, 0x69, 0x42, 0x12, 0xa4, 0x20, 0x7a, 0x29, 0xa8, 0x20, 0x81, 0xa8, 0x10, 0x06, 0x09, 0x89,
	0x9b, 0x93, 0x6c, 0xc0, 0x6a, 0x12, 0x1b, 0x7b, 0x53, 0x35, 0x1f, 0xc3, 0x1f, 0xf1, 0x03, 0xfc,
	0x0d, 0xda, 0xdd, 0x59, 0xc7, 0x4e, 0x5b, 0x0e, 0xa5, 0x17, 0xcb, 0x33, 0xbb, 0xf3, 0xe6, 0xbd,
	0x37, 0x93, 0x18, 0xf6, 0x52, 0x1e, 0x25, 0xac, 0x27, 0x9f, 0xdd, 0x38, 0x89, 0x78, 0x44, 0x2c,
	0x19, 0xd0, 0x3e, 0x58, 0x6f, 0x43, 0x36, 0x9f, 0x12, 0x02, 0x26, 0x5f, 0xc7, 0xcc, 0x35, 0x3a,
	0x86, 0xe7, 0xf8, 0xf2, 0x9d, 0xb4, 0xc0, 0xba, 0x08, 0xe6, 0x2b, 0xe6, 0x96, 0x65, 0x52, 0x05,
	0xf4, 0xb7, 0x01, 0xb6, 0xcf, 0x26, 0x51, 0x32, 0x25, 0x4d, 0xa8, 0x9c, 0xb3, 0x35, 0xd6, 0x88,
	0xd7, 0x62, 0x49, 0x1d, 0x4b, 0xc8, 0x01, 0xd8, 0xec, 0x32, 0x0e, 0x93, 0xb5, 0x5b, 0xe9, 0x18,
	0x5e, 0xc5, 0xc7, 0x88, 0x8c, 0xa0, 0xba, 0x60, 0x3c, 0x98, 0x06, 0x3c, 0x70, 0xcd, 0x4e, 0xc5,
	0xab, 0x0d, 0x1e, 0x76, 0x15, 0x49, 0xd5, 0xa0, 0x7b, 0x86, 0xa7, 0x6f, 0x96, 0x3c, 0x59, 0xfb,
	0xd9, 0xe5, 0xf6, 0x3b, 0x68, 0x14, 0x8e, 0xae, 0x61, 0x42, 0xf3, 0x4c, 0x6a, 0x83, 0x3a, 0x02,
	0x4b, 0xb5, 0xc8, 0xeb, 0x65, 0xf9, 0x85, 0x41, 0x7f, 0x19, 0x50, 0xf3, 0x59, 0x30, 0xfd, 0x18,
	0xf3, 0x30, 0x5a, 0xa6, 0xa4, 0x0d, 0x55, 0x01, 0x3b, 0x0e, 0x52, 0x6d, 0x46, 0x16, 0x0b, 0x75,
	0x3c, 0x18, 0xcf, 0x33, 0x43, 0x64, 0x20, 0xd4, 0xc5, 0x09, 0x9b, 0x85, 0x97, 0x52, 0x5d, 0xd5,
	0xc7, 0x48, 0xe4, 0xd3, 0xd5, 0x4c, 0xe4, 0x4d, 0x95, 0x57, 0x91, 0x40, 0x99, 0x87, 0x8b, 0x90,
	0xbb, 0x56, 0xc7, 0xf0, 0x4c, 0x5f, 0x05, 0xe2, 0x76, 0x34, 0x9b, 0xa5, 0x8c, 0xbb, 0xb6, 0x4c,
	0x63, 0x44, 0xcf, 0x14, 0x3d, 0x9f, 0xfd, 0x5c, 0xb1, 0x94, 0x5f, 0x23, 0xf4, 0x29, 0xec, 0x44,
	0x8a, 0x3b, 0x4a, 0x25, 0x99, 0x87, 0x99, 0x2a, 0x5f, 0x5f, 0xa1, 0x23, 0xa8, 0x2b, 0xb8, 0x34,
	0x8e, 0x96, 0x29, 0x23, 0x4f, 0x60, 0x27, 0x91, 0x5e, 0xa7, 0xae, 0x21, 0x27, 0xd0, 0x28, 0x4c,
	0xc0, 0xd7, 0xa7, 0xf4, 0x04, 0xea, 0x5f, 0x93, 0x90, 0xb3, 0x5b, 0xfb, 0x44, 0xa7, 0x88, 0xa0,
	0xa5, 0x3c, 0x06, 0x5b, 0x81, 0xcb, 0xfa, 0x2b, 0x9d, 0xf1, 0x90, 0x1c, 0x6d, 0xeb, 0xdb, 0xc7,
	0x7b, 0x79, 0x3a, 0x1b, 0x81, 0xf7, 0xa0, 0x81, 0x5d, 0x94, 0x42, 0xfa, 0x0a, 0x1a, 0xa7, 0x6c,
	0xce, 0xfe, 0x87, 0xf9, 0x27, 0x0d, 0x71, 0xf3, 0x14, 0xba, 0xdb, 0x2c, 0x5b, 0xc8, 0xb2, 0xd0,
	0x7b, 0x43, 0xb3, 0x09, 0xbb, 0x1a, 0x12, 0x79, 0x8a, 0x45, 0xfc, 0x10, 0xa6, 0xfc, 0xae, 0x16,
	0xd1, 0xb9, 0x61, 0x11, 0x9d, 0x5b, 0x2e, 0xe2, 0xb1, 0xa2, 0xa7, 0x2d, 0xc8, 0xad, 0x9d, 0x51,
	0x58, 0xbb, 0x9c, 0x86, 0x8d, 0x5c, 0x0f, 0xea, 0xaa, 0x18, 0xd7, 0x8e, 0x80, 0x79, 0xce, 0xd6,
	0xc2, 0xab, 0x8a, 0xf8, 0xbb, 0x11, 0xef, 0xef, 0xcd, 0xaa, 0xd1, 0x2c, 0x53, 0x02, 0xcd, 0x53,
	0x94, 0x99, 0x62, 0x2f, 0xda, 0x87, 0xbd, 0x5c, 0x0e, 0x21, 0x1e, 0x81, 0xa3, 0xfd, 0x50, 0xbb,
	0xeb, 0xf8, 0x9b, 0x04, 0x3d, 0x84, 0xc6, 0x17, 0x61, 0x8a, 0xc6, 0xf8, 0x97, 0x9d, 0xd4, 0x83,
	0x5d, 0x7d, 0x19, 0xc1, 0x0f, 0xc0, 0x96, 0x9e, 0x6a, 0x64, 0x8c, 0x06, 0x7f, 0xca, 0x60, 0x7d,
	0x16, 0x32, 0x49, 0x1f, 0x4c, 0xf1, 0x43, 0x22, 0xf9, 0x5f, 0x1b, 0xf6, 0x6a, 0xef, 0x17, 0x72,
	0x38, 0xdf, 0x12, 0x79, 0x0e, 0x96, 0x5c, 0x4d, 0x52, 0xd8, 0x60, 0x5d, 0xd4, 0x2a, 0x26, 0xb3,
	0xaa, 0x11, 0xd8, 0x6a, 0x53, 0x48, 0x71, 0xa5, 0x74, 0xdd, 0xfd, 0xad, 0x6c, 0x56, 0x38, 0x04,
	0x53, 0x78, 0x4e, 0xf2, 0x83, 0xd9, 0x66, 0x98, 0x1f, 0x0a, 0x2d, 0x3d, 0x33, 0xc8, 0x09, 0x38,
	0x99, 0xd5, 0xe4, 0x81, 0x86, 0xde, 0x1a, 0x48, 0xdb, 0xbd, 0x7a, 0x90, 0xe7, 0xab, 0xcc, 0xcc,
	0xf8, 0x16, 0x06, 0x91, 0xf1, 0x2d, 0x3a, 0x4e, 0x4b, 0xaf, 0x8f, 0xbe, 0x1d, 0x7e, 0x0f, 0xf9,
	0x8f, 0xd5, 0xb8, 0x3b, 0x89, 0x16, 0xbd, 0x45, 0x38, 0x49, 0x22, 0x7c, 0x5e, 0x0c, 0x7b, 0xf2,
	0xab, 0xa5, 0xbe, 0x60, 0xc7, 0xf2, 0x39, 0xb6, 0x65, 0x6a, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x9c, 0x61, 0x2f, 0xdd, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Store_ListClient, error)
	Databases(ctx context.Context, in *DatabasesRequest, opts ...grpc.CallOption) (*DatabasesResponse, error)
	Tables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/store.Store/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/store.Store/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/store.Store/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Store_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Store_serviceDesc.Streams[0], "/store.Store/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Store_ListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type storeListClient struct {
	grpc.ClientStream
}

func (x *storeListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeClient) Databases(ctx context.Context, in *DatabasesRequest, opts ...grpc.CallOption) (*DatabasesResponse, error) {
	out := new(DatabasesResponse)
	err := c.cc.Invoke(ctx, "/store.Store/Databases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Tables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesResponse, error) {
	out := new(TablesResponse)
	err := c.cc.Invoke(ctx, "/store.Store/Tables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(*ListRequest, Store_ListServer) error
	Databases(context.Context, *DatabasesRequest) (*DatabasesResponse, error)
	Tables(context.Context, *TablesRequest) (*TablesResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreServer).List(m, &storeListServer{stream})
}

type Store_ListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type storeListServer struct {
	grpc.ServerStream
}

func (x *storeListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Store_Databases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Databases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/Databases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Databases(ctx, req.(*DatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Tables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Tables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.Store/Tables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Tables(ctx, req.(*TablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Store_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Store_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Store_Delete_Handler,
		},
		{
			MethodName: "Databases",
			Handler:    _Store_Databases_Handler,
		},
		{
			MethodName: "Tables",
			Handler:    _Store_Tables_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Store_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "store/store.proto",
}