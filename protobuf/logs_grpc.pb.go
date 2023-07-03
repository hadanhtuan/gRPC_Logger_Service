// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: protobuf/logs.proto

package protobuf

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

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	WriteLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	GetLogByName(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogService_GetLogByNameClient, error)
	CountLogByName(ctx context.Context, opts ...grpc.CallOption) (LogService_CountLogByNameClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) WriteLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/protobuf.LogService/WriteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetLogByName(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogService_GetLogByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], "/protobuf.LogService/GetLogByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceGetLogByNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_GetLogByNameClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type logServiceGetLogByNameClient struct {
	grpc.ClientStream
}

func (x *logServiceGetLogByNameClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) CountLogByName(ctx context.Context, opts ...grpc.CallOption) (LogService_CountLogByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[1], "/protobuf.LogService/CountLogByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceCountLogByNameClient{stream}
	return x, nil
}

type LogService_CountLogByNameClient interface {
	Send(*LogRequest) error
	CloseAndRecv() (*CountLog, error)
	grpc.ClientStream
}

type logServiceCountLogByNameClient struct {
	grpc.ClientStream
}

func (x *logServiceCountLogByNameClient) Send(m *LogRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logServiceCountLogByNameClient) CloseAndRecv() (*CountLog, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CountLog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	WriteLog(context.Context, *LogRequest) (*LogResponse, error)
	GetLogByName(*LogRequest, LogService_GetLogByNameServer) error
	CountLogByName(LogService_CountLogByNameServer) error
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) WriteLog(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (UnimplementedLogServiceServer) GetLogByName(*LogRequest, LogService_GetLogByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogByName not implemented")
}
func (UnimplementedLogServiceServer) CountLogByName(LogService_CountLogByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method CountLogByName not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_WriteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).WriteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.LogService/WriteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).WriteLog(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetLogByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).GetLogByName(m, &logServiceGetLogByNameServer{stream})
}

type LogService_GetLogByNameServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type logServiceGetLogByNameServer struct {
	grpc.ServerStream
}

func (x *logServiceGetLogByNameServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_CountLogByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).CountLogByName(&logServiceCountLogByNameServer{stream})
}

type LogService_CountLogByNameServer interface {
	SendAndClose(*CountLog) error
	Recv() (*LogRequest, error)
	grpc.ServerStream
}

type logServiceCountLogByNameServer struct {
	grpc.ServerStream
}

func (x *logServiceCountLogByNameServer) SendAndClose(m *CountLog) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logServiceCountLogByNameServer) Recv() (*LogRequest, error) {
	m := new(LogRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteLog",
			Handler:    _LogService_WriteLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogByName",
			Handler:       _LogService_GetLogByName_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CountLogByName",
			Handler:       _LogService_CountLogByName_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/logs.proto",
}