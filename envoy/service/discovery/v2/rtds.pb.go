// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/service/discovery/v2/rtds.proto

package discoveryv2

import (
	context "context"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type RtdsDummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RtdsDummy) Reset() {
	*x = RtdsDummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v2_rtds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtdsDummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtdsDummy) ProtoMessage() {}

func (x *RtdsDummy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v2_rtds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtdsDummy.ProtoReflect.Descriptor instead.
func (*RtdsDummy) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v2_rtds_proto_rawDescGZIP(), []int{0}
}

// RTDS resource type. This describes a layer in the runtime virtual filesystem.
type Runtime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Runtime resource name. This makes the Runtime a self-describing xDS
	// resource.
	Name  string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layer *_struct.Struct `protobuf:"bytes,2,opt,name=layer,proto3" json:"layer,omitempty"`
}

func (x *Runtime) Reset() {
	*x = Runtime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v2_rtds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runtime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runtime) ProtoMessage() {}

func (x *Runtime) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v2_rtds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runtime.ProtoReflect.Descriptor instead.
func (*Runtime) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v2_rtds_proto_rawDescGZIP(), []int{1}
}

func (x *Runtime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Runtime) GetLayer() *_struct.Struct {
	if x != nil {
		return x.Layer
	}
	return nil
}

var File_envoy_service_discovery_v2_rtds_proto protoreflect.FileDescriptor

var file_envoy_service_discovery_v2_rtds_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x74, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x74, 0x64, 0x73, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x22, 0x55, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0xf1, 0x02, 0x0a, 0x17,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x5f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x71, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76,
	0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x1a, 0x2a, 0x8a, 0xa4, 0x96, 0xf3, 0x07, 0x24, 0x0a, 0x22, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0xae, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x1a, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x52, 0x74, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_discovery_v2_rtds_proto_rawDescOnce sync.Once
	file_envoy_service_discovery_v2_rtds_proto_rawDescData = file_envoy_service_discovery_v2_rtds_proto_rawDesc
)

func file_envoy_service_discovery_v2_rtds_proto_rawDescGZIP() []byte {
	file_envoy_service_discovery_v2_rtds_proto_rawDescOnce.Do(func() {
		file_envoy_service_discovery_v2_rtds_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_discovery_v2_rtds_proto_rawDescData)
	})
	return file_envoy_service_discovery_v2_rtds_proto_rawDescData
}

var file_envoy_service_discovery_v2_rtds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_service_discovery_v2_rtds_proto_goTypes = []interface{}{
	(*RtdsDummy)(nil),                 // 0: envoy.service.discovery.v2.RtdsDummy
	(*Runtime)(nil),                   // 1: envoy.service.discovery.v2.Runtime
	(*_struct.Struct)(nil),            // 2: google.protobuf.Struct
	(*v2.DiscoveryRequest)(nil),       // 3: envoy.api.v2.DiscoveryRequest
	(*v2.DeltaDiscoveryRequest)(nil),  // 4: envoy.api.v2.DeltaDiscoveryRequest
	(*v2.DiscoveryResponse)(nil),      // 5: envoy.api.v2.DiscoveryResponse
	(*v2.DeltaDiscoveryResponse)(nil), // 6: envoy.api.v2.DeltaDiscoveryResponse
}
var file_envoy_service_discovery_v2_rtds_proto_depIdxs = []int32{
	2, // 0: envoy.service.discovery.v2.Runtime.layer:type_name -> google.protobuf.Struct
	3, // 1: envoy.service.discovery.v2.RuntimeDiscoveryService.StreamRuntime:input_type -> envoy.api.v2.DiscoveryRequest
	4, // 2: envoy.service.discovery.v2.RuntimeDiscoveryService.DeltaRuntime:input_type -> envoy.api.v2.DeltaDiscoveryRequest
	3, // 3: envoy.service.discovery.v2.RuntimeDiscoveryService.FetchRuntime:input_type -> envoy.api.v2.DiscoveryRequest
	5, // 4: envoy.service.discovery.v2.RuntimeDiscoveryService.StreamRuntime:output_type -> envoy.api.v2.DiscoveryResponse
	6, // 5: envoy.service.discovery.v2.RuntimeDiscoveryService.DeltaRuntime:output_type -> envoy.api.v2.DeltaDiscoveryResponse
	5, // 6: envoy.service.discovery.v2.RuntimeDiscoveryService.FetchRuntime:output_type -> envoy.api.v2.DiscoveryResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_service_discovery_v2_rtds_proto_init() }
func file_envoy_service_discovery_v2_rtds_proto_init() {
	if File_envoy_service_discovery_v2_rtds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_discovery_v2_rtds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RtdsDummy); i {
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
		file_envoy_service_discovery_v2_rtds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runtime); i {
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
			RawDescriptor: file_envoy_service_discovery_v2_rtds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_discovery_v2_rtds_proto_goTypes,
		DependencyIndexes: file_envoy_service_discovery_v2_rtds_proto_depIdxs,
		MessageInfos:      file_envoy_service_discovery_v2_rtds_proto_msgTypes,
	}.Build()
	File_envoy_service_discovery_v2_rtds_proto = out.File
	file_envoy_service_discovery_v2_rtds_proto_rawDesc = nil
	file_envoy_service_discovery_v2_rtds_proto_goTypes = nil
	file_envoy_service_discovery_v2_rtds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RuntimeDiscoveryServiceClient is the client API for RuntimeDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeDiscoveryServiceClient(cc grpc.ClientConnInterface) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.RuntimeDiscoveryService/StreamRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.RuntimeDiscoveryService/DeltaRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeDiscoveryServiceServer is the server API for RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedRuntimeDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeDiscoveryServiceServer struct {
}

func (*UnimplementedRuntimeDiscoveryServiceServer) StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) FetchRuntime(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRuntime not implemented")
}

func RegisterRuntimeDiscoveryServiceServer(s *grpc.Server, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&_RuntimeDiscoveryService_serviceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/rtds.proto",
}
