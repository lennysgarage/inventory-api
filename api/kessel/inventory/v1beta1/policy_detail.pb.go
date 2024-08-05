// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: kessel/inventory/v1beta1/policy_detail.proto

package v1beta1

import (
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

// The kind of policy
type PolicyDetail_SeverityEnum int32

const (
	PolicyDetail_SeverityEnum_LOW      PolicyDetail_SeverityEnum = 0
	PolicyDetail_SeverityEnum_MEDIUM   PolicyDetail_SeverityEnum = 1
	PolicyDetail_SeverityEnum_HIGH     PolicyDetail_SeverityEnum = 2
	PolicyDetail_SeverityEnum_CRITICAL PolicyDetail_SeverityEnum = 3
	PolicyDetail_SeverityEnum_OTHER    PolicyDetail_SeverityEnum = 4
)

// Enum value maps for PolicyDetail_SeverityEnum.
var (
	PolicyDetail_SeverityEnum_name = map[int32]string{
		0: "SeverityEnum_LOW",
		1: "SeverityEnum_MEDIUM",
		2: "SeverityEnum_HIGH",
		3: "SeverityEnum_CRITICAL",
		4: "SeverityEnum_OTHER",
	}
	PolicyDetail_SeverityEnum_value = map[string]int32{
		"SeverityEnum_LOW":      0,
		"SeverityEnum_MEDIUM":   1,
		"SeverityEnum_HIGH":     2,
		"SeverityEnum_CRITICAL": 3,
		"SeverityEnum_OTHER":    4,
	}
)

func (x PolicyDetail_SeverityEnum) Enum() *PolicyDetail_SeverityEnum {
	p := new(PolicyDetail_SeverityEnum)
	*p = x
	return p
}

func (x PolicyDetail_SeverityEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyDetail_SeverityEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_kessel_inventory_v1beta1_policy_detail_proto_enumTypes[0].Descriptor()
}

func (PolicyDetail_SeverityEnum) Type() protoreflect.EnumType {
	return &file_kessel_inventory_v1beta1_policy_detail_proto_enumTypes[0]
}

func (x PolicyDetail_SeverityEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyDetail_SeverityEnum.Descriptor instead.
func (PolicyDetail_SeverityEnum) EnumDescriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_policy_detail_proto_rawDescGZIP(), []int{0, 0}
}

type PolicyDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines if the policy is currently enabled or disabled across all targets
	Disabled bool                      `protobuf:"varint,270940796,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Severity PolicyDetail_SeverityEnum `protobuf:"varint,404558591,opt,name=severity,proto3,enum=api.kessel.inventory.v1beta1.PolicyDetail_SeverityEnum" json:"severity,omitempty"`
}

func (x *PolicyDetail) Reset() {
	*x = PolicyDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_policy_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyDetail) ProtoMessage() {}

func (x *PolicyDetail) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_policy_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyDetail.ProtoReflect.Descriptor instead.
func (*PolicyDetail) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_policy_detail_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyDetail) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *PolicyDetail) GetSeverity() PolicyDetail_SeverityEnum {
	if x != nil {
		return x.Severity
	}
	return PolicyDetail_SeverityEnum_LOW
}

var File_kessel_inventory_v1beta1_policy_detail_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_policy_detail_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x91, 0x02, 0x0a,
	0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0xfc, 0xf4, 0x98, 0x81, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x57, 0x0a,
	0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0xff, 0xa5, 0xf4, 0xc0, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x4d, 0x45,
	0x44, 0x49, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x12, 0x19, 0x0a,
	0x15, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x43, 0x52,
	0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x04,
	0x42, 0x46, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x50, 0x01, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_policy_detail_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_policy_detail_proto_rawDescData = file_kessel_inventory_v1beta1_policy_detail_proto_rawDesc
)

func file_kessel_inventory_v1beta1_policy_detail_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_policy_detail_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_policy_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_policy_detail_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_policy_detail_proto_rawDescData
}

var file_kessel_inventory_v1beta1_policy_detail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kessel_inventory_v1beta1_policy_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_policy_detail_proto_goTypes = []any{
	(PolicyDetail_SeverityEnum)(0), // 0: api.kessel.inventory.v1beta1.PolicyDetail.SeverityEnum
	(*PolicyDetail)(nil),           // 1: api.kessel.inventory.v1beta1.PolicyDetail
}
var file_kessel_inventory_v1beta1_policy_detail_proto_depIdxs = []int32{
	0, // 0: api.kessel.inventory.v1beta1.PolicyDetail.severity:type_name -> api.kessel.inventory.v1beta1.PolicyDetail.SeverityEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_policy_detail_proto_init() }
func file_kessel_inventory_v1beta1_policy_detail_proto_init() {
	if File_kessel_inventory_v1beta1_policy_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_policy_detail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PolicyDetail); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_policy_detail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_policy_detail_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_policy_detail_proto_depIdxs,
		EnumInfos:         file_kessel_inventory_v1beta1_policy_detail_proto_enumTypes,
		MessageInfos:      file_kessel_inventory_v1beta1_policy_detail_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_policy_detail_proto = out.File
	file_kessel_inventory_v1beta1_policy_detail_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_policy_detail_proto_goTypes = nil
	file_kessel_inventory_v1beta1_policy_detail_proto_depIdxs = nil
}
