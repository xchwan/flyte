// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// SyncAgentService_ExecuteTaskSyncServer is an autogenerated mock type for the SyncAgentService_ExecuteTaskSyncServer type
type SyncAgentService_ExecuteTaskSyncServer struct {
	mock.Mock
}

type SyncAgentService_ExecuteTaskSyncServer_Context struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_Context) Return(_a0 context.Context) *SyncAgentService_ExecuteTaskSyncServer_Context {
	return &SyncAgentService_ExecuteTaskSyncServer_Context{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnContext() *SyncAgentService_ExecuteTaskSyncServer_Context {
	c_call := _m.On("Context")
	return &SyncAgentService_ExecuteTaskSyncServer_Context{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnContextMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_Context {
	c_call := _m.On("Context", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_Context{Call: c_call}
}

// Context provides a mock function with given fields:
func (_m *SyncAgentService_ExecuteTaskSyncServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

type SyncAgentService_ExecuteTaskSyncServer_Recv struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_Recv) Return(_a0 *admin.ExecuteTaskSyncRequest, _a1 error) *SyncAgentService_ExecuteTaskSyncServer_Recv {
	return &SyncAgentService_ExecuteTaskSyncServer_Recv{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnRecv() *SyncAgentService_ExecuteTaskSyncServer_Recv {
	c_call := _m.On("Recv")
	return &SyncAgentService_ExecuteTaskSyncServer_Recv{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnRecvMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_Recv {
	c_call := _m.On("Recv", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_Recv{Call: c_call}
}

// Recv provides a mock function with given fields:
func (_m *SyncAgentService_ExecuteTaskSyncServer) Recv() (*admin.ExecuteTaskSyncRequest, error) {
	ret := _m.Called()

	var r0 *admin.ExecuteTaskSyncRequest
	if rf, ok := ret.Get(0).(func() *admin.ExecuteTaskSyncRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecuteTaskSyncRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SyncAgentService_ExecuteTaskSyncServer_RecvMsg struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_RecvMsg) Return(_a0 error) *SyncAgentService_ExecuteTaskSyncServer_RecvMsg {
	return &SyncAgentService_ExecuteTaskSyncServer_RecvMsg{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnRecvMsg(m interface{}) *SyncAgentService_ExecuteTaskSyncServer_RecvMsg {
	c_call := _m.On("RecvMsg", m)
	return &SyncAgentService_ExecuteTaskSyncServer_RecvMsg{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnRecvMsgMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_RecvMsg {
	c_call := _m.On("RecvMsg", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_RecvMsg{Call: c_call}
}

// RecvMsg provides a mock function with given fields: m
func (_m *SyncAgentService_ExecuteTaskSyncServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SyncAgentService_ExecuteTaskSyncServer_Send struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_Send) Return(_a0 error) *SyncAgentService_ExecuteTaskSyncServer_Send {
	return &SyncAgentService_ExecuteTaskSyncServer_Send{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSend(_a0 *admin.ExecuteTaskSyncResponse) *SyncAgentService_ExecuteTaskSyncServer_Send {
	c_call := _m.On("Send", _a0)
	return &SyncAgentService_ExecuteTaskSyncServer_Send{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSendMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_Send {
	c_call := _m.On("Send", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_Send{Call: c_call}
}

// Send provides a mock function with given fields: _a0
func (_m *SyncAgentService_ExecuteTaskSyncServer) Send(_a0 *admin.ExecuteTaskSyncResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*admin.ExecuteTaskSyncResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SyncAgentService_ExecuteTaskSyncServer_SendHeader struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_SendHeader) Return(_a0 error) *SyncAgentService_ExecuteTaskSyncServer_SendHeader {
	return &SyncAgentService_ExecuteTaskSyncServer_SendHeader{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSendHeader(_a0 metadata.MD) *SyncAgentService_ExecuteTaskSyncServer_SendHeader {
	c_call := _m.On("SendHeader", _a0)
	return &SyncAgentService_ExecuteTaskSyncServer_SendHeader{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSendHeaderMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_SendHeader {
	c_call := _m.On("SendHeader", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_SendHeader{Call: c_call}
}

// SendHeader provides a mock function with given fields: _a0
func (_m *SyncAgentService_ExecuteTaskSyncServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SyncAgentService_ExecuteTaskSyncServer_SendMsg struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_SendMsg) Return(_a0 error) *SyncAgentService_ExecuteTaskSyncServer_SendMsg {
	return &SyncAgentService_ExecuteTaskSyncServer_SendMsg{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSendMsg(m interface{}) *SyncAgentService_ExecuteTaskSyncServer_SendMsg {
	c_call := _m.On("SendMsg", m)
	return &SyncAgentService_ExecuteTaskSyncServer_SendMsg{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSendMsgMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_SendMsg {
	c_call := _m.On("SendMsg", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_SendMsg{Call: c_call}
}

// SendMsg provides a mock function with given fields: m
func (_m *SyncAgentService_ExecuteTaskSyncServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SyncAgentService_ExecuteTaskSyncServer_SetHeader struct {
	*mock.Call
}

func (_m SyncAgentService_ExecuteTaskSyncServer_SetHeader) Return(_a0 error) *SyncAgentService_ExecuteTaskSyncServer_SetHeader {
	return &SyncAgentService_ExecuteTaskSyncServer_SetHeader{Call: _m.Call.Return(_a0)}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSetHeader(_a0 metadata.MD) *SyncAgentService_ExecuteTaskSyncServer_SetHeader {
	c_call := _m.On("SetHeader", _a0)
	return &SyncAgentService_ExecuteTaskSyncServer_SetHeader{Call: c_call}
}

func (_m *SyncAgentService_ExecuteTaskSyncServer) OnSetHeaderMatch(matchers ...interface{}) *SyncAgentService_ExecuteTaskSyncServer_SetHeader {
	c_call := _m.On("SetHeader", matchers...)
	return &SyncAgentService_ExecuteTaskSyncServer_SetHeader{Call: c_call}
}

// SetHeader provides a mock function with given fields: _a0
func (_m *SyncAgentService_ExecuteTaskSyncServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *SyncAgentService_ExecuteTaskSyncServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}