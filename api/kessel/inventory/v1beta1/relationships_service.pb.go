// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: kessel/inventory/v1beta1/relationships_service.proto

package v1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatePolicyRelationshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource relationship to create in Kessel Asset Inventory
	PolicyRelationship *PolicyRelationship `protobuf:"bytes,1,opt,name=policyRelationship,proto3" json:"policyRelationship,omitempty"`
}

func (x *CreatePolicyRelationshipRequest) Reset() {
	*x = CreatePolicyRelationshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRelationshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRelationshipRequest) ProtoMessage() {}

func (x *CreatePolicyRelationshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRelationshipRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRelationshipRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyRelationshipRequest) GetPolicyRelationship() *PolicyRelationship {
	if x != nil {
		return x.PolicyRelationship
	}
	return nil
}

type CreatePolicyRelationshipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource relationship created in Kessel Asset Inventory
	PolicyRelationship *PolicyRelationship `protobuf:"bytes,1,opt,name=policyRelationship,proto3" json:"policyRelationship,omitempty"`
}

func (x *CreatePolicyRelationshipResponse) Reset() {
	*x = CreatePolicyRelationshipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRelationshipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRelationshipResponse) ProtoMessage() {}

func (x *CreatePolicyRelationshipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRelationshipResponse.ProtoReflect.Descriptor instead.
func (*CreatePolicyRelationshipResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyRelationshipResponse) GetPolicyRelationship() *PolicyRelationship {
	if x != nil {
		return x.PolicyRelationship
	}
	return nil
}

type UpdateResourceRelationshipByUrnHsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subject {resource} and the object {resource}. A relationship is between
	// a Subject and an Object, e.g. a Policy (subject) has a propagated
	// relationship to an k8s cluster (object)
	Resources *UpdateResourceRelationshipByUrnHsResourcesParameter `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
	// Update a resource relationship in Kessel Asset Inventory by the {subject
	// resource, object resource}. A relationship is between a Subject and an
	// Object, e.g. a Policy (subject) has a propagated relationship to an k8s
	// cluster (object).. The {resource} format
	// \"<reporter_data.reporter_type>:<reporter_data.reporter_id>::<reporter_data.resourceId_alias>\".
	PolicyRelationship *PolicyRelationship `protobuf:"bytes,2,opt,name=policyRelationship,proto3" json:"policyRelationship,omitempty"`
}

func (x *UpdateResourceRelationshipByUrnHsRequest) Reset() {
	*x = UpdateResourceRelationshipByUrnHsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRelationshipByUrnHsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRelationshipByUrnHsRequest) ProtoMessage() {}

func (x *UpdateResourceRelationshipByUrnHsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRelationshipByUrnHsRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceRelationshipByUrnHsRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateResourceRelationshipByUrnHsRequest) GetResources() *UpdateResourceRelationshipByUrnHsResourcesParameter {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *UpdateResourceRelationshipByUrnHsRequest) GetPolicyRelationship() *PolicyRelationship {
	if x != nil {
		return x.PolicyRelationship
	}
	return nil
}

type UpdateResourceRelationshipByUrnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResourceRelationshipByUrnResponse) Reset() {
	*x = UpdateResourceRelationshipByUrnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceRelationshipByUrnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceRelationshipByUrnResponse) ProtoMessage() {}

func (x *UpdateResourceRelationshipByUrnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceRelationshipByUrnResponse.ProtoReflect.Descriptor instead.
func (*UpdateResourceRelationshipByUrnResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{3}
}

type DeleteResourceRelationshipByUrnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subject {resource} and the object {resource}. A relationship is between
	// a Subject and an Object, e.g. a Policy (subject) has a propagated
	// relationship to an k8s cluster (object)
	Resources *UpdateResourceRelationshipByUrnHsResourcesParameter `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *DeleteResourceRelationshipByUrnRequest) Reset() {
	*x = DeleteResourceRelationshipByUrnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRelationshipByUrnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRelationshipByUrnRequest) ProtoMessage() {}

func (x *DeleteResourceRelationshipByUrnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRelationshipByUrnRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceRelationshipByUrnRequest) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResourceRelationshipByUrnRequest) GetResources() *UpdateResourceRelationshipByUrnHsResourcesParameter {
	if x != nil {
		return x.Resources
	}
	return nil
}

type DeleteResourceRelationshipByUrnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResourceRelationshipByUrnResponse) Reset() {
	*x = DeleteResourceRelationshipByUrnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceRelationshipByUrnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceRelationshipByUrnResponse) ProtoMessage() {}

func (x *DeleteResourceRelationshipByUrnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceRelationshipByUrnResponse.ProtoReflect.Descriptor instead.
func (*DeleteResourceRelationshipByUrnResponse) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP(), []int{5}
}

var File_kessel_inventory_v1beta1_relationships_service_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_relationships_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x62, 0x79,
	0x5f, 0x75, 0x72, 0x6e, 0x68, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0x84, 0x01, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x12, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0xfd, 0x01,
	0x0a, 0x28, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72,
	0x6e, 0x48, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x48, 0x73, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x12, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x22, 0x29, 0x0a,
	0x27, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x26, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x42, 0x79, 0x55, 0x72, 0x6e, 0x48, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x27, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xdf, 0x05, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe2, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x12, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x22, 0x2b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0xfb, 0x01,
	0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72,
	0x6e, 0x48, 0x73, 0x12, 0x46, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55,
	0x72, 0x6e, 0x48, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a, 0x12, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x2b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0xe3, 0x01, 0x0a, 0x1f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x12,
	0x44, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42, 0x79, 0x55, 0x72, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x42,
	0x79, 0x55, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2d, 0x2a, 0x2b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x42, 0x46, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x50, 0x01, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x50, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_relationships_service_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_relationships_service_proto_rawDescData = file_kessel_inventory_v1beta1_relationships_service_proto_rawDesc
)

func file_kessel_inventory_v1beta1_relationships_service_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_relationships_service_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_relationships_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_relationships_service_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_relationships_service_proto_rawDescData
}

var file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kessel_inventory_v1beta1_relationships_service_proto_goTypes = []any{
	(*CreatePolicyRelationshipRequest)(nil),                     // 0: api.kessel.inventory.v1beta1.CreatePolicyRelationshipRequest
	(*CreatePolicyRelationshipResponse)(nil),                    // 1: api.kessel.inventory.v1beta1.CreatePolicyRelationshipResponse
	(*UpdateResourceRelationshipByUrnHsRequest)(nil),            // 2: api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsRequest
	(*UpdateResourceRelationshipByUrnResponse)(nil),             // 3: api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnResponse
	(*DeleteResourceRelationshipByUrnRequest)(nil),              // 4: api.kessel.inventory.v1beta1.DeleteResourceRelationshipByUrnRequest
	(*DeleteResourceRelationshipByUrnResponse)(nil),             // 5: api.kessel.inventory.v1beta1.DeleteResourceRelationshipByUrnResponse
	(*PolicyRelationship)(nil),                                  // 6: api.kessel.inventory.v1beta1.PolicyRelationship
	(*UpdateResourceRelationshipByUrnHsResourcesParameter)(nil), // 7: api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsResourcesParameter
}
var file_kessel_inventory_v1beta1_relationships_service_proto_depIdxs = []int32{
	6, // 0: api.kessel.inventory.v1beta1.CreatePolicyRelationshipRequest.policyRelationship:type_name -> api.kessel.inventory.v1beta1.PolicyRelationship
	6, // 1: api.kessel.inventory.v1beta1.CreatePolicyRelationshipResponse.policyRelationship:type_name -> api.kessel.inventory.v1beta1.PolicyRelationship
	7, // 2: api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsRequest.resources:type_name -> api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsResourcesParameter
	6, // 3: api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsRequest.policyRelationship:type_name -> api.kessel.inventory.v1beta1.PolicyRelationship
	7, // 4: api.kessel.inventory.v1beta1.DeleteResourceRelationshipByUrnRequest.resources:type_name -> api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsResourcesParameter
	0, // 5: api.kessel.inventory.v1beta1.RelationshipsService.CreatePolicyRelationship:input_type -> api.kessel.inventory.v1beta1.CreatePolicyRelationshipRequest
	2, // 6: api.kessel.inventory.v1beta1.RelationshipsService.UpdateResourceRelationshipByUrnHs:input_type -> api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnHsRequest
	4, // 7: api.kessel.inventory.v1beta1.RelationshipsService.DeleteResourceRelationshipByUrn:input_type -> api.kessel.inventory.v1beta1.DeleteResourceRelationshipByUrnRequest
	1, // 8: api.kessel.inventory.v1beta1.RelationshipsService.CreatePolicyRelationship:output_type -> api.kessel.inventory.v1beta1.CreatePolicyRelationshipResponse
	3, // 9: api.kessel.inventory.v1beta1.RelationshipsService.UpdateResourceRelationshipByUrnHs:output_type -> api.kessel.inventory.v1beta1.UpdateResourceRelationshipByUrnResponse
	5, // 10: api.kessel.inventory.v1beta1.RelationshipsService.DeleteResourceRelationshipByUrn:output_type -> api.kessel.inventory.v1beta1.DeleteResourceRelationshipByUrnResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_relationships_service_proto_init() }
func file_kessel_inventory_v1beta1_relationships_service_proto_init() {
	if File_kessel_inventory_v1beta1_relationships_service_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_policy_relationship_proto_init()
	file_kessel_inventory_v1beta1_update_resource_relationship_by_urnhs_resources_parameter_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePolicyRelationshipRequest); i {
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
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePolicyRelationshipResponse); i {
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
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResourceRelationshipByUrnHsRequest); i {
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
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResourceRelationshipByUrnResponse); i {
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
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResourceRelationshipByUrnRequest); i {
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
		file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteResourceRelationshipByUrnResponse); i {
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
			RawDescriptor: file_kessel_inventory_v1beta1_relationships_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kessel_inventory_v1beta1_relationships_service_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_relationships_service_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_relationships_service_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_relationships_service_proto = out.File
	file_kessel_inventory_v1beta1_relationships_service_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_relationships_service_proto_goTypes = nil
	file_kessel_inventory_v1beta1_relationships_service_proto_depIdxs = nil
}
