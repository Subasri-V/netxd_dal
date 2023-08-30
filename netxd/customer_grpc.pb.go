// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: netxd_customer/netxd/customer.proto

package netxd_customer

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

// CustomerDetailsClient is the client API for CustomerDetails service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerDetailsClient interface {
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
}

type customerDetailsClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerDetailsClient(cc grpc.ClientConnInterface) CustomerDetailsClient {
	return &customerDetailsClient{cc}
}

func (c *customerDetailsClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/CustomerDetails/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerDetailsServer is the server API for CustomerDetails service.
// All implementations must embed UnimplementedCustomerDetailsServer
// for forward compatibility
type CustomerDetailsServer interface {
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	mustEmbedUnimplementedCustomerDetailsServer()
}

// UnimplementedCustomerDetailsServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerDetailsServer struct {
}

func (UnimplementedCustomerDetailsServer) CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerDetailsServer) mustEmbedUnimplementedCustomerDetailsServer() {}

// UnsafeCustomerDetailsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerDetailsServer will
// result in compilation errors.
type UnsafeCustomerDetailsServer interface {
	mustEmbedUnimplementedCustomerDetailsServer()
}

func RegisterCustomerDetailsServer(s grpc.ServiceRegistrar, srv CustomerDetailsServer) {
	s.RegisterService(&CustomerDetails_ServiceDesc, srv)
}

func _CustomerDetails_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDetailsServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerDetails/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDetailsServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerDetails_ServiceDesc is the grpc.ServiceDesc for CustomerDetails service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerDetails_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CustomerDetails",
	HandlerType: (*CustomerDetailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerDetails_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netxd_customer/netxd/customer.proto",
}
