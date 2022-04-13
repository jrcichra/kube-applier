// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jrcichra/kube-applier/git (interfaces: GitUtilInterface)

package git

import (
	gomock "github.com/golang/mock/gomock"
)

// MockGitUtilInterface is a mock of GitUtilInterface interface
type MockGitUtilInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGitUtilInterfaceMockRecorder
}

// MockGitUtilInterfaceMockRecorder is the mock recorder for MockGitUtilInterface
type MockGitUtilInterfaceMockRecorder struct {
	mock *MockGitUtilInterface
}

// NewMockGitUtilInterface creates a new mock instance
func NewMockGitUtilInterface(ctrl *gomock.Controller) *MockGitUtilInterface {
	mock := &MockGitUtilInterface{ctrl: ctrl}
	mock.recorder = &MockGitUtilInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockGitUtilInterface) EXPECT() *MockGitUtilInterfaceMockRecorder {
	return _m.recorder
}

// CommitLog mocks base method
func (_m *MockGitUtilInterface) CommitLog(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "CommitLog", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitLog indicates an expected call of CommitLog
func (_mr *MockGitUtilInterfaceMockRecorder) CommitLog(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CommitLog", arg0)
}

// HeadHash mocks base method
func (_m *MockGitUtilInterface) HeadHash() (string, error) {
	ret := _m.ctrl.Call(_m, "HeadHash")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadHash indicates an expected call of HeadHash
func (_mr *MockGitUtilInterfaceMockRecorder) HeadHash() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HeadHash")
}

// ListAllFiles mocks base method
func (_m *MockGitUtilInterface) ListAllFiles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListAllFiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllFiles indicates an expected call of ListAllFiles
func (_mr *MockGitUtilInterfaceMockRecorder) ListAllFiles() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllFiles")
}

// ListDiffFiles mocks base method
func (_m *MockGitUtilInterface) ListDiffFiles(_param0 string, _param1 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListDiffFiles", _param0, _param1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDiffFiles indicates an expected call of ListDiffFiles
func (_mr *MockGitUtilInterfaceMockRecorder) ListDiffFiles(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListDiffFiles", arg0, arg1)
}
