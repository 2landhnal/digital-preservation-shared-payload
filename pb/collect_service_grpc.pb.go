// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.0--rc1
// source: collect_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CollectService_UpdateTmpObject_FullMethodName             = "/pb.CollectService/UpdateTmpObject"
	CollectService_CreateTmpObject_FullMethodName             = "/pb.CollectService/CreateTmpObject"
	CollectService_CreateVirussScanResult_FullMethodName      = "/pb.CollectService/CreateVirussScanResult"
	CollectService_CreateExtractMetadataResult_FullMethodName = "/pb.CollectService/CreateExtractMetadataResult"
)

// CollectServiceClient is the client API for CollectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectServiceClient interface {
	UpdateTmpObject(ctx context.Context, in *UpdateTmpObjectRequest, opts ...grpc.CallOption) (*UpdateTmpObjectResponse, error)
	CreateTmpObject(ctx context.Context, in *CreateTmpObjectRequest, opts ...grpc.CallOption) (*CreateTmpObjectResponse, error)
	CreateVirussScanResult(ctx context.Context, in *CreateVirussScanResultRequest, opts ...grpc.CallOption) (*CreateVirussScanResultResponse, error)
	CreateExtractMetadataResult(ctx context.Context, in *CreateExtractMetadataResultRequest, opts ...grpc.CallOption) (*CreateExtractMetadataResultResponse, error)
}

type collectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectServiceClient(cc grpc.ClientConnInterface) CollectServiceClient {
	return &collectServiceClient{cc}
}

func (c *collectServiceClient) UpdateTmpObject(ctx context.Context, in *UpdateTmpObjectRequest, opts ...grpc.CallOption) (*UpdateTmpObjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTmpObjectResponse)
	err := c.cc.Invoke(ctx, CollectService_UpdateTmpObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateTmpObject(ctx context.Context, in *CreateTmpObjectRequest, opts ...grpc.CallOption) (*CreateTmpObjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTmpObjectResponse)
	err := c.cc.Invoke(ctx, CollectService_CreateTmpObject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateVirussScanResult(ctx context.Context, in *CreateVirussScanResultRequest, opts ...grpc.CallOption) (*CreateVirussScanResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVirussScanResultResponse)
	err := c.cc.Invoke(ctx, CollectService_CreateVirussScanResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CreateExtractMetadataResult(ctx context.Context, in *CreateExtractMetadataResultRequest, opts ...grpc.CallOption) (*CreateExtractMetadataResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateExtractMetadataResultResponse)
	err := c.cc.Invoke(ctx, CollectService_CreateExtractMetadataResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectServiceServer is the server API for CollectService service.
// All implementations must embed UnimplementedCollectServiceServer
// for forward compatibility.
type CollectServiceServer interface {
	UpdateTmpObject(context.Context, *UpdateTmpObjectRequest) (*UpdateTmpObjectResponse, error)
	CreateTmpObject(context.Context, *CreateTmpObjectRequest) (*CreateTmpObjectResponse, error)
	CreateVirussScanResult(context.Context, *CreateVirussScanResultRequest) (*CreateVirussScanResultResponse, error)
	CreateExtractMetadataResult(context.Context, *CreateExtractMetadataResultRequest) (*CreateExtractMetadataResultResponse, error)
	mustEmbedUnimplementedCollectServiceServer()
}

// UnimplementedCollectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCollectServiceServer struct{}

func (UnimplementedCollectServiceServer) UpdateTmpObject(context.Context, *UpdateTmpObjectRequest) (*UpdateTmpObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTmpObject not implemented")
}
func (UnimplementedCollectServiceServer) CreateTmpObject(context.Context, *CreateTmpObjectRequest) (*CreateTmpObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTmpObject not implemented")
}
func (UnimplementedCollectServiceServer) CreateVirussScanResult(context.Context, *CreateVirussScanResultRequest) (*CreateVirussScanResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVirussScanResult not implemented")
}
func (UnimplementedCollectServiceServer) CreateExtractMetadataResult(context.Context, *CreateExtractMetadataResultRequest) (*CreateExtractMetadataResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExtractMetadataResult not implemented")
}
func (UnimplementedCollectServiceServer) mustEmbedUnimplementedCollectServiceServer() {}
func (UnimplementedCollectServiceServer) testEmbeddedByValue()                        {}

// UnsafeCollectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectServiceServer will
// result in compilation errors.
type UnsafeCollectServiceServer interface {
	mustEmbedUnimplementedCollectServiceServer()
}

func RegisterCollectServiceServer(s grpc.ServiceRegistrar, srv CollectServiceServer) {
	// If the following call pancis, it indicates UnimplementedCollectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CollectService_ServiceDesc, srv)
}

func _CollectService_UpdateTmpObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTmpObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).UpdateTmpObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_UpdateTmpObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).UpdateTmpObject(ctx, req.(*UpdateTmpObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateTmpObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTmpObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateTmpObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_CreateTmpObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateTmpObject(ctx, req.(*CreateTmpObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateVirussScanResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVirussScanResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateVirussScanResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_CreateVirussScanResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateVirussScanResult(ctx, req.(*CreateVirussScanResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CreateExtractMetadataResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExtractMetadataResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CreateExtractMetadataResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_CreateExtractMetadataResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CreateExtractMetadataResult(ctx, req.(*CreateExtractMetadataResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectService_ServiceDesc is the grpc.ServiceDesc for CollectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CollectService",
	HandlerType: (*CollectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTmpObject",
			Handler:    _CollectService_UpdateTmpObject_Handler,
		},
		{
			MethodName: "CreateTmpObject",
			Handler:    _CollectService_CreateTmpObject_Handler,
		},
		{
			MethodName: "CreateVirussScanResult",
			Handler:    _CollectService_CreateVirussScanResult_Handler,
		},
		{
			MethodName: "CreateExtractMetadataResult",
			Handler:    _CollectService_CreateExtractMetadataResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collect_service.proto",
}
