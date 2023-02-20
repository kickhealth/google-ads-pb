// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: google/ads/googleads/v13/services/campaign_customizer_service.proto

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

// CampaignCustomizerServiceClient is the client API for CampaignCustomizerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignCustomizerServiceClient interface {
	// Creates, updates or removes campaign customizers. Operation statuses are
	// returned.
	MutateCampaignCustomizers(ctx context.Context, in *MutateCampaignCustomizersRequest, opts ...grpc.CallOption) (*MutateCampaignCustomizersResponse, error)
}

type campaignCustomizerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignCustomizerServiceClient(cc grpc.ClientConnInterface) CampaignCustomizerServiceClient {
	return &campaignCustomizerServiceClient{cc}
}

func (c *campaignCustomizerServiceClient) MutateCampaignCustomizers(ctx context.Context, in *MutateCampaignCustomizersRequest, opts ...grpc.CallOption) (*MutateCampaignCustomizersResponse, error) {
	out := new(MutateCampaignCustomizersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v13.services.CampaignCustomizerService/MutateCampaignCustomizers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCustomizerServiceServer is the server API for CampaignCustomizerService service.
// All implementations must embed UnimplementedCampaignCustomizerServiceServer
// for forward compatibility
type CampaignCustomizerServiceServer interface {
	// Creates, updates or removes campaign customizers. Operation statuses are
	// returned.
	MutateCampaignCustomizers(context.Context, *MutateCampaignCustomizersRequest) (*MutateCampaignCustomizersResponse, error)
	mustEmbedUnimplementedCampaignCustomizerServiceServer()
}

// UnimplementedCampaignCustomizerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignCustomizerServiceServer struct {
}

func (UnimplementedCampaignCustomizerServiceServer) MutateCampaignCustomizers(context.Context, *MutateCampaignCustomizersRequest) (*MutateCampaignCustomizersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignCustomizers not implemented")
}
func (UnimplementedCampaignCustomizerServiceServer) mustEmbedUnimplementedCampaignCustomizerServiceServer() {
}

// UnsafeCampaignCustomizerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignCustomizerServiceServer will
// result in compilation errors.
type UnsafeCampaignCustomizerServiceServer interface {
	mustEmbedUnimplementedCampaignCustomizerServiceServer()
}

func RegisterCampaignCustomizerServiceServer(s grpc.ServiceRegistrar, srv CampaignCustomizerServiceServer) {
	s.RegisterService(&CampaignCustomizerService_ServiceDesc, srv)
}

func _CampaignCustomizerService_MutateCampaignCustomizers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCustomizersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCustomizerServiceServer).MutateCampaignCustomizers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v13.services.CampaignCustomizerService/MutateCampaignCustomizers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCustomizerServiceServer).MutateCampaignCustomizers(ctx, req.(*MutateCampaignCustomizersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignCustomizerService_ServiceDesc is the grpc.ServiceDesc for CampaignCustomizerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignCustomizerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CampaignCustomizerService",
	HandlerType: (*CampaignCustomizerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignCustomizers",
			Handler:    _CampaignCustomizerService_MutateCampaignCustomizers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/campaign_customizer_service.proto",
}
