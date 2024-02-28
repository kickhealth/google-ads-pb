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
// source: google/ads/googleads/v16/services/customer_lifecycle_goal_service.proto

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
	CustomerLifecycleGoalService_ConfigureCustomerLifecycleGoals_FullMethodName = "/google.ads.googleads.v16.services.CustomerLifecycleGoalService/ConfigureCustomerLifecycleGoals"
)

// CustomerLifecycleGoalServiceClient is the client API for CustomerLifecycleGoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerLifecycleGoalServiceClient interface {
	// Process the given customer lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCustomerLifecycleGoals(ctx context.Context, in *ConfigureCustomerLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCustomerLifecycleGoalsResponse, error)
}

type customerLifecycleGoalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerLifecycleGoalServiceClient(cc grpc.ClientConnInterface) CustomerLifecycleGoalServiceClient {
	return &customerLifecycleGoalServiceClient{cc}
}

func (c *customerLifecycleGoalServiceClient) ConfigureCustomerLifecycleGoals(ctx context.Context, in *ConfigureCustomerLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCustomerLifecycleGoalsResponse, error) {
	out := new(ConfigureCustomerLifecycleGoalsResponse)
	err := c.cc.Invoke(ctx, CustomerLifecycleGoalService_ConfigureCustomerLifecycleGoals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerLifecycleGoalServiceServer is the server API for CustomerLifecycleGoalService service.
// All implementations must embed UnimplementedCustomerLifecycleGoalServiceServer
// for forward compatibility
type CustomerLifecycleGoalServiceServer interface {
	// Process the given customer lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCustomerLifecycleGoals(context.Context, *ConfigureCustomerLifecycleGoalsRequest) (*ConfigureCustomerLifecycleGoalsResponse, error)
	mustEmbedUnimplementedCustomerLifecycleGoalServiceServer()
}

// UnimplementedCustomerLifecycleGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerLifecycleGoalServiceServer struct {
}

func (UnimplementedCustomerLifecycleGoalServiceServer) ConfigureCustomerLifecycleGoals(context.Context, *ConfigureCustomerLifecycleGoalsRequest) (*ConfigureCustomerLifecycleGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCustomerLifecycleGoals not implemented")
}
func (UnimplementedCustomerLifecycleGoalServiceServer) mustEmbedUnimplementedCustomerLifecycleGoalServiceServer() {
}

// UnsafeCustomerLifecycleGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerLifecycleGoalServiceServer will
// result in compilation errors.
type UnsafeCustomerLifecycleGoalServiceServer interface {
	mustEmbedUnimplementedCustomerLifecycleGoalServiceServer()
}

func RegisterCustomerLifecycleGoalServiceServer(s grpc.ServiceRegistrar, srv CustomerLifecycleGoalServiceServer) {
	s.RegisterService(&CustomerLifecycleGoalService_ServiceDesc, srv)
}

func _CustomerLifecycleGoalService_ConfigureCustomerLifecycleGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureCustomerLifecycleGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLifecycleGoalServiceServer).ConfigureCustomerLifecycleGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerLifecycleGoalService_ConfigureCustomerLifecycleGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLifecycleGoalServiceServer).ConfigureCustomerLifecycleGoals(ctx, req.(*ConfigureCustomerLifecycleGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerLifecycleGoalService_ServiceDesc is the grpc.ServiceDesc for CustomerLifecycleGoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerLifecycleGoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v16.services.CustomerLifecycleGoalService",
	HandlerType: (*CustomerLifecycleGoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureCustomerLifecycleGoals",
			Handler:    _CustomerLifecycleGoalService_ConfigureCustomerLifecycleGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v16/services/customer_lifecycle_goal_service.proto",
}
