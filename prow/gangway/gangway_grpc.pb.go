// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: gangway.proto

package gangway

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

// ProwClient is the client API for Prow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProwClient interface {
	// FIXME: In the future we can just return a unique token (only), in the same
	// way that GCB returns immediately with the globally-unique BuildId. That is,
	// in the future the response will be a union of either the full JobExecution
	// message or a single JobExecutionToken (string). See
	// https://docs.google.com/document/d/1v77jp1Nb5C2C2-PdV02SGViO9CyZ9SvNxCPOHyIUQeo/edit#bookmark=id.q68srxklvpt4.
	CreateJobExecution(ctx context.Context, in *CreateJobExecutionRequest, opts ...grpc.CallOption) (*JobExecution, error)
	GetJobExecution(ctx context.Context, in *GetJobExecutionRequest, opts ...grpc.CallOption) (*JobExecution, error)
	ListJobExecutions(ctx context.Context, in *ListJobExecutionsRequest, opts ...grpc.CallOption) (*JobExecutions, error)
}

type prowClient struct {
	cc grpc.ClientConnInterface
}

func NewProwClient(cc grpc.ClientConnInterface) ProwClient {
	return &prowClient{cc}
}

func (c *prowClient) CreateJobExecution(ctx context.Context, in *CreateJobExecutionRequest, opts ...grpc.CallOption) (*JobExecution, error) {
	out := new(JobExecution)
	err := c.cc.Invoke(ctx, "/Prow/CreateJobExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prowClient) GetJobExecution(ctx context.Context, in *GetJobExecutionRequest, opts ...grpc.CallOption) (*JobExecution, error) {
	out := new(JobExecution)
	err := c.cc.Invoke(ctx, "/Prow/GetJobExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prowClient) ListJobExecutions(ctx context.Context, in *ListJobExecutionsRequest, opts ...grpc.CallOption) (*JobExecutions, error) {
	out := new(JobExecutions)
	err := c.cc.Invoke(ctx, "/Prow/ListJobExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProwServer is the server API for Prow service.
// All implementations must embed UnimplementedProwServer
// for forward compatibility
type ProwServer interface {
	// FIXME: In the future we can just return a unique token (only), in the same
	// way that GCB returns immediately with the globally-unique BuildId. That is,
	// in the future the response will be a union of either the full JobExecution
	// message or a single JobExecutionToken (string). See
	// https://docs.google.com/document/d/1v77jp1Nb5C2C2-PdV02SGViO9CyZ9SvNxCPOHyIUQeo/edit#bookmark=id.q68srxklvpt4.
	CreateJobExecution(context.Context, *CreateJobExecutionRequest) (*JobExecution, error)
	GetJobExecution(context.Context, *GetJobExecutionRequest) (*JobExecution, error)
	ListJobExecutions(context.Context, *ListJobExecutionsRequest) (*JobExecutions, error)
	mustEmbedUnimplementedProwServer()
}

// UnimplementedProwServer must be embedded to have forward compatible implementations.
type UnimplementedProwServer struct {
}

func (UnimplementedProwServer) CreateJobExecution(context.Context, *CreateJobExecutionRequest) (*JobExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJobExecution not implemented")
}
func (UnimplementedProwServer) GetJobExecution(context.Context, *GetJobExecutionRequest) (*JobExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobExecution not implemented")
}
func (UnimplementedProwServer) ListJobExecutions(context.Context, *ListJobExecutionsRequest) (*JobExecutions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobExecutions not implemented")
}
func (UnimplementedProwServer) mustEmbedUnimplementedProwServer() {}

// UnsafeProwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProwServer will
// result in compilation errors.
type UnsafeProwServer interface {
	mustEmbedUnimplementedProwServer()
}

func RegisterProwServer(s grpc.ServiceRegistrar, srv ProwServer) {
	s.RegisterService(&Prow_ServiceDesc, srv)
}

func _Prow_CreateJobExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProwServer).CreateJobExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Prow/CreateJobExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProwServer).CreateJobExecution(ctx, req.(*CreateJobExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Prow_GetJobExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProwServer).GetJobExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Prow/GetJobExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProwServer).GetJobExecution(ctx, req.(*GetJobExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Prow_ListJobExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProwServer).ListJobExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Prow/ListJobExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProwServer).ListJobExecutions(ctx, req.(*ListJobExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Prow_ServiceDesc is the grpc.ServiceDesc for Prow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Prow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Prow",
	HandlerType: (*ProwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobExecution",
			Handler:    _Prow_CreateJobExecution_Handler,
		},
		{
			MethodName: "GetJobExecution",
			Handler:    _Prow_GetJobExecution_Handler,
		},
		{
			MethodName: "ListJobExecutions",
			Handler:    _Prow_ListJobExecutions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gangway.proto",
}