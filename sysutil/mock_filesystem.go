// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/jrcichra/kube-applier/sysutil

package sysutil

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of FileSystemInterface interface
type MockFileSystemInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockFileSystemInterfaceRecorder
}

// Recorder for MockFileSystemInterface (not exported)
type _MockFileSystemInterfaceRecorder struct {
	mock *MockFileSystemInterface
}

func NewMockFileSystemInterface(ctrl *gomock.Controller) *MockFileSystemInterface {
	mock := &MockFileSystemInterface{ctrl: ctrl}
	mock.recorder = &_MockFileSystemInterfaceRecorder{mock}
	return mock
}

func (_m *MockFileSystemInterface) EXPECT() *_MockFileSystemInterfaceRecorder {
	return _m.recorder
}

func (_m *MockFileSystemInterface) ListAllFiles(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListAllFiles", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemInterfaceRecorder) ListAllFiles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllFiles", arg0)
}

func (_m *MockFileSystemInterface) ReadLines(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "ReadLines", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemInterfaceRecorder) ReadLines(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadLines", arg0)
}
