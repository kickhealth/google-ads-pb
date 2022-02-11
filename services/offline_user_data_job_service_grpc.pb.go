// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OfflineUserDataJobServiceClient is the client API for OfflineUserDataJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfflineUserDataJobServiceClient interface {
	// Creates an offline user data job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [NotAllowlistedError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	CreateOfflineUserDataJob(ctx context.Context, in *CreateOfflineUserDataJobRequest, opts ...grpc.CallOption) (*CreateOfflineUserDataJobResponse, error)
	// Adds operations to the offline user data job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	AddOfflineUserDataJobOperations(ctx context.Context, in *AddOfflineUserDataJobOperationsRequest, opts ...grpc.CallOption) (*AddOfflineUserDataJobOperationsResponse, error)
	// Runs the offline user data job.
	//
	// When finished, the long running operation will contain the processing
	// result or failure information, if any.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	RunOfflineUserDataJob(ctx context.Context, in *RunOfflineUserDataJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type offlineUserDataJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOfflineUserDataJobServiceClient(cc grpc.ClientConnInterface) OfflineUserDataJobServiceClient {
	return &offlineUserDataJobServiceClient{cc}
}

func (c *offlineUserDataJobServiceClient) CreateOfflineUserDataJob(ctx context.Context, in *CreateOfflineUserDataJobRequest, opts ...grpc.CallOption) (*CreateOfflineUserDataJobResponse, error) {
	out := new(CreateOfflineUserDataJobResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.OfflineUserDataJobService/CreateOfflineUserDataJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineUserDataJobServiceClient) AddOfflineUserDataJobOperations(ctx context.Context, in *AddOfflineUserDataJobOperationsRequest, opts ...grpc.CallOption) (*AddOfflineUserDataJobOperationsResponse, error) {
	out := new(AddOfflineUserDataJobOperationsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.OfflineUserDataJobService/AddOfflineUserDataJobOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineUserDataJobServiceClient) RunOfflineUserDataJob(ctx context.Context, in *RunOfflineUserDataJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.OfflineUserDataJobService/RunOfflineUserDataJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfflineUserDataJobServiceServer is the server API for OfflineUserDataJobService service.
// All implementations must embed UnimplementedOfflineUserDataJobServiceServer
// for forward compatibility
type OfflineUserDataJobServiceServer interface {
	// Creates an offline user data job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [NotAllowlistedError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	CreateOfflineUserDataJob(context.Context, *CreateOfflineUserDataJobRequest) (*CreateOfflineUserDataJobResponse, error)
	// Adds operations to the offline user data job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	AddOfflineUserDataJobOperations(context.Context, *AddOfflineUserDataJobOperationsRequest) (*AddOfflineUserDataJobOperationsResponse, error)
	// Runs the offline user data job.
	//
	// When finished, the long running operation will contain the processing
	// result or failure information, if any.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [OfflineUserDataJobError]()
	//   [QuotaError]()
	//   [RequestError]()
	RunOfflineUserDataJob(context.Context, *RunOfflineUserDataJobRequest) (*longrunning.Operation, error)
	mustEmbedUnimplementedOfflineUserDataJobServiceServer()
}

// UnimplementedOfflineUserDataJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOfflineUserDataJobServiceServer struct {
}

func (UnimplementedOfflineUserDataJobServiceServer) CreateOfflineUserDataJob(context.Context, *CreateOfflineUserDataJobRequest) (*CreateOfflineUserDataJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOfflineUserDataJob not implemented")
}
func (UnimplementedOfflineUserDataJobServiceServer) AddOfflineUserDataJobOperations(context.Context, *AddOfflineUserDataJobOperationsRequest) (*AddOfflineUserDataJobOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOfflineUserDataJobOperations not implemented")
}
func (UnimplementedOfflineUserDataJobServiceServer) RunOfflineUserDataJob(context.Context, *RunOfflineUserDataJobRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunOfflineUserDataJob not implemented")
}
func (UnimplementedOfflineUserDataJobServiceServer) mustEmbedUnimplementedOfflineUserDataJobServiceServer() {
}

// UnsafeOfflineUserDataJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfflineUserDataJobServiceServer will
// result in compilation errors.
type UnsafeOfflineUserDataJobServiceServer interface {
	mustEmbedUnimplementedOfflineUserDataJobServiceServer()
}

func RegisterOfflineUserDataJobServiceServer(s grpc.ServiceRegistrar, srv OfflineUserDataJobServiceServer) {
	s.RegisterService(&OfflineUserDataJobService_ServiceDesc, srv)
}

func _OfflineUserDataJobService_CreateOfflineUserDataJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfflineUserDataJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).CreateOfflineUserDataJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.OfflineUserDataJobService/CreateOfflineUserDataJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).CreateOfflineUserDataJob(ctx, req.(*CreateOfflineUserDataJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfflineUserDataJobService_AddOfflineUserDataJobOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOfflineUserDataJobOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).AddOfflineUserDataJobOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.OfflineUserDataJobService/AddOfflineUserDataJobOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).AddOfflineUserDataJobOperations(ctx, req.(*AddOfflineUserDataJobOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfflineUserDataJobService_RunOfflineUserDataJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunOfflineUserDataJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).RunOfflineUserDataJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.OfflineUserDataJobService/RunOfflineUserDataJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).RunOfflineUserDataJob(ctx, req.(*RunOfflineUserDataJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OfflineUserDataJobService_ServiceDesc is the grpc.ServiceDesc for OfflineUserDataJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OfflineUserDataJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.OfflineUserDataJobService",
	HandlerType: (*OfflineUserDataJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOfflineUserDataJob",
			Handler:    _OfflineUserDataJobService_CreateOfflineUserDataJob_Handler,
		},
		{
			MethodName: "AddOfflineUserDataJobOperations",
			Handler:    _OfflineUserDataJobService_AddOfflineUserDataJobOperations_Handler,
		},
		{
			MethodName: "RunOfflineUserDataJob",
			Handler:    _OfflineUserDataJobService_RunOfflineUserDataJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/offline_user_data_job_service.proto",
}