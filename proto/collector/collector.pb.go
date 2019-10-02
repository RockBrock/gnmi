// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/collector/collector.proto

package gnmi

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

type ReconnectRequest struct {
	Target               []string `protobuf:"bytes,1,rep,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReconnectRequest) Reset()         { *m = ReconnectRequest{} }
func (m *ReconnectRequest) String() string { return proto.CompactTextString(m) }
func (*ReconnectRequest) ProtoMessage()    {}
func (*ReconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32029eaac52468e, []int{0}
}

func (m *ReconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReconnectRequest.Unmarshal(m, b)
}
func (m *ReconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReconnectRequest.Marshal(b, m, deterministic)
}
func (m *ReconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReconnectRequest.Merge(m, src)
}
func (m *ReconnectRequest) XXX_Size() int {
	return xxx_messageInfo_ReconnectRequest.Size(m)
}
func (m *ReconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReconnectRequest proto.InternalMessageInfo

func (m *ReconnectRequest) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

type Nil struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nil) Reset()         { *m = Nil{} }
func (m *Nil) String() string { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()    {}
func (*Nil) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32029eaac52468e, []int{1}
}

func (m *Nil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nil.Unmarshal(m, b)
}
func (m *Nil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nil.Marshal(b, m, deterministic)
}
func (m *Nil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nil.Merge(m, src)
}
func (m *Nil) XXX_Size() int {
	return xxx_messageInfo_Nil.Size(m)
}
func (m *Nil) XXX_DiscardUnknown() {
	xxx_messageInfo_Nil.DiscardUnknown(m)
}

var xxx_messageInfo_Nil proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReconnectRequest)(nil), "gnmi.ReconnectRequest")
	proto.RegisterType((*Nil)(nil), "gnmi.Nil")
}

func init() { proto.RegisterFile("proto/collector/collector.proto", fileDescriptor_f32029eaac52468e) }

var fileDescriptor_f32029eaac52468e = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0x42, 0xb0, 0xf4, 0xc0, 0x32,
	0x42, 0x2c, 0xe9, 0x79, 0xb9, 0x99, 0x4a, 0x5a, 0x5c, 0x02, 0x41, 0xa9, 0xc9, 0xf9, 0x79, 0x79,
	0xa9, 0xc9, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x25, 0x89,
	0x45, 0xe9, 0xa9, 0x25, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x12, 0x2b, 0x17,
	0xb3, 0x5f, 0x66, 0x8e, 0x91, 0x2d, 0x17, 0xa7, 0x33, 0xcc, 0x2c, 0x21, 0x03, 0x2e, 0x4e, 0xb8,
	0x7e, 0x21, 0x31, 0x3d, 0x90, 0x99, 0x7a, 0xe8, 0x06, 0x4a, 0x71, 0x42, 0xc4, 0xfd, 0x32, 0x73,
	0x94, 0x18, 0x92, 0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x2b, 0x57,
	0x3a, 0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectorClient interface {
	// Reconnect requests that the existing connections for one or more specified
	// targets will be stopped and new connections established.
	Reconnect(ctx context.Context, in *ReconnectRequest, opts ...grpc.CallOption) (*Nil, error)
}

type collectorClient struct {
	cc *grpc.ClientConn
}

func NewCollectorClient(cc *grpc.ClientConn) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) Reconnect(ctx context.Context, in *ReconnectRequest, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/gnmi.Collector/Reconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServer is the server API for Collector service.
type CollectorServer interface {
	// Reconnect requests that the existing connections for one or more specified
	// targets will be stopped and new connections established.
	Reconnect(context.Context, *ReconnectRequest) (*Nil, error)
}

// UnimplementedCollectorServer can be embedded to have forward compatible implementations.
type UnimplementedCollectorServer struct {
}

func (*UnimplementedCollectorServer) Reconnect(ctx context.Context, req *ReconnectRequest) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reconnect not implemented")
}

func RegisterCollectorServer(s *grpc.Server, srv CollectorServer) {
	s.RegisterService(&_Collector_serviceDesc, srv)
}

func _Collector_Reconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).Reconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnmi.Collector/Reconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).Reconnect(ctx, req.(*ReconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Collector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnmi.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reconnect",
			Handler:    _Collector_Reconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/collector/collector.proto",
}