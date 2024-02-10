// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: envoy/extensions/load_balancing_policies/least_request/v3/least_request.proto

package least_requestv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/load_balancing_policies/common/v3"
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

// Available methods for selecting the host set from which to return the host with the
// fewest active requests.
type LeastRequest_SelectionMethod int32

const (
	// Return host with fewest requests from a set of “choice_count“ randomly selected hosts.
	// Best selection method for most scenarios.
	LeastRequest_N_CHOICES LeastRequest_SelectionMethod = 0
	// Return host with fewest requests from all hosts.
	// Useful in some niche use cases involving low request rates and one of:
	// (example 1) low request limits on workloads, or (example 2) few hosts.
	//
	// Example 1: Consider a workload type that can only accept one connection at a time.
	// If such workloads are deployed across many hosts, only a small percentage of those
	// workloads have zero connections at any given time, and the rate of new connections is low,
	// the “FULL_SCAN“ method is more likely to select a suitable host than “N_CHOICES“.
	//
	// Example 2: Consider a workload type that is only deployed on 2 hosts. With default settings,
	// the “N_CHOICES“ method will return the host with more active requests 25% of the time.
	// If the request rate is sufficiently low, the behavior of always selecting the host with least
	// requests as of the last metrics refresh may be preferable.
	LeastRequest_FULL_SCAN LeastRequest_SelectionMethod = 1
)

// Enum value maps for LeastRequest_SelectionMethod.
var (
	LeastRequest_SelectionMethod_name = map[int32]string{
		0: "N_CHOICES",
		1: "FULL_SCAN",
	}
	LeastRequest_SelectionMethod_value = map[string]int32{
		"N_CHOICES": 0,
		"FULL_SCAN": 1,
	}
)

func (x LeastRequest_SelectionMethod) Enum() *LeastRequest_SelectionMethod {
	p := new(LeastRequest_SelectionMethod)
	*p = x
	return p
}

func (x LeastRequest_SelectionMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LeastRequest_SelectionMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_enumTypes[0].Descriptor()
}

func (LeastRequest_SelectionMethod) Type() protoreflect.EnumType {
	return &file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_enumTypes[0]
}

func (x LeastRequest_SelectionMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LeastRequest_SelectionMethod.Descriptor instead.
func (LeastRequest_SelectionMethod) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescGZIP(), []int{0, 0}
}

// This configuration allows the built-in LEAST_REQUEST LB policy to be configured via the LB policy
// extension point. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` for more information.
// [#next-free-field: 7]
type LeastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of random healthy hosts from which the host with the fewest active requests will
	// be chosen. Defaults to 2 so that we perform two-choice selection if the field is not set.
	// Only applies to the “N_CHOICES“ selection method.
	ChoiceCount *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	// The following formula is used to calculate the dynamic weights when hosts have different load
	// balancing weights:
	//
	// “weight = load_balancing_weight / (active_requests + 1)^active_request_bias“
	//
	// The larger the active request bias is, the more aggressively active requests will lower the
	// effective weight when all host weights are not equal.
	//
	// “active_request_bias“ must be greater than or equal to 0.0.
	//
	// When “active_request_bias == 0.0“ the Least Request Load Balancer doesn't consider the number
	// of active requests at the time it picks a host and behaves like the Round Robin Load
	// Balancer.
	//
	// When “active_request_bias > 0.0“ the Least Request Load Balancer scales the load balancing
	// weight by the number of active requests at the time it does a pick.
	//
	// The value is cached for performance reasons and refreshed whenever one of the Load Balancer's
	// host sets changes, e.g., whenever there is a host membership update or a host load balancing
	// weight change.
	//
	// .. note::
	//
	//	This setting only takes effect if all host weights are not equal.
	ActiveRequestBias *v3.RuntimeDouble `protobuf:"bytes,2,opt,name=active_request_bias,json=activeRequestBias,proto3" json:"active_request_bias,omitempty"`
	// Configuration for slow start mode.
	// If this configuration is not set, slow start will not be not enabled.
	SlowStartConfig *v31.SlowStartConfig `protobuf:"bytes,3,opt,name=slow_start_config,json=slowStartConfig,proto3" json:"slow_start_config,omitempty"`
	// Configuration for local zone aware load balancing or locality weighted load balancing.
	LocalityLbConfig *v31.LocalityLbConfig `protobuf:"bytes,4,opt,name=locality_lb_config,json=localityLbConfig,proto3" json:"locality_lb_config,omitempty"`
	// [#not-implemented-hide:]
	// Unused. Replaced by the `selection_method` enum for better extensibility.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/load_balancing_policies/least_request/v3/least_request.proto.
	EnableFullScan *wrappers.BoolValue `protobuf:"bytes,5,opt,name=enable_full_scan,json=enableFullScan,proto3" json:"enable_full_scan,omitempty"`
	// Method for selecting the host set from which to return the host with the fewest active requests.
	//
	// Defaults to “N_CHOICES“.
	SelectionMethod LeastRequest_SelectionMethod `protobuf:"varint,6,opt,name=selection_method,json=selectionMethod,proto3,enum=envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest_SelectionMethod" json:"selection_method,omitempty"`
}

func (x *LeastRequest) Reset() {
	*x = LeastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeastRequest) ProtoMessage() {}

func (x *LeastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeastRequest.ProtoReflect.Descriptor instead.
func (*LeastRequest) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescGZIP(), []int{0}
}

func (x *LeastRequest) GetChoiceCount() *wrappers.UInt32Value {
	if x != nil {
		return x.ChoiceCount
	}
	return nil
}

func (x *LeastRequest) GetActiveRequestBias() *v3.RuntimeDouble {
	if x != nil {
		return x.ActiveRequestBias
	}
	return nil
}

func (x *LeastRequest) GetSlowStartConfig() *v31.SlowStartConfig {
	if x != nil {
		return x.SlowStartConfig
	}
	return nil
}

func (x *LeastRequest) GetLocalityLbConfig() *v31.LocalityLbConfig {
	if x != nil {
		return x.LocalityLbConfig
	}
	return nil
}

// Deprecated: Marked as deprecated in envoy/extensions/load_balancing_policies/least_request/v3/least_request.proto.
func (x *LeastRequest) GetEnableFullScan() *wrappers.BoolValue {
	if x != nil {
		return x.EnableFullScan
	}
	return nil
}

func (x *LeastRequest) GetSelectionMethod() LeastRequest_SelectionMethod {
	if x != nil {
		return x.SelectionMethod
	}
	return LeastRequest_N_CHOICES
}

var File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x05, 0x0a, 0x0c, 0x4c, 0x65,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x69, 0x61, 0x73, 0x12, 0x6f, 0x0a, 0x11, 0x73, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6c, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x73, 0x6c, 0x6f, 0x77, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x72, 0x0a, 0x12, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51,
	0x0a, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63,
	0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x18,
	0x01, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x61,
	0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x53,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10,
	0x01, 0x42, 0xd8, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x47, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x70, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData = file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes = []interface{}{
	(LeastRequest_SelectionMethod)(0), // 0: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.SelectionMethod
	(*LeastRequest)(nil),              // 1: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest
	(*wrappers.UInt32Value)(nil),      // 2: google.protobuf.UInt32Value
	(*v3.RuntimeDouble)(nil),          // 3: envoy.config.core.v3.RuntimeDouble
	(*v31.SlowStartConfig)(nil),       // 4: envoy.extensions.load_balancing_policies.common.v3.SlowStartConfig
	(*v31.LocalityLbConfig)(nil),      // 5: envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig
	(*wrappers.BoolValue)(nil),        // 6: google.protobuf.BoolValue
}
var file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.choice_count:type_name -> google.protobuf.UInt32Value
	3, // 1: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.active_request_bias:type_name -> envoy.config.core.v3.RuntimeDouble
	4, // 2: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.slow_start_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.SlowStartConfig
	5, // 3: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.locality_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig
	6, // 4: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.enable_full_scan:type_name -> google.protobuf.BoolValue
	0, // 5: envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.selection_method:type_name -> envoy.extensions.load_balancing_policies.least_request.v3.LeastRequest.SelectionMethod
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_init()
}
func file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_init() {
	if File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeastRequest); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto = out.File
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_least_request_v3_least_request_proto_depIdxs = nil
}
