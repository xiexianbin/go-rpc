// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: demo/v1/stream.proto

package demov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StreamService_List_FullMethodName   = "/demo.v1.StreamService/List"
	StreamService_Record_FullMethodName = "/demo.v1.StreamService/Record"
	StreamService_Route_FullMethodName  = "/demo.v1.StreamService/Route"
)

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	// 服务器端流式 RPC
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (StreamService_ListClient, error)
	// 客户端流式 RPC
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error)
	// 双向流式 RPC
	Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], StreamService_List_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], StreamService_Record_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordClient{ClientStream: stream}
	return x, nil
}

type StreamService_RecordClient interface {
	Send(*RecordRequest) error
	CloseAndRecv() (*RecordResponse, error)
	grpc.ClientStream
}

type streamServiceRecordClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordClient) Send(m *RecordRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordClient) CloseAndRecv() (*RecordResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RecordResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], StreamService_Route_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRouteClient{ClientStream: stream}
	return x, nil
}

type StreamService_RouteClient interface {
	Send(*RouteRequest) error
	Recv() (*RouteResponse, error)
	grpc.ClientStream
}

type streamServiceRouteClient struct {
	grpc.ClientStream
}

func (x *streamServiceRouteClient) Send(m *RouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRouteClient) Recv() (*RouteResponse, error) {
	m := new(RouteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	// 服务器端流式 RPC
	List(*ListRequest, StreamService_ListServer) error
	// 客户端流式 RPC
	Record(StreamService_RecordServer) error
	// 双向流式 RPC
	Route(StreamService_RouteServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) List(*ListRequest, StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStreamServiceServer) Record(StreamService_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (UnimplementedStreamServiceServer) Route(StreamService_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{ServerStream: stream})
}

type StreamService_ListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Record(&streamServiceRecordServer{ServerStream: stream})
}

type StreamService_RecordServer interface {
	SendAndClose(*RecordResponse) error
	Recv() (*RecordRequest, error)
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) SendAndClose(m *RecordResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*RecordRequest, error) {
	m := new(RecordRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Route(&streamServiceRouteServer{ServerStream: stream})
}

type StreamService_RouteServer interface {
	Send(*RouteResponse) error
	Recv() (*RouteRequest, error)
	grpc.ServerStream
}

type streamServiceRouteServer struct {
	grpc.ServerStream
}

func (x *streamServiceRouteServer) Send(m *RouteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRouteServer) Recv() (*RouteRequest, error) {
	m := new(RouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.v1.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _StreamService_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _StreamService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo/v1/stream.proto",
}