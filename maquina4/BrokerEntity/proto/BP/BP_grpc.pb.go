// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package broker

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

// BrokerPrincesaServiceClient is the client API for BrokerPrincesaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerPrincesaServiceClient interface {
	RequestRebeldes(ctx context.Context, in *RebeldesReq, opts ...grpc.CallOption) (*RebeldesResp, error)
}

type brokerPrincesaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerPrincesaServiceClient(cc grpc.ClientConnInterface) BrokerPrincesaServiceClient {
	return &brokerPrincesaServiceClient{cc}
}

func (c *brokerPrincesaServiceClient) RequestRebeldes(ctx context.Context, in *RebeldesReq, opts ...grpc.CallOption) (*RebeldesResp, error) {
	out := new(RebeldesResp)
	err := c.cc.Invoke(ctx, "/BP.BrokerPrincesaService/RequestRebeldes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerPrincesaServiceServer is the server API for BrokerPrincesaService service.
// All implementations must embed UnimplementedBrokerPrincesaServiceServer
// for forward compatibility
type BrokerPrincesaServiceServer interface {
	RequestRebeldes(context.Context, *RebeldesReq) (*RebeldesResp, error)
	mustEmbedUnimplementedBrokerPrincesaServiceServer()
}

// UnimplementedBrokerPrincesaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerPrincesaServiceServer struct {
}

func (UnimplementedBrokerPrincesaServiceServer) RequestRebeldes(context.Context, *RebeldesReq) (*RebeldesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRebeldes not implemented")
}
func (UnimplementedBrokerPrincesaServiceServer) mustEmbedUnimplementedBrokerPrincesaServiceServer() {}

// UnsafeBrokerPrincesaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerPrincesaServiceServer will
// result in compilation errors.
type UnsafeBrokerPrincesaServiceServer interface {
	mustEmbedUnimplementedBrokerPrincesaServiceServer()
}

func RegisterBrokerPrincesaServiceServer(s grpc.ServiceRegistrar, srv BrokerPrincesaServiceServer) {
	s.RegisterService(&BrokerPrincesaService_ServiceDesc, srv)
}

func _BrokerPrincesaService_RequestRebeldes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebeldesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerPrincesaServiceServer).RequestRebeldes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BP.BrokerPrincesaService/RequestRebeldes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerPrincesaServiceServer).RequestRebeldes(ctx, req.(*RebeldesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerPrincesaService_ServiceDesc is the grpc.ServiceDesc for BrokerPrincesaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerPrincesaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BP.BrokerPrincesaService",
	HandlerType: (*BrokerPrincesaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRebeldes",
			Handler:    _BrokerPrincesaService_RequestRebeldes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "BP.proto",
}
