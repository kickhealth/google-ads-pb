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
// source: google/ads/googleads/v16/services/campaign_lifecycle_goal_service.proto

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
	CampaignLifecycleGoalService_ConfigureCampaignLifecycleGoals_FullMethodName = "/google.ads.googleads.v16.services.CampaignLifecycleGoalService/ConfigureCampaignLifecycleGoals"
)

// CampaignLifecycleGoalServiceClient is the client API for CampaignLifecycleGoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignLifecycleGoalServiceClient interface {
	// Process the given campaign lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCampaignLifecycleGoals(ctx context.Context, in *ConfigureCampaignLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCampaignLifecycleGoalsResponse, error)
}

type campaignLifecycleGoalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLifecycleGoalServiceClient(cc grpc.ClientConnInterface) CampaignLifecycleGoalServiceClient {
	return &campaignLifecycleGoalServiceClient{cc}
}

func (c *campaignLifecycleGoalServiceClient) ConfigureCampaignLifecycleGoals(ctx context.Context, in *ConfigureCampaignLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCampaignLifecycleGoalsResponse, error) {
	out := new(ConfigureCampaignLifecycleGoalsResponse)
	err := c.cc.Invoke(ctx, CampaignLifecycleGoalService_ConfigureCampaignLifecycleGoals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLifecycleGoalServiceServer is the server API for CampaignLifecycleGoalService service.
// All implementations must embed UnimplementedCampaignLifecycleGoalServiceServer
// for forward compatibility
type CampaignLifecycleGoalServiceServer interface {
	// Process the given campaign lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCampaignLifecycleGoals(context.Context, *ConfigureCampaignLifecycleGoalsRequest) (*ConfigureCampaignLifecycleGoalsResponse, error)
	mustEmbedUnimplementedCampaignLifecycleGoalServiceServer()
}

// UnimplementedCampaignLifecycleGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignLifecycleGoalServiceServer struct {
}

func (UnimplementedCampaignLifecycleGoalServiceServer) ConfigureCampaignLifecycleGoals(context.Context, *ConfigureCampaignLifecycleGoalsRequest) (*ConfigureCampaignLifecycleGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCampaignLifecycleGoals not implemented")
}
func (UnimplementedCampaignLifecycleGoalServiceServer) mustEmbedUnimplementedCampaignLifecycleGoalServiceServer() {
}

// UnsafeCampaignLifecycleGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignLifecycleGoalServiceServer will
// result in compilation errors.
type UnsafeCampaignLifecycleGoalServiceServer interface {
	mustEmbedUnimplementedCampaignLifecycleGoalServiceServer()
}

func RegisterCampaignLifecycleGoalServiceServer(s grpc.ServiceRegistrar, srv CampaignLifecycleGoalServiceServer) {
	s.RegisterService(&CampaignLifecycleGoalService_ServiceDesc, srv)
}

func _CampaignLifecycleGoalService_ConfigureCampaignLifecycleGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureCampaignLifecycleGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLifecycleGoalServiceServer).ConfigureCampaignLifecycleGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignLifecycleGoalService_ConfigureCampaignLifecycleGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLifecycleGoalServiceServer).ConfigureCampaignLifecycleGoals(ctx, req.(*ConfigureCampaignLifecycleGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignLifecycleGoalService_ServiceDesc is the grpc.ServiceDesc for CampaignLifecycleGoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignLifecycleGoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v16.services.CampaignLifecycleGoalService",
	HandlerType: (*CampaignLifecycleGoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureCampaignLifecycleGoals",
			Handler:    _CampaignLifecycleGoalService_ConfigureCampaignLifecycleGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v16/services/campaign_lifecycle_goal_service.proto",
}