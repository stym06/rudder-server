// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/config/backend-config (interfaces: BackendConfig)

// Package mock_backendconfig is a generated GoMock package.
package mock_backendconfig

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	utils "github.com/rudderlabs/rudder-server/utils"
)

// MockBackendConfig is a mock of BackendConfig interface.
type MockBackendConfig struct {
	ctrl     *gomock.Controller
	recorder *MockBackendConfigMockRecorder
}

// MockBackendConfigMockRecorder is the mock recorder for MockBackendConfig.
type MockBackendConfigMockRecorder struct {
	mock *MockBackendConfig
}

// NewMockBackendConfig creates a new mock instance.
func NewMockBackendConfig(ctrl *gomock.Controller) *MockBackendConfig {
	mock := &MockBackendConfig{ctrl: ctrl}
	mock.recorder = &MockBackendConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendConfig) EXPECT() *MockBackendConfigMockRecorder {
	return m.recorder
}

// EnhanceConfig mocks base method.
func (m *MockBackendConfig) EnhanceConfig(arg0, arg1 *backendconfig.ConfigT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnhanceConfig", arg0, arg1)
}

// EnhanceConfig indicates an expected call of EnhanceConfig.
func (mr *MockBackendConfigMockRecorder) EnhanceConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnhanceConfig", reflect.TypeOf((*MockBackendConfig)(nil).EnhanceConfig), arg0, arg1)
}

// Get mocks base method.
func (m *MockBackendConfig) Get(arg0 bool, arg1 time.Duration) (backendconfig.ConfigT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(backendconfig.ConfigT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBackendConfigMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBackendConfig)(nil).Get), arg0, arg1)
}

// GetRegulations mocks base method.
func (m *MockBackendConfig) GetRegulations() (backendconfig.RegulationsT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegulations")
	ret0, _ := ret[0].(backendconfig.RegulationsT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetRegulations indicates an expected call of GetRegulations.
func (mr *MockBackendConfigMockRecorder) GetRegulations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegulations", reflect.TypeOf((*MockBackendConfig)(nil).GetRegulations))
}

// GetWorkspaceIDForWriteKey mocks base method.
func (m *MockBackendConfig) GetWorkspaceIDForWriteKey(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceIDForWriteKey", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspaceIDForWriteKey indicates an expected call of GetWorkspaceIDForWriteKey.
func (mr *MockBackendConfigMockRecorder) GetWorkspaceIDForWriteKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceIDForWriteKey", reflect.TypeOf((*MockBackendConfig)(nil).GetWorkspaceIDForWriteKey), arg0)
}

// GetWorkspaceLibrariesForWorkspaceID mocks base method.
func (m *MockBackendConfig) GetWorkspaceLibrariesForWorkspaceID(arg0 string) backendconfig.LibrariesT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceLibrariesForWorkspaceID", arg0)
	ret0, _ := ret[0].(backendconfig.LibrariesT)
	return ret0
}

// GetWorkspaceLibrariesForWorkspaceID indicates an expected call of GetWorkspaceLibrariesForWorkspaceID.
func (mr *MockBackendConfigMockRecorder) GetWorkspaceLibrariesForWorkspaceID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceLibrariesForWorkspaceID", reflect.TypeOf((*MockBackendConfig)(nil).GetWorkspaceLibrariesForWorkspaceID), arg0)
}

// SetUp mocks base method.
func (m *MockBackendConfig) SetUp() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUp")
}

// SetUp indicates an expected call of SetUp.
func (mr *MockBackendConfigMockRecorder) SetUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUp", reflect.TypeOf((*MockBackendConfig)(nil).SetUp))
}

// Subscribe mocks base method.
func (m *MockBackendConfig) Subscribe(arg0 chan utils.DataEvent, arg1 backendconfig.Topic) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Subscribe", arg0, arg1)
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBackendConfigMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBackendConfig)(nil).Subscribe), arg0, arg1)
}

// WaitForConfig mocks base method.
func (m *MockBackendConfig) WaitForConfig() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitForConfig")
}

// WaitForConfig indicates an expected call of WaitForConfig.
func (mr *MockBackendConfigMockRecorder) WaitForConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForConfig", reflect.TypeOf((*MockBackendConfig)(nil).WaitForConfig))
}
