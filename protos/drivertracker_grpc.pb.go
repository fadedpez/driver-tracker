// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// DriverTrackerAPIClient is the client API for DriverTrackerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverTrackerAPIClient interface {
	StoreDriver(ctx context.Context, in *StoreDriverRequest, opts ...grpc.CallOption) (*StoreDriverResponse, error)
}

type driverTrackerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverTrackerAPIClient(cc grpc.ClientConnInterface) DriverTrackerAPIClient {
	return &driverTrackerAPIClient{cc}
}

func (c *driverTrackerAPIClient) StoreDriver(ctx context.Context, in *StoreDriverRequest, opts ...grpc.CallOption) (*StoreDriverResponse, error) {
	out := new(StoreDriverResponse)
	err := c.cc.Invoke(ctx, "/drivertracker.api.alpha.DriverTrackerAPI/StoreDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverTrackerAPIServer is the server API for DriverTrackerAPI service.
// All implementations must embed UnimplementedDriverTrackerAPIServer
// for forward compatibility
type DriverTrackerAPIServer interface {
	StoreDriver(context.Context, *StoreDriverRequest) (*StoreDriverResponse, error)
}

// UnimplementedDriverTrackerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedDriverTrackerAPIServer struct {
}

func (UnimplementedDriverTrackerAPIServer) StoreDriver(context.Context, *StoreDriverRequest) (*StoreDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDriver not implemented")
}
func (UnimplementedDriverTrackerAPIServer) mustEmbedUnimplementedDriverTrackerAPIServer() {}

// UnsafeDriverTrackerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverTrackerAPIServer will
// result in compilation errors.
type UnsafeDriverTrackerAPIServer interface {
	mustEmbedUnimplementedDriverTrackerAPIServer()
}

func RegisterDriverTrackerAPIServer(s grpc.ServiceRegistrar, srv DriverTrackerAPIServer) {
	s.RegisterService(&DriverTrackerAPI_ServiceDesc, srv)
}

func _DriverTrackerAPI_StoreDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverTrackerAPIServer).StoreDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drivertracker.api.alpha.DriverTrackerAPI/StoreDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverTrackerAPIServer).StoreDriver(ctx, req.(*StoreDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverTrackerAPI_ServiceDesc is the grpc.ServiceDesc for DriverTrackerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverTrackerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drivertracker.api.alpha.DriverTrackerAPI",
	HandlerType: (*DriverTrackerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreDriver",
			Handler:    _DriverTrackerAPI_StoreDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drivertracker.proto",
}
