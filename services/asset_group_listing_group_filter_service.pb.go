// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/ads/googleads/v12/services/asset_group_listing_group_filter_service.proto

package services

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters][google.ads.googleads.v12.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters].
// partial_failure is not supported because the tree needs to be validated
// together.
type MutateAssetGroupListingGroupFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose asset group listing group filters are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual asset group listing group
	// filters.
	Operations []*AssetGroupListingGroupFilterOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,4,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v12.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateAssetGroupListingGroupFiltersRequest) Reset() {
	*x = MutateAssetGroupListingGroupFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupListingGroupFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFiltersRequest) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFiltersRequest.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFiltersRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetOperations() []*AssetGroupListingGroupFilterOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove) on an asset group listing group filter.
type AssetGroupListingGroupFilterOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*AssetGroupListingGroupFilterOperation_Create
	//	*AssetGroupListingGroupFilterOperation_Update
	//	*AssetGroupListingGroupFilterOperation_Remove
	Operation isAssetGroupListingGroupFilterOperation_Operation `protobuf_oneof:"operation"`
}

func (x *AssetGroupListingGroupFilterOperation) Reset() {
	*x = AssetGroupListingGroupFilterOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupListingGroupFilterOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupListingGroupFilterOperation) ProtoMessage() {}

func (x *AssetGroupListingGroupFilterOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupListingGroupFilterOperation.ProtoReflect.Descriptor instead.
func (*AssetGroupListingGroupFilterOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{1}
}

func (x *AssetGroupListingGroupFilterOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *AssetGroupListingGroupFilterOperation) GetOperation() isAssetGroupListingGroupFilterOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetCreate() *resources.AssetGroupListingGroupFilter {
	if x, ok := x.GetOperation().(*AssetGroupListingGroupFilterOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetUpdate() *resources.AssetGroupListingGroupFilter {
	if x, ok := x.GetOperation().(*AssetGroupListingGroupFilterOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*AssetGroupListingGroupFilterOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isAssetGroupListingGroupFilterOperation_Operation interface {
	isAssetGroupListingGroupFilterOperation_Operation()
}

type AssetGroupListingGroupFilterOperation_Create struct {
	// Create operation: No resource name is expected for the new asset group
	// listing group filter.
	Create *resources.AssetGroupListingGroupFilter `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AssetGroupListingGroupFilterOperation_Update struct {
	// Update operation: The asset group listing group filter is expected to
	// have a valid resource name.
	Update *resources.AssetGroupListingGroupFilter `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AssetGroupListingGroupFilterOperation_Remove struct {
	// Remove operation: A resource name for the removed asset group listing
	// group filter is expected, in this format:
	// `customers/{customer_id}/assetGroupListingGroupFilters/{asset_group_id}~{listing_group_filter_id}`
	// An entity can be removed only if it's not referenced by other
	// parent_listing_group_id. If multiple entities are being deleted, the
	// mutates must be in the correct order.
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AssetGroupListingGroupFilterOperation_Create) isAssetGroupListingGroupFilterOperation_Operation() {
}

func (*AssetGroupListingGroupFilterOperation_Update) isAssetGroupListingGroupFilterOperation_Operation() {
}

func (*AssetGroupListingGroupFilterOperation_Remove) isAssetGroupListingGroupFilterOperation_Operation() {
}

// Response message for an asset group listing group filter mutate.
type MutateAssetGroupListingGroupFiltersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Results []*MutateAssetGroupListingGroupFilterResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateAssetGroupListingGroupFiltersResponse) Reset() {
	*x = MutateAssetGroupListingGroupFiltersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupListingGroupFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFiltersResponse) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFiltersResponse.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFiltersResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAssetGroupListingGroupFiltersResponse) GetResults() []*MutateAssetGroupListingGroupFilterResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the asset group listing group filter mutate.
type MutateAssetGroupListingGroupFilterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated AssetGroupListingGroupFilter with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	AssetGroupListingGroupFilter *resources.AssetGroupListingGroupFilter `protobuf:"bytes,2,opt,name=asset_group_listing_group_filter,json=assetGroupListingGroupFilter,proto3" json:"asset_group_listing_group_filter,omitempty"`
}

func (x *MutateAssetGroupListingGroupFilterResult) Reset() {
	*x = MutateAssetGroupListingGroupFilterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateAssetGroupListingGroupFilterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFilterResult) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFilterResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFilterResult.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFilterResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAssetGroupListingGroupFilterResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateAssetGroupListingGroupFilterResult) GetAssetGroupListingGroupFilter() *resources.AssetGroupListingGroupFilter {
	if x != nil {
		return x.AssetGroupListingGroupFilter
	}
	return nil
}

var File_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x2a, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xff, 0x02, 0x0a, 0x25,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x5a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5a,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a,
	0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01,
	0x0a, 0x2b, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x96, 0x02, 0x0a, 0x28, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x1c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x32, 0x9c, 0x03,
	0x0a, 0x23, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x02, 0x0a, 0x23, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x48, 0x22, 0x43, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x94, 0x02, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x28, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x32, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescData = file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDesc
)

func file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDescData
}

var file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_goTypes = []interface{}{
	(*MutateAssetGroupListingGroupFiltersRequest)(nil),     // 0: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersRequest
	(*AssetGroupListingGroupFilterOperation)(nil),          // 1: google.ads.googleads.v12.services.AssetGroupListingGroupFilterOperation
	(*MutateAssetGroupListingGroupFiltersResponse)(nil),    // 2: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersResponse
	(*MutateAssetGroupListingGroupFilterResult)(nil),       // 3: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFilterResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.AssetGroupListingGroupFilter)(nil),         // 6: google.ads.googleads.v12.resources.AssetGroupListingGroupFilter
}
var file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersRequest.operations:type_name -> google.ads.googleads.v12.services.AssetGroupListingGroupFilterOperation
	4, // 1: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersRequest.response_content_type:type_name -> google.ads.googleads.v12.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v12.services.AssetGroupListingGroupFilterOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v12.services.AssetGroupListingGroupFilterOperation.create:type_name -> google.ads.googleads.v12.resources.AssetGroupListingGroupFilter
	6, // 4: google.ads.googleads.v12.services.AssetGroupListingGroupFilterOperation.update:type_name -> google.ads.googleads.v12.resources.AssetGroupListingGroupFilter
	3, // 5: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersResponse.results:type_name -> google.ads.googleads.v12.services.MutateAssetGroupListingGroupFilterResult
	6, // 6: google.ads.googleads.v12.services.MutateAssetGroupListingGroupFilterResult.asset_group_listing_group_filter:type_name -> google.ads.googleads.v12.resources.AssetGroupListingGroupFilter
	0, // 7: google.ads.googleads.v12.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters:input_type -> google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersRequest
	2, // 8: google.ads.googleads.v12.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters:output_type -> google.ads.googleads.v12.services.MutateAssetGroupListingGroupFiltersResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_init()
}
func file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_init() {
	if File_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupListingGroupFiltersRequest); i {
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
		file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupListingGroupFilterOperation); i {
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
		file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupListingGroupFiltersResponse); i {
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
		file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateAssetGroupListingGroupFilterResult); i {
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
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AssetGroupListingGroupFilterOperation_Create)(nil),
		(*AssetGroupListingGroupFilterOperation_Update)(nil),
		(*AssetGroupListingGroupFilterOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto = out.File
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_rawDesc = nil
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_goTypes = nil
	file_google_ads_googleads_v12_services_asset_group_listing_group_filter_service_proto_depIdxs = nil
}
