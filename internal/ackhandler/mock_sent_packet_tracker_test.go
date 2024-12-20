// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/ackhandler (interfaces: SentPacketTracker)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package ackhandler -destination mock_sent_packet_tracker_test.go github.com/quic-go/quic-go/internal/ackhandler SentPacketTracker
//

// Package ackhandler is a generated GoMock package.
package ackhandler

import (
	reflect "reflect"

	protocol "github.com/xtls/quic-go/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockSentPacketTracker is a mock of SentPacketTracker interface.
type MockSentPacketTracker struct {
	ctrl     *gomock.Controller
	recorder *MockSentPacketTrackerMockRecorder
}

// MockSentPacketTrackerMockRecorder is the mock recorder for MockSentPacketTracker.
type MockSentPacketTrackerMockRecorder struct {
	mock *MockSentPacketTracker
}

// NewMockSentPacketTracker creates a new mock instance.
func NewMockSentPacketTracker(ctrl *gomock.Controller) *MockSentPacketTracker {
	mock := &MockSentPacketTracker{ctrl: ctrl}
	mock.recorder = &MockSentPacketTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentPacketTracker) EXPECT() *MockSentPacketTrackerMockRecorder {
	return m.recorder
}

// GetLowestPacketNotConfirmedAcked mocks base method.
func (m *MockSentPacketTracker) GetLowestPacketNotConfirmedAcked() protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLowestPacketNotConfirmedAcked")
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// GetLowestPacketNotConfirmedAcked indicates an expected call of GetLowestPacketNotConfirmedAcked.
func (mr *MockSentPacketTrackerMockRecorder) GetLowestPacketNotConfirmedAcked() *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLowestPacketNotConfirmedAcked", reflect.TypeOf((*MockSentPacketTracker)(nil).GetLowestPacketNotConfirmedAcked))
	return &MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall{Call: call}
}

// MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall wrap *gomock.Call
type MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall) Return(arg0 protocol.PacketNumber) *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall) Do(f func() protocol.PacketNumber) *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall) DoAndReturn(f func() protocol.PacketNumber) *MockSentPacketTrackerGetLowestPacketNotConfirmedAckedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReceivedPacket mocks base method.
func (m *MockSentPacketTracker) ReceivedPacket(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedPacket", arg0)
}

// ReceivedPacket indicates an expected call of ReceivedPacket.
func (mr *MockSentPacketTrackerMockRecorder) ReceivedPacket(arg0 any) *MockSentPacketTrackerReceivedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockSentPacketTracker)(nil).ReceivedPacket), arg0)
	return &MockSentPacketTrackerReceivedPacketCall{Call: call}
}

// MockSentPacketTrackerReceivedPacketCall wrap *gomock.Call
type MockSentPacketTrackerReceivedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentPacketTrackerReceivedPacketCall) Return() *MockSentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentPacketTrackerReceivedPacketCall) Do(f func(protocol.EncryptionLevel)) *MockSentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentPacketTrackerReceivedPacketCall) DoAndReturn(f func(protocol.EncryptionLevel)) *MockSentPacketTrackerReceivedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
