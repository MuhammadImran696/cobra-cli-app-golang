// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todo

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

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Response, error)
	GetTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error)
	DeleteTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*List, error)
	UpdateTask(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error)
	ListTask(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListResponse, error)
	MarkTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Todo/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) GetTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Todo/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) DeleteTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/proto.Todo/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) UpdateTask(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Todo/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) ListTask(ctx context.Context, in *List, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Todo/ListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) MarkTask(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Todo/MarkTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	AddTask(context.Context, *Task) (*Response, error)
	GetTask(context.Context, *Args) (*Response, error)
	DeleteTask(context.Context, *Args) (*List, error)
	UpdateTask(context.Context, *Item) (*Response, error)
	ListTask(context.Context, *List) (*ListResponse, error)
	MarkTask(context.Context, *Args) (*Response, error)
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) AddTask(context.Context, *Task) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedTodoServer) GetTask(context.Context, *Args) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTodoServer) DeleteTask(context.Context, *Args) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTodoServer) UpdateTask(context.Context, *Item) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTodoServer) ListTask(context.Context, *List) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTask not implemented")
}
func (UnimplementedTodoServer) MarkTask(context.Context, *Args) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkTask not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).AddTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTask(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).DeleteTask(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).UpdateTask(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_ListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).ListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/ListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).ListTask(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_MarkTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).MarkTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Todo/MarkTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).MarkTask(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTask",
			Handler:    _Todo_AddTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Todo_GetTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Todo_DeleteTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Todo_UpdateTask_Handler,
		},
		{
			MethodName: "ListTask",
			Handler:    _Todo_ListTask_Handler,
		},
		{
			MethodName: "MarkTask",
			Handler:    _Todo_MarkTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
