// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/curatedpackages/kubectlrunner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	bytes "bytes"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKubectlRunner is a mock of KubectlRunner interface.
type MockKubectlRunner struct {
	ctrl     *gomock.Controller
	recorder *MockKubectlRunnerMockRecorder
}

// MockKubectlRunnerMockRecorder is the mock recorder for MockKubectlRunner.
type MockKubectlRunnerMockRecorder struct {
	mock *MockKubectlRunner
}

// NewMockKubectlRunner creates a new mock instance.
func NewMockKubectlRunner(ctrl *gomock.Controller) *MockKubectlRunner {
	mock := &MockKubectlRunner{ctrl: ctrl}
	mock.recorder = &MockKubectlRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubectlRunner) EXPECT() *MockKubectlRunnerMockRecorder {
	return m.recorder
}

// ExecuteCommand mocks base method.
func (m *MockKubectlRunner) ExecuteCommand(ctx context.Context, opts ...string) (bytes.Buffer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteCommand", varargs...)
	ret0, _ := ret[0].(bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteCommand indicates an expected call of ExecuteCommand.
func (mr *MockKubectlRunnerMockRecorder) ExecuteCommand(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCommand", reflect.TypeOf((*MockKubectlRunner)(nil).ExecuteCommand), varargs...)
}

// ExecuteFromYaml mocks base method.
func (m *MockKubectlRunner) ExecuteFromYaml(ctx context.Context, yaml []byte, opts ...string) (bytes.Buffer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, yaml}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteFromYaml", varargs...)
	ret0, _ := ret[0].(bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteFromYaml indicates an expected call of ExecuteFromYaml.
func (mr *MockKubectlRunnerMockRecorder) ExecuteFromYaml(ctx, yaml interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, yaml}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteFromYaml", reflect.TypeOf((*MockKubectlRunner)(nil).ExecuteFromYaml), varargs...)
}

// HasResource mocks base method.
func (m *MockKubectlRunner) HasResource(ctx context.Context, resourceType, name, kubeconfig, namespace string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasResource", ctx, resourceType, name, kubeconfig, namespace)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasResource indicates an expected call of HasResource.
func (mr *MockKubectlRunnerMockRecorder) HasResource(ctx, resourceType, name, kubeconfig, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasResource", reflect.TypeOf((*MockKubectlRunner)(nil).HasResource), ctx, resourceType, name, kubeconfig, namespace)
}
