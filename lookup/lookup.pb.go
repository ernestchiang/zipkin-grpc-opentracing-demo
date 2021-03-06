// Code generated by protoc-gen-go.
// source: lookup.proto
// DO NOT EDIT!

/*
Package lookup is a generated protocol buffer package.

It is generated from these files:
	lookup.proto

It has these top-level messages:
	SearchRequest
	SearchReply
*/
package lookup

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

type SearchRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchReply struct {
	Memory string `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
}

func (m *SearchReply) Reset()                    { *m = SearchReply{} }
func (m *SearchReply) String() string            { return proto.CompactTextString(m) }
func (*SearchReply) ProtoMessage()               {}
func (*SearchReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchReply) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "lookup.SearchRequest")
	proto.RegisterType((*SearchReply)(nil), "lookup.SearchReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LookUp service

type LookUpClient interface {
	ElementarySchool(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	HighSchool(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	College(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	Work(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
}

type lookUpClient struct {
	cc *grpc.ClientConn
}

func NewLookUpClient(cc *grpc.ClientConn) LookUpClient {
	return &lookUpClient{cc}
}

func (c *lookUpClient) ElementarySchool(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := grpc.Invoke(ctx, "/lookup.LookUp/ElementarySchool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookUpClient) HighSchool(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := grpc.Invoke(ctx, "/lookup.LookUp/HighSchool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookUpClient) College(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := grpc.Invoke(ctx, "/lookup.LookUp/College", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lookUpClient) Work(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := grpc.Invoke(ctx, "/lookup.LookUp/Work", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LookUp service

type LookUpServer interface {
	ElementarySchool(context.Context, *SearchRequest) (*SearchReply, error)
	HighSchool(context.Context, *SearchRequest) (*SearchReply, error)
	College(context.Context, *SearchRequest) (*SearchReply, error)
	Work(context.Context, *SearchRequest) (*SearchReply, error)
}

func RegisterLookUpServer(s *grpc.Server, srv LookUpServer) {
	s.RegisterService(&_LookUp_serviceDesc, srv)
}

func _LookUp_ElementarySchool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookUpServer).ElementarySchool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lookup.LookUp/ElementarySchool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookUpServer).ElementarySchool(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookUp_HighSchool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookUpServer).HighSchool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lookup.LookUp/HighSchool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookUpServer).HighSchool(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookUp_College_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookUpServer).College(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lookup.LookUp/College",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookUpServer).College(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LookUp_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookUpServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lookup.LookUp/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookUpServer).Work(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LookUp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lookup.LookUp",
	HandlerType: (*LookUpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ElementarySchool",
			Handler:    _LookUp_ElementarySchool_Handler,
		},
		{
			MethodName: "HighSchool",
			Handler:    _LookUp_HighSchool_Handler,
		},
		{
			MethodName: "College",
			Handler:    _LookUp_College_Handler,
		},
		{
			MethodName: "Work",
			Handler:    _LookUp_Work_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lookup.proto",
}

func init() { proto.RegisterFile("lookup.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0xcf, 0xcf,
	0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x94, 0xb9, 0x78,
	0x83, 0x53, 0x13, 0x8b, 0x92, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8,
	0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x55,
	0x2e, 0x6e, 0x98, 0xa2, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xdc, 0xd4, 0xdc, 0xfc, 0xa2,
	0x4a, 0x09, 0x26, 0xb0, 0x22, 0x28, 0xcf, 0xe8, 0x2b, 0x23, 0x17, 0x9b, 0x4f, 0x7e, 0x7e, 0x76,
	0x68, 0x81, 0x90, 0x03, 0x97, 0x80, 0x6b, 0x4e, 0x6a, 0x6e, 0x6a, 0x5e, 0x49, 0x62, 0x51, 0x65,
	0x70, 0x72, 0x46, 0x7e, 0x7e, 0x8e, 0x90, 0xa8, 0x1e, 0xd4, 0x05, 0x28, 0x16, 0x4a, 0x09, 0xa3,
	0x0b, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x59, 0x71, 0x71, 0x79, 0x64, 0xa6, 0x67, 0x90, 0xa5,
	0xd7, 0x9c, 0x8b, 0xdd, 0x39, 0x3f, 0x27, 0x27, 0x35, 0x3d, 0x95, 0x44, 0x8d, 0x26, 0x5c, 0x2c,
	0xe1, 0xf9, 0x45, 0xd9, 0xa4, 0xe9, 0x72, 0xd2, 0xe0, 0x12, 0xcf, 0xcc, 0xd7, 0x4b, 0x2f, 0x2a,
	0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x86, 0x2a, 0x74, 0xe2, 0xf6, 0x01,
	0xd3, 0x01, 0xa0, 0x30, 0x0f, 0x60, 0x4c, 0x62, 0x03, 0x07, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xbe, 0x9b, 0x21, 0x5d, 0x8c, 0x01, 0x00, 0x00,
}
