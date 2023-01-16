// Code generated by MockGen. DO NOT EDIT.
// Source: ../scheduler_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	common "d7y.io/api/pkg/apis/common/v2"
	scheduler "d7y.io/api/pkg/apis/scheduler/v2"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockSchedulerClient is a mock of SchedulerClient interface.
type MockSchedulerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerClientMockRecorder
}

// MockSchedulerClientMockRecorder is the mock recorder for MockSchedulerClient.
type MockSchedulerClientMockRecorder struct {
	mock *MockSchedulerClient
}

// NewMockSchedulerClient creates a new mock instance.
func NewMockSchedulerClient(ctrl *gomock.Controller) *MockSchedulerClient {
	mock := &MockSchedulerClient{ctrl: ctrl}
	mock.recorder = &MockSchedulerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerClient) EXPECT() *MockSchedulerClientMockRecorder {
	return m.recorder
}

// AnnounceHost mocks base method.
func (m *MockSchedulerClient) AnnounceHost(ctx context.Context, in *scheduler.AnnounceHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceHost", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceHost indicates an expected call of AnnounceHost.
func (mr *MockSchedulerClientMockRecorder) AnnounceHost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceHost", reflect.TypeOf((*MockSchedulerClient)(nil).AnnounceHost), varargs...)
}

// AnnouncePeer mocks base method.
func (m *MockSchedulerClient) AnnouncePeer(ctx context.Context, opts ...grpc.CallOption) (scheduler.Scheduler_AnnouncePeerClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnouncePeer", varargs...)
	ret0, _ := ret[0].(scheduler.Scheduler_AnnouncePeerClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnouncePeer indicates an expected call of AnnouncePeer.
func (mr *MockSchedulerClientMockRecorder) AnnouncePeer(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnouncePeer", reflect.TypeOf((*MockSchedulerClient)(nil).AnnouncePeer), varargs...)
}

// ExchangePeer mocks base method.
func (m *MockSchedulerClient) ExchangePeer(ctx context.Context, in *scheduler.ExchangePeerRequest, opts ...grpc.CallOption) (*scheduler.ExchangePeerResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExchangePeer", varargs...)
	ret0, _ := ret[0].(*scheduler.ExchangePeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangePeer indicates an expected call of ExchangePeer.
func (mr *MockSchedulerClientMockRecorder) ExchangePeer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangePeer", reflect.TypeOf((*MockSchedulerClient)(nil).ExchangePeer), varargs...)
}

// LeaveHost mocks base method.
func (m *MockSchedulerClient) LeaveHost(ctx context.Context, in *scheduler.LeaveHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeaveHost", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveHost indicates an expected call of LeaveHost.
func (mr *MockSchedulerClientMockRecorder) LeaveHost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveHost", reflect.TypeOf((*MockSchedulerClient)(nil).LeaveHost), varargs...)
}

// LeavePeer mocks base method.
func (m *MockSchedulerClient) LeavePeer(ctx context.Context, in *scheduler.LeavePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeavePeer", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeavePeer indicates an expected call of LeavePeer.
func (mr *MockSchedulerClientMockRecorder) LeavePeer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeavePeer", reflect.TypeOf((*MockSchedulerClient)(nil).LeavePeer), varargs...)
}

// StatPeer mocks base method.
func (m *MockSchedulerClient) StatPeer(ctx context.Context, in *scheduler.StatPeerRequest, opts ...grpc.CallOption) (*common.Peer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatPeer", varargs...)
	ret0, _ := ret[0].(*common.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatPeer indicates an expected call of StatPeer.
func (mr *MockSchedulerClientMockRecorder) StatPeer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatPeer", reflect.TypeOf((*MockSchedulerClient)(nil).StatPeer), varargs...)
}

// StatTask mocks base method.
func (m *MockSchedulerClient) StatTask(ctx context.Context, in *scheduler.StatTaskRequest, opts ...grpc.CallOption) (*common.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatTask", varargs...)
	ret0, _ := ret[0].(*common.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockSchedulerClientMockRecorder) StatTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockSchedulerClient)(nil).StatTask), varargs...)
}

// SyncProbes mocks base method.
func (m *MockSchedulerClient) SyncProbes(ctx context.Context, opts ...grpc.CallOption) (scheduler.Scheduler_SyncProbesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncProbes", varargs...)
	ret0, _ := ret[0].(scheduler.Scheduler_SyncProbesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncProbes indicates an expected call of SyncProbes.
func (mr *MockSchedulerClientMockRecorder) SyncProbes(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProbes", reflect.TypeOf((*MockSchedulerClient)(nil).SyncProbes), varargs...)
}

// MockScheduler_AnnouncePeerClient is a mock of Scheduler_AnnouncePeerClient interface.
type MockScheduler_AnnouncePeerClient struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_AnnouncePeerClientMockRecorder
}

// MockScheduler_AnnouncePeerClientMockRecorder is the mock recorder for MockScheduler_AnnouncePeerClient.
type MockScheduler_AnnouncePeerClientMockRecorder struct {
	mock *MockScheduler_AnnouncePeerClient
}

// NewMockScheduler_AnnouncePeerClient creates a new mock instance.
func NewMockScheduler_AnnouncePeerClient(ctrl *gomock.Controller) *MockScheduler_AnnouncePeerClient {
	mock := &MockScheduler_AnnouncePeerClient{ctrl: ctrl}
	mock.recorder = &MockScheduler_AnnouncePeerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_AnnouncePeerClient) EXPECT() *MockScheduler_AnnouncePeerClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockScheduler_AnnouncePeerClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockScheduler_AnnouncePeerClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).Context))
}

// Header mocks base method.
func (m *MockScheduler_AnnouncePeerClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockScheduler_AnnouncePeerClient) Recv() (*scheduler.AnnouncePeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*scheduler.AnnouncePeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_AnnouncePeerClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_AnnouncePeerClient) Send(arg0 *scheduler.AnnouncePeerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_AnnouncePeerClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockScheduler_AnnouncePeerClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockScheduler_AnnouncePeerClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockScheduler_AnnouncePeerClient)(nil).Trailer))
}

// MockScheduler_SyncProbesClient is a mock of Scheduler_SyncProbesClient interface.
type MockScheduler_SyncProbesClient struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_SyncProbesClientMockRecorder
}

// MockScheduler_SyncProbesClientMockRecorder is the mock recorder for MockScheduler_SyncProbesClient.
type MockScheduler_SyncProbesClientMockRecorder struct {
	mock *MockScheduler_SyncProbesClient
}

// NewMockScheduler_SyncProbesClient creates a new mock instance.
func NewMockScheduler_SyncProbesClient(ctrl *gomock.Controller) *MockScheduler_SyncProbesClient {
	mock := &MockScheduler_SyncProbesClient{ctrl: ctrl}
	mock.recorder = &MockScheduler_SyncProbesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_SyncProbesClient) EXPECT() *MockScheduler_SyncProbesClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockScheduler_SyncProbesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockScheduler_SyncProbesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockScheduler_SyncProbesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_SyncProbesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).Context))
}

// Header mocks base method.
func (m *MockScheduler_SyncProbesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockScheduler_SyncProbesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockScheduler_SyncProbesClient) Recv() (*scheduler.SyncProbesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*scheduler.SyncProbesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_SyncProbesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_SyncProbesClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_SyncProbesClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_SyncProbesClient) Send(arg0 *scheduler.SyncProbesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_SyncProbesClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_SyncProbesClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_SyncProbesClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockScheduler_SyncProbesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockScheduler_SyncProbesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockScheduler_SyncProbesClient)(nil).Trailer))
}

// MockSchedulerServer is a mock of SchedulerServer interface.
type MockSchedulerServer struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerServerMockRecorder
}

// MockSchedulerServerMockRecorder is the mock recorder for MockSchedulerServer.
type MockSchedulerServerMockRecorder struct {
	mock *MockSchedulerServer
}

// NewMockSchedulerServer creates a new mock instance.
func NewMockSchedulerServer(ctrl *gomock.Controller) *MockSchedulerServer {
	mock := &MockSchedulerServer{ctrl: ctrl}
	mock.recorder = &MockSchedulerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerServer) EXPECT() *MockSchedulerServerMockRecorder {
	return m.recorder
}

// AnnounceHost mocks base method.
func (m *MockSchedulerServer) AnnounceHost(arg0 context.Context, arg1 *scheduler.AnnounceHostRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnounceHost", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceHost indicates an expected call of AnnounceHost.
func (mr *MockSchedulerServerMockRecorder) AnnounceHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceHost", reflect.TypeOf((*MockSchedulerServer)(nil).AnnounceHost), arg0, arg1)
}

// AnnouncePeer mocks base method.
func (m *MockSchedulerServer) AnnouncePeer(arg0 scheduler.Scheduler_AnnouncePeerServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnouncePeer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnouncePeer indicates an expected call of AnnouncePeer.
func (mr *MockSchedulerServerMockRecorder) AnnouncePeer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnouncePeer", reflect.TypeOf((*MockSchedulerServer)(nil).AnnouncePeer), arg0)
}

// ExchangePeer mocks base method.
func (m *MockSchedulerServer) ExchangePeer(arg0 context.Context, arg1 *scheduler.ExchangePeerRequest) (*scheduler.ExchangePeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangePeer", arg0, arg1)
	ret0, _ := ret[0].(*scheduler.ExchangePeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangePeer indicates an expected call of ExchangePeer.
func (mr *MockSchedulerServerMockRecorder) ExchangePeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangePeer", reflect.TypeOf((*MockSchedulerServer)(nil).ExchangePeer), arg0, arg1)
}

// LeaveHost mocks base method.
func (m *MockSchedulerServer) LeaveHost(arg0 context.Context, arg1 *scheduler.LeaveHostRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveHost", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveHost indicates an expected call of LeaveHost.
func (mr *MockSchedulerServerMockRecorder) LeaveHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveHost", reflect.TypeOf((*MockSchedulerServer)(nil).LeaveHost), arg0, arg1)
}

// LeavePeer mocks base method.
func (m *MockSchedulerServer) LeavePeer(arg0 context.Context, arg1 *scheduler.LeavePeerRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeavePeer", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeavePeer indicates an expected call of LeavePeer.
func (mr *MockSchedulerServerMockRecorder) LeavePeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeavePeer", reflect.TypeOf((*MockSchedulerServer)(nil).LeavePeer), arg0, arg1)
}

// StatPeer mocks base method.
func (m *MockSchedulerServer) StatPeer(arg0 context.Context, arg1 *scheduler.StatPeerRequest) (*common.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatPeer", arg0, arg1)
	ret0, _ := ret[0].(*common.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatPeer indicates an expected call of StatPeer.
func (mr *MockSchedulerServerMockRecorder) StatPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatPeer", reflect.TypeOf((*MockSchedulerServer)(nil).StatPeer), arg0, arg1)
}

// StatTask mocks base method.
func (m *MockSchedulerServer) StatTask(arg0 context.Context, arg1 *scheduler.StatTaskRequest) (*common.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatTask", arg0, arg1)
	ret0, _ := ret[0].(*common.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockSchedulerServerMockRecorder) StatTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockSchedulerServer)(nil).StatTask), arg0, arg1)
}

// SyncProbes mocks base method.
func (m *MockSchedulerServer) SyncProbes(arg0 scheduler.Scheduler_SyncProbesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncProbes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncProbes indicates an expected call of SyncProbes.
func (mr *MockSchedulerServerMockRecorder) SyncProbes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncProbes", reflect.TypeOf((*MockSchedulerServer)(nil).SyncProbes), arg0)
}

// MockUnsafeSchedulerServer is a mock of UnsafeSchedulerServer interface.
type MockUnsafeSchedulerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSchedulerServerMockRecorder
}

// MockUnsafeSchedulerServerMockRecorder is the mock recorder for MockUnsafeSchedulerServer.
type MockUnsafeSchedulerServerMockRecorder struct {
	mock *MockUnsafeSchedulerServer
}

// NewMockUnsafeSchedulerServer creates a new mock instance.
func NewMockUnsafeSchedulerServer(ctrl *gomock.Controller) *MockUnsafeSchedulerServer {
	mock := &MockUnsafeSchedulerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSchedulerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSchedulerServer) EXPECT() *MockUnsafeSchedulerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSchedulerServer mocks base method.
func (m *MockUnsafeSchedulerServer) mustEmbedUnimplementedSchedulerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSchedulerServer")
}

// mustEmbedUnimplementedSchedulerServer indicates an expected call of mustEmbedUnimplementedSchedulerServer.
func (mr *MockUnsafeSchedulerServerMockRecorder) mustEmbedUnimplementedSchedulerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSchedulerServer", reflect.TypeOf((*MockUnsafeSchedulerServer)(nil).mustEmbedUnimplementedSchedulerServer))
}

// MockScheduler_AnnouncePeerServer is a mock of Scheduler_AnnouncePeerServer interface.
type MockScheduler_AnnouncePeerServer struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_AnnouncePeerServerMockRecorder
}

// MockScheduler_AnnouncePeerServerMockRecorder is the mock recorder for MockScheduler_AnnouncePeerServer.
type MockScheduler_AnnouncePeerServerMockRecorder struct {
	mock *MockScheduler_AnnouncePeerServer
}

// NewMockScheduler_AnnouncePeerServer creates a new mock instance.
func NewMockScheduler_AnnouncePeerServer(ctrl *gomock.Controller) *MockScheduler_AnnouncePeerServer {
	mock := &MockScheduler_AnnouncePeerServer{ctrl: ctrl}
	mock.recorder = &MockScheduler_AnnouncePeerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_AnnouncePeerServer) EXPECT() *MockScheduler_AnnouncePeerServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockScheduler_AnnouncePeerServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockScheduler_AnnouncePeerServer) Recv() (*scheduler.AnnouncePeerRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*scheduler.AnnouncePeerRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_AnnouncePeerServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_AnnouncePeerServer) Send(arg0 *scheduler.AnnouncePeerResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockScheduler_AnnouncePeerServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_AnnouncePeerServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockScheduler_AnnouncePeerServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockScheduler_AnnouncePeerServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockScheduler_AnnouncePeerServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockScheduler_AnnouncePeerServer)(nil).SetTrailer), arg0)
}

// MockScheduler_SyncProbesServer is a mock of Scheduler_SyncProbesServer interface.
type MockScheduler_SyncProbesServer struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_SyncProbesServerMockRecorder
}

// MockScheduler_SyncProbesServerMockRecorder is the mock recorder for MockScheduler_SyncProbesServer.
type MockScheduler_SyncProbesServerMockRecorder struct {
	mock *MockScheduler_SyncProbesServer
}

// NewMockScheduler_SyncProbesServer creates a new mock instance.
func NewMockScheduler_SyncProbesServer(ctrl *gomock.Controller) *MockScheduler_SyncProbesServer {
	mock := &MockScheduler_SyncProbesServer{ctrl: ctrl}
	mock.recorder = &MockScheduler_SyncProbesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_SyncProbesServer) EXPECT() *MockScheduler_SyncProbesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockScheduler_SyncProbesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_SyncProbesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockScheduler_SyncProbesServer) Recv() (*scheduler.SyncProbesRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*scheduler.SyncProbesRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_SyncProbesServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_SyncProbesServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_SyncProbesServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_SyncProbesServer) Send(arg0 *scheduler.SyncProbesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_SyncProbesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockScheduler_SyncProbesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockScheduler_SyncProbesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_SyncProbesServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_SyncProbesServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockScheduler_SyncProbesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockScheduler_SyncProbesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockScheduler_SyncProbesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockScheduler_SyncProbesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockScheduler_SyncProbesServer)(nil).SetTrailer), arg0)
}
