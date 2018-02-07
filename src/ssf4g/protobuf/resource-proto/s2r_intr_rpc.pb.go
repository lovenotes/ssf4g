// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2r_intr_rpc.proto

/*
Package resourceproto is a generated protocol buffer package.

It is generated from these files:
	s2r_intr_rpc.proto

It has these top-level messages:
	S2RZoneStatusGetReqProto
	S2RZoneStatusGetRespProto
*/
package resourceproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type S2RZoneStatusGetReqProto struct {
	ZoneID uint32 `protobuf:"varint,1,opt,name=zoneID" json:"zoneID,omitempty"`
}

func (m *S2RZoneStatusGetReqProto) Reset()                    { *m = S2RZoneStatusGetReqProto{} }
func (m *S2RZoneStatusGetReqProto) String() string            { return proto.CompactTextString(m) }
func (*S2RZoneStatusGetReqProto) ProtoMessage()               {}
func (*S2RZoneStatusGetReqProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *S2RZoneStatusGetReqProto) GetZoneID() uint32 {
	if m != nil {
		return m.ZoneID
	}
	return 0
}

type S2RZoneStatusGetRespProto struct {
	ZoneID     uint32 `protobuf:"varint,1,opt,name=zoneID" json:"zoneID,omitempty"`
	ZoneStatus int32  `protobuf:"varint,2,opt,name=zoneStatus" json:"zoneStatus,omitempty"`
}

func (m *S2RZoneStatusGetRespProto) Reset()                    { *m = S2RZoneStatusGetRespProto{} }
func (m *S2RZoneStatusGetRespProto) String() string            { return proto.CompactTextString(m) }
func (*S2RZoneStatusGetRespProto) ProtoMessage()               {}
func (*S2RZoneStatusGetRespProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2RZoneStatusGetRespProto) GetZoneID() uint32 {
	if m != nil {
		return m.ZoneID
	}
	return 0
}

func (m *S2RZoneStatusGetRespProto) GetZoneStatus() int32 {
	if m != nil {
		return m.ZoneStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*S2RZoneStatusGetReqProto)(nil), "proto.resource.S2RZoneStatusGetReqProto")
	proto.RegisterType((*S2RZoneStatusGetRespProto)(nil), "proto.resource.S2RZoneStatusGetRespProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceIntrService service

type ResourceIntrServiceClient interface {
	S2RZoneStatusGet(ctx context.Context, in *S2RZoneStatusGetReqProto, opts ...grpc.CallOption) (*S2RZoneStatusGetRespProto, error)
}

type resourceIntrServiceClient struct {
	cc *grpc.ClientConn
}

func NewResourceIntrServiceClient(cc *grpc.ClientConn) ResourceIntrServiceClient {
	return &resourceIntrServiceClient{cc}
}

func (c *resourceIntrServiceClient) S2RZoneStatusGet(ctx context.Context, in *S2RZoneStatusGetReqProto, opts ...grpc.CallOption) (*S2RZoneStatusGetRespProto, error) {
	out := new(S2RZoneStatusGetRespProto)
	err := grpc.Invoke(ctx, "/proto.resource.ResourceIntrService/S2RZoneStatusGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceIntrService service

type ResourceIntrServiceServer interface {
	S2RZoneStatusGet(context.Context, *S2RZoneStatusGetReqProto) (*S2RZoneStatusGetRespProto, error)
}

func RegisterResourceIntrServiceServer(s *grpc.Server, srv ResourceIntrServiceServer) {
	s.RegisterService(&_ResourceIntrService_serviceDesc, srv)
}

func _ResourceIntrService_S2RZoneStatusGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(S2RZoneStatusGetReqProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceIntrServiceServer).S2RZoneStatusGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.resource.ResourceIntrService/S2RZoneStatusGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceIntrServiceServer).S2RZoneStatusGet(ctx, req.(*S2RZoneStatusGetReqProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceIntrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.resource.ResourceIntrService",
	HandlerType: (*ResourceIntrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "S2RZoneStatusGet",
			Handler:    _ResourceIntrService_S2RZoneStatusGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s2r_intr_rpc.proto",
}

func init() { proto.RegisterFile("s2r_intr_rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x36, 0x2a, 0x8a,
	0xcf, 0xcc, 0x2b, 0x29, 0x8a, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x03, 0x53, 0x7a, 0x45, 0xa9, 0xc5, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x4a, 0x46, 0x5c, 0x12, 0xc1,
	0x46, 0x41, 0x51, 0xf9, 0x79, 0xa9, 0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0xee, 0xa9, 0x25, 0x41,
	0xa9, 0x85, 0x01, 0x60, 0xb5, 0x62, 0x5c, 0x6c, 0x55, 0xf9, 0x79, 0xa9, 0x9e, 0x2e, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x50, 0x9e, 0x52, 0x30, 0x97, 0x24, 0xa6, 0x9e, 0xe2, 0x02, 0xbc,
	0x9a, 0x84, 0xe4, 0xb8, 0xb8, 0xaa, 0xe0, 0x3a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x90,
	0x44, 0x8c, 0xea, 0xb8, 0x84, 0x83, 0xa0, 0x8e, 0xf2, 0xcc, 0x2b, 0x29, 0x0a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0x4a, 0xe7, 0x12, 0x40, 0xb7, 0x4b, 0x48, 0x43, 0x0f, 0xd5, 0x13, 0x7a,
	0xb8, 0x7c, 0x20, 0xa5, 0x49, 0x58, 0x25, 0xd4, 0xdd, 0x4e, 0xfc, 0x51, 0xbc, 0x30, 0x55, 0x60,
	0x3d, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x13, 0x19, 0x56, 0x43, 0x46,
	0x01, 0x00, 0x00,
}
