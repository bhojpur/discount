// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// DiscountUIClient is the client API for DiscountUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountUIClient interface {
	// ListEngineSpecs returns a list of Discount Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (DiscountUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type discountUIClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountUIClient(cc grpc.ClientConnInterface) DiscountUIClient {
	return &discountUIClient{cc}
}

func (c *discountUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (DiscountUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &DiscountUI_ServiceDesc.Streams[0], "/v1.DiscountUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &discountUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DiscountUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type discountUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *discountUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *discountUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.DiscountUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountUIServer is the server API for DiscountUI service.
// All implementations must embed UnimplementedDiscountUIServer
// for forward compatibility
type DiscountUIServer interface {
	// ListEngineSpecs returns a list of Discount Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, DiscountUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedDiscountUIServer()
}

// UnimplementedDiscountUIServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountUIServer struct {
}

func (UnimplementedDiscountUIServer) ListEngineSpecs(*ListEngineSpecsRequest, DiscountUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedDiscountUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedDiscountUIServer) mustEmbedUnimplementedDiscountUIServer() {}

// UnsafeDiscountUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountUIServer will
// result in compilation errors.
type UnsafeDiscountUIServer interface {
	mustEmbedUnimplementedDiscountUIServer()
}

func RegisterDiscountUIServer(s grpc.ServiceRegistrar, srv DiscountUIServer) {
	s.RegisterService(&DiscountUI_ServiceDesc, srv)
}

func _DiscountUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiscountUIServer).ListEngineSpecs(m, &discountUIListEngineSpecsServer{stream})
}

type DiscountUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type discountUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *discountUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DiscountUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DiscountUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscountUI_ServiceDesc is the grpc.ServiceDesc for DiscountUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscountUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DiscountUI",
	HandlerType: (*DiscountUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _DiscountUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _DiscountUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "discount-ui.proto",
}
