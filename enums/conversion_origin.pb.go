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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v16/enums/conversion_origin.proto

package enums

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

// The possible places where a conversion can occur.
type ConversionOriginEnum_ConversionOrigin int32

const (
	// The conversion origin has not been specified.
	ConversionOriginEnum_UNSPECIFIED ConversionOriginEnum_ConversionOrigin = 0
	// The conversion origin is not known in this version.
	ConversionOriginEnum_UNKNOWN ConversionOriginEnum_ConversionOrigin = 1
	// Conversion that occurs when a user visits a website or takes an action
	// there after viewing an ad.
	ConversionOriginEnum_WEBSITE ConversionOriginEnum_ConversionOrigin = 2
	// Conversions reported by an offline pipeline which collects local actions
	// from Google-hosted pages (for example, Google Maps, Google Place Page,
	// etc) and attributes them to relevant ad events.
	ConversionOriginEnum_GOOGLE_HOSTED ConversionOriginEnum_ConversionOrigin = 3
	// Conversion that occurs when a user performs an action through any app
	// platforms.
	ConversionOriginEnum_APP ConversionOriginEnum_ConversionOrigin = 4
	// Conversion that occurs when a user makes a call from ads.
	ConversionOriginEnum_CALL_FROM_ADS ConversionOriginEnum_ConversionOrigin = 5
	// Conversion that occurs when a user visits or makes a purchase at a
	// physical store.
	ConversionOriginEnum_STORE ConversionOriginEnum_ConversionOrigin = 6
	// Conversion that occurs on YouTube.
	ConversionOriginEnum_YOUTUBE_HOSTED ConversionOriginEnum_ConversionOrigin = 7
)

// Enum value maps for ConversionOriginEnum_ConversionOrigin.
var (
	ConversionOriginEnum_ConversionOrigin_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "WEBSITE",
		3: "GOOGLE_HOSTED",
		4: "APP",
		5: "CALL_FROM_ADS",
		6: "STORE",
		7: "YOUTUBE_HOSTED",
	}
	ConversionOriginEnum_ConversionOrigin_value = map[string]int32{
		"UNSPECIFIED":    0,
		"UNKNOWN":        1,
		"WEBSITE":        2,
		"GOOGLE_HOSTED":  3,
		"APP":            4,
		"CALL_FROM_ADS":  5,
		"STORE":          6,
		"YOUTUBE_HOSTED": 7,
	}
)

func (x ConversionOriginEnum_ConversionOrigin) Enum() *ConversionOriginEnum_ConversionOrigin {
	p := new(ConversionOriginEnum_ConversionOrigin)
	*p = x
	return p
}

func (x ConversionOriginEnum_ConversionOrigin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionOriginEnum_ConversionOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_enums_conversion_origin_proto_enumTypes[0].Descriptor()
}

func (ConversionOriginEnum_ConversionOrigin) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_enums_conversion_origin_proto_enumTypes[0]
}

func (x ConversionOriginEnum_ConversionOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionOriginEnum_ConversionOrigin.Descriptor instead.
func (ConversionOriginEnum_ConversionOrigin) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion origins.
type ConversionOriginEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionOriginEnum) Reset() {
	*x = ConversionOriginEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_enums_conversion_origin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionOriginEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionOriginEnum) ProtoMessage() {}

func (x *ConversionOriginEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_enums_conversion_origin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionOriginEnum.ProtoReflect.Descriptor instead.
func (*ConversionOriginEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_enums_conversion_origin_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x10, 0x04, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x05,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x59,
	0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x45, 0x44, 0x10, 0x07, 0x42,
	0xef, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescData = file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDesc
)

func file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDescData
}

var file_google_ads_googleads_v16_enums_conversion_origin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_enums_conversion_origin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_enums_conversion_origin_proto_goTypes = []interface{}{
	(ConversionOriginEnum_ConversionOrigin)(0), // 0: google.ads.googleads.v16.enums.ConversionOriginEnum.ConversionOrigin
	(*ConversionOriginEnum)(nil),               // 1: google.ads.googleads.v16.enums.ConversionOriginEnum
}
var file_google_ads_googleads_v16_enums_conversion_origin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_enums_conversion_origin_proto_init() }
func file_google_ads_googleads_v16_enums_conversion_origin_proto_init() {
	if File_google_ads_googleads_v16_enums_conversion_origin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_enums_conversion_origin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionOriginEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_enums_conversion_origin_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_enums_conversion_origin_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_enums_conversion_origin_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_enums_conversion_origin_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_enums_conversion_origin_proto = out.File
	file_google_ads_googleads_v16_enums_conversion_origin_proto_rawDesc = nil
	file_google_ads_googleads_v16_enums_conversion_origin_proto_goTypes = nil
	file_google_ads_googleads_v16_enums_conversion_origin_proto_depIdxs = nil
}
