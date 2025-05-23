// Code generated by MockGen. DO NOT EDIT.
// Source: inventory/inventory_grpc.pb.go

// Package inventory is a generated GoMock package.
package inventory

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockInventoryServiceClient is a mock of InventoryServiceClient interface.
type MockInventoryServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServiceClientMockRecorder
}

// MockInventoryServiceClientMockRecorder is the mock recorder for MockInventoryServiceClient.
type MockInventoryServiceClientMockRecorder struct {
	mock *MockInventoryServiceClient
}

// NewMockInventoryServiceClient creates a new mock instance.
func NewMockInventoryServiceClient(ctrl *gomock.Controller) *MockInventoryServiceClient {
	mock := &MockInventoryServiceClient{ctrl: ctrl}
	mock.recorder = &MockInventoryServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryServiceClient) EXPECT() *MockInventoryServiceClientMockRecorder {
	return m.recorder
}

// CreateInventory mocks base method.
func (m *MockInventoryServiceClient) CreateInventory(ctx context.Context, in *CreateInventoryRequest, opts ...grpc.CallOption) (*CreateInventoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInventory", varargs...)
	ret0, _ := ret[0].(*CreateInventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInventory indicates an expected call of CreateInventory.
func (mr *MockInventoryServiceClientMockRecorder) CreateInventory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInventory", reflect.TypeOf((*MockInventoryServiceClient)(nil).CreateInventory), varargs...)
}

// GetCategories mocks base method.
func (m *MockInventoryServiceClient) GetCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllCategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCategories", varargs...)
	ret0, _ := ret[0].(*AllCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockInventoryServiceClientMockRecorder) GetCategories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockInventoryServiceClient)(nil).GetCategories), varargs...)
}

// GetCategory mocks base method.
func (m *MockInventoryServiceClient) GetCategory(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*CategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCategory", varargs...)
	ret0, _ := ret[0].(*CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockInventoryServiceClientMockRecorder) GetCategory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockInventoryServiceClient)(nil).GetCategory), varargs...)
}

// GetCategorySubcategories mocks base method.
func (m *MockInventoryServiceClient) GetCategorySubcategories(ctx context.Context, in *ResourceId, opts ...grpc.CallOption) (*AllSubCategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCategorySubcategories", varargs...)
	ret0, _ := ret[0].(*AllSubCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategorySubcategories indicates an expected call of GetCategorySubcategories.
func (mr *MockInventoryServiceClientMockRecorder) GetCategorySubcategories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategorySubcategories", reflect.TypeOf((*MockInventoryServiceClient)(nil).GetCategorySubcategories), varargs...)
}

// GetSubCategories mocks base method.
func (m *MockInventoryServiceClient) GetSubCategories(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*AllSubCategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubCategories", varargs...)
	ret0, _ := ret[0].(*AllSubCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCategories indicates an expected call of GetSubCategories.
func (mr *MockInventoryServiceClientMockRecorder) GetSubCategories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCategories", reflect.TypeOf((*MockInventoryServiceClient)(nil).GetSubCategories), varargs...)
}

// GetUsers mocks base method.
func (m *MockInventoryServiceClient) GetUsers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsers", varargs...)
	ret0, _ := ret[0].(*UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockInventoryServiceClientMockRecorder) GetUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockInventoryServiceClient)(nil).GetUsers), varargs...)
}

// RateInventory mocks base method.
func (m *MockInventoryServiceClient) RateInventory(ctx context.Context, in *InventoryRatingRequest, opts ...grpc.CallOption) (*InventoryRatingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RateInventory", varargs...)
	ret0, _ := ret[0].(*InventoryRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RateInventory indicates an expected call of RateInventory.
func (mr *MockInventoryServiceClientMockRecorder) RateInventory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateInventory", reflect.TypeOf((*MockInventoryServiceClient)(nil).RateInventory), varargs...)
}

// RateUser mocks base method.
func (m *MockInventoryServiceClient) RateUser(ctx context.Context, in *UserRatingRequest, opts ...grpc.CallOption) (*UserRatingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RateUser", varargs...)
	ret0, _ := ret[0].(*UserRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RateUser indicates an expected call of RateUser.
func (mr *MockInventoryServiceClientMockRecorder) RateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateUser", reflect.TypeOf((*MockInventoryServiceClient)(nil).RateUser), varargs...)
}

// MockInventoryServiceServer is a mock of InventoryServiceServer interface.
type MockInventoryServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServiceServerMockRecorder
}

// MockInventoryServiceServerMockRecorder is the mock recorder for MockInventoryServiceServer.
type MockInventoryServiceServerMockRecorder struct {
	mock *MockInventoryServiceServer
}

// NewMockInventoryServiceServer creates a new mock instance.
func NewMockInventoryServiceServer(ctrl *gomock.Controller) *MockInventoryServiceServer {
	mock := &MockInventoryServiceServer{ctrl: ctrl}
	mock.recorder = &MockInventoryServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryServiceServer) EXPECT() *MockInventoryServiceServerMockRecorder {
	return m.recorder
}

// CreateInventory mocks base method.
func (m *MockInventoryServiceServer) CreateInventory(arg0 context.Context, arg1 *CreateInventoryRequest) (*CreateInventoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInventory", arg0, arg1)
	ret0, _ := ret[0].(*CreateInventoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInventory indicates an expected call of CreateInventory.
func (mr *MockInventoryServiceServerMockRecorder) CreateInventory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInventory", reflect.TypeOf((*MockInventoryServiceServer)(nil).CreateInventory), arg0, arg1)
}

// GetCategories mocks base method.
func (m *MockInventoryServiceServer) GetCategories(arg0 context.Context, arg1 *EmptyRequest) (*AllCategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", arg0, arg1)
	ret0, _ := ret[0].(*AllCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockInventoryServiceServerMockRecorder) GetCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockInventoryServiceServer)(nil).GetCategories), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockInventoryServiceServer) GetCategory(arg0 context.Context, arg1 *ResourceId) (*CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(*CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockInventoryServiceServerMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockInventoryServiceServer)(nil).GetCategory), arg0, arg1)
}

// GetCategorySubcategories mocks base method.
func (m *MockInventoryServiceServer) GetCategorySubcategories(arg0 context.Context, arg1 *ResourceId) (*AllSubCategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategorySubcategories", arg0, arg1)
	ret0, _ := ret[0].(*AllSubCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategorySubcategories indicates an expected call of GetCategorySubcategories.
func (mr *MockInventoryServiceServerMockRecorder) GetCategorySubcategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategorySubcategories", reflect.TypeOf((*MockInventoryServiceServer)(nil).GetCategorySubcategories), arg0, arg1)
}

// GetSubCategories mocks base method.
func (m *MockInventoryServiceServer) GetSubCategories(arg0 context.Context, arg1 *EmptyRequest) (*AllSubCategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubCategories", arg0, arg1)
	ret0, _ := ret[0].(*AllSubCategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubCategories indicates an expected call of GetSubCategories.
func (mr *MockInventoryServiceServerMockRecorder) GetSubCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubCategories", reflect.TypeOf((*MockInventoryServiceServer)(nil).GetSubCategories), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockInventoryServiceServer) GetUsers(arg0 context.Context, arg1 *EmptyRequest) (*UserListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(*UserListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockInventoryServiceServerMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockInventoryServiceServer)(nil).GetUsers), arg0, arg1)
}

// RateInventory mocks base method.
func (m *MockInventoryServiceServer) RateInventory(arg0 context.Context, arg1 *InventoryRatingRequest) (*InventoryRatingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RateInventory", arg0, arg1)
	ret0, _ := ret[0].(*InventoryRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RateInventory indicates an expected call of RateInventory.
func (mr *MockInventoryServiceServerMockRecorder) RateInventory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateInventory", reflect.TypeOf((*MockInventoryServiceServer)(nil).RateInventory), arg0, arg1)
}

// RateUser mocks base method.
func (m *MockInventoryServiceServer) RateUser(arg0 context.Context, arg1 *UserRatingRequest) (*UserRatingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RateUser", arg0, arg1)
	ret0, _ := ret[0].(*UserRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RateUser indicates an expected call of RateUser.
func (mr *MockInventoryServiceServerMockRecorder) RateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateUser", reflect.TypeOf((*MockInventoryServiceServer)(nil).RateUser), arg0, arg1)
}

// mustEmbedUnimplementedInventoryServiceServer mocks base method.
func (m *MockInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInventoryServiceServer")
}

// mustEmbedUnimplementedInventoryServiceServer indicates an expected call of mustEmbedUnimplementedInventoryServiceServer.
func (mr *MockInventoryServiceServerMockRecorder) mustEmbedUnimplementedInventoryServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInventoryServiceServer", reflect.TypeOf((*MockInventoryServiceServer)(nil).mustEmbedUnimplementedInventoryServiceServer))
}

// MockUnsafeInventoryServiceServer is a mock of UnsafeInventoryServiceServer interface.
type MockUnsafeInventoryServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeInventoryServiceServerMockRecorder
}

// MockUnsafeInventoryServiceServerMockRecorder is the mock recorder for MockUnsafeInventoryServiceServer.
type MockUnsafeInventoryServiceServerMockRecorder struct {
	mock *MockUnsafeInventoryServiceServer
}

// NewMockUnsafeInventoryServiceServer creates a new mock instance.
func NewMockUnsafeInventoryServiceServer(ctrl *gomock.Controller) *MockUnsafeInventoryServiceServer {
	mock := &MockUnsafeInventoryServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeInventoryServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeInventoryServiceServer) EXPECT() *MockUnsafeInventoryServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedInventoryServiceServer mocks base method.
func (m *MockUnsafeInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInventoryServiceServer")
}

// mustEmbedUnimplementedInventoryServiceServer indicates an expected call of mustEmbedUnimplementedInventoryServiceServer.
func (mr *MockUnsafeInventoryServiceServerMockRecorder) mustEmbedUnimplementedInventoryServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInventoryServiceServer", reflect.TypeOf((*MockUnsafeInventoryServiceServer)(nil).mustEmbedUnimplementedInventoryServiceServer))
}
