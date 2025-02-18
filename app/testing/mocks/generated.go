// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pjb7687/phabulous/app/interfaces (interfaces: Bot,Message,HandlerTuple,Module,Connector)

// Package mocks is a generated GoMock package.
package mocks

import (
	gonduit "github.com/etcinit/gonduit"
	interfaces "github.com/pjb7687/phabulous/app/interfaces"
	messages "github.com/pjb7687/phabulous/app/messages"
	gomock "github.com/golang/mock/gomock"
	confer "github.com/jacobstr/confer"
	reflect "reflect"
	regexp "regexp"
)

// MockBot is a mock of Bot interface
type MockBot struct {
	ctrl     *gomock.Controller
	recorder *MockBotMockRecorder
}

// MockBotMockRecorder is the mock recorder for MockBot
type MockBotMockRecorder struct {
	mock *MockBot
}

// NewMockBot creates a new mock instance
func NewMockBot(ctrl *gomock.Controller) *MockBot {
	mock := &MockBot{ctrl: ctrl}
	mock.recorder = &MockBotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBot) EXPECT() *MockBotMockRecorder {
	return m.recorder
}

// Excuse mocks base method
func (m *MockBot) Excuse(arg0 interfaces.Message, arg1 error) {
	m.ctrl.Call(m, "Excuse", arg0, arg1)
}

// Excuse indicates an expected call of Excuse
func (mr *MockBotMockRecorder) Excuse(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Excuse", reflect.TypeOf((*MockBot)(nil).Excuse), arg0, arg1)
}

// GetConfig mocks base method
func (m *MockBot) GetConfig() *confer.Config {
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*confer.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockBotMockRecorder) GetConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBot)(nil).GetConfig))
}

// GetGonduit mocks base method
func (m *MockBot) GetGonduit() (*gonduit.Conn, error) {
	ret := m.ctrl.Call(m, "GetGonduit")
	ret0, _ := ret[0].(*gonduit.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGonduit indicates an expected call of GetGonduit
func (mr *MockBotMockRecorder) GetGonduit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGonduit", reflect.TypeOf((*MockBot)(nil).GetGonduit))
}

// GetHandlers mocks base method
func (m *MockBot) GetHandlers() []interfaces.HandlerTuple {
	ret := m.ctrl.Call(m, "GetHandlers")
	ret0, _ := ret[0].([]interfaces.HandlerTuple)
	return ret0
}

// GetHandlers indicates an expected call of GetHandlers
func (mr *MockBotMockRecorder) GetHandlers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandlers", reflect.TypeOf((*MockBot)(nil).GetHandlers))
}

// GetIMHandlers mocks base method
func (m *MockBot) GetIMHandlers() []interfaces.HandlerTuple {
	ret := m.ctrl.Call(m, "GetIMHandlers")
	ret0, _ := ret[0].([]interfaces.HandlerTuple)
	return ret0
}

// GetIMHandlers indicates an expected call of GetIMHandlers
func (mr *MockBotMockRecorder) GetIMHandlers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIMHandlers", reflect.TypeOf((*MockBot)(nil).GetIMHandlers))
}

// GetModules mocks base method
func (m *MockBot) GetModules() []interfaces.Module {
	ret := m.ctrl.Call(m, "GetModules")
	ret0, _ := ret[0].([]interfaces.Module)
	return ret0
}

// GetModules indicates an expected call of GetModules
func (mr *MockBotMockRecorder) GetModules() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModules", reflect.TypeOf((*MockBot)(nil).GetModules))
}

// GetUsageHandler mocks base method
func (m *MockBot) GetUsageHandler() interfaces.Handler {
	ret := m.ctrl.Call(m, "GetUsageHandler")
	ret0, _ := ret[0].(interfaces.Handler)
	return ret0
}

// GetUsageHandler indicates an expected call of GetUsageHandler
func (mr *MockBotMockRecorder) GetUsageHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageHandler", reflect.TypeOf((*MockBot)(nil).GetUsageHandler))
}

// GetUsername mocks base method
func (m *MockBot) GetUsername(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetUsername", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsername indicates an expected call of GetUsername
func (mr *MockBotMockRecorder) GetUsername(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsername", reflect.TypeOf((*MockBot)(nil).GetUsername), arg0)
}

// Post mocks base method
func (m *MockBot) Post(arg0, arg1 string, arg2 messages.Icon, arg3 bool) {
	m.ctrl.Call(m, "Post", arg0, arg1, arg2, arg3)
}

// Post indicates an expected call of Post
func (mr *MockBotMockRecorder) Post(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockBot)(nil).Post), arg0, arg1, arg2, arg3)
}

// PostImage mocks base method
func (m *MockBot) PostImage(arg0, arg1, arg2 string, arg3 messages.Icon, arg4 bool) {
	m.ctrl.Call(m, "PostImage", arg0, arg1, arg2, arg3, arg4)
}

// PostImage indicates an expected call of PostImage
func (mr *MockBotMockRecorder) PostImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostImage", reflect.TypeOf((*MockBot)(nil).PostImage), arg0, arg1, arg2, arg3, arg4)
}

// PostOnFeed mocks base method
func (m *MockBot) PostOnFeed(arg0 string) error {
	ret := m.ctrl.Call(m, "PostOnFeed", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostOnFeed indicates an expected call of PostOnFeed
func (mr *MockBotMockRecorder) PostOnFeed(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOnFeed", reflect.TypeOf((*MockBot)(nil).PostOnFeed), arg0)
}

// StartTyping mocks base method
func (m *MockBot) StartTyping(arg0 string) {
	m.ctrl.Call(m, "StartTyping", arg0)
}

// StartTyping indicates an expected call of StartTyping
func (mr *MockBotMockRecorder) StartTyping(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTyping", reflect.TypeOf((*MockBot)(nil).StartTyping), arg0)
}

// MockMessage is a mock of Message interface
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// GetChannel mocks base method
func (m *MockMessage) GetChannel() string {
	ret := m.ctrl.Call(m, "GetChannel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetChannel indicates an expected call of GetChannel
func (mr *MockMessageMockRecorder) GetChannel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockMessage)(nil).GetChannel))
}

// GetContent mocks base method
func (m *MockMessage) GetContent() string {
	ret := m.ctrl.Call(m, "GetContent")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContent indicates an expected call of GetContent
func (mr *MockMessageMockRecorder) GetContent() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockMessage)(nil).GetContent))
}

// GetProviderName mocks base method
func (m *MockMessage) GetProviderName() string {
	ret := m.ctrl.Call(m, "GetProviderName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProviderName indicates an expected call of GetProviderName
func (mr *MockMessageMockRecorder) GetProviderName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviderName", reflect.TypeOf((*MockMessage)(nil).GetProviderName))
}

// GetUserID mocks base method
func (m *MockMessage) GetUserID() string {
	ret := m.ctrl.Call(m, "GetUserID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserID indicates an expected call of GetUserID
func (mr *MockMessageMockRecorder) GetUserID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockMessage)(nil).GetUserID))
}

// HasUser mocks base method
func (m *MockMessage) HasUser() bool {
	ret := m.ctrl.Call(m, "HasUser")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUser indicates an expected call of HasUser
func (mr *MockMessageMockRecorder) HasUser() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUser", reflect.TypeOf((*MockMessage)(nil).HasUser))
}

// IsIM mocks base method
func (m *MockMessage) IsIM() bool {
	ret := m.ctrl.Call(m, "IsIM")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIM indicates an expected call of IsIM
func (mr *MockMessageMockRecorder) IsIM() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIM", reflect.TypeOf((*MockMessage)(nil).IsIM))
}

// IsSelf mocks base method
func (m *MockMessage) IsSelf() bool {
	ret := m.ctrl.Call(m, "IsSelf")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSelf indicates an expected call of IsSelf
func (mr *MockMessageMockRecorder) IsSelf() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSelf", reflect.TypeOf((*MockMessage)(nil).IsSelf))
}

// MockHandlerTuple is a mock of HandlerTuple interface
type MockHandlerTuple struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerTupleMockRecorder
}

// MockHandlerTupleMockRecorder is the mock recorder for MockHandlerTuple
type MockHandlerTupleMockRecorder struct {
	mock *MockHandlerTuple
}

// NewMockHandlerTuple creates a new mock instance
func NewMockHandlerTuple(ctrl *gomock.Controller) *MockHandlerTuple {
	mock := &MockHandlerTuple{ctrl: ctrl}
	mock.recorder = &MockHandlerTupleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandlerTuple) EXPECT() *MockHandlerTupleMockRecorder {
	return m.recorder
}

// GetHandler mocks base method
func (m *MockHandlerTuple) GetHandler() interfaces.Handler {
	ret := m.ctrl.Call(m, "GetHandler")
	ret0, _ := ret[0].(interfaces.Handler)
	return ret0
}

// GetHandler indicates an expected call of GetHandler
func (mr *MockHandlerTupleMockRecorder) GetHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandler", reflect.TypeOf((*MockHandlerTuple)(nil).GetHandler))
}

// GetPattern mocks base method
func (m *MockHandlerTuple) GetPattern() *regexp.Regexp {
	ret := m.ctrl.Call(m, "GetPattern")
	ret0, _ := ret[0].(*regexp.Regexp)
	return ret0
}

// GetPattern indicates an expected call of GetPattern
func (mr *MockHandlerTupleMockRecorder) GetPattern() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPattern", reflect.TypeOf((*MockHandlerTuple)(nil).GetPattern))
}

// MockModule is a mock of Module interface
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// GetCommands mocks base method
func (m *MockModule) GetCommands() []interfaces.Command {
	ret := m.ctrl.Call(m, "GetCommands")
	ret0, _ := ret[0].([]interfaces.Command)
	return ret0
}

// GetCommands indicates an expected call of GetCommands
func (mr *MockModuleMockRecorder) GetCommands() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommands", reflect.TypeOf((*MockModule)(nil).GetCommands))
}

// GetName mocks base method
func (m *MockModule) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockModuleMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockModule)(nil).GetName))
}

// MockConnector is a mock of Connector interface
type MockConnector struct {
	ctrl     *gomock.Controller
	recorder *MockConnectorMockRecorder
}

// MockConnectorMockRecorder is the mock recorder for MockConnector
type MockConnectorMockRecorder struct {
	mock *MockConnector
}

// NewMockConnector creates a new mock instance
func NewMockConnector(ctrl *gomock.Controller) *MockConnector {
	mock := &MockConnector{ctrl: ctrl}
	mock.recorder = &MockConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnector) EXPECT() *MockConnectorMockRecorder {
	return m.recorder
}

// Boot mocks base method
func (m *MockConnector) Boot() error {
	ret := m.ctrl.Call(m, "Boot")
	ret0, _ := ret[0].(error)
	return ret0
}

// Boot indicates an expected call of Boot
func (mr *MockConnectorMockRecorder) Boot() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockConnector)(nil).Boot))
}

// LoadModules mocks base method
func (m *MockConnector) LoadModules(arg0 []interfaces.Module) {
	m.ctrl.Call(m, "LoadModules", arg0)
}

// LoadModules indicates an expected call of LoadModules
func (mr *MockConnectorMockRecorder) LoadModules(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadModules", reflect.TypeOf((*MockConnector)(nil).LoadModules), arg0)
}

// Post mocks base method
func (m *MockConnector) Post(arg0, arg1 string, arg2 messages.Icon, arg3 bool) {
	m.ctrl.Call(m, "Post", arg0, arg1, arg2, arg3)
}

// Post indicates an expected call of Post
func (mr *MockConnectorMockRecorder) Post(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockConnector)(nil).Post), arg0, arg1, arg2, arg3)
}

// PostImage mocks base method
func (m *MockConnector) PostImage(arg0, arg1, arg2 string, arg3 messages.Icon, arg4 bool) {
	m.ctrl.Call(m, "PostImage", arg0, arg1, arg2, arg3, arg4)
}

// PostImage indicates an expected call of PostImage
func (mr *MockConnectorMockRecorder) PostImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostImage", reflect.TypeOf((*MockConnector)(nil).PostImage), arg0, arg1, arg2, arg3, arg4)
}

// PostOnFeed mocks base method
func (m *MockConnector) PostOnFeed(arg0 string) error {
	ret := m.ctrl.Call(m, "PostOnFeed", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostOnFeed indicates an expected call of PostOnFeed
func (mr *MockConnectorMockRecorder) PostOnFeed(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOnFeed", reflect.TypeOf((*MockConnector)(nil).PostOnFeed), arg0)
}
