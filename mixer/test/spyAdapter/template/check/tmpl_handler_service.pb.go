// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/test/spyAdapter/template/check/tmpl_handler_service.proto

package samplecheck

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	v1beta1 "istio.io/api/mixer/adapter/model/v1beta1"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for HandleSampleCheck method.
type HandleSampleCheckRequest struct {
	// 'samplecheck' instance.
	Instance *InstanceMsg `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *types.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig,proto3" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleSampleCheckRequest) Reset()      { *m = HandleSampleCheckRequest{} }
func (*HandleSampleCheckRequest) ProtoMessage() {}
func (*HandleSampleCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cdff2ff38982ad, []int{0}
}
func (m *HandleSampleCheckRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandleSampleCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandleSampleCheckRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandleSampleCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleSampleCheckRequest.Merge(m, src)
}
func (m *HandleSampleCheckRequest) XXX_Size() int {
	return m.Size()
}
func (m *HandleSampleCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleSampleCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleSampleCheckRequest proto.InternalMessageInfo

// Contains instance payload for 'samplecheck' template. This is passed to infrastructure backends during request-time
// through HandleSampleCheckService.HandleSampleCheck.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name            string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
	StringPrimitive string `protobuf:"bytes,4,opt,name=stringPrimitive,proto3" json:"stringPrimitive,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cdff2ff38982ad, []int{1}
}
func (m *InstanceMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstanceMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstanceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceMsg.Merge(m, src)
}
func (m *InstanceMsg) XXX_Size() int {
	return m.Size()
}
func (m *InstanceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceMsg proto.InternalMessageInfo

// Contains inferred type information about specific instance of 'samplecheck' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()      { *m = Type{} }
func (*Type) ProtoMessage() {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cdff2ff38982ad, []int{2}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Type.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return m.Size()
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

// Represents instance configuration schema for 'samplecheck' template.
type InstanceParam struct {
	StringPrimitive string `protobuf:"bytes,4,opt,name=stringPrimitive,proto3" json:"stringPrimitive,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cdff2ff38982ad, []int{3}
}
func (m *InstanceParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstanceParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstanceParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceParam.Merge(m, src)
}
func (m *InstanceParam) XXX_Size() int {
	return m.Size()
}
func (m *InstanceParam) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceParam.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceParam proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HandleSampleCheckRequest)(nil), "samplecheck.HandleSampleCheckRequest")
	proto.RegisterType((*InstanceMsg)(nil), "samplecheck.InstanceMsg")
	proto.RegisterType((*Type)(nil), "samplecheck.Type")
	proto.RegisterType((*InstanceParam)(nil), "samplecheck.InstanceParam")
}

func init() {
	proto.RegisterFile("mixer/test/spyAdapter/template/check/tmpl_handler_service.proto", fileDescriptor_d4cdff2ff38982ad)
}

var fileDescriptor_d4cdff2ff38982ad = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x6f, 0x13, 0x31,
	0x18, 0xc6, 0xed, 0x10, 0x95, 0xe2, 0xa8, 0x20, 0x4e, 0x45, 0xba, 0x66, 0xb0, 0xaa, 0x93, 0x10,
	0x19, 0x90, 0xad, 0x16, 0x16, 0xc4, 0x80, 0x4a, 0x17, 0x3a, 0x00, 0x55, 0xca, 0x1e, 0x39, 0x77,
	0x6f, 0xaf, 0x86, 0x3b, 0xfb, 0x38, 0x3b, 0x51, 0xb3, 0x21, 0x06, 0x66, 0xa4, 0x7e, 0x89, 0x6e,
	0x7c, 0x01, 0x3e, 0x40, 0xc5, 0x14, 0x31, 0x75, 0x41, 0xe2, 0x2e, 0x1d, 0x18, 0x3b, 0x32, 0xa2,
	0xfa, 0x1c, 0x14, 0xfe, 0xaa, 0xdb, 0xbd, 0x7e, 0x7e, 0x7e, 0xdf, 0xe7, 0xde, 0xc7, 0xe4, 0x51,
	0x2e, 0x0f, 0xa1, 0xe4, 0x16, 0x8c, 0xe5, 0xa6, 0x98, 0x6c, 0x25, 0xa2, 0xb0, 0xae, 0xce, 0x8b,
	0x4c, 0x58, 0xe0, 0xf1, 0x01, 0xc4, 0xaf, 0xb8, 0xcd, 0x8b, 0x6c, 0x70, 0x20, 0x54, 0x92, 0x41,
	0x39, 0x30, 0x50, 0x8e, 0x65, 0x0c, 0xac, 0x28, 0xb5, 0xd5, 0x41, 0xc7, 0x88, 0xbc, 0xc8, 0xc0,
	0x71, 0xdd, 0xd5, 0x54, 0xa7, 0xda, 0x9d, 0xf3, 0x8b, 0xaf, 0x06, 0xe9, 0xde, 0x6d, 0x66, 0x08,
	0xdf, 0x3b, 0xd7, 0x09, 0x64, 0x7c, 0xbc, 0x31, 0x04, 0x2b, 0x36, 0x38, 0x1c, 0x5a, 0x50, 0x46,
	0x6a, 0x65, 0x3c, 0xbd, 0x96, 0x6a, 0x9d, 0x66, 0xc0, 0x5d, 0x35, 0x1c, 0xed, 0x73, 0xa1, 0x26,
	0x5e, 0xba, 0xf3, 0xbf, 0x46, 0xce, 0x41, 0x03, 0x46, 0xc7, 0x98, 0x84, 0x4f, 0x9c, 0xdd, 0x3d,
	0xe7, 0x6e, 0xfb, 0x42, 0xeb, 0xc3, 0xeb, 0x11, 0x18, 0x1b, 0xdc, 0x27, 0xcb, 0x52, 0x19, 0x2b,
	0x54, 0x0c, 0x21, 0x5e, 0xc7, 0xbd, 0xce, 0x66, 0xc8, 0x16, 0x7e, 0x82, 0xed, 0x78, 0xf1, 0xa9,
	0x49, 0xfb, 0x3f, 0xc9, 0xe0, 0x21, 0xb9, 0xee, 0xe7, 0x0e, 0x62, 0xad, 0xf6, 0x65, 0x1a, 0xb6,
	0xdc, 0xdd, 0x55, 0xd6, 0xf8, 0x65, 0x73, 0xbf, 0x6c, 0x4b, 0x4d, 0xfa, 0x2b, 0x9e, 0xdd, 0x76,
	0x68, 0xb0, 0x46, 0x96, 0x13, 0x48, 0x46, 0xc5, 0x40, 0x26, 0xe1, 0x95, 0x75, 0xdc, 0xbb, 0xd6,
	0xbf, 0xea, 0xea, 0x9d, 0x24, 0x7a, 0x46, 0x3a, 0x0b, 0x03, 0x83, 0x5b, 0xa4, 0xad, 0x44, 0x0e,
	0xe1, 0x87, 0x4f, 0x1f, 0x23, 0x07, 0xba, 0x32, 0xe8, 0x91, 0x1b, 0xc6, 0x96, 0x52, 0xa5, 0xbb,
	0xa5, 0xcc, 0xa5, 0x95, 0x63, 0x08, 0xdb, 0x4e, 0xfe, 0xfd, 0x38, 0x5a, 0x22, 0xed, 0x17, 0x93,
	0x02, 0xa2, 0x07, 0x64, 0x65, 0xde, 0x77, 0x57, 0x94, 0x22, 0xbf, 0x7c, 0x8b, 0xcd, 0x77, 0x7f,
	0xdb, 0xde, 0x5e, 0x93, 0x7a, 0xf0, 0x92, 0xdc, 0xfc, 0x43, 0x0b, 0x6e, 0xff, 0xb2, 0xc0, 0x7f,
	0x6d, 0xbe, 0xcb, 0x98, 0x34, 0x56, 0x6a, 0xe6, 0x62, 0x64, 0x7e, 0x45, 0xcc, 0xc5, 0xc8, 0x7c,
	0x8c, 0xcc, 0x5f, 0x30, 0xa3, 0xcc, 0x3e, 0x7e, 0x7e, 0x52, 0x51, 0x34, 0xad, 0x28, 0x3a, 0xad,
	0x28, 0x3a, 0xaf, 0x28, 0x7a, 0x53, 0x53, 0x7c, 0x5c, 0x53, 0x74, 0x52, 0x53, 0x3c, 0xad, 0x29,
	0xfe, 0x5a, 0x53, 0xfc, 0xad, 0xa6, 0xe8, 0xbc, 0xa6, 0xf8, 0xfd, 0x8c, 0xa2, 0xe9, 0x8c, 0xa2,
	0xd3, 0x19, 0x45, 0xdf, 0x3f, 0x9f, 0x1d, 0xb5, 0xd0, 0xdb, 0x2f, 0x67, 0x47, 0xad, 0xc5, 0xf7,
	0x39, 0x5c, 0x72, 0x21, 0xdd, 0xfb, 0x11, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x23, 0xd8, 0x06, 0xf6,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandleSampleCheckServiceClient is the client API for HandleSampleCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleSampleCheckServiceClient interface {
	// HandleSampleCheck is called by Mixer at request-time to deliver 'samplecheck' instances to the backend.
	HandleSampleCheck(ctx context.Context, in *HandleSampleCheckRequest, opts ...grpc.CallOption) (*v1beta1.CheckResult, error)
}

type handleSampleCheckServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleSampleCheckServiceClient(cc *grpc.ClientConn) HandleSampleCheckServiceClient {
	return &handleSampleCheckServiceClient{cc}
}

func (c *handleSampleCheckServiceClient) HandleSampleCheck(ctx context.Context, in *HandleSampleCheckRequest, opts ...grpc.CallOption) (*v1beta1.CheckResult, error) {
	out := new(v1beta1.CheckResult)
	err := c.cc.Invoke(ctx, "/samplecheck.HandleSampleCheckService/HandleSampleCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleSampleCheckServiceServer is the server API for HandleSampleCheckService service.
type HandleSampleCheckServiceServer interface {
	// HandleSampleCheck is called by Mixer at request-time to deliver 'samplecheck' instances to the backend.
	HandleSampleCheck(context.Context, *HandleSampleCheckRequest) (*v1beta1.CheckResult, error)
}

// UnimplementedHandleSampleCheckServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHandleSampleCheckServiceServer struct {
}

func (*UnimplementedHandleSampleCheckServiceServer) HandleSampleCheck(ctx context.Context, req *HandleSampleCheckRequest) (*v1beta1.CheckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSampleCheck not implemented")
}

func RegisterHandleSampleCheckServiceServer(s *grpc.Server, srv HandleSampleCheckServiceServer) {
	s.RegisterService(&_HandleSampleCheckService_serviceDesc, srv)
}

func _HandleSampleCheckService_HandleSampleCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleSampleCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleSampleCheckServiceServer).HandleSampleCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/samplecheck.HandleSampleCheckService/HandleSampleCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleSampleCheckServiceServer).HandleSampleCheck(ctx, req.(*HandleSampleCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleSampleCheckService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "samplecheck.HandleSampleCheckService",
	HandlerType: (*HandleSampleCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSampleCheck",
			Handler:    _HandleSampleCheckService_HandleSampleCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/test/spyAdapter/template/check/tmpl_handler_service.proto",
}

func (m *HandleSampleCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleSampleCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandleSampleCheckRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DedupId) > 0 {
		i -= len(m.DedupId)
		copy(dAtA[i:], m.DedupId)
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AdapterConfig != nil {
		{
			size, err := m.AdapterConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTmplHandlerService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Instance != nil {
		{
			size, err := m.Instance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTmplHandlerService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstanceMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2
		i--
		dAtA[i] = 0x93
		i--
		dAtA[i] = 0xe4
		i--
		dAtA[i] = 0xd2
		i--
		dAtA[i] = 0xfa
	}
	if len(m.StringPrimitive) > 0 {
		i -= len(m.StringPrimitive)
		copy(dAtA[i:], m.StringPrimitive)
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.StringPrimitive)))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Type) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstanceParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StringPrimitive) > 0 {
		i -= len(m.StringPrimitive)
		copy(dAtA[i:], m.StringPrimitive)
		i = encodeVarintTmplHandlerService(dAtA, i, uint64(len(m.StringPrimitive)))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func encodeVarintTmplHandlerService(dAtA []byte, offset int, v uint64) int {
	offset -= sovTmplHandlerService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HandleSampleCheckRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Instance != nil {
		l = m.Instance.Size()
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringPrimitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTmplHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringPrimitive)
	if l > 0 {
		n += 1 + l + sovTmplHandlerService(uint64(l))
	}
	return n
}

func sovTmplHandlerService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTmplHandlerService(x uint64) (n int) {
	return sovTmplHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleSampleCheckRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleSampleCheckRequest{`,
		`Instance:` + strings.Replace(this.Instance.String(), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "types.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`StringPrimitive:` + fmt.Sprintf("%v", this.StringPrimitive) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`StringPrimitive:` + fmt.Sprintf("%v", this.StringPrimitive) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTmplHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleSampleCheckRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmplHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandleSampleCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleSampleCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Instance == nil {
				m.Instance = &InstanceMsg{}
			}
			if err := m.Instance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &types.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmplHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmplHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTmplHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringPrimitive", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringPrimitive = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTmplHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTmplHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTmplHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTmplHandlerService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTmplHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTmplHandlerService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTmplHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTmplHandlerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTmplHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTmplHandlerService
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTmplHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTmplHandlerService   = fmt.Errorf("proto: integer overflow")
)
