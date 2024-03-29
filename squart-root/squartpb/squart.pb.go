// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SquartRoot/squartpb/squart.proto

package squartpb

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

type SquartRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquartRootRequest) Reset()         { *m = SquartRootRequest{} }
func (m *SquartRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquartRootRequest) ProtoMessage()    {}
func (*SquartRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_911ac8e2e90f512a, []int{0}
}

func (m *SquartRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquartRootRequest.Unmarshal(m, b)
}
func (m *SquartRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquartRootRequest.Marshal(b, m, deterministic)
}
func (m *SquartRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquartRootRequest.Merge(m, src)
}
func (m *SquartRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquartRootRequest.Size(m)
}
func (m *SquartRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquartRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquartRootRequest proto.InternalMessageInfo

func (m *SquartRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquartRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquartRootResponse) Reset()         { *m = SquartRootResponse{} }
func (m *SquartRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquartRootResponse) ProtoMessage()    {}
func (*SquartRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_911ac8e2e90f512a, []int{1}
}

func (m *SquartRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquartRootResponse.Unmarshal(m, b)
}
func (m *SquartRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquartRootResponse.Marshal(b, m, deterministic)
}
func (m *SquartRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquartRootResponse.Merge(m, src)
}
func (m *SquartRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquartRootResponse.Size(m)
}
func (m *SquartRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquartRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquartRootResponse proto.InternalMessageInfo

func (m *SquartRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*SquartRootRequest)(nil), "squart.SquartRootRequest")
	proto.RegisterType((*SquartRootResponse)(nil), "squart.SquartRootResponse")
}

func init() { proto.RegisterFile("SquartRoot/squartpb/squart.proto", fileDescriptor_911ac8e2e90f512a) }

var fileDescriptor_911ac8e2e90f512a = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x08, 0x2e, 0x2c, 0x4d,
	0x2c, 0x2a, 0x09, 0xca, 0xcf, 0x2f, 0xd1, 0x2f, 0x06, 0x33, 0x0b, 0x92, 0xa0, 0x0c, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x36, 0x08, 0x4f, 0x49, 0x9b, 0x4b, 0x10, 0xa1, 0x36, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0x32, 0xe5, 0x12, 0x42, 0x56, 0x5c, 0x5c, 0x90,
	0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xcf, 0xc5, 0x0d, 0x91, 0x8f, 0x2f, 0xca, 0xcf, 0x2f, 0x01, 0x6b,
	0x61, 0x0c, 0xe2, 0x82, 0x08, 0x81, 0x14, 0x1a, 0x05, 0x72, 0x71, 0x21, 0xb4, 0x09, 0x39, 0xa3,
	0xf0, 0x24, 0xf5, 0xa0, 0xce, 0xc2, 0x70, 0x85, 0x94, 0x14, 0x36, 0x29, 0x88, 0x9d, 0x4e, 0x5c,
	0x51, 0x1c, 0x30, 0x7f, 0x25, 0xb1, 0x81, 0x7d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x48,
	0xaa, 0xea, 0x6b, 0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SquartRootClient is the client API for SquartRoot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquartRootClient interface {
	SquartRoot(ctx context.Context, in *SquartRootRequest, opts ...grpc.CallOption) (*SquartRootResponse, error)
}

type squartRootClient struct {
	cc *grpc.ClientConn
}

func NewSquartRootClient(cc *grpc.ClientConn) SquartRootClient {
	return &squartRootClient{cc}
}

func (c *squartRootClient) SquartRoot(ctx context.Context, in *SquartRootRequest, opts ...grpc.CallOption) (*SquartRootResponse, error) {
	out := new(SquartRootResponse)
	err := c.cc.Invoke(ctx, "/squart.SquartRoot/SquartRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquartRootServer is the server API for SquartRoot service.
type SquartRootServer interface {
	SquartRoot(context.Context, *SquartRootRequest) (*SquartRootResponse, error)
}

// UnimplementedSquartRootServer can be embedded to have forward compatible implementations.
type UnimplementedSquartRootServer struct {
}

func (*UnimplementedSquartRootServer) SquartRoot(ctx context.Context, req *SquartRootRequest) (*SquartRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquartRoot not implemented")
}

func RegisterSquartRootServer(s *grpc.Server, srv SquartRootServer) {
	s.RegisterService(&_SquartRoot_serviceDesc, srv)
}

func _SquartRoot_SquartRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquartRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquartRootServer).SquartRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squart.SquartRoot/SquartRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquartRootServer).SquartRoot(ctx, req.(*SquartRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquartRoot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squart.SquartRoot",
	HandlerType: (*SquartRootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquartRoot",
			Handler:    _SquartRoot_SquartRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SquartRoot/squartpb/squart.proto",
}
