// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google/ads/googleads/v12/services/customer_label_service.proto

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

// CustomerLabelServiceClient is the client API for CustomerLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerLabelServiceClient interface {
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error)
}

type customerLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerLabelServiceClient(cc grpc.ClientConnInterface) CustomerLabelServiceClient {
	return &customerLabelServiceClient{cc}
}

func (c *customerLabelServiceClient) MutateCustomerLabels(ctx context.Context, in *MutateCustomerLabelsRequest, opts ...grpc.CallOption) (*MutateCustomerLabelsResponse, error) {
	out := new(MutateCustomerLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.CustomerLabelService/MutateCustomerLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerLabelServiceServer is the server API for CustomerLabelService service.
// All implementations must embed UnimplementedCustomerLabelServiceServer
// for forward compatibility
type CustomerLabelServiceServer interface {
	// Creates and removes customer-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [LabelError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerLabels(context.Context, *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error)
	mustEmbedUnimplementedCustomerLabelServiceServer()
}

// UnimplementedCustomerLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerLabelServiceServer struct {
}

func (UnimplementedCustomerLabelServiceServer) MutateCustomerLabels(context.Context, *MutateCustomerLabelsRequest) (*MutateCustomerLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerLabels not implemented")
}
func (UnimplementedCustomerLabelServiceServer) mustEmbedUnimplementedCustomerLabelServiceServer() {}

// UnsafeCustomerLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerLabelServiceServer will
// result in compilation errors.
type UnsafeCustomerLabelServiceServer interface {
	mustEmbedUnimplementedCustomerLabelServiceServer()
}

func RegisterCustomerLabelServiceServer(s grpc.ServiceRegistrar, srv CustomerLabelServiceServer) {
	s.RegisterService(&CustomerLabelService_ServiceDesc, srv)
}

func _CustomerLabelService_MutateCustomerLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.CustomerLabelService/MutateCustomerLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLabelServiceServer).MutateCustomerLabels(ctx, req.(*MutateCustomerLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerLabelService_ServiceDesc is the grpc.ServiceDesc for CustomerLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.CustomerLabelService",
	HandlerType: (*CustomerLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerLabels",
			Handler:    _CustomerLabelService_MutateCustomerLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/customer_label_service.proto",
}
