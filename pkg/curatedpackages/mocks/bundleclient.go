// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/curatedpackages/bundleclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockBundleClient is a mock of BundleClient interface.
type MockBundleClient struct {
	ctrl     *gomock.Controller
	recorder *MockBundleClientMockRecorder
}

// MockBundleClientMockRecorder is the mock recorder for MockBundleClient.
type MockBundleClientMockRecorder struct {
	mock *MockBundleClient
}

// NewMockBundleClient creates a new mock instance.
func NewMockBundleClient(ctrl *gomock.Controller) *MockBundleClient {
	mock := &MockBundleClient{ctrl: ctrl}
	mock.recorder = &MockBundleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBundleClient) EXPECT() *MockBundleClientMockRecorder {
	return m.recorder
}

// ActiveOrLatest mocks base method.
func (m *MockBundleClient) ActiveOrLatest(ctx context.Context) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveOrLatest", ctx)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveOrLatest indicates an expected call of ActiveOrLatest.
func (mr *MockBundleClientMockRecorder) ActiveOrLatest(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveOrLatest", reflect.TypeOf((*MockBundleClient)(nil).ActiveOrLatest), ctx)
}
