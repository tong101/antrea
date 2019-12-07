// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/agent/openflow (interfaces: Client)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/vmware-tanzu/antrea/pkg/agent/types"
	openflow "github.com/vmware-tanzu/antrea/pkg/ovs/openflow"
	net "net"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddPolicyRuleAddress mocks base method
func (m *MockClient) AddPolicyRuleAddress(arg0 uint32, arg1 types.AddressType, arg2 []types.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicyRuleAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPolicyRuleAddress indicates an expected call of AddPolicyRuleAddress
func (mr *MockClientMockRecorder) AddPolicyRuleAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicyRuleAddress", reflect.TypeOf((*MockClient)(nil).AddPolicyRuleAddress), arg0, arg1, arg2)
}

// DeletePolicyRuleAddress mocks base method
func (m *MockClient) DeletePolicyRuleAddress(arg0 uint32, arg1 types.AddressType, arg2 []types.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicyRuleAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicyRuleAddress indicates an expected call of DeletePolicyRuleAddress
func (mr *MockClientMockRecorder) DeletePolicyRuleAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicyRuleAddress", reflect.TypeOf((*MockClient)(nil).DeletePolicyRuleAddress), arg0, arg1, arg2)
}

// Disconnect mocks base method
func (m *MockClient) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockClientMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClient)(nil).Disconnect))
}

// GetFlowTableStatus mocks base method
func (m *MockClient) GetFlowTableStatus() []openflow.TableStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowTableStatus")
	ret0, _ := ret[0].([]openflow.TableStatus)
	return ret0
}

// GetFlowTableStatus indicates an expected call of GetFlowTableStatus
func (mr *MockClientMockRecorder) GetFlowTableStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowTableStatus", reflect.TypeOf((*MockClient)(nil).GetFlowTableStatus))
}

// Initialize mocks base method
func (m *MockClient) Initialize() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockClientMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockClient)(nil).Initialize))
}

// InstallClusterServiceCIDRFlows mocks base method
func (m *MockClient) InstallClusterServiceCIDRFlows(arg0 *net.IPNet, arg1 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallClusterServiceCIDRFlows", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallClusterServiceCIDRFlows indicates an expected call of InstallClusterServiceCIDRFlows
func (mr *MockClientMockRecorder) InstallClusterServiceCIDRFlows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallClusterServiceCIDRFlows", reflect.TypeOf((*MockClient)(nil).InstallClusterServiceCIDRFlows), arg0, arg1)
}

// InstallDefaultTunnelFlows mocks base method
func (m *MockClient) InstallDefaultTunnelFlows(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallDefaultTunnelFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDefaultTunnelFlows indicates an expected call of InstallDefaultTunnelFlows
func (mr *MockClientMockRecorder) InstallDefaultTunnelFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDefaultTunnelFlows", reflect.TypeOf((*MockClient)(nil).InstallDefaultTunnelFlows), arg0)
}

// InstallGatewayFlows mocks base method
func (m *MockClient) InstallGatewayFlows(arg0 net.IP, arg1 net.HardwareAddr, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallGatewayFlows", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallGatewayFlows indicates an expected call of InstallGatewayFlows
func (mr *MockClientMockRecorder) InstallGatewayFlows(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallGatewayFlows", reflect.TypeOf((*MockClient)(nil).InstallGatewayFlows), arg0, arg1, arg2)
}

// InstallNodeFlows mocks base method
func (m *MockClient) InstallNodeFlows(arg0 string, arg1 net.HardwareAddr, arg2 net.IP, arg3 net.IPNet, arg4 net.IP, arg5 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallNodeFlows", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallNodeFlows indicates an expected call of InstallNodeFlows
func (mr *MockClientMockRecorder) InstallNodeFlows(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallNodeFlows", reflect.TypeOf((*MockClient)(nil).InstallNodeFlows), arg0, arg1, arg2, arg3, arg4, arg5)
}

// InstallPodFlows mocks base method
func (m *MockClient) InstallPodFlows(arg0 string, arg1 net.IP, arg2, arg3 net.HardwareAddr, arg4 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPodFlows", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallPodFlows indicates an expected call of InstallPodFlows
func (mr *MockClientMockRecorder) InstallPodFlows(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPodFlows", reflect.TypeOf((*MockClient)(nil).InstallPodFlows), arg0, arg1, arg2, arg3, arg4)
}

// InstallPolicyRuleFlows mocks base method
func (m *MockClient) InstallPolicyRuleFlows(arg0 *types.PolicyRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallPolicyRuleFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallPolicyRuleFlows indicates an expected call of InstallPolicyRuleFlows
func (mr *MockClientMockRecorder) InstallPolicyRuleFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallPolicyRuleFlows", reflect.TypeOf((*MockClient)(nil).InstallPolicyRuleFlows), arg0)
}

// UninstallNodeFlows mocks base method
func (m *MockClient) UninstallNodeFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallNodeFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallNodeFlows indicates an expected call of UninstallNodeFlows
func (mr *MockClientMockRecorder) UninstallNodeFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallNodeFlows", reflect.TypeOf((*MockClient)(nil).UninstallNodeFlows), arg0)
}

// UninstallPodFlows mocks base method
func (m *MockClient) UninstallPodFlows(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPodFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallPodFlows indicates an expected call of UninstallPodFlows
func (mr *MockClientMockRecorder) UninstallPodFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPodFlows", reflect.TypeOf((*MockClient)(nil).UninstallPodFlows), arg0)
}

// UninstallPolicyRuleFlows mocks base method
func (m *MockClient) UninstallPolicyRuleFlows(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallPolicyRuleFlows", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallPolicyRuleFlows indicates an expected call of UninstallPolicyRuleFlows
func (mr *MockClientMockRecorder) UninstallPolicyRuleFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallPolicyRuleFlows", reflect.TypeOf((*MockClient)(nil).UninstallPolicyRuleFlows), arg0)
}
