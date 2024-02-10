// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/admin/v3/listeners.proto

package adminv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

// Admin endpoint uses this wrapper for “/listeners“ to display listener status information.
// See :ref:`/listeners <operations_admin_interface_listeners>` for more information.
type Listeners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of listener statuses.
	ListenerStatuses []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
}

func (x *Listeners) Reset() {
	*x = Listeners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_listeners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listeners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listeners) ProtoMessage() {}

func (x *Listeners) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_listeners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listeners.ProtoReflect.Descriptor instead.
func (*Listeners) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_listeners_proto_rawDescGZIP(), []int{0}
}

func (x *Listeners) GetListenerStatuses() []*ListenerStatus {
	if x != nil {
		return x.ListenerStatuses
	}
	return nil
}

// Details an individual listener's current status.
type ListenerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the listener
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual local address that the listener is listening on. If a listener was configured
	// to listen on port 0, then this address has the port that was allocated by the OS.
	LocalAddress *v3.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	// The additional addresses the listener is listening on as specified via the :ref:`additional_addresses <envoy_v3_api_field_config.listener.v3.Listener.additional_addresses>`
	// configuration.
	AdditionalLocalAddresses []*v3.Address `protobuf:"bytes,3,rep,name=additional_local_addresses,json=additionalLocalAddresses,proto3" json:"additional_local_addresses,omitempty"`
}

func (x *ListenerStatus) Reset() {
	*x = ListenerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_listeners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerStatus) ProtoMessage() {}

func (x *ListenerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_listeners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerStatus.ProtoReflect.Descriptor instead.
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_listeners_proto_rawDescGZIP(), []int{1}
}

func (x *ListenerStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListenerStatus) GetLocalAddress() *v3.Address {
	if x != nil {
		return x.LocalAddress
	}
	return nil
}

func (x *ListenerStatus) GetAdditionalLocalAddresses() []*v3.Address {
	if x != nil {
		return x.AdditionalLocalAddresses
	}
	return nil
}

var File_envoy_admin_v3_listeners_proto protoreflect.FileDescriptor

var file_envoy_admin_v3_listeners_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33,
	0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x3a, 0x24, 0x9a, 0xc5, 0x88, 0x1e, 0x1f, 0x0a, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a,
	0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x5b, 0x0a, 0x1a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x3a, 0x29,
	0x9a, 0xc5, 0x88, 0x1e, 0x24, 0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x77, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x33, 0x42, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v3_listeners_proto_rawDescOnce sync.Once
	file_envoy_admin_v3_listeners_proto_rawDescData = file_envoy_admin_v3_listeners_proto_rawDesc
)

func file_envoy_admin_v3_listeners_proto_rawDescGZIP() []byte {
	file_envoy_admin_v3_listeners_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v3_listeners_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v3_listeners_proto_rawDescData)
	})
	return file_envoy_admin_v3_listeners_proto_rawDescData
}

var file_envoy_admin_v3_listeners_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_admin_v3_listeners_proto_goTypes = []interface{}{
	(*Listeners)(nil),      // 0: envoy.admin.v3.Listeners
	(*ListenerStatus)(nil), // 1: envoy.admin.v3.ListenerStatus
	(*v3.Address)(nil),     // 2: envoy.config.core.v3.Address
}
var file_envoy_admin_v3_listeners_proto_depIdxs = []int32{
	1, // 0: envoy.admin.v3.Listeners.listener_statuses:type_name -> envoy.admin.v3.ListenerStatus
	2, // 1: envoy.admin.v3.ListenerStatus.local_address:type_name -> envoy.config.core.v3.Address
	2, // 2: envoy.admin.v3.ListenerStatus.additional_local_addresses:type_name -> envoy.config.core.v3.Address
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_admin_v3_listeners_proto_init() }
func file_envoy_admin_v3_listeners_proto_init() {
	if File_envoy_admin_v3_listeners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v3_listeners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listeners); i {
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
		file_envoy_admin_v3_listeners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerStatus); i {
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
			RawDescriptor: file_envoy_admin_v3_listeners_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v3_listeners_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v3_listeners_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v3_listeners_proto_msgTypes,
	}.Build()
	File_envoy_admin_v3_listeners_proto = out.File
	file_envoy_admin_v3_listeners_proto_rawDesc = nil
	file_envoy_admin_v3_listeners_proto_goTypes = nil
	file_envoy_admin_v3_listeners_proto_depIdxs = nil
}
