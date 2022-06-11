// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/streamtest.proto

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

// StreamTestClient is the client API for StreamTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamTestClient interface {
	List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamTest_ListClient, error)
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamTest_RecordClient, error)
	Route(ctx context.Context, opts ...grpc.CallOption) (StreamTest_RouteClient, error)
	Route2(ctx context.Context, opts ...grpc.CallOption) (StreamTest_Route2Client, error)
}

type streamTestClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamTestClient(cc grpc.ClientConnInterface) StreamTestClient {
	return &streamTestClient{cc}
}

func (c *streamTestClient) List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamTest_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamTest_ServiceDesc.Streams[0], "/streamtest.StreamTest/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamTestListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamTest_ListClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamTestListClient struct {
	grpc.ClientStream
}

func (x *streamTestListClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamTestClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamTest_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamTest_ServiceDesc.Streams[1], "/streamtest.StreamTest/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamTestRecordClient{stream}
	return x, nil
}

type StreamTest_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamTestRecordClient struct {
	grpc.ClientStream
}

func (x *streamTestRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamTestRecordClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamTestClient) Route(ctx context.Context, opts ...grpc.CallOption) (StreamTest_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamTest_ServiceDesc.Streams[2], "/streamtest.StreamTest/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamTestRouteClient{stream}
	return x, nil
}

type StreamTest_RouteClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamTestRouteClient struct {
	grpc.ClientStream
}

func (x *streamTestRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamTestRouteClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamTestClient) Route2(ctx context.Context, opts ...grpc.CallOption) (StreamTest_Route2Client, error) {
	stream, err := c.cc.NewStream(ctx, &StreamTest_ServiceDesc.Streams[3], "/streamtest.StreamTest/Route2", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamTestRoute2Client{stream}
	return x, nil
}

type StreamTest_Route2Client interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamTestRoute2Client struct {
	grpc.ClientStream
}

func (x *streamTestRoute2Client) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamTestRoute2Client) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamTestServer is the server API for StreamTest service.
// All implementations must embed UnimplementedStreamTestServer
// for forward compatibility
type StreamTestServer interface {
	List(*StreamRequest, StreamTest_ListServer) error
	Record(StreamTest_RecordServer) error
	Route(StreamTest_RouteServer) error
	Route2(StreamTest_Route2Server) error
	mustEmbedUnimplementedStreamTestServer()
}

// UnimplementedStreamTestServer must be embedded to have forward compatible implementations.
type UnimplementedStreamTestServer struct {
}

func (UnimplementedStreamTestServer) List(*StreamRequest, StreamTest_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStreamTestServer) Record(StreamTest_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (UnimplementedStreamTestServer) Route(StreamTest_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedStreamTestServer) Route2(StreamTest_Route2Server) error {
	return status.Errorf(codes.Unimplemented, "method Route2 not implemented")
}
func (UnimplementedStreamTestServer) mustEmbedUnimplementedStreamTestServer() {}

// UnsafeStreamTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamTestServer will
// result in compilation errors.
type UnsafeStreamTestServer interface {
	mustEmbedUnimplementedStreamTestServer()
}

func RegisterStreamTestServer(s grpc.ServiceRegistrar, srv StreamTestServer) {
	s.RegisterService(&StreamTest_ServiceDesc, srv)
}

func _StreamTest_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamTestServer).List(m, &streamTestListServer{stream})
}

type StreamTest_ListServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamTestListServer struct {
	grpc.ServerStream
}

func (x *streamTestListServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamTest_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamTestServer).Record(&streamTestRecordServer{stream})
}

type StreamTest_RecordServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamTestRecordServer struct {
	grpc.ServerStream
}

func (x *streamTestRecordServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamTestRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamTest_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamTestServer).Route(&streamTestRouteServer{stream})
}

type StreamTest_RouteServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamTestRouteServer struct {
	grpc.ServerStream
}

func (x *streamTestRouteServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamTestRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamTest_Route2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamTestServer).Route2(&streamTestRoute2Server{stream})
}

type StreamTest_Route2Server interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamTestRoute2Server struct {
	grpc.ServerStream
}

func (x *streamTestRoute2Server) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamTestRoute2Server) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamTest_ServiceDesc is the grpc.ServiceDesc for StreamTest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamTest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streamtest.StreamTest",
	HandlerType: (*StreamTestServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _StreamTest_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _StreamTest_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _StreamTest_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Route2",
			Handler:       _StreamTest_Route2_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/streamtest.proto",
}
