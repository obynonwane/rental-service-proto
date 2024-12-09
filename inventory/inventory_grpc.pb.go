// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: inventory/inventory.proto

package inventory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InventoryService_GetUsers_FullMethodName                 = "/inventory.InventoryService/GetUsers"
	InventoryService_CreateInventory_FullMethodName          = "/inventory.InventoryService/CreateInventory"
	InventoryService_GetCategories_FullMethodName            = "/inventory.InventoryService/GetCategories"
	InventoryService_GetSubCategories_FullMethodName         = "/inventory.InventoryService/GetSubCategories"
	InventoryService_GetCategory_FullMethodName              = "/inventory.InventoryService/GetCategory"
	InventoryService_GetCategorySubcategories_FullMethodName = "/inventory.InventoryService/GetCategorySubcategories"
	InventoryService_RateInventory_FullMethodName            = "/inventory.InventoryService/RateInventory"
	InventoryService_RateUser_FullMethodName                 = "/inventory.InventoryService/RateUser"
	InventoryService_GetInventoryByID_FullMethodName         = "/inventory.InventoryService/GetInventoryByID"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	CreateInventory(ctx context.Context, in *CreateInventoryRequest, opts ...grpc.CallOption) (*CreateInventoryResponse, error)
	GetCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllCategoryResponse, error)
	GetSubCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllSubCategoryResponse, error)
	GetCategory(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategorySubcategories(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*AllSubCategoryResponse, error)
	RateInventory(ctx context.Context, in *InventoryRatingRequest, opts ...grpc.CallOption) (*InventoryRatingResponse, error)
	RateUser(ctx context.Context, in *UserRatingRequest, opts ...grpc.CallOption) (*UserRatingResponse, error)
	GetInventoryByID(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*InventoryResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) CreateInventory(ctx context.Context, in *CreateInventoryRequest, opts ...grpc.CallOption) (*CreateInventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateInventoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_CreateInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllCategoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetSubCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllSubCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllSubCategoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetSubCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetCategory(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetCategorySubcategories(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*AllSubCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllSubCategoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetCategorySubcategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) RateInventory(ctx context.Context, in *InventoryRatingRequest, opts ...grpc.CallOption) (*InventoryRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InventoryRatingResponse)
	err := c.cc.Invoke(ctx, InventoryService_RateInventory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) RateUser(ctx context.Context, in *UserRatingRequest, opts ...grpc.CallOption) (*UserRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserRatingResponse)
	err := c.cc.Invoke(ctx, InventoryService_RateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetInventoryByID(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*InventoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InventoryResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetInventoryByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility.
type InventoryServiceServer interface {
	GetUsers(context.Context, *EmptyRequest) (*UserListResponse, error)
	CreateInventory(context.Context, *CreateInventoryRequest) (*CreateInventoryResponse, error)
	GetCategories(context.Context, *EmptyRequest) (*AllCategoryResponse, error)
	GetSubCategories(context.Context, *EmptyRequest) (*AllSubCategoryResponse, error)
	GetCategory(context.Context, *ResourceId) (*CategoryResponse, error)
	GetCategorySubcategories(context.Context, *ResourceId) (*AllSubCategoryResponse, error)
	RateInventory(context.Context, *InventoryRatingRequest) (*InventoryRatingResponse, error)
	RateUser(context.Context, *UserRatingRequest) (*UserRatingResponse, error)
	GetInventoryByID(context.Context, *ResourceId) (*InventoryResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) GetUsers(context.Context, *EmptyRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedInventoryServiceServer) CreateInventory(context.Context, *CreateInventoryRequest) (*CreateInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInventory not implemented")
}
func (UnimplementedInventoryServiceServer) GetCategories(context.Context, *EmptyRequest) (*AllCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedInventoryServiceServer) GetSubCategories(context.Context, *EmptyRequest) (*AllSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategories not implemented")
}
func (UnimplementedInventoryServiceServer) GetCategory(context.Context, *ResourceId) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedInventoryServiceServer) GetCategorySubcategories(context.Context, *ResourceId) (*AllSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategorySubcategories not implemented")
}
func (UnimplementedInventoryServiceServer) RateInventory(context.Context, *InventoryRatingRequest) (*InventoryRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateInventory not implemented")
}
func (UnimplementedInventoryServiceServer) RateUser(context.Context, *UserRatingRequest) (*UserRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateUser not implemented")
}
func (UnimplementedInventoryServiceServer) GetInventoryByID(context.Context, *ResourceId) (*InventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInventoryByID not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedInventoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetUsers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_CreateInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CreateInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CreateInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CreateInventory(ctx, req.(*CreateInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetCategories(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetSubCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetSubCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetSubCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetSubCategories(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetCategory(ctx, req.(*ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetCategorySubcategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetCategorySubcategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetCategorySubcategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetCategorySubcategories(ctx, req.(*ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_RateInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).RateInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_RateInventory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).RateInventory(ctx, req.(*InventoryRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_RateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).RateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_RateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).RateUser(ctx, req.(*UserRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetInventoryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetInventoryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetInventoryByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetInventoryByID(ctx, req.(*ResourceId))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _InventoryService_GetUsers_Handler,
		},
		{
			MethodName: "CreateInventory",
			Handler:    _InventoryService_CreateInventory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _InventoryService_GetCategories_Handler,
		},
		{
			MethodName: "GetSubCategories",
			Handler:    _InventoryService_GetSubCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _InventoryService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategorySubcategories",
			Handler:    _InventoryService_GetCategorySubcategories_Handler,
		},
		{
			MethodName: "RateInventory",
			Handler:    _InventoryService_RateInventory_Handler,
		},
		{
			MethodName: "RateUser",
			Handler:    _InventoryService_RateUser_Handler,
		},
		{
			MethodName: "GetInventoryByID",
			Handler:    _InventoryService_GetInventoryByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/inventory.proto",
}
