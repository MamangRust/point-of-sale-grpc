// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: transaction.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransactionService_FindAll_FullMethodName                       = "/pb.TransactionService/FindAll"
	TransactionService_FindByMerchant_FullMethodName                = "/pb.TransactionService/FindByMerchant"
	TransactionService_FindById_FullMethodName                      = "/pb.TransactionService/FindById"
	TransactionService_FindByActive_FullMethodName                  = "/pb.TransactionService/FindByActive"
	TransactionService_FindByTrashed_FullMethodName                 = "/pb.TransactionService/FindByTrashed"
	TransactionService_Create_FullMethodName                        = "/pb.TransactionService/Create"
	TransactionService_Update_FullMethodName                        = "/pb.TransactionService/Update"
	TransactionService_TrashedTransaction_FullMethodName            = "/pb.TransactionService/TrashedTransaction"
	TransactionService_RestoreTransaction_FullMethodName            = "/pb.TransactionService/RestoreTransaction"
	TransactionService_DeleteTransactionPermanent_FullMethodName    = "/pb.TransactionService/DeleteTransactionPermanent"
	TransactionService_RestoreAllTransaction_FullMethodName         = "/pb.TransactionService/RestoreAllTransaction"
	TransactionService_DeleteAllTransactionPermanent_FullMethodName = "/pb.TransactionService/DeleteAllTransactionPermanent"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	FindAll(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransaction, error)
	FindByMerchant(ctx context.Context, in *FindAllTransactionMerchantRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransaction, error)
	FindById(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error)
	FindByActive(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransactionDeleteAt, error)
	FindByTrashed(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransactionDeleteAt, error)
	Create(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error)
	Update(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error)
	TrashedTransaction(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDeleteAt, error)
	RestoreTransaction(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDeleteAt, error)
	DeleteTransactionPermanent(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDelete, error)
	RestoreAllTransaction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseTransactionAll, error)
	DeleteAllTransactionPermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseTransactionAll, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) FindAll(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationTransaction)
	err := c.cc.Invoke(ctx, TransactionService_FindAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) FindByMerchant(ctx context.Context, in *FindAllTransactionMerchantRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationTransaction)
	err := c.cc.Invoke(ctx, TransactionService_FindByMerchant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) FindById(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransaction)
	err := c.cc.Invoke(ctx, TransactionService_FindById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) FindByActive(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransactionDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationTransactionDeleteAt)
	err := c.cc.Invoke(ctx, TransactionService_FindByActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) FindByTrashed(ctx context.Context, in *FindAllTransactionRequest, opts ...grpc.CallOption) (*ApiResponsePaginationTransactionDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponsePaginationTransactionDeleteAt)
	err := c.cc.Invoke(ctx, TransactionService_FindByTrashed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Create(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransaction)
	err := c.cc.Invoke(ctx, TransactionService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) Update(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransaction, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransaction)
	err := c.cc.Invoke(ctx, TransactionService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TrashedTransaction(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransactionDeleteAt)
	err := c.cc.Invoke(ctx, TransactionService_TrashedTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) RestoreTransaction(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDeleteAt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransactionDeleteAt)
	err := c.cc.Invoke(ctx, TransactionService_RestoreTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteTransactionPermanent(ctx context.Context, in *FindByIdTransactionRequest, opts ...grpc.CallOption) (*ApiResponseTransactionDelete, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransactionDelete)
	err := c.cc.Invoke(ctx, TransactionService_DeleteTransactionPermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) RestoreAllTransaction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseTransactionAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransactionAll)
	err := c.cc.Invoke(ctx, TransactionService_RestoreAllTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteAllTransactionPermanent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ApiResponseTransactionAll, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiResponseTransactionAll)
	err := c.cc.Invoke(ctx, TransactionService_DeleteAllTransactionPermanent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility.
type TransactionServiceServer interface {
	FindAll(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransaction, error)
	FindByMerchant(context.Context, *FindAllTransactionMerchantRequest) (*ApiResponsePaginationTransaction, error)
	FindById(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransaction, error)
	FindByActive(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransactionDeleteAt, error)
	FindByTrashed(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransactionDeleteAt, error)
	Create(context.Context, *CreateTransactionRequest) (*ApiResponseTransaction, error)
	Update(context.Context, *UpdateTransactionRequest) (*ApiResponseTransaction, error)
	TrashedTransaction(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDeleteAt, error)
	RestoreTransaction(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDeleteAt, error)
	DeleteTransactionPermanent(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDelete, error)
	RestoreAllTransaction(context.Context, *emptypb.Empty) (*ApiResponseTransactionAll, error)
	DeleteAllTransactionPermanent(context.Context, *emptypb.Empty) (*ApiResponseTransactionAll, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}

func (UnimplementedTransactionServiceServer) FindAll(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedTransactionServiceServer) FindByMerchant(context.Context, *FindAllTransactionMerchantRequest) (*ApiResponsePaginationTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByMerchant not implemented")
}
func (UnimplementedTransactionServiceServer) FindById(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedTransactionServiceServer) FindByActive(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransactionDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByActive not implemented")
}
func (UnimplementedTransactionServiceServer) FindByTrashed(context.Context, *FindAllTransactionRequest) (*ApiResponsePaginationTransactionDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByTrashed not implemented")
}
func (UnimplementedTransactionServiceServer) Create(context.Context, *CreateTransactionRequest) (*ApiResponseTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransactionServiceServer) Update(context.Context, *UpdateTransactionRequest) (*ApiResponseTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransactionServiceServer) TrashedTransaction(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrashedTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) RestoreTransaction(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDeleteAt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteTransactionPermanent(context.Context, *FindByIdTransactionRequest) (*ApiResponseTransactionDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransactionPermanent not implemented")
}
func (UnimplementedTransactionServiceServer) RestoreAllTransaction(context.Context, *emptypb.Empty) (*ApiResponseTransactionAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreAllTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteAllTransactionPermanent(context.Context, *emptypb.Empty) (*ApiResponseTransactionAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllTransactionPermanent not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}
func (UnimplementedTransactionServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FindAll(ctx, req.(*FindAllTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_FindByMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllTransactionMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FindByMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_FindByMerchant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FindByMerchant(ctx, req.(*FindAllTransactionMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_FindById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FindById(ctx, req.(*FindByIdTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_FindByActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FindByActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_FindByActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FindByActive(ctx, req.(*FindAllTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_FindByTrashed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).FindByTrashed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_FindByTrashed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).FindByTrashed(ctx, req.(*FindAllTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Create(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).Update(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_TrashedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).TrashedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_TrashedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).TrashedTransaction(ctx, req.(*FindByIdTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_RestoreTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).RestoreTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_RestoreTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).RestoreTransaction(ctx, req.(*FindByIdTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteTransactionPermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteTransactionPermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_DeleteTransactionPermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteTransactionPermanent(ctx, req.(*FindByIdTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_RestoreAllTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).RestoreAllTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_RestoreAllTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).RestoreAllTransaction(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteAllTransactionPermanent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteAllTransactionPermanent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_DeleteAllTransactionPermanent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteAllTransactionPermanent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _TransactionService_FindAll_Handler,
		},
		{
			MethodName: "FindByMerchant",
			Handler:    _TransactionService_FindByMerchant_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _TransactionService_FindById_Handler,
		},
		{
			MethodName: "FindByActive",
			Handler:    _TransactionService_FindByActive_Handler,
		},
		{
			MethodName: "FindByTrashed",
			Handler:    _TransactionService_FindByTrashed_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TransactionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransactionService_Update_Handler,
		},
		{
			MethodName: "TrashedTransaction",
			Handler:    _TransactionService_TrashedTransaction_Handler,
		},
		{
			MethodName: "RestoreTransaction",
			Handler:    _TransactionService_RestoreTransaction_Handler,
		},
		{
			MethodName: "DeleteTransactionPermanent",
			Handler:    _TransactionService_DeleteTransactionPermanent_Handler,
		},
		{
			MethodName: "RestoreAllTransaction",
			Handler:    _TransactionService_RestoreAllTransaction_Handler,
		},
		{
			MethodName: "DeleteAllTransactionPermanent",
			Handler:    _TransactionService_DeleteAllTransactionPermanent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
