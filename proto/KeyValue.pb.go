// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: KeyValue.proto

package KeyValue

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeyValue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_KeyValue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_KeyValue_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *KeyResponse) Reset() {
	*x = KeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeyValue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyResponse) ProtoMessage() {}

func (x *KeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_KeyValue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyResponse.ProtoReflect.Descriptor instead.
func (*KeyResponse) Descriptor() ([]byte, []int) {
	return file_KeyValue_proto_rawDescGZIP(), []int{1}
}

func (x *KeyResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KeyValue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_KeyValue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_KeyValue_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_KeyValue_proto protoreflect.FileDescriptor

var file_KeyValue_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x23, 0x0a, 0x0b, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x6f, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x31, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x12, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KeyValue_proto_rawDescOnce sync.Once
	file_KeyValue_proto_rawDescData = file_KeyValue_proto_rawDesc
)

func file_KeyValue_proto_rawDescGZIP() []byte {
	file_KeyValue_proto_rawDescOnce.Do(func() {
		file_KeyValue_proto_rawDescData = protoimpl.X.CompressGZIP(file_KeyValue_proto_rawDescData)
	})
	return file_KeyValue_proto_rawDescData
}

var file_KeyValue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_KeyValue_proto_goTypes = []interface{}{
	(*KeyRequest)(nil),  // 0: KeyValue.KeyRequest
	(*KeyResponse)(nil), // 1: KeyValue.KeyResponse
	(*KeyValue)(nil),    // 2: KeyValue.KeyValue
}
var file_KeyValue_proto_depIdxs = []int32{
	0, // 0: KeyValue.KV.Get:input_type -> KeyValue.KeyRequest
	2, // 1: KeyValue.KV.Add:input_type -> KeyValue.KeyValue
	1, // 2: KeyValue.KV.Get:output_type -> KeyValue.KeyResponse
	2, // 3: KeyValue.KV.Add:output_type -> KeyValue.KeyValue
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KeyValue_proto_init() }
func file_KeyValue_proto_init() {
	if File_KeyValue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KeyValue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_KeyValue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_KeyValue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_KeyValue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_KeyValue_proto_goTypes,
		DependencyIndexes: file_KeyValue_proto_depIdxs,
		MessageInfos:      file_KeyValue_proto_msgTypes,
	}.Build()
	File_KeyValue_proto = out.File
	file_KeyValue_proto_rawDesc = nil
	file_KeyValue_proto_goTypes = nil
	file_KeyValue_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (KV_GetClient, error)
	Add(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (KV_AddClient, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (KV_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KV_serviceDesc.Streams[0], "/KeyValue.KV/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_GetClient interface {
	Recv() (*KeyResponse, error)
	grpc.ClientStream
}

type kVGetClient struct {
	grpc.ClientStream
}

func (x *kVGetClient) Recv() (*KeyResponse, error) {
	m := new(KeyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) Add(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (KV_AddClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KV_serviceDesc.Streams[1], "/KeyValue.KV/Add", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVAddClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_AddClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type kVAddClient struct {
	grpc.ClientStream
}

func (x *kVAddClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	Get(*KeyRequest, KV_GetServer) error
	Add(*KeyValue, KV_AddServer) error
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) Get(*KeyRequest, KV_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVServer) Add(*KeyValue, KV_AddServer) error {
	return status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).Get(m, &kVGetServer{stream})
}

type KV_GetServer interface {
	Send(*KeyResponse) error
	grpc.ServerStream
}

type kVGetServer struct {
	grpc.ServerStream
}

func (x *kVGetServer) Send(m *KeyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_Add_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).Add(m, &kVAddServer{stream})
}

type KV_AddServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type kVAddServer struct {
	grpc.ServerStream
}

func (x *kVAddServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KeyValue.KV",
	HandlerType: (*KVServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _KV_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Add",
			Handler:       _KV_Add_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "KeyValue.proto",
}
