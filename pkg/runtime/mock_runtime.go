// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/arana-db/arana/pkg/runtime (interfaces: Runtime)

// Package runtime is a generated GoMock package.
package runtime

import (
	context "context"
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

import (
	proto "github.com/arana-db/arana/pkg/proto"
	namespace "github.com/arana-db/arana/pkg/runtime/namespace"
)

// MockRuntime is a mock of Runtime interface.
type MockRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockRuntimeMockRecorder
}

// MockRuntimeMockRecorder is the mock recorder for MockRuntime.
type MockRuntimeMockRecorder struct {
	mock *MockRuntime
}

// NewMockRuntime creates a new mock instance.
func NewMockRuntime(ctrl *gomock.Controller) *MockRuntime {
	mock := &MockRuntime{ctrl: ctrl}
	mock.recorder = &MockRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuntime) EXPECT() *MockRuntimeMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockRuntime) Begin(arg0 context.Context, arg1 ...TxHook) (proto.Tx, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Begin", varargs...)
	ret0, _ := ret[0].(proto.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockRuntimeMockRecorder) Begin(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockRuntime)(nil).Begin), varargs...)
}

// Exec mocks base method.
func (m *MockRuntime) Exec(arg0 context.Context, arg1, arg2 string, arg3 ...proto.Value) (proto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockRuntimeMockRecorder) Exec(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockRuntime)(nil).Exec), varargs...)
}

// Execute mocks base method.
func (m *MockRuntime) Execute(arg0 *proto.Context) (proto.Result, uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(uint16)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Execute indicates an expected call of Execute.
func (mr *MockRuntimeMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRuntime)(nil).Execute), arg0)
}

// Namespace mocks base method.
func (m *MockRuntime) Namespace() *namespace.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(*namespace.Namespace)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockRuntimeMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockRuntime)(nil).Namespace))
}

// Query mocks base method.
func (m *MockRuntime) Query(arg0 context.Context, arg1, arg2 string, arg3 ...proto.Value) (proto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockRuntimeMockRecorder) Query(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRuntime)(nil).Query), varargs...)
}
