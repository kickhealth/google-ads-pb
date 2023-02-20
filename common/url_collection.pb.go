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
// 	protoc        v3.21.9
// source: google/ads/googleads/v13/common/url_collection.proto

package common

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

// Collection of urls that is tagged with a unique identifier.
type UrlCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for this UrlCollection instance.
	UrlCollectionId *string `protobuf:"bytes,5,opt,name=url_collection_id,json=urlCollectionId,proto3,oneof" json:"url_collection_id,omitempty"`
	// A list of possible final URLs.
	FinalUrls []string `protobuf:"bytes,6,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of possible final mobile URLs.
	FinalMobileUrls []string `protobuf:"bytes,7,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// URL template for constructing a tracking URL.
	TrackingUrlTemplate *string `protobuf:"bytes,8,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3,oneof" json:"tracking_url_template,omitempty"`
}

func (x *UrlCollection) Reset() {
	*x = UrlCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_common_url_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlCollection) ProtoMessage() {}

func (x *UrlCollection) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_common_url_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlCollection.ProtoReflect.Descriptor instead.
func (*UrlCollection) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_common_url_collection_proto_rawDescGZIP(), []int{0}
}

func (x *UrlCollection) GetUrlCollectionId() string {
	if x != nil && x.UrlCollectionId != nil {
		return *x.UrlCollectionId
	}
	return ""
}

func (x *UrlCollection) GetFinalUrls() []string {
	if x != nil {
		return x.FinalUrls
	}
	return nil
}

func (x *UrlCollection) GetFinalMobileUrls() []string {
	if x != nil {
		return x.FinalMobileUrls
	}
	return nil
}

func (x *UrlCollection) GetTrackingUrlTemplate() string {
	if x != nil && x.TrackingUrlTemplate != nil {
		return *x.TrackingUrlTemplate
	}
	return ""
}

var File_google_ads_googleads_v13_common_url_collection_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_common_url_collection_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x75, 0x72, 0x6c, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xf4, 0x01, 0x0a, 0x0d, 0x55, 0x72, 0x6c, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x11, 0x75, 0x72, 0x6c,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x75, 0x72, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x55, 0x72, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0xf2,
	0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x12, 0x55, 0x72, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_common_url_collection_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_common_url_collection_proto_rawDescData = file_google_ads_googleads_v13_common_url_collection_proto_rawDesc
)

func file_google_ads_googleads_v13_common_url_collection_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_common_url_collection_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_common_url_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_common_url_collection_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_common_url_collection_proto_rawDescData
}

var file_google_ads_googleads_v13_common_url_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_common_url_collection_proto_goTypes = []interface{}{
	(*UrlCollection)(nil), // 0: google.ads.googleads.v13.common.UrlCollection
}
var file_google_ads_googleads_v13_common_url_collection_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_common_url_collection_proto_init() }
func file_google_ads_googleads_v13_common_url_collection_proto_init() {
	if File_google_ads_googleads_v13_common_url_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_common_url_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlCollection); i {
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
	file_google_ads_googleads_v13_common_url_collection_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v13_common_url_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_common_url_collection_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_common_url_collection_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v13_common_url_collection_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_common_url_collection_proto = out.File
	file_google_ads_googleads_v13_common_url_collection_proto_rawDesc = nil
	file_google_ads_googleads_v13_common_url_collection_proto_goTypes = nil
	file_google_ads_googleads_v13_common_url_collection_proto_depIdxs = nil
}
