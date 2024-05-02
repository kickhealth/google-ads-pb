// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v16/resources/asset_group_top_combination_view.proto

package resources

import (
	common "github.com/shenzhencenter/google-ads-pb/common"
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

// A view on the usage of ad group ad asset combination.
type AssetGroupTopCombinationView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the asset group top combination view.
	// AssetGroup Top Combination view resource names have the form:
	// `"customers/{customer_id}/assetGroupTopCombinationViews/{asset_group_id}~{asset_combination_category}"
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The top combinations of assets that served together.
	AssetGroupTopCombinations []*AssetGroupAssetCombinationData `protobuf:"bytes,2,rep,name=asset_group_top_combinations,json=assetGroupTopCombinations,proto3" json:"asset_group_top_combinations,omitempty"`
}

func (x *AssetGroupTopCombinationView) Reset() {
	*x = AssetGroupTopCombinationView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupTopCombinationView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupTopCombinationView) ProtoMessage() {}

func (x *AssetGroupTopCombinationView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupTopCombinationView.ProtoReflect.Descriptor instead.
func (*AssetGroupTopCombinationView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescGZIP(), []int{0}
}

func (x *AssetGroupTopCombinationView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetGroupTopCombinationView) GetAssetGroupTopCombinations() []*AssetGroupAssetCombinationData {
	if x != nil {
		return x.AssetGroupTopCombinations
	}
	return nil
}

// Asset group asset combination data
type AssetGroupAssetCombinationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Served assets.
	AssetCombinationServedAssets []*common.AssetUsage `protobuf:"bytes,1,rep,name=asset_combination_served_assets,json=assetCombinationServedAssets,proto3" json:"asset_combination_served_assets,omitempty"`
}

func (x *AssetGroupAssetCombinationData) Reset() {
	*x = AssetGroupAssetCombinationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupAssetCombinationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupAssetCombinationData) ProtoMessage() {}

func (x *AssetGroupAssetCombinationData) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupAssetCombinationData.ProtoReflect.Descriptor instead.
func (*AssetGroupAssetCombinationData) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescGZIP(), []int{1}
}

func (x *AssetGroupAssetCombinationData) GetAssetCombinationServedAssets() []*common.AssetUsage {
	if x != nil {
		return x.AssetCombinationServedAssets
	}
	return nil
}

var File_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x03, 0x0a, 0x1c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x70,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x62, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3d, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x37, 0x0a, 0x35,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x1c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x19, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54,
	0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0xa0,
	0x01, 0xea, 0x41, 0x9c, 0x01, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x12, 0x63, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69,
	0x65, 0x77, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x7d, 0x22, 0x99, 0x01, 0x0a, 0x1e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x77, 0x0a, 0x1f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x1c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x42, 0x93, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x21, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36,
	0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescData = file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDesc
)

func file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDescData
}

var file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_goTypes = []interface{}{
	(*AssetGroupTopCombinationView)(nil),   // 0: google.ads.googleads.v16.resources.AssetGroupTopCombinationView
	(*AssetGroupAssetCombinationData)(nil), // 1: google.ads.googleads.v16.resources.AssetGroupAssetCombinationData
	(*common.AssetUsage)(nil),              // 2: google.ads.googleads.v16.common.AssetUsage
}
var file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v16.resources.AssetGroupTopCombinationView.asset_group_top_combinations:type_name -> google.ads.googleads.v16.resources.AssetGroupAssetCombinationData
	2, // 1: google.ads.googleads.v16.resources.AssetGroupAssetCombinationData.asset_combination_served_assets:type_name -> google.ads.googleads.v16.common.AssetUsage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_init() }
func file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_init() {
	if File_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupTopCombinationView); i {
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
		file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupAssetCombinationData); i {
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
			RawDescriptor: file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto = out.File
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_rawDesc = nil
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_goTypes = nil
	file_google_ads_googleads_v16_resources_asset_group_top_combination_view_proto_depIdxs = nil
}