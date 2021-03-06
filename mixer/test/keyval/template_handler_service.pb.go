// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/test/keyval/template_handler_service.proto

//
// Input template

package keyval

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	io "io"
	v1beta1 "istio.io/api/mixer/adapter/model/v1beta1"
	math "math"
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

// Request message for HandleKeyval method.
type HandleKeyvalRequest struct {
	// 'keyval' instance.
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

func (m *HandleKeyvalRequest) Reset()      { *m = HandleKeyvalRequest{} }
func (*HandleKeyvalRequest) ProtoMessage() {}
func (*HandleKeyvalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{0}
}
func (m *HandleKeyvalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandleKeyvalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandleKeyvalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandleKeyvalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleKeyvalRequest.Merge(m, src)
}
func (m *HandleKeyvalRequest) XXX_Size() int {
	return m.Size()
}
func (m *HandleKeyvalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleKeyvalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleKeyvalRequest proto.InternalMessageInfo

type HandleKeyvalResponse struct {
	Result *v1beta1.CheckResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Output *OutputMsg           `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *HandleKeyvalResponse) Reset()      { *m = HandleKeyvalResponse{} }
func (*HandleKeyvalResponse) ProtoMessage() {}
func (*HandleKeyvalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{1}
}
func (m *HandleKeyvalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandleKeyvalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandleKeyvalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandleKeyvalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleKeyvalResponse.Merge(m, src)
}
func (m *HandleKeyvalResponse) XXX_Size() int {
	return m.Size()
}
func (m *HandleKeyvalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleKeyvalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandleKeyvalResponse proto.InternalMessageInfo

// Contains output payload for 'keyval' template.
type OutputMsg struct {
	// value to return
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *OutputMsg) Reset()      { *m = OutputMsg{} }
func (*OutputMsg) ProtoMessage() {}
func (*OutputMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{2}
}
func (m *OutputMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutputMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutputMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutputMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputMsg.Merge(m, src)
}
func (m *OutputMsg) XXX_Size() int {
	return m.Size()
}
func (m *OutputMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputMsg.DiscardUnknown(m)
}

var xxx_messageInfo_OutputMsg proto.InternalMessageInfo

// Contains instance payload for 'keyval' template. This is passed to infrastructure backends during request-time
// through HandleKeyvalService.HandleKeyval.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
	// lookup key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{3}
}
func (m *InstanceMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstanceMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
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

// Contains inferred type information about specific instance of 'keyval' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
}

func (m *Type) Reset()      { *m = Type{} }
func (*Type) ProtoMessage() {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{4}
}
func (m *Type) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Type.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
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

// Represents instance configuration schema for 'keyval' template.
type InstanceParam struct {
	// lookup key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_021857cd14e5196f, []int{5}
}
func (m *InstanceParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstanceParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstanceParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
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
	proto.RegisterType((*HandleKeyvalRequest)(nil), "keyval.HandleKeyvalRequest")
	proto.RegisterType((*HandleKeyvalResponse)(nil), "keyval.HandleKeyvalResponse")
	proto.RegisterType((*OutputMsg)(nil), "keyval.OutputMsg")
	proto.RegisterType((*InstanceMsg)(nil), "keyval.InstanceMsg")
	proto.RegisterType((*Type)(nil), "keyval.Type")
	proto.RegisterType((*InstanceParam)(nil), "keyval.InstanceParam")
}

func init() {
	proto.RegisterFile("mixer/test/keyval/template_handler_service.proto", fileDescriptor_021857cd14e5196f)
}

var fileDescriptor_021857cd14e5196f = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0x3e, 0xb7, 0x21, 0x34, 0x2e, 0x45, 0xe0, 0x06, 0x29, 0x0d, 0xc8, 0x6a, 0x6f, 0xa1, 0x48,
	0xc8, 0xa6, 0x45, 0x62, 0x61, 0x82, 0x4a, 0x88, 0x80, 0x10, 0xe8, 0x60, 0x0f, 0x4e, 0xee, 0xf5,
	0x7a, 0xca, 0xc5, 0x3e, 0xce, 0xbe, 0xa8, 0xb7, 0x21, 0x36, 0x36, 0xa4, 0xee, 0xcc, 0x6c, 0xfc,
	0x01, 0x7e, 0x40, 0xc5, 0x14, 0x31, 0x75, 0x41, 0x22, 0x97, 0x0e, 0x8c, 0x1d, 0x19, 0x51, 0x7c,
	0x4e, 0xd5, 0x8a, 0x8a, 0xcd, 0xef, 0x7b, 0xdf, 0xf7, 0xfc, 0x7d, 0x7e, 0xc6, 0xf7, 0x86, 0xf1,
	0x3e, 0x64, 0xdc, 0x80, 0x36, 0x7c, 0x00, 0xc5, 0x48, 0x24, 0xdc, 0xc0, 0x30, 0x4d, 0x84, 0x81,
	0xee, 0x9e, 0x90, 0x61, 0x02, 0x59, 0x57, 0x43, 0x36, 0x8a, 0xfb, 0xc0, 0xd2, 0x4c, 0x19, 0x45,
	0xea, 0x15, 0xad, 0xdd, 0x8c, 0x54, 0xa4, 0x2c, 0xc4, 0x67, 0xa7, 0xaa, 0xdb, 0xbe, 0x5b, 0xcd,
	0x13, 0xa1, 0x48, 0x0d, 0x64, 0x7c, 0xa8, 0x42, 0x48, 0xf8, 0x68, 0xab, 0x07, 0x46, 0x6c, 0x71,
	0xd8, 0x37, 0x20, 0x75, 0xac, 0xa4, 0x76, 0xec, 0xb5, 0x48, 0xa9, 0x28, 0x01, 0x6e, 0xab, 0x5e,
	0xbe, 0xcb, 0x85, 0x2c, 0x5c, 0xeb, 0xf6, 0xff, 0x06, 0xf5, 0xf7, 0xa0, 0x3f, 0xa8, 0x88, 0xfe,
	0x67, 0x84, 0x57, 0x9f, 0x5a, 0xa7, 0xcf, 0xad, 0xb1, 0x00, 0xde, 0xe5, 0xa0, 0x0d, 0xe1, 0x78,
	0x29, 0x96, 0xda, 0x08, 0xd9, 0x87, 0x16, 0x5a, 0x47, 0x9b, 0xcb, 0xdb, 0xab, 0xac, 0xb2, 0xce,
	0x3a, 0x0e, 0x7f, 0xa1, 0xa3, 0xe0, 0x94, 0x44, 0x1e, 0xe2, 0xab, 0xee, 0xb6, 0x6e, 0x5f, 0xc9,
	0xdd, 0x38, 0x6a, 0x2d, 0x58, 0x59, 0x93, 0x55, 0x2e, 0xd9, 0xdc, 0x25, 0x7b, 0x24, 0x8b, 0x60,
	0xc5, 0x71, 0x77, 0x2c, 0x95, 0xac, 0xe1, 0xa5, 0x10, 0xc2, 0x3c, 0xed, 0xc6, 0x61, 0x6b, 0x71,
	0x1d, 0x6d, 0x36, 0x82, 0xcb, 0xb6, 0xee, 0x84, 0xfe, 0x47, 0x84, 0x9b, 0xe7, 0x0d, 0xea, 0x54,
	0x49, 0x0d, 0xe4, 0x09, 0xae, 0x67, 0xa0, 0xf3, 0xc4, 0x38, 0x7f, 0x8c, 0xc5, 0xda, 0xc4, 0x8a,
	0xd9, 0xe4, 0xcc, 0xcd, 0x67, 0x36, 0x39, 0x73, 0xc9, 0xd9, 0xce, 0x2c, 0x79, 0x60, 0x55, 0x81,
	0x53, 0x93, 0x3b, 0xb8, 0xae, 0x72, 0x93, 0xe6, 0xc6, 0x19, 0xbe, 0x3e, 0xcf, 0xf9, 0xd2, 0xa2,
	0xb3, 0x94, 0x8e, 0xe0, 0x6f, 0xe0, 0xc6, 0x29, 0x48, 0x9a, 0xf8, 0xd2, 0x48, 0x24, 0x39, 0x58,
	0x59, 0x23, 0xa8, 0x0a, 0xff, 0x01, 0x5e, 0x3e, 0xf3, 0x3e, 0xe4, 0x06, 0xae, 0x49, 0x31, 0x84,
	0xd6, 0xd7, 0xef, 0xdf, 0x7c, 0x4b, 0xb3, 0x25, 0xb9, 0x86, 0x17, 0x07, 0x50, 0x58, 0xe3, 0x8d,
	0x60, 0x76, 0xf4, 0xeb, 0xb8, 0xf6, 0xa6, 0x48, 0xc1, 0xdf, 0xc0, 0x2b, 0x73, 0xfd, 0x2b, 0x91,
	0x89, 0xe1, 0xbf, 0xd4, 0xed, 0xb7, 0xe7, 0x37, 0xf6, 0xba, 0xfa, 0x5f, 0xa4, 0x83, 0xaf, 0x9c,
	0x85, 0xc9, 0xcd, 0x79, 0x8e, 0x0b, 0xd6, 0xdb, 0xbe, 0x75, 0x71, 0xb3, 0x7a, 0xda, 0xc7, 0xcf,
	0x0e, 0x27, 0xd4, 0x1b, 0x4f, 0xa8, 0x77, 0x34, 0xa1, 0xde, 0xc9, 0x84, 0x7a, 0xef, 0x4b, 0x8a,
	0xbe, 0x94, 0xd4, 0x3b, 0x2c, 0x29, 0x1a, 0x97, 0x14, 0xfd, 0x2a, 0x29, 0xfa, 0x5d, 0x52, 0xef,
	0xa4, 0xa4, 0xe8, 0xd3, 0x94, 0x7a, 0xe3, 0x29, 0xf5, 0x8e, 0xa6, 0xd4, 0xfb, 0xf3, 0xe3, 0xf8,
	0x60, 0xa1, 0xf6, 0xe1, 0xe7, 0xf1, 0xc1, 0x82, 0xfb, 0xe8, 0xbd, 0xba, 0xdd, 0xfb, 0xfd, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0xf2, 0xfe, 0xcf, 0x2b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandleKeyvalServiceClient is the client API for HandleKeyvalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandleKeyvalServiceClient interface {
	// HandleKeyval is called by Mixer at request-time to deliver 'keyval' instances to the backend.
	HandleKeyval(ctx context.Context, in *HandleKeyvalRequest, opts ...grpc.CallOption) (*HandleKeyvalResponse, error)
}

type handleKeyvalServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleKeyvalServiceClient(cc *grpc.ClientConn) HandleKeyvalServiceClient {
	return &handleKeyvalServiceClient{cc}
}

func (c *handleKeyvalServiceClient) HandleKeyval(ctx context.Context, in *HandleKeyvalRequest, opts ...grpc.CallOption) (*HandleKeyvalResponse, error) {
	out := new(HandleKeyvalResponse)
	err := c.cc.Invoke(ctx, "/keyval.HandleKeyvalService/HandleKeyval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleKeyvalServiceServer is the server API for HandleKeyvalService service.
type HandleKeyvalServiceServer interface {
	// HandleKeyval is called by Mixer at request-time to deliver 'keyval' instances to the backend.
	HandleKeyval(context.Context, *HandleKeyvalRequest) (*HandleKeyvalResponse, error)
}

func RegisterHandleKeyvalServiceServer(s *grpc.Server, srv HandleKeyvalServiceServer) {
	s.RegisterService(&_HandleKeyvalService_serviceDesc, srv)
}

func _HandleKeyvalService_HandleKeyval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleKeyvalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleKeyvalServiceServer).HandleKeyval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyval.HandleKeyvalService/HandleKeyval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleKeyvalServiceServer).HandleKeyval(ctx, req.(*HandleKeyvalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleKeyvalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyval.HandleKeyvalService",
	HandlerType: (*HandleKeyvalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleKeyval",
			Handler:    _HandleKeyvalService_HandleKeyval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/test/keyval/template_handler_service.proto",
}

func (m *HandleKeyvalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleKeyvalRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Instance != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Instance.Size()))
		n1, err := m.Instance.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AdapterConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.AdapterConfig.Size()))
		n2, err := m.AdapterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.DedupId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i += copy(dAtA[i:], m.DedupId)
	}
	return i, nil
}

func (m *HandleKeyvalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleKeyvalResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Result.Size()))
		n3, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Output != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Output.Size()))
		n4, err := m.Output.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *OutputMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutputMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0xe4
		i++
		dAtA[i] = 0x93
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	return i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HandleKeyvalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Instance != nil {
		l = m.Instance.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *HandleKeyvalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.Output != nil {
		l = m.Output.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *OutputMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
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
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleKeyvalRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleKeyvalRequest{`,
		`Instance:` + strings.Replace(fmt.Sprintf("%v", this.Instance), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "types.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HandleKeyvalResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleKeyvalResponse{`,
		`Result:` + strings.Replace(fmt.Sprintf("%v", this.Result), "CheckResult", "v1beta1.CheckResult", 1) + `,`,
		`Output:` + strings.Replace(fmt.Sprintf("%v", this.Output), "OutputMsg", "OutputMsg", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OutputMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OutputMsg{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
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
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleKeyvalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: HandleKeyvalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleKeyvalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *HandleKeyvalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: HandleKeyvalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleKeyvalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &v1beta1.CheckResult{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Output == nil {
				m.Output = &OutputMsg{}
			}
			if err := m.Output.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func (m *OutputMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
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
			return fmt.Errorf("proto: OutputMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutputMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
				return ErrIntOverflowTemplateHandlerService
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
				return ErrIntOverflowTemplateHandlerService
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
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
				return ErrIntOverflowTemplateHandlerService
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
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
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemplateHandlerService
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
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
					return 0, ErrIntOverflowTemplateHandlerService
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
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
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
				next, err := skipTemplateHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTemplateHandlerService
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
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)
