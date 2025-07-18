// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: context.proto

package mcp

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
	ModelContextService_GetContext_FullMethodName = "/mcp.ModelContextService/GetContext"
)

// ModelContextServiceClient is the client API for ModelContextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelContextServiceClient interface {
	GetContext(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*ContextResponse, error)
}

type modelContextServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelContextServiceClient(cc grpc.ClientConnInterface) ModelContextServiceClient {
	return &modelContextServiceClient{cc}
}

func (c *modelContextServiceClient) GetContext(ctx context.Context, in *ContextRequest, opts ...grpc.CallOption) (*ContextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ContextResponse)
	err := c.cc.Invoke(ctx, ModelContextService_GetContext_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelContextServiceServer is the server API for ModelContextService service.
// All implementations must embed UnimplementedModelContextServiceServer
// for forward compatibility.
type ModelContextServiceServer interface {
	GetContext(context.Context, *ContextRequest) (*ContextResponse, error)
	mustEmbedUnimplementedModelContextServiceServer()
}

// UnimplementedModelContextServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModelContextServiceServer struct{}

func (UnimplementedModelContextServiceServer) GetContext(context.Context, *ContextRequest) (*ContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContext not implemented")
}
func (UnimplementedModelContextServiceServer) mustEmbedUnimplementedModelContextServiceServer() {}
func (UnimplementedModelContextServiceServer) testEmbeddedByValue()                             {}

// UnsafeModelContextServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelContextServiceServer will
// result in compilation errors.
type UnsafeModelContextServiceServer interface {
	mustEmbedUnimplementedModelContextServiceServer()
}

func RegisterModelContextServiceServer(s grpc.ServiceRegistrar, srv ModelContextServiceServer) {
	// If the following call pancis, it indicates UnimplementedModelContextServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModelContextService_ServiceDesc, srv)
}

func _ModelContextService_GetContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelContextServiceServer).GetContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelContextService_GetContext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelContextServiceServer).GetContext(ctx, req.(*ContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelContextService_ServiceDesc is the grpc.ServiceDesc for ModelContextService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelContextService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcp.ModelContextService",
	HandlerType: (*ModelContextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContext",
			Handler:    _ModelContextService_GetContext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "context.proto",
}
