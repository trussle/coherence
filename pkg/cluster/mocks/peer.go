// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SimonRichardson/coherence/pkg/cluster (interfaces: Peer)

// Package mocks is a generated GoMock package.
package mocks

import (
	cluster "github.com/SimonRichardson/coherence/pkg/cluster"
	members "github.com/SimonRichardson/coherence/pkg/cluster/members"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPeer is a mock of Peer interface
type MockPeer struct {
	ctrl     *gomock.Controller
	recorder *MockPeerMockRecorder
}

// MockPeerMockRecorder is the mock recorder for MockPeer
type MockPeerMockRecorder struct {
	mock *MockPeer
}

// NewMockPeer creates a new mock instance
func NewMockPeer(ctrl *gomock.Controller) *MockPeer {
	mock := &MockPeer{ctrl: ctrl}
	mock.recorder = &MockPeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeer) EXPECT() *MockPeerMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockPeer) Address() string {
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockPeerMockRecorder) Address() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockPeer)(nil).Address))
}

// Close mocks base method
func (m *MockPeer) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockPeerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPeer)(nil).Close))
}

// ClusterSize mocks base method
func (m *MockPeer) ClusterSize() int {
	ret := m.ctrl.Call(m, "ClusterSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// ClusterSize indicates an expected call of ClusterSize
func (mr *MockPeerMockRecorder) ClusterSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSize", reflect.TypeOf((*MockPeer)(nil).ClusterSize))
}

// Current mocks base method
func (m *MockPeer) Current(arg0 members.PeerType, arg1 bool) ([]string, error) {
	ret := m.ctrl.Call(m, "Current", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Current indicates an expected call of Current
func (mr *MockPeerMockRecorder) Current(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockPeer)(nil).Current), arg0, arg1)
}

// Join mocks base method
func (m *MockPeer) Join() (int, error) {
	ret := m.ctrl.Call(m, "Join")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join
func (mr *MockPeerMockRecorder) Join() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockPeer)(nil).Join))
}

// Leave mocks base method
func (m *MockPeer) Leave() error {
	ret := m.ctrl.Call(m, "Leave")
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave
func (mr *MockPeerMockRecorder) Leave() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockPeer)(nil).Leave))
}

// Listen mocks base method
func (m *MockPeer) Listen(arg0 func(cluster.Reason)) error {
	ret := m.ctrl.Call(m, "Listen", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Listen indicates an expected call of Listen
func (mr *MockPeerMockRecorder) Listen(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockPeer)(nil).Listen), arg0)
}

// Name mocks base method
func (m *MockPeer) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockPeerMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPeer)(nil).Name))
}

// State mocks base method
func (m *MockPeer) State() map[string]interface{} {
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// State indicates an expected call of State
func (mr *MockPeerMockRecorder) State() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockPeer)(nil).State))
}
