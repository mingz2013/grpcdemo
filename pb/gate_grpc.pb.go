// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/gate.proto

package pb

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

// GateClient is the client API for Gate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GateClient interface {
	// Sends a greeting
	Route(ctx context.Context, opts ...grpc.CallOption) (Gate_RouteClient, error)
}

type gateClient struct {
	cc grpc.ClientConnInterface
}

func NewGateClient(cc grpc.ClientConnInterface) GateClient {
	return &gateClient{cc}
}

func (c *gateClient) Route(ctx context.Context, opts ...grpc.CallOption) (Gate_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gate_ServiceDesc.Streams[0], "/gate.Gate/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &gateRouteClient{stream}
	return x, nil
}

type Gate_RouteClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type gateRouteClient struct {
	grpc.ClientStream
}

func (x *gateRouteClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gateRouteClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GateServer is the server API for Gate service.
// All implementations must embed UnimplementedGateServer
// for forward compatibility
type GateServer interface {
	// Sends a greeting
	Route(Gate_RouteServer) error
	mustEmbedUnimplementedGateServer()
}

// UnimplementedGateServer must be embedded to have forward compatible implementations.
type UnimplementedGateServer struct {
}

func (UnimplementedGateServer) Route(Gate_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedGateServer) mustEmbedUnimplementedGateServer() {}

// UnsafeGateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GateServer will
// result in compilation errors.
type UnsafeGateServer interface {
	mustEmbedUnimplementedGateServer()
}

func RegisterGateServer(s grpc.ServiceRegistrar, srv GateServer) {
	s.RegisterService(&Gate_ServiceDesc, srv)
}

func _Gate_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GateServer).Route(&gateRouteServer{stream})
}

type Gate_RouteServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type gateRouteServer struct {
	grpc.ServerStream
}

func (x *gateRouteServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gateRouteServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Gate_ServiceDesc is the grpc.ServiceDesc for Gate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gate.Gate",
	HandlerType: (*GateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Route",
			Handler:       _Gate_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/gate.proto",
}