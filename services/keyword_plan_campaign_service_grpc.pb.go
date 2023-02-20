// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: google/ads/googleads/v13/services/keyword_plan_campaign_service.proto

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

// KeywordPlanCampaignServiceClient is the client API for KeywordPlanCampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordPlanCampaignServiceClient interface {
	// Creates, updates, or removes Keyword Plan campaigns. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanCampaignError]()
	//	[KeywordPlanError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanCampaigns(ctx context.Context, in *MutateKeywordPlanCampaignsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignsResponse, error)
}

type keywordPlanCampaignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanCampaignServiceClient(cc grpc.ClientConnInterface) KeywordPlanCampaignServiceClient {
	return &keywordPlanCampaignServiceClient{cc}
}

func (c *keywordPlanCampaignServiceClient) MutateKeywordPlanCampaigns(ctx context.Context, in *MutateKeywordPlanCampaignsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignsResponse, error) {
	out := new(MutateKeywordPlanCampaignsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.KeywordPlanCampaignService/MutateKeywordPlanCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanCampaignServiceServer is the server API for KeywordPlanCampaignService service.
// All implementations must embed UnimplementedKeywordPlanCampaignServiceServer
// for forward compatibility
type KeywordPlanCampaignServiceServer interface {
	// Creates, updates, or removes Keyword Plan campaigns. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanCampaignError]()
	//	[KeywordPlanError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanCampaigns(context.Context, *MutateKeywordPlanCampaignsRequest) (*MutateKeywordPlanCampaignsResponse, error)
	mustEmbedUnimplementedKeywordPlanCampaignServiceServer()
}

// UnimplementedKeywordPlanCampaignServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanCampaignServiceServer struct {
}

func (UnimplementedKeywordPlanCampaignServiceServer) MutateKeywordPlanCampaigns(context.Context, *MutateKeywordPlanCampaignsRequest) (*MutateKeywordPlanCampaignsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlanCampaigns not implemented")
}
func (UnimplementedKeywordPlanCampaignServiceServer) mustEmbedUnimplementedKeywordPlanCampaignServiceServer() {
}

// UnsafeKeywordPlanCampaignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanCampaignServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanCampaignServiceServer interface {
	mustEmbedUnimplementedKeywordPlanCampaignServiceServer()
}

func RegisterKeywordPlanCampaignServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanCampaignServiceServer) {
	s.RegisterService(&KeywordPlanCampaignService_ServiceDesc, srv)
}

func _KeywordPlanCampaignService_MutateKeywordPlanCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanCampaignServiceServer).MutateKeywordPlanCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.KeywordPlanCampaignService/MutateKeywordPlanCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanCampaignServiceServer).MutateKeywordPlanCampaigns(ctx, req.(*MutateKeywordPlanCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanCampaignService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanCampaignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanCampaignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.KeywordPlanCampaignService",
	HandlerType: (*KeywordPlanCampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlanCampaigns",
			Handler:    _KeywordPlanCampaignService_MutateKeywordPlanCampaigns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/keyword_plan_campaign_service.proto",
}
