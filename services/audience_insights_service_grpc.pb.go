// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: google/ads/googleads/v13/services/audience_insights_service.proto

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

// AudienceInsightsServiceClient is the client API for AudienceInsightsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudienceInsightsServiceClient interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error)
	// Lists date ranges for which audience insights data can be requested.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ListInsightsEligibleDates(ctx context.Context, in *ListInsightsEligibleDatesRequest, opts ...grpc.CallOption) (*ListInsightsEligibleDatesResponse, error)
	// Returns a collection of attributes that are represented in an audience of
	// interest, with metrics that compare each attribute's share of the audience
	// with its share of a baseline audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceCompositionInsights(ctx context.Context, in *GenerateAudienceCompositionInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceCompositionInsightsResponse, error)
}

type audienceInsightsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudienceInsightsServiceClient(cc grpc.ClientConnInterface) AudienceInsightsServiceClient {
	return &audienceInsightsServiceClient{cc}
}

func (c *audienceInsightsServiceClient) GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error) {
	out := new(GenerateInsightsFinderReportResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AudienceInsightsService/GenerateInsightsFinderReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error) {
	out := new(ListAudienceInsightsAttributesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AudienceInsightsService/ListAudienceInsightsAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) ListInsightsEligibleDates(ctx context.Context, in *ListInsightsEligibleDatesRequest, opts ...grpc.CallOption) (*ListInsightsEligibleDatesResponse, error) {
	out := new(ListInsightsEligibleDatesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AudienceInsightsService/ListInsightsEligibleDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) GenerateAudienceCompositionInsights(ctx context.Context, in *GenerateAudienceCompositionInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceCompositionInsightsResponse, error) {
	out := new(GenerateAudienceCompositionInsightsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.AudienceInsightsService/GenerateAudienceCompositionInsights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudienceInsightsServiceServer is the server API for AudienceInsightsService service.
// All implementations must embed UnimplementedAudienceInsightsServiceServer
// for forward compatibility
type AudienceInsightsServiceServer interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error)
	// Lists date ranges for which audience insights data can be requested.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ListInsightsEligibleDates(context.Context, *ListInsightsEligibleDatesRequest) (*ListInsightsEligibleDatesResponse, error)
	// Returns a collection of attributes that are represented in an audience of
	// interest, with metrics that compare each attribute's share of the audience
	// with its share of a baseline audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceCompositionInsights(context.Context, *GenerateAudienceCompositionInsightsRequest) (*GenerateAudienceCompositionInsightsResponse, error)
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

// UnimplementedAudienceInsightsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAudienceInsightsServiceServer struct {
}

func (UnimplementedAudienceInsightsServiceServer) GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInsightsFinderReport not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAudienceInsightsAttributes not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) ListInsightsEligibleDates(context.Context, *ListInsightsEligibleDatesRequest) (*ListInsightsEligibleDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInsightsEligibleDates not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) GenerateAudienceCompositionInsights(context.Context, *GenerateAudienceCompositionInsightsRequest) (*GenerateAudienceCompositionInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAudienceCompositionInsights not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) mustEmbedUnimplementedAudienceInsightsServiceServer() {
}

// UnsafeAudienceInsightsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudienceInsightsServiceServer will
// result in compilation errors.
type UnsafeAudienceInsightsServiceServer interface {
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

func RegisterAudienceInsightsServiceServer(s grpc.ServiceRegistrar, srv AudienceInsightsServiceServer) {
	s.RegisterService(&AudienceInsightsService_ServiceDesc, srv)
}

func _AudienceInsightsService_GenerateInsightsFinderReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInsightsFinderReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateInsightsFinderReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AudienceInsightsService/GenerateInsightsFinderReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateInsightsFinderReport(ctx, req.(*GenerateInsightsFinderReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_ListAudienceInsightsAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAudienceInsightsAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).ListAudienceInsightsAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AudienceInsightsService/ListAudienceInsightsAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).ListAudienceInsightsAttributes(ctx, req.(*ListAudienceInsightsAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_ListInsightsEligibleDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInsightsEligibleDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).ListInsightsEligibleDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AudienceInsightsService/ListInsightsEligibleDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).ListInsightsEligibleDates(ctx, req.(*ListInsightsEligibleDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_GenerateAudienceCompositionInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAudienceCompositionInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceCompositionInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.AudienceInsightsService/GenerateAudienceCompositionInsights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceCompositionInsights(ctx, req.(*GenerateAudienceCompositionInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudienceInsightsService_ServiceDesc is the grpc.ServiceDesc for AudienceInsightsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudienceInsightsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.AudienceInsightsService",
	HandlerType: (*AudienceInsightsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInsightsFinderReport",
			Handler:    _AudienceInsightsService_GenerateInsightsFinderReport_Handler,
		},
		{
			MethodName: "ListAudienceInsightsAttributes",
			Handler:    _AudienceInsightsService_ListAudienceInsightsAttributes_Handler,
		},
		{
			MethodName: "ListInsightsEligibleDates",
			Handler:    _AudienceInsightsService_ListInsightsEligibleDates_Handler,
		},
		{
			MethodName: "GenerateAudienceCompositionInsights",
			Handler:    _AudienceInsightsService_GenerateAudienceCompositionInsights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/audience_insights_service.proto",
}
