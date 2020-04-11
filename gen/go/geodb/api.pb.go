// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type StreamObjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamObjectsRequest) Reset()         { *m = StreamObjectsRequest{} }
func (m *StreamObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamObjectsRequest) ProtoMessage()    {}
func (*StreamObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *StreamObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamObjectsRequest.Unmarshal(m, b)
}
func (m *StreamObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamObjectsRequest.Marshal(b, m, deterministic)
}
func (m *StreamObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamObjectsRequest.Merge(m, src)
}
func (m *StreamObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamObjectsRequest.Size(m)
}
func (m *StreamObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamObjectsRequest proto.InternalMessageInfo

type StreamObjectsResponse struct {
	Object               *Object  `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamObjectsResponse) Reset()         { *m = StreamObjectsResponse{} }
func (m *StreamObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamObjectsResponse) ProtoMessage()    {}
func (*StreamObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *StreamObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamObjectsResponse.Unmarshal(m, b)
}
func (m *StreamObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamObjectsResponse.Marshal(b, m, deterministic)
}
func (m *StreamObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamObjectsResponse.Merge(m, src)
}
func (m *StreamObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamObjectsResponse.Size(m)
}
func (m *StreamObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamObjectsResponse proto.InternalMessageInfo

func (m *StreamObjectsResponse) GetObject() *Object {
	if m != nil {
		return m.Object
	}
	return nil
}

type UpsertObjectsRequest struct {
	Data                 map[string]*Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpsertObjectsRequest) Reset()         { *m = UpsertObjectsRequest{} }
func (m *UpsertObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertObjectsRequest) ProtoMessage()    {}
func (*UpsertObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *UpsertObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertObjectsRequest.Unmarshal(m, b)
}
func (m *UpsertObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertObjectsRequest.Marshal(b, m, deterministic)
}
func (m *UpsertObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertObjectsRequest.Merge(m, src)
}
func (m *UpsertObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertObjectsRequest.Size(m)
}
func (m *UpsertObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertObjectsRequest proto.InternalMessageInfo

func (m *UpsertObjectsRequest) GetData() map[string]*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpsertObjectsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertObjectsResponse) Reset()         { *m = UpsertObjectsResponse{} }
func (m *UpsertObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertObjectsResponse) ProtoMessage()    {}
func (*UpsertObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *UpsertObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertObjectsResponse.Unmarshal(m, b)
}
func (m *UpsertObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertObjectsResponse.Marshal(b, m, deterministic)
}
func (m *UpsertObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertObjectsResponse.Merge(m, src)
}
func (m *UpsertObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertObjectsResponse.Size(m)
}
func (m *UpsertObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertObjectsResponse proto.InternalMessageInfo

type GetObjectsRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectsRequest) Reset()         { *m = GetObjectsRequest{} }
func (m *GetObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectsRequest) ProtoMessage()    {}
func (*GetObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectsRequest.Unmarshal(m, b)
}
func (m *GetObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectsRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectsRequest.Merge(m, src)
}
func (m *GetObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectsRequest.Size(m)
}
func (m *GetObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectsRequest proto.InternalMessageInfo

func (m *GetObjectsRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetObjectsResponse struct {
	Objects              map[string]*Object `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetObjectsResponse) Reset()         { *m = GetObjectsResponse{} }
func (m *GetObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*GetObjectsResponse) ProtoMessage()    {}
func (*GetObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *GetObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectsResponse.Unmarshal(m, b)
}
func (m *GetObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectsResponse.Marshal(b, m, deterministic)
}
func (m *GetObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectsResponse.Merge(m, src)
}
func (m *GetObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_GetObjectsResponse.Size(m)
}
func (m *GetObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectsResponse proto.InternalMessageInfo

func (m *GetObjectsResponse) GetObjects() map[string]*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type DeleteObjectsRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectsRequest) Reset()         { *m = DeleteObjectsRequest{} }
func (m *DeleteObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectsRequest) ProtoMessage()    {}
func (*DeleteObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *DeleteObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectsRequest.Unmarshal(m, b)
}
func (m *DeleteObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectsRequest.Merge(m, src)
}
func (m *DeleteObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectsRequest.Size(m)
}
func (m *DeleteObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectsRequest proto.InternalMessageInfo

func (m *DeleteObjectsRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteObjectsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectsResponse) Reset()         { *m = DeleteObjectsResponse{} }
func (m *DeleteObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectsResponse) ProtoMessage()    {}
func (*DeleteObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *DeleteObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectsResponse.Unmarshal(m, b)
}
func (m *DeleteObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectsResponse.Marshal(b, m, deterministic)
}
func (m *DeleteObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectsResponse.Merge(m, src)
}
func (m *DeleteObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectsResponse.Size(m)
}
func (m *DeleteObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectsResponse proto.InternalMessageInfo

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type Point struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Point) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Data struct {
	Point                *Point            `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
	Radius               int64             `protobuf:"varint,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UpdatedUnix          int64             `protobuf:"varint,5,opt,name=updated_unix,json=updatedUnix,proto3" json:"updated_unix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Data) GetRadius() int64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Data) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Data) GetUpdatedUnix() int64 {
	if m != nil {
		return m.UpdatedUnix
	}
	return 0
}

type Object struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data                 *Data    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Object) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamObjectsRequest)(nil), "api.StreamObjectsRequest")
	proto.RegisterType((*StreamObjectsResponse)(nil), "api.StreamObjectsResponse")
	proto.RegisterType((*UpsertObjectsRequest)(nil), "api.UpsertObjectsRequest")
	proto.RegisterMapType((map[string]*Data)(nil), "api.UpsertObjectsRequest.DataEntry")
	proto.RegisterType((*UpsertObjectsResponse)(nil), "api.UpsertObjectsResponse")
	proto.RegisterType((*GetObjectsRequest)(nil), "api.GetObjectsRequest")
	proto.RegisterType((*GetObjectsResponse)(nil), "api.GetObjectsResponse")
	proto.RegisterMapType((map[string]*Object)(nil), "api.GetObjectsResponse.ObjectsEntry")
	proto.RegisterType((*DeleteObjectsRequest)(nil), "api.DeleteObjectsRequest")
	proto.RegisterType((*DeleteObjectsResponse)(nil), "api.DeleteObjectsResponse")
	proto.RegisterType((*PingRequest)(nil), "api.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "api.PingResponse")
	proto.RegisterType((*Point)(nil), "api.Point")
	proto.RegisterType((*Data)(nil), "api.Data")
	proto.RegisterMapType((map[string]string)(nil), "api.Data.MetadataEntry")
	proto.RegisterType((*Object)(nil), "api.Object")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xa5, 0x40, 0x71, 0x7b, 0x01, 0xb3, 0x3b, 0xe1, 0xcb, 0x26, 0x2a, 0xdb, 0x35, 0x91, 0x68,
	0x00, 0xc3, 0x26, 0x7e, 0x47, 0x13, 0x82, 0xc1, 0x98, 0x18, 0x37, 0x35, 0xfb, 0xe2, 0x8b, 0x19,
	0x96, 0x09, 0x56, 0xa0, 0x53, 0xdb, 0xe9, 0xee, 0xf2, 0x3b, 0x7c, 0xf6, 0x67, 0x19, 0xff, 0x8e,
	0x99, 0x3b, 0xb3, 0xdd, 0x16, 0xba, 0xc9, 0xbe, 0x0d, 0xf7, 0xab, 0xe7, 0x9c, 0x7b, 0xb8, 0x60,
	0xd1, 0xc0, 0x1b, 0x04, 0x21, 0x17, 0x9c, 0x94, 0x68, 0xe0, 0xd9, 0xcf, 0x17, 0x9e, 0xf8, 0x11,
	0xcf, 0x06, 0x67, 0x7c, 0x3d, 0x5c, 0x5f, 0x78, 0x62, 0xc9, 0x2f, 0x86, 0x0b, 0xde, 0xc7, 0x8a,
	0xfe, 0x39, 0x5d, 0x79, 0x73, 0x2a, 0x78, 0x18, 0x0d, 0x93, 0xa7, 0x6a, 0x76, 0x5a, 0xd0, 0xf8,
	0x2a, 0x42, 0x46, 0xd7, 0x5f, 0x66, 0x3f, 0xd9, 0x99, 0x88, 0x5c, 0xf6, 0x2b, 0x66, 0x91, 0x70,
	0xde, 0x42, 0x73, 0x2b, 0x1e, 0x05, 0xdc, 0x8f, 0x18, 0x39, 0x82, 0x0a, 0xc7, 0x50, 0xc7, 0xe8,
	0x1a, 0xbd, 0xea, 0xa8, 0x3a, 0x90, 0x48, 0x54, 0x95, 0xab, 0x53, 0xce, 0x6f, 0x03, 0x1a, 0xa7,
	0x41, 0xc4, 0x42, 0x91, 0x1d, 0x4b, 0x5e, 0x40, 0x79, 0x4e, 0x05, 0xed, 0x18, 0xdd, 0x52, 0xaf,
	0x3a, 0x3a, 0xc2, 0xde, 0xbc, 0xc2, 0xc1, 0x84, 0x0a, 0xfa, 0xc1, 0x17, 0xe1, 0xc6, 0xc5, 0x06,
	0x7b, 0x0c, 0x56, 0x12, 0x22, 0xfb, 0x50, 0x5a, 0xb2, 0x0d, 0x02, 0xb0, 0x5c, 0xf9, 0x24, 0x0f,
	0xc1, 0x3c, 0xa7, 0xab, 0x98, 0x75, 0x8a, 0x08, 0xca, 0xc2, 0xc1, 0xb2, 0xc1, 0x55, 0xf1, 0xd7,
	0xc5, 0x97, 0x86, 0xd3, 0x86, 0xe6, 0xd6, 0xb7, 0x14, 0x27, 0xe7, 0x31, 0x1c, 0x4c, 0xd9, 0x36,
	0x54, 0x02, 0xe5, 0x25, 0xdb, 0x44, 0x08, 0xd5, 0x72, 0xf1, 0xed, 0xfc, 0x31, 0x80, 0xa4, 0x2b,
	0xb5, 0x26, 0xef, 0xe0, 0x8e, 0x22, 0x1e, 0x69, 0x62, 0x8f, 0xf0, 0xfb, 0xbb, 0x95, 0x5a, 0xa7,
	0x48, 0x31, 0xbb, 0x6a, 0xb2, 0xa7, 0x50, 0x4b, 0x27, 0x72, 0xf8, 0x1d, 0x66, 0xf9, 0x65, 0x44,
	0x4f, 0x31, 0x7c, 0x02, 0x8d, 0x09, 0x5b, 0x31, 0xc1, 0x6e, 0xc1, 0xa5, 0x0d, 0xcd, 0xad, 0x5a,
	0xad, 0x46, 0x1d, 0xaa, 0x27, 0x9e, 0xbf, 0xb8, 0x72, 0xc2, 0x03, 0xa8, 0xa9, 0x9f, 0x9a, 0xec,
	0x5d, 0x28, 0xf2, 0x25, 0x62, 0xdb, 0x73, 0x8b, 0x7c, 0xe9, 0x3c, 0x05, 0xf3, 0x84, 0x7b, 0xbe,
	0x90, 0xa8, 0x57, 0x54, 0xd9, 0xc2, 0x70, 0xe5, 0x13, 0x23, 0xdc, 0x47, 0xcc, 0x32, 0xc2, 0x7d,
	0xe7, 0xaf, 0x01, 0x65, 0xb9, 0x16, 0xd2, 0x05, 0x33, 0x90, 0x5d, 0x9a, 0x10, 0x20, 0x21, 0x9c,
	0xe3, 0xaa, 0x04, 0x69, 0x41, 0x25, 0xa4, 0x73, 0x2f, 0x8e, 0x3a, 0xa5, 0xae, 0xd1, 0x2b, 0xb9,
	0xfa, 0x17, 0x39, 0x86, 0xbd, 0x35, 0x13, 0x14, 0x6d, 0x54, 0x46, 0xb5, 0xdb, 0xc9, 0xb6, 0x07,
	0x9f, 0x75, 0x46, 0x09, 0x9c, 0x14, 0x92, 0x43, 0xa8, 0xc5, 0xc1, 0x9c, 0x0a, 0x36, 0xff, 0x1e,
	0xfb, 0xde, 0x65, 0xc7, 0xc4, 0x91, 0x55, 0x1d, 0x3b, 0xf5, 0xbd, 0x4b, 0xfb, 0x0d, 0xd4, 0x33,
	0xdd, 0x39, 0x5b, 0x68, 0xa4, 0xb7, 0x60, 0xa5, 0x85, 0x7f, 0x05, 0x15, 0x25, 0x63, 0x4e, 0xd7,
	0x7d, 0xed, 0xf9, 0x1d, 0x6b, 0x62, 0x78, 0xf4, 0xaf, 0x08, 0xe6, 0x94, 0xf1, 0xc9, 0x98, 0xf4,
	0xa1, 0x2c, 0x95, 0x26, 0xfb, 0x4a, 0x8c, 0xeb, 0x1d, 0xd8, 0x07, 0xa9, 0x88, 0xde, 0x52, 0x81,
	0x7c, 0x84, 0x7a, 0xc6, 0xce, 0xe4, 0xde, 0x8d, 0x7f, 0x27, 0xdb, 0xce, 0x4b, 0x25, 0x93, 0xde,
	0x03, 0x5c, 0x7b, 0x95, 0xb4, 0x76, 0xcc, 0xab, 0x66, 0xb4, 0x6f, 0x30, 0xb5, 0x82, 0x92, 0xf1,
	0x92, 0x86, 0x92, 0xe7, 0x45, 0x0d, 0x25, 0xdf, 0x7a, 0x05, 0xf2, 0x09, 0xea, 0x99, 0xbb, 0xa3,
	0x27, 0xe5, 0xdd, 0x28, 0x3d, 0x29, 0xf7, 0x4c, 0x39, 0x85, 0x67, 0xc6, 0xd8, 0xfc, 0x26, 0x4f,
	0xe3, 0xac, 0x82, 0x97, 0xee, 0xf8, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x01, 0xb2, 0xea,
	0x33, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoDBClient is the client API for GeoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoDBClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	UpsertObjects(ctx context.Context, in *UpsertObjectsRequest, opts ...grpc.CallOption) (*UpsertObjectsResponse, error)
	GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error)
	DeleteObjects(ctx context.Context, in *DeleteObjectsRequest, opts ...grpc.CallOption) (*DeleteObjectsResponse, error)
	StreamObjects(ctx context.Context, in *StreamObjectsRequest, opts ...grpc.CallOption) (GeoDB_StreamObjectsClient, error)
}

type geoDBClient struct {
	cc *grpc.ClientConn
}

func NewGeoDBClient(cc *grpc.ClientConn) GeoDBClient {
	return &geoDBClient{cc}
}

func (c *geoDBClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) UpsertObjects(ctx context.Context, in *UpsertObjectsRequest, opts ...grpc.CallOption) (*UpsertObjectsResponse, error) {
	out := new(UpsertObjectsResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/UpsertObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error) {
	out := new(GetObjectsResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/GetObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) DeleteObjects(ctx context.Context, in *DeleteObjectsRequest, opts ...grpc.CallOption) (*DeleteObjectsResponse, error) {
	out := new(DeleteObjectsResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/DeleteObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) StreamObjects(ctx context.Context, in *StreamObjectsRequest, opts ...grpc.CallOption) (GeoDB_StreamObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GeoDB_serviceDesc.Streams[0], "/api.GeoDB/StreamObjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &geoDBStreamObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GeoDB_StreamObjectsClient interface {
	Recv() (*StreamObjectsResponse, error)
	grpc.ClientStream
}

type geoDBStreamObjectsClient struct {
	grpc.ClientStream
}

func (x *geoDBStreamObjectsClient) Recv() (*StreamObjectsResponse, error) {
	m := new(StreamObjectsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeoDBServer is the server API for GeoDB service.
type GeoDBServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	UpsertObjects(context.Context, *UpsertObjectsRequest) (*UpsertObjectsResponse, error)
	GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error)
	DeleteObjects(context.Context, *DeleteObjectsRequest) (*DeleteObjectsResponse, error)
	StreamObjects(*StreamObjectsRequest, GeoDB_StreamObjectsServer) error
}

// UnimplementedGeoDBServer can be embedded to have forward compatible implementations.
type UnimplementedGeoDBServer struct {
}

func (*UnimplementedGeoDBServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGeoDBServer) UpsertObjects(ctx context.Context, req *UpsertObjectsRequest) (*UpsertObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertObjects not implemented")
}
func (*UnimplementedGeoDBServer) GetObjects(ctx context.Context, req *GetObjectsRequest) (*GetObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (*UnimplementedGeoDBServer) DeleteObjects(ctx context.Context, req *DeleteObjectsRequest) (*DeleteObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjects not implemented")
}
func (*UnimplementedGeoDBServer) StreamObjects(req *StreamObjectsRequest, srv GeoDB_StreamObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamObjects not implemented")
}

func RegisterGeoDBServer(s *grpc.Server, srv GeoDBServer) {
	s.RegisterService(&_GeoDB_serviceDesc, srv)
}

func _GeoDB_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_UpsertObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).UpsertObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/UpsertObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).UpsertObjects(ctx, req.(*UpsertObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/GetObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).GetObjects(ctx, req.(*GetObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_DeleteObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).DeleteObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/DeleteObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).DeleteObjects(ctx, req.(*DeleteObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_StreamObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamObjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeoDBServer).StreamObjects(m, &geoDBStreamObjectsServer{stream})
}

type GeoDB_StreamObjectsServer interface {
	Send(*StreamObjectsResponse) error
	grpc.ServerStream
}

type geoDBStreamObjectsServer struct {
	grpc.ServerStream
}

func (x *geoDBStreamObjectsServer) Send(m *StreamObjectsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GeoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GeoDB",
	HandlerType: (*GeoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GeoDB_Ping_Handler,
		},
		{
			MethodName: "UpsertObjects",
			Handler:    _GeoDB_UpsertObjects_Handler,
		},
		{
			MethodName: "GetObjects",
			Handler:    _GeoDB_GetObjects_Handler,
		},
		{
			MethodName: "DeleteObjects",
			Handler:    _GeoDB_DeleteObjects_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamObjects",
			Handler:       _GeoDB_StreamObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
