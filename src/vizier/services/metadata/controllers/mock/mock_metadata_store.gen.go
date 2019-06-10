// Code generated by MockGen. DO NOT EDIT.
// Source: pixielabs.ai/pixielabs/src/vizier/services/metadata/controllers (interfaces: MetadataStore)

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	gomock "github.com/golang/mock/gomock"
	metadatapb "pixielabs.ai/pixielabs/src/shared/k8s/metadatapb"
	reflect "reflect"
)

// MockMetadataStore is a mock of MetadataStore interface
type MockMetadataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataStoreMockRecorder
}

// MockMetadataStoreMockRecorder is the mock recorder for MockMetadataStore
type MockMetadataStoreMockRecorder struct {
	mock *MockMetadataStore
}

// NewMockMetadataStore creates a new mock instance
func NewMockMetadataStore(ctrl *gomock.Controller) *MockMetadataStore {
	mock := &MockMetadataStore{ctrl: ctrl}
	mock.recorder = &MockMetadataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadataStore) EXPECT() *MockMetadataStoreMockRecorder {
	return m.recorder
}

// UpdateEndpoints mocks base method
func (m *MockMetadataStore) UpdateEndpoints(arg0 *metadatapb.Endpoints) error {
	ret := m.ctrl.Call(m, "UpdateEndpoints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEndpoints indicates an expected call of UpdateEndpoints
func (mr *MockMetadataStoreMockRecorder) UpdateEndpoints(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEndpoints", reflect.TypeOf((*MockMetadataStore)(nil).UpdateEndpoints), arg0)
}

// UpdatePod mocks base method
func (m *MockMetadataStore) UpdatePod(arg0 *metadatapb.Pod) error {
	ret := m.ctrl.Call(m, "UpdatePod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePod indicates an expected call of UpdatePod
func (mr *MockMetadataStoreMockRecorder) UpdatePod(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePod", reflect.TypeOf((*MockMetadataStore)(nil).UpdatePod), arg0)
}

// UpdateService mocks base method
func (m *MockMetadataStore) UpdateService(arg0 *metadatapb.Service) error {
	ret := m.ctrl.Call(m, "UpdateService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockMetadataStoreMockRecorder) UpdateService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockMetadataStore)(nil).UpdateService), arg0)
}
