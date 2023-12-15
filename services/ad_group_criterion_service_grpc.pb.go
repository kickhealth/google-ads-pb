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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: google/ads/googleads/v15/services/ad_group_criterion_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdGroupCriterionService_MutateAdGroupCriteria_FullMethodName = "/google.ads.googleads.v15.services.AdGroupCriterionService/MutateAdGroupCriteria"
)

// AdGroupCriterionServiceClient is the client API for AdGroupCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupCriterionServiceClient interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdGroupCriterionError]()
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MultiplierError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error)
}

type adGroupCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupCriterionServiceClient(cc grpc.ClientConnInterface) AdGroupCriterionServiceClient {
	return &adGroupCriterionServiceClient{cc}
}

func (c *adGroupCriterionServiceClient) MutateAdGroupCriteria(ctx context.Context, in *MutateAdGroupCriteriaRequest, opts ...grpc.CallOption) (*MutateAdGroupCriteriaResponse, error) {
	out := new(MutateAdGroupCriteriaResponse)
	err := c.cc.Invoke(ctx, AdGroupCriterionService_MutateAdGroupCriteria_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupCriterionServiceServer is the server API for AdGroupCriterionService service.
// All implementations must embed UnimplementedAdGroupCriterionServiceServer
// for forward compatibility
type AdGroupCriterionServiceServer interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdGroupCriterionError]()
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MultiplierError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAdGroupCriteria(context.Context, *MutateAdGroupCriteriaRequest) (*MutateAdGroupCriteriaResponse, error)
	mustEmbedUnimplementedAdGroupCriterionServiceServer()
}

// UnimplementedAdGroupCriterionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupCriterionServiceServer struct {
}

func (UnimplementedAdGroupCriterionServiceServer) MutateAdGroupCriteria(context.Context, *MutateAdGroupCriteriaRequest) (*MutateAdGroupCriteriaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupCriteria not implemented")
}
func (UnimplementedAdGroupCriterionServiceServer) mustEmbedUnimplementedAdGroupCriterionServiceServer() {
}

// UnsafeAdGroupCriterionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupCriterionServiceServer will
// result in compilation errors.
type UnsafeAdGroupCriterionServiceServer interface {
	mustEmbedUnimplementedAdGroupCriterionServiceServer()
}

func RegisterAdGroupCriterionServiceServer(s grpc.ServiceRegistrar, srv AdGroupCriterionServiceServer) {
	s.RegisterService(&AdGroupCriterionService_ServiceDesc, srv)
}

func _AdGroupCriterionService_MutateAdGroupCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdGroupCriterionService_MutateAdGroupCriteria_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupCriterionServiceServer).MutateAdGroupCriteria(ctx, req.(*MutateAdGroupCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupCriterionService_ServiceDesc is the grpc.ServiceDesc for AdGroupCriterionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupCriterionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.AdGroupCriterionService",
	HandlerType: (*AdGroupCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupCriteria",
			Handler:    _AdGroupCriterionService_MutateAdGroupCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/ad_group_criterion_service.proto",
}
