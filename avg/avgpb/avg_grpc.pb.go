// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package avgpb

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

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	Calculate(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalculateClient, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calculate(ctx context.Context, opts ...grpc.CallOption) (CalcService_CalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcService_ServiceDesc.Streams[0], "/avgpb.CalcService/Calculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceCalculateClient{stream}
	return x, nil
}

type CalcService_CalculateClient interface {
	Send(*CalculateRequest) error
	CloseAndRecv() (*CalculateResponse, error)
	grpc.ClientStream
}

type calcServiceCalculateClient struct {
	grpc.ClientStream
}

func (x *calcServiceCalculateClient) Send(m *CalculateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceCalculateClient) CloseAndRecv() (*CalculateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	Calculate(CalcService_CalculateServer) error
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) Calculate(CalcService_CalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_Calculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).Calculate(&calcServiceCalculateServer{stream})
}

type CalcService_CalculateServer interface {
	SendAndClose(*CalculateResponse) error
	Recv() (*CalculateRequest, error)
	grpc.ServerStream
}

type calcServiceCalculateServer struct {
	grpc.ServerStream
}

func (x *calcServiceCalculateServer) SendAndClose(m *CalculateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceCalculateServer) Recv() (*CalculateRequest, error) {
	m := new(CalculateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "avgpb.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Calculate",
			Handler:       _CalcService_Calculate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "avg/avgpb/avg.proto",
}
