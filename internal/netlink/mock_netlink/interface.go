// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/gluetun/internal/netlink (interfaces: NetLinker)

// Package mock_netlink is a generated GoMock package.
package mock_netlink

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
)

// MockNetLinker is a mock of NetLinker interface.
type MockNetLinker struct {
	ctrl     *gomock.Controller
	recorder *MockNetLinkerMockRecorder
}

// MockNetLinkerMockRecorder is the mock recorder for MockNetLinker.
type MockNetLinkerMockRecorder struct {
	mock *MockNetLinker
}

// NewMockNetLinker creates a new mock instance.
func NewMockNetLinker(ctrl *gomock.Controller) *MockNetLinker {
	mock := &MockNetLinker{ctrl: ctrl}
	mock.recorder = &MockNetLinkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetLinker) EXPECT() *MockNetLinkerMockRecorder {
	return m.recorder
}

// AddrAdd mocks base method.
func (m *MockNetLinker) AddrAdd(arg0 netlink.Link, arg1 *netlink.Addr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd.
func (mr *MockNetLinkerMockRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockNetLinker)(nil).AddrAdd), arg0, arg1)
}

// AddrList mocks base method.
func (m *MockNetLinker) AddrList(arg0 netlink.Link, arg1 int) ([]netlink.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddrList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddrList indicates an expected call of AddrList.
func (mr *MockNetLinkerMockRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockNetLinker)(nil).AddrList), arg0, arg1)
}

// LinkAdd mocks base method.
func (m *MockNetLinker) LinkAdd(arg0 netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkAdd indicates an expected call of LinkAdd.
func (mr *MockNetLinkerMockRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkAdd", reflect.TypeOf((*MockNetLinker)(nil).LinkAdd), arg0)
}

// LinkByIndex mocks base method.
func (m *MockNetLinker) LinkByIndex(arg0 int) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByIndex", arg0)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByIndex indicates an expected call of LinkByIndex.
func (mr *MockNetLinkerMockRecorder) LinkByIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByIndex", reflect.TypeOf((*MockNetLinker)(nil).LinkByIndex), arg0)
}

// LinkByName mocks base method.
func (m *MockNetLinker) LinkByName(arg0 string) (netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByName", arg0)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByName indicates an expected call of LinkByName.
func (mr *MockNetLinkerMockRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByName", reflect.TypeOf((*MockNetLinker)(nil).LinkByName), arg0)
}

// LinkDel mocks base method.
func (m *MockNetLinker) LinkDel(arg0 netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDel indicates an expected call of LinkDel.
func (mr *MockNetLinkerMockRecorder) LinkDel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDel", reflect.TypeOf((*MockNetLinker)(nil).LinkDel), arg0)
}

// LinkList mocks base method.
func (m *MockNetLinker) LinkList() ([]netlink.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkList")
	ret0, _ := ret[0].([]netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkList indicates an expected call of LinkList.
func (mr *MockNetLinkerMockRecorder) LinkList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkList", reflect.TypeOf((*MockNetLinker)(nil).LinkList))
}

// LinkSetUp mocks base method.
func (m *MockNetLinker) LinkSetUp(arg0 netlink.Link) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkSetUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp.
func (mr *MockNetLinkerMockRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockNetLinker)(nil).LinkSetUp), arg0)
}

// RouteAdd mocks base method.
func (m *MockNetLinker) RouteAdd(arg0 *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteAdd indicates an expected call of RouteAdd.
func (mr *MockNetLinkerMockRecorder) RouteAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteAdd", reflect.TypeOf((*MockNetLinker)(nil).RouteAdd), arg0)
}

// RouteDel mocks base method.
func (m *MockNetLinker) RouteDel(arg0 *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteDel indicates an expected call of RouteDel.
func (mr *MockNetLinkerMockRecorder) RouteDel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDel", reflect.TypeOf((*MockNetLinker)(nil).RouteDel), arg0)
}

// RouteList mocks base method.
func (m *MockNetLinker) RouteList(arg0 netlink.Link, arg1 int) ([]netlink.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteList indicates an expected call of RouteList.
func (mr *MockNetLinkerMockRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteList", reflect.TypeOf((*MockNetLinker)(nil).RouteList), arg0, arg1)
}

// RouteReplace mocks base method.
func (m *MockNetLinker) RouteReplace(arg0 *netlink.Route) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteReplace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteReplace indicates an expected call of RouteReplace.
func (mr *MockNetLinkerMockRecorder) RouteReplace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteReplace", reflect.TypeOf((*MockNetLinker)(nil).RouteReplace), arg0)
}

// RuleAdd mocks base method.
func (m *MockNetLinker) RuleAdd(arg0 *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd.
func (mr *MockNetLinkerMockRecorder) RuleAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockNetLinker)(nil).RuleAdd), arg0)
}

// RuleDel mocks base method.
func (m *MockNetLinker) RuleDel(arg0 *netlink.Rule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleDel indicates an expected call of RuleDel.
func (mr *MockNetLinkerMockRecorder) RuleDel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleDel", reflect.TypeOf((*MockNetLinker)(nil).RuleDel), arg0)
}

// RuleList mocks base method.
func (m *MockNetLinker) RuleList(arg0 int) ([]netlink.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuleList", arg0)
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RuleList indicates an expected call of RuleList.
func (mr *MockNetLinkerMockRecorder) RuleList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleList", reflect.TypeOf((*MockNetLinker)(nil).RuleList), arg0)
}
