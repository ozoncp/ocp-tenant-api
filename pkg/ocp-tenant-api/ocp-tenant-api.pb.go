// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/ocp-tenant-api/ocp-tenant-api.proto

package ocp_tenant_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Tenant description
type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type uint32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CreateTenantV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type uint32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *CreateTenantV1Request) Reset() {
	*x = CreateTenantV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantV1Request) ProtoMessage() {}

func (x *CreateTenantV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantV1Request.ProtoReflect.Descriptor instead.
func (*CreateTenantV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTenantV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTenantV1Request) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CreateTenantV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId uint64 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *CreateTenantV1Response) Reset() {
	*x = CreateTenantV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantV1Response) ProtoMessage() {}

func (x *CreateTenantV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantV1Response.ProtoReflect.Descriptor instead.
func (*CreateTenantV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTenantV1Response) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

type DescribeTenantV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId uint64 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *DescribeTenantV1Request) Reset() {
	*x = DescribeTenantV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTenantV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTenantV1Request) ProtoMessage() {}

func (x *DescribeTenantV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTenantV1Request.ProtoReflect.Descriptor instead.
func (*DescribeTenantV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeTenantV1Request) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

type DescribeTenantV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *DescribeTenantV1Response) Reset() {
	*x = DescribeTenantV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTenantV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTenantV1Response) ProtoMessage() {}

func (x *DescribeTenantV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTenantV1Response.ProtoReflect.Descriptor instead.
func (*DescribeTenantV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTenantV1Response) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type UpdateTenantV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *UpdateTenantV1Request) Reset() {
	*x = UpdateTenantV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantV1Request) ProtoMessage() {}

func (x *UpdateTenantV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantV1Request.ProtoReflect.Descriptor instead.
func (*UpdateTenantV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTenantV1Request) GetTenant() *Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type UpdateTenantV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated bool `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateTenantV1Response) Reset() {
	*x = UpdateTenantV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantV1Response) ProtoMessage() {}

func (x *UpdateTenantV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantV1Response.ProtoReflect.Descriptor instead.
func (*UpdateTenantV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTenantV1Response) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

type ListTenantsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListTenantsV1Request) Reset() {
	*x = ListTenantsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTenantsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsV1Request) ProtoMessage() {}

func (x *ListTenantsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsV1Request.ProtoReflect.Descriptor instead.
func (*ListTenantsV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListTenantsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTenantsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTenantsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenants []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
}

func (x *ListTenantsV1Response) Reset() {
	*x = ListTenantsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTenantsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsV1Response) ProtoMessage() {}

func (x *ListTenantsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsV1Response.ProtoReflect.Descriptor instead.
func (*ListTenantsV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListTenantsV1Response) GetTenants() []*Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type RemoveTenantV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId uint64 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *RemoveTenantV1Request) Reset() {
	*x = RemoveTenantV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTenantV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTenantV1Request) ProtoMessage() {}

func (x *RemoveTenantV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTenantV1Request.ProtoReflect.Descriptor instead.
func (*RemoveTenantV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveTenantV1Request) GetTenantId() uint64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

type RemoveTenantV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Removed bool `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *RemoveTenantV1Response) Reset() {
	*x = RemoveTenantV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTenantV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTenantV1Response) ProtoMessage() {}

func (x *RemoveTenantV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTenantV1Response.ProtoReflect.Descriptor instead.
func (*RemoveTenantV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveTenantV1Response) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

var File_api_ocp_tenant_api_ocp_tenant_api_proto protoreflect.FileDescriptor

var file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x63, 0x70, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x06, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a,
	0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4a,
	0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x32, 0xe8, 0x04,
	0x0a, 0x0c, 0x4f, 0x63, 0x70, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x12, 0x73,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31,
	0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x27, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x12, 0x25, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x1a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x73, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x56, 0x31, 0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x71, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x12, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63,
	0x70, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x3b,
	0x6f, 0x63, 0x70, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescOnce sync.Once
	file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescData = file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDesc
)

func file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescGZIP() []byte {
	file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescData)
	})
	return file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDescData
}

var file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_ocp_tenant_api_ocp_tenant_api_proto_goTypes = []interface{}{
	(*Tenant)(nil),                   // 0: ocp.tenant.api.Tenant
	(*CreateTenantV1Request)(nil),    // 1: ocp.tenant.api.CreateTenantV1Request
	(*CreateTenantV1Response)(nil),   // 2: ocp.tenant.api.CreateTenantV1Response
	(*DescribeTenantV1Request)(nil),  // 3: ocp.tenant.api.DescribeTenantV1Request
	(*DescribeTenantV1Response)(nil), // 4: ocp.tenant.api.DescribeTenantV1Response
	(*UpdateTenantV1Request)(nil),    // 5: ocp.tenant.api.UpdateTenantV1Request
	(*UpdateTenantV1Response)(nil),   // 6: ocp.tenant.api.UpdateTenantV1Response
	(*ListTenantsV1Request)(nil),     // 7: ocp.tenant.api.ListTenantsV1Request
	(*ListTenantsV1Response)(nil),    // 8: ocp.tenant.api.ListTenantsV1Response
	(*RemoveTenantV1Request)(nil),    // 9: ocp.tenant.api.RemoveTenantV1Request
	(*RemoveTenantV1Response)(nil),   // 10: ocp.tenant.api.RemoveTenantV1Response
}
var file_api_ocp_tenant_api_ocp_tenant_api_proto_depIdxs = []int32{
	0,  // 0: ocp.tenant.api.DescribeTenantV1Response.tenant:type_name -> ocp.tenant.api.Tenant
	0,  // 1: ocp.tenant.api.UpdateTenantV1Request.tenant:type_name -> ocp.tenant.api.Tenant
	0,  // 2: ocp.tenant.api.ListTenantsV1Response.tenants:type_name -> ocp.tenant.api.Tenant
	1,  // 3: ocp.tenant.api.OcpTenantApi.CreateTenantV1:input_type -> ocp.tenant.api.CreateTenantV1Request
	3,  // 4: ocp.tenant.api.OcpTenantApi.DescribeTenantV1:input_type -> ocp.tenant.api.DescribeTenantV1Request
	5,  // 5: ocp.tenant.api.OcpTenantApi.UpdateTenantV1:input_type -> ocp.tenant.api.UpdateTenantV1Request
	9,  // 6: ocp.tenant.api.OcpTenantApi.RemoveTenantV1:input_type -> ocp.tenant.api.RemoveTenantV1Request
	7,  // 7: ocp.tenant.api.OcpTenantApi.ListTenantsV1:input_type -> ocp.tenant.api.ListTenantsV1Request
	2,  // 8: ocp.tenant.api.OcpTenantApi.CreateTenantV1:output_type -> ocp.tenant.api.CreateTenantV1Response
	4,  // 9: ocp.tenant.api.OcpTenantApi.DescribeTenantV1:output_type -> ocp.tenant.api.DescribeTenantV1Response
	6,  // 10: ocp.tenant.api.OcpTenantApi.UpdateTenantV1:output_type -> ocp.tenant.api.UpdateTenantV1Response
	10, // 11: ocp.tenant.api.OcpTenantApi.RemoveTenantV1:output_type -> ocp.tenant.api.RemoveTenantV1Response
	8,  // 12: ocp.tenant.api.OcpTenantApi.ListTenantsV1:output_type -> ocp.tenant.api.ListTenantsV1Response
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_ocp_tenant_api_ocp_tenant_api_proto_init() }
func file_api_ocp_tenant_api_ocp_tenant_api_proto_init() {
	if File_api_ocp_tenant_api_ocp_tenant_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantV1Request); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantV1Response); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTenantV1Request); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTenantV1Response); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantV1Request); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantV1Response); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTenantsV1Request); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTenantsV1Response); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTenantV1Request); i {
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
		file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTenantV1Response); i {
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
			RawDescriptor: file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_tenant_api_ocp_tenant_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_tenant_api_ocp_tenant_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_tenant_api_ocp_tenant_api_proto_msgTypes,
	}.Build()
	File_api_ocp_tenant_api_ocp_tenant_api_proto = out.File
	file_api_ocp_tenant_api_ocp_tenant_api_proto_rawDesc = nil
	file_api_ocp_tenant_api_ocp_tenant_api_proto_goTypes = nil
	file_api_ocp_tenant_api_ocp_tenant_api_proto_depIdxs = nil
}
