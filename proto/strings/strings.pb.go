// Code generated by protoc-gen-go. DO NOT EDIT.
// source: strings.proto

package strings

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

type StringsRequest struct {
	// ID of file.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringsRequest) Reset()         { *m = StringsRequest{} }
func (m *StringsRequest) String() string { return proto.CompactTextString(m) }
func (*StringsRequest) ProtoMessage()    {}
func (*StringsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af3cde72035e609, []int{0}
}

func (m *StringsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringsRequest.Unmarshal(m, b)
}
func (m *StringsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringsRequest.Marshal(b, m, deterministic)
}
func (m *StringsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringsRequest.Merge(m, src)
}
func (m *StringsRequest) XXX_Size() int {
	return xxx_messageInfo_StringsRequest.Size(m)
}
func (m *StringsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StringsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StringsRequest proto.InternalMessageInfo

func (m *StringsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StringInfo struct {
	Location             uint64   `protobuf:"varint,1,opt,name=location,proto3" json:"location,omitempty"`
	RawString            string   `protobuf:"bytes,2,opt,name=raw_string,json=rawString,proto3" json:"raw_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringInfo) Reset()         { *m = StringInfo{} }
func (m *StringInfo) String() string { return proto.CompactTextString(m) }
func (*StringInfo) ProtoMessage()    {}
func (*StringInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af3cde72035e609, []int{1}
}

func (m *StringInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringInfo.Unmarshal(m, b)
}
func (m *StringInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringInfo.Marshal(b, m, deterministic)
}
func (m *StringInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringInfo.Merge(m, src)
}
func (m *StringInfo) XXX_Size() int {
	return xxx_messageInfo_StringInfo.Size(m)
}
func (m *StringInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StringInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StringInfo proto.InternalMessageInfo

func (m *StringInfo) GetLocation() uint64 {
	if m != nil {
		return m.Location
	}
	return 0
}

func (m *StringInfo) GetRawString() string {
	if m != nil {
		return m.RawString
	}
	return ""
}

type StringsReply struct {
	Strings              []*StringInfo `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StringsReply) Reset()         { *m = StringsReply{} }
func (m *StringsReply) String() string { return proto.CompactTextString(m) }
func (*StringsReply) ProtoMessage()    {}
func (*StringsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af3cde72035e609, []int{2}
}

func (m *StringsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringsReply.Unmarshal(m, b)
}
func (m *StringsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringsReply.Marshal(b, m, deterministic)
}
func (m *StringsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringsReply.Merge(m, src)
}
func (m *StringsReply) XXX_Size() int {
	return xxx_messageInfo_StringsReply.Size(m)
}
func (m *StringsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StringsReply.DiscardUnknown(m)
}

var xxx_messageInfo_StringsReply proto.InternalMessageInfo

func (m *StringsReply) GetStrings() []*StringInfo {
	if m != nil {
		return m.Strings
	}
	return nil
}

func init() {
	proto.RegisterType((*StringsRequest)(nil), "strings.StringsRequest")
	proto.RegisterType((*StringInfo)(nil), "strings.StringInfo")
	proto.RegisterType((*StringsReply)(nil), "strings.StringsReply")
}

func init() { proto.RegisterFile("strings.proto", fileDescriptor_0af3cde72035e609) }

var fileDescriptor_0af3cde72035e609 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x14, 0xb8,
	0xf8, 0x82, 0x21, 0xcc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x14, 0x25, 0x77, 0x2e, 0x2e, 0x88,
	0x0a, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xfc, 0xe4, 0xc4, 0x92, 0xcc, 0xfc,
	0x3c, 0xb0, 0x1a, 0x96, 0x20, 0x38, 0x5f, 0x48, 0x96, 0x8b, 0xab, 0x28, 0xb1, 0x3c, 0x1e, 0x62,
	0xb4, 0x04, 0x13, 0xd8, 0x04, 0xce, 0xa2, 0xc4, 0x72, 0x88, 0x76, 0x25, 0x5b, 0x2e, 0x1e, 0xb8,
	0x55, 0x05, 0x39, 0x95, 0x42, 0xba, 0x5c, 0x30, 0x57, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b,
	0x09, 0xeb, 0xc1, 0x1c, 0x89, 0xb0, 0x30, 0x08, 0xa6, 0xc6, 0x28, 0x8c, 0x4b, 0x00, 0xaa, 0xdd,
	0xb5, 0xa2, 0xa4, 0x28, 0x31, 0xb9, 0x24, 0xbf, 0x48, 0xc8, 0x89, 0x8b, 0x0f, 0xca, 0x81, 0x4a,
	0x09, 0x89, 0xa3, 0x99, 0x01, 0xf3, 0x96, 0x94, 0x28, 0xa6, 0x44, 0x41, 0x4e, 0xa5, 0x12, 0x43,
	0x12, 0x1b, 0x38, 0x44, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x7e, 0x91, 0xed, 0x22,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringsExtractorClient is the client API for StringsExtractor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringsExtractorClient interface {
	ExtractStrings(ctx context.Context, in *StringsRequest, opts ...grpc.CallOption) (*StringsReply, error)
}

type stringsExtractorClient struct {
	cc *grpc.ClientConn
}

func NewStringsExtractorClient(cc *grpc.ClientConn) StringsExtractorClient {
	return &stringsExtractorClient{cc}
}

func (c *stringsExtractorClient) ExtractStrings(ctx context.Context, in *StringsRequest, opts ...grpc.CallOption) (*StringsReply, error) {
	out := new(StringsReply)
	err := c.cc.Invoke(ctx, "/strings.StringsExtractor/ExtractStrings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringsExtractorServer is the server API for StringsExtractor service.
type StringsExtractorServer interface {
	ExtractStrings(context.Context, *StringsRequest) (*StringsReply, error)
}

// UnimplementedStringsExtractorServer can be embedded to have forward compatible implementations.
type UnimplementedStringsExtractorServer struct {
}

func (*UnimplementedStringsExtractorServer) ExtractStrings(ctx context.Context, req *StringsRequest) (*StringsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractStrings not implemented")
}

func RegisterStringsExtractorServer(s *grpc.Server, srv StringsExtractorServer) {
	s.RegisterService(&_StringsExtractor_serviceDesc, srv)
}

func _StringsExtractor_ExtractStrings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringsExtractorServer).ExtractStrings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strings.StringsExtractor/ExtractStrings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringsExtractorServer).ExtractStrings(ctx, req.(*StringsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringsExtractor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "strings.StringsExtractor",
	HandlerType: (*StringsExtractorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExtractStrings",
			Handler:    _StringsExtractor_ExtractStrings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strings.proto",
}
