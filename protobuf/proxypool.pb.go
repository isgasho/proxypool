// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/proxypool.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Protocol int32

const (
	Protocol_UNKNOWN Protocol = 0
	Protocol_HTTPS   Protocol = 1
	Protocol_SOCKS   Protocol = 2
	Protocol_HTTP    Protocol = 3
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "HTTPS",
	2: "SOCKS",
	3: "HTTP",
}

var Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"HTTPS":   1,
	"SOCKS":   2,
	"HTTP":    3,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e479f9deb75e7ea6, []int{0}
}

type Proxy struct {
	Address              string               `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
	Port                 uint64               `protobuf:"varint,2,opt,name=Port,json=port,proto3" json:"Port,omitempty"`
	Protocol             Protocol             `protobuf:"varint,3,opt,name=Protocol,json=protocol,proto3,enum=rpc.Protocol" json:"Protocol,omitempty"`
	LastCheck            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=LastCheck,json=last_check,proto3" json:"LastCheck,omitempty"`
	Speed                uint64               `protobuf:"varint,5,opt,name=Speed,json=speed,proto3" json:"Speed,omitempty"`
	From                 string               `protobuf:"bytes,6,opt,name=From,json=from,proto3" json:"From,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e479f9deb75e7ea6, []int{0}
}

func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (m *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(m, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Proxy) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Proxy) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

func (m *Proxy) GetLastCheck() *timestamp.Timestamp {
	if m != nil {
		return m.LastCheck
	}
	return nil
}

func (m *Proxy) GetSpeed() uint64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Proxy) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

type Proxies struct {
	Counts               uint64   `protobuf:"varint,1,opt,name=Counts,json=counts,proto3" json:"Counts,omitempty"`
	Proxies              []*Proxy `protobuf:"bytes,2,rep,name=proxies,proto3" json:"proxies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proxies) Reset()         { *m = Proxies{} }
func (m *Proxies) String() string { return proto.CompactTextString(m) }
func (*Proxies) ProtoMessage()    {}
func (*Proxies) Descriptor() ([]byte, []int) {
	return fileDescriptor_e479f9deb75e7ea6, []int{1}
}

func (m *Proxies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxies.Unmarshal(m, b)
}
func (m *Proxies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxies.Marshal(b, m, deterministic)
}
func (m *Proxies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxies.Merge(m, src)
}
func (m *Proxies) XXX_Size() int {
	return xxx_messageInfo_Proxies.Size(m)
}
func (m *Proxies) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxies.DiscardUnknown(m)
}

var xxx_messageInfo_Proxies proto.InternalMessageInfo

func (m *Proxies) GetCounts() uint64 {
	if m != nil {
		return m.Counts
	}
	return 0
}

func (m *Proxies) GetProxies() []*Proxy {
	if m != nil {
		return m.Proxies
	}
	return nil
}

func init() {
	proto.RegisterEnum("rpc.Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*Proxy)(nil), "rpc.Proxy")
	proto.RegisterType((*Proxies)(nil), "rpc.Proxies")
}

func init() { proto.RegisterFile("protobuf/proxypool.proto", fileDescriptor_e479f9deb75e7ea6) }

var fileDescriptor_e479f9deb75e7ea6 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0x9d, 0xcd, 0xbf, 0xcd, 0x5d, 0x95, 0x70, 0x91, 0x32, 0xc4, 0x07, 0x43, 0xf1, 0x21,
	0x55, 0x49, 0x21, 0x3e, 0x88, 0x8f, 0xcb, 0xe2, 0x1f, 0xa8, 0x6c, 0x43, 0xb2, 0xe2, 0xa3, 0xa4,
	0xc9, 0xb4, 0x2e, 0x26, 0xbd, 0xc3, 0xcc, 0x14, 0xba, 0x1f, 0xc3, 0x0f, 0xe5, 0xf7, 0x92, 0x49,
	0xcc, 0xba, 0x58, 0xf6, 0x6d, 0xee, 0x3d, 0x3f, 0xce, 0x3d, 0x87, 0x01, 0x2e, 0x15, 0x19, 0xba,
	0xba, 0xbb, 0x3e, 0x97, 0x8a, 0xee, 0x77, 0x92, 0xa8, 0xcb, 0x86, 0x15, 0x3a, 0x4a, 0x36, 0xf1,
	0x8b, 0x1b, 0xa2, 0x9b, 0x4e, 0x9c, 0xef, 0x29, 0xb3, 0xed, 0x85, 0x36, 0x75, 0x2f, 0x47, 0x2a,
	0x7e, 0xfe, 0x3f, 0x20, 0x7a, 0x69, 0x76, 0xa3, 0x78, 0xfa, 0x9b, 0x81, 0x57, 0x58, 0x5b, 0xe4,
	0x10, 0x2c, 0xdb, 0x56, 0x09, 0xad, 0x39, 0x4b, 0x58, 0x1a, 0x96, 0x41, 0x3d, 0x8e, 0x88, 0xe0,
	0x16, 0xa4, 0x0c, 0x9f, 0x25, 0x2c, 0x75, 0x4b, 0x57, 0x92, 0x32, 0x78, 0x06, 0xf3, 0xc2, 0x1a,
	0x34, 0xd4, 0x71, 0x27, 0x61, 0xe9, 0xd3, 0xfc, 0x49, 0xa6, 0x64, 0x93, 0x4d, 0xcb, 0x72, 0x2e,
	0xff, 0xbe, 0xf0, 0x3d, 0x84, 0x5f, 0x6a, 0x6d, 0x56, 0x3f, 0x44, 0xf3, 0x93, 0xbb, 0x09, 0x4b,
	0x17, 0x79, 0x9c, 0x8d, 0x99, 0xb2, 0x29, 0x53, 0xb6, 0x99, 0x42, 0x97, 0xd0, 0xd5, 0xda, 0x7c,
	0x6f, 0x2c, 0x8d, 0xcf, 0xc0, 0xab, 0xa4, 0x10, 0x2d, 0xf7, 0x86, 0xd3, 0x9e, 0xb6, 0x83, 0xcd,
	0xf3, 0x51, 0x51, 0xcf, 0xfd, 0x21, 0xa6, 0x7b, 0xad, 0xa8, 0x3f, 0xfd, 0x04, 0x81, 0xad, 0xb1,
	0x15, 0x1a, 0x4f, 0xc0, 0x5f, 0xd1, 0xdd, 0xad, 0x19, 0x7b, 0xb8, 0xa5, 0xdf, 0x0c, 0x13, 0xbe,
	0x84, 0x40, 0x8e, 0x08, 0x9f, 0x25, 0x4e, 0xba, 0xc8, 0x61, 0x4a, 0x7c, 0xbf, 0x2b, 0x27, 0xe9,
	0xd5, 0xbb, 0x7f, 0xc5, 0x70, 0x01, 0xc1, 0xd7, 0xf5, 0xc5, 0xfa, 0xf2, 0xdb, 0x3a, 0x7a, 0x84,
	0x21, 0x78, 0x9f, 0x37, 0x9b, 0xa2, 0x8a, 0x98, 0x7d, 0x56, 0x97, 0xab, 0x8b, 0x2a, 0x9a, 0xe1,
	0x1c, 0x5c, 0xbb, 0x8d, 0x9c, 0xfc, 0x17, 0x83, 0x70, 0xf0, 0x2a, 0x88, 0x3a, 0x3c, 0x03, 0x67,
	0xd9, 0xb6, 0x78, 0x70, 0x22, 0x3e, 0x79, 0x50, 0xfa, 0x83, 0xfd, 0x08, 0x7c, 0x03, 0x7e, 0x59,
	0xdf, 0xb6, 0xd4, 0xe3, 0x11, 0x22, 0x3e, 0x70, 0xc1, 0xd7, 0xe0, 0x2c, 0xbb, 0xee, 0x28, 0xfa,
	0x78, 0x8f, 0x6e, 0x85, 0xbe, 0xf2, 0x07, 0xf5, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0x48, 0x28, 0x11, 0x43, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyPoolClient is the client API for ProxyPool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyPoolClient interface {
	Add(ctx context.Context, in *Proxy, opts ...grpc.CallOption) (*empty.Empty, error)
	Random(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Proxy, error)
	All(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Proxies, error)
}

type proxyPoolClient struct {
	cc *grpc.ClientConn
}

func NewProxyPoolClient(cc *grpc.ClientConn) ProxyPoolClient {
	return &proxyPoolClient{cc}
}

func (c *proxyPoolClient) Add(ctx context.Context, in *Proxy, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.ProxyPool/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyPoolClient) Random(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Proxy, error) {
	out := new(Proxy)
	err := c.cc.Invoke(ctx, "/rpc.ProxyPool/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyPoolClient) All(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Proxies, error) {
	out := new(Proxies)
	err := c.cc.Invoke(ctx, "/rpc.ProxyPool/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyPoolServer is the server API for ProxyPool service.
type ProxyPoolServer interface {
	Add(context.Context, *Proxy) (*empty.Empty, error)
	Random(context.Context, *empty.Empty) (*Proxy, error)
	All(context.Context, *empty.Empty) (*Proxies, error)
}

// UnimplementedProxyPoolServer can be embedded to have forward compatible implementations.
type UnimplementedProxyPoolServer struct {
}

func (*UnimplementedProxyPoolServer) Add(ctx context.Context, req *Proxy) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedProxyPoolServer) Random(ctx context.Context, req *empty.Empty) (*Proxy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (*UnimplementedProxyPoolServer) All(ctx context.Context, req *empty.Empty) (*Proxies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}

func RegisterProxyPoolServer(s *grpc.Server, srv ProxyPoolServer) {
	s.RegisterService(&_ProxyPool_serviceDesc, srv)
}

func _ProxyPool_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proxy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyPoolServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ProxyPool/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyPoolServer).Add(ctx, req.(*Proxy))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyPool_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyPoolServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ProxyPool/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyPoolServer).Random(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyPool_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyPoolServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ProxyPool/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyPoolServer).All(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyPool_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ProxyPool",
	HandlerType: (*ProxyPoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ProxyPool_Add_Handler,
		},
		{
			MethodName: "Random",
			Handler:    _ProxyPool_Random_Handler,
		},
		{
			MethodName: "All",
			Handler:    _ProxyPool_All_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/proxypool.proto",
}
