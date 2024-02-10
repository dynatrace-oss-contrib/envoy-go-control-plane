// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/compression/gzip/decompressor/v3/gzip.proto

package decompressorv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Gzip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value from 9 to 15 that represents the base two logarithmic of the decompressor's window size.
	// The decompression window size needs to be equal or larger than the compression window size.
	// The default window size is 15.
	// This is so that the decompressor can decompress a response compressed by a compressor with any compression window size.
	// For more details about this parameter, please refer to `zlib manual <https://www.zlib.net/manual.html>`_ > inflateInit2.
	WindowBits *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Value for zlib's decompressor output buffer. If not set, defaults to 4096.
	// See https://www.zlib.net/manual.html for more details.
	ChunkSize *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	// An upper bound to the number of times the output buffer is allowed to be bigger than the size of
	// the accumulated input. This value is used to prevent decompression bombs. If not set, defaults to 100.
	// [#comment:TODO(rojkov): Re-design the Decompressor interface to handle compression bombs gracefully instead of this quick solution.
	// See https://github.com/envoyproxy/envoy/commit/d4c39e635603e2f23e1e08ddecf5a5fb5a706338 for details.]
	MaxInflateRatio *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_inflate_ratio,json=maxInflateRatio,proto3" json:"max_inflate_ratio,omitempty"`
}

func (x *Gzip) Reset() {
	*x = Gzip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gzip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gzip) ProtoMessage() {}

func (x *Gzip) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gzip.ProtoReflect.Descriptor instead.
func (*Gzip) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescGZIP(), []int{0}
}

func (x *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if x != nil {
		return x.WindowBits
	}
	return nil
}

func (x *Gzip) GetChunkSize() *wrappers.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

func (x *Gzip) GetMaxInflateRatio() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxInflateRatio
	}
	return nil
}

var File_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto protoreflect.FileDescriptor

var file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x7a, 0x69, 0x70, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x2f, 0x76, 0x33, 0x2f, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69,
	0x70, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x04, 0x47, 0x7a,
	0x69, 0x70, 0x12, 0x48, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x62, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x0f, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x69, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x0a,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c,
	0xfa, 0x42, 0x09, 0x2a, 0x07, 0x18, 0x80, 0x80, 0x04, 0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x69,
	0x6e, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x2a, 0x05, 0x18, 0x88, 0x08, 0x28, 0x01, 0x52, 0x0f, 0x6d, 0x61,
	0x78, 0x49, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0xbf, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x09, 0x47, 0x7a, 0x69, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x7a, 0x69, 0x70,
	0x2f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33,
	0x3b, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescOnce sync.Once
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescData = file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDesc
)

func file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescGZIP() []byte {
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescData)
	})
	return file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDescData
}

var file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_goTypes = []interface{}{
	(*Gzip)(nil),                 // 0: envoy.extensions.compression.gzip.decompressor.v3.Gzip
	(*wrappers.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
}
var file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.compression.gzip.decompressor.v3.Gzip.window_bits:type_name -> google.protobuf.UInt32Value
	1, // 1: envoy.extensions.compression.gzip.decompressor.v3.Gzip.chunk_size:type_name -> google.protobuf.UInt32Value
	1, // 2: envoy.extensions.compression.gzip.decompressor.v3.Gzip.max_inflate_ratio:type_name -> google.protobuf.UInt32Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_init() }
func file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_init() {
	if File_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gzip); i {
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
			RawDescriptor: file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_msgTypes,
	}.Build()
	File_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto = out.File
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_rawDesc = nil
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_goTypes = nil
	file_envoy_extensions_compression_gzip_decompressor_v3_gzip_proto_depIdxs = nil
}
