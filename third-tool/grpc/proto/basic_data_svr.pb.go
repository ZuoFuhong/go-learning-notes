// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: basic_data_svr.proto

package grpc_example_basic_data_svr

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

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_data_svr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_basic_data_svr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_basic_data_svr_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  uint32 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *GetUserInfoRsp) Reset() {
	*x = GetUserInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_data_svr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRsp) ProtoMessage() {}

func (x *GetUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_basic_data_svr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRsp.ProtoReflect.Descriptor instead.
func (*GetUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_basic_data_svr_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoRsp) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *GetUserInfoRsp) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_basic_data_svr_proto protoreflect.FileDescriptor

var file_basic_data_svr_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x76, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x32, 0x3f, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x76, 0x72, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x76, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_data_svr_proto_rawDescOnce sync.Once
	file_basic_data_svr_proto_rawDescData = file_basic_data_svr_proto_rawDesc
)

func file_basic_data_svr_proto_rawDescGZIP() []byte {
	file_basic_data_svr_proto_rawDescOnce.Do(func() {
		file_basic_data_svr_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_data_svr_proto_rawDescData)
	})
	return file_basic_data_svr_proto_rawDescData
}

var file_basic_data_svr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_basic_data_svr_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil), // 0: GetUserInfoReq
	(*GetUserInfoRsp)(nil), // 1: GetUserInfoRsp
}
var file_basic_data_svr_proto_depIdxs = []int32{
	0, // 0: BasicDataSvr.GetUserInfo:input_type -> GetUserInfoReq
	1, // 1: BasicDataSvr.GetUserInfo:output_type -> GetUserInfoRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basic_data_svr_proto_init() }
func file_basic_data_svr_proto_init() {
	if File_basic_data_svr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_data_svr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_basic_data_svr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRsp); i {
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
			RawDescriptor: file_basic_data_svr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_data_svr_proto_goTypes,
		DependencyIndexes: file_basic_data_svr_proto_depIdxs,
		MessageInfos:      file_basic_data_svr_proto_msgTypes,
	}.Build()
	File_basic_data_svr_proto = out.File
	file_basic_data_svr_proto_rawDesc = nil
	file_basic_data_svr_proto_goTypes = nil
	file_basic_data_svr_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BasicDataSvrClient is the client API for BasicDataSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasicDataSvrClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error)
}

type basicDataSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicDataSvrClient(cc grpc.ClientConnInterface) BasicDataSvrClient {
	return &basicDataSvrClient{cc}
}

func (c *basicDataSvrClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error) {
	out := new(GetUserInfoRsp)
	err := c.cc.Invoke(ctx, "/BasicDataSvr/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicDataSvrServer is the server API for BasicDataSvr service.
type BasicDataSvrServer interface {
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error)
}

// UnimplementedBasicDataSvrServer can be embedded to have forward compatible implementations.
type UnimplementedBasicDataSvrServer struct {
}

func (*UnimplementedBasicDataSvrServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterBasicDataSvrServer(s *grpc.Server, srv BasicDataSvrServer) {
	s.RegisterService(&_BasicDataSvr_serviceDesc, srv)
}

func _BasicDataSvr_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicDataSvrServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BasicDataSvr/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicDataSvrServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasicDataSvr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BasicDataSvr",
	HandlerType: (*BasicDataSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _BasicDataSvr_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic_data_svr.proto",
}
