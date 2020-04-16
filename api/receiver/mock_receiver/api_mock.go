// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/web-platform-tests/wpt.fyi/api/receiver (interfaces: API)

// Package mock_receiver is a generated GoMock package.
package mock_receiver

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v31/github"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
	taskqueue "google.golang.org/appengine/taskqueue"
	io "io"
	http "net/http"
	url "net/url"
	reflect "reflect"
	time "time"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// AddTestRun mocks base method
func (m *MockAPI) AddTestRun(arg0 *shared.TestRun) (shared.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTestRun", arg0)
	ret0, _ := ret[0].(shared.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTestRun indicates an expected call of AddTestRun
func (mr *MockAPIMockRecorder) AddTestRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTestRun", reflect.TypeOf((*MockAPI)(nil).AddTestRun), arg0)
}

// Context mocks base method
func (m *MockAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAPI)(nil).Context))
}

// GetGitHubClient mocks base method
func (m *MockAPI) GetGitHubClient() (*github.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitHubClient")
	ret0, _ := ret[0].(*github.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGitHubClient indicates an expected call of GetGitHubClient
func (mr *MockAPIMockRecorder) GetGitHubClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitHubClient", reflect.TypeOf((*MockAPI)(nil).GetGitHubClient))
}

// GetHTTPClient mocks base method
func (m *MockAPI) GetHTTPClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClient indicates an expected call of GetHTTPClient
func (mr *MockAPIMockRecorder) GetHTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClient", reflect.TypeOf((*MockAPI)(nil).GetHTTPClient))
}

// GetHTTPClientWithTimeout mocks base method
func (m *MockAPI) GetHTTPClientWithTimeout(arg0 time.Duration) *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClientWithTimeout", arg0)
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClientWithTimeout indicates an expected call of GetHTTPClientWithTimeout
func (mr *MockAPIMockRecorder) GetHTTPClientWithTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClientWithTimeout", reflect.TypeOf((*MockAPI)(nil).GetHTTPClientWithTimeout), arg0)
}

// GetHostname mocks base method
func (m *MockAPI) GetHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHostname indicates an expected call of GetHostname
func (mr *MockAPIMockRecorder) GetHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostname", reflect.TypeOf((*MockAPI)(nil).GetHostname))
}

// GetResultsURL mocks base method
func (m *MockAPI) GetResultsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsURL indicates an expected call of GetResultsURL
func (mr *MockAPIMockRecorder) GetResultsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsURL", reflect.TypeOf((*MockAPI)(nil).GetResultsURL), arg0)
}

// GetResultsUploadURL mocks base method
func (m *MockAPI) GetResultsUploadURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResultsUploadURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetResultsUploadURL indicates an expected call of GetResultsUploadURL
func (mr *MockAPIMockRecorder) GetResultsUploadURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResultsUploadURL", reflect.TypeOf((*MockAPI)(nil).GetResultsUploadURL))
}

// GetRunsURL mocks base method
func (m *MockAPI) GetRunsURL(arg0 shared.TestRunFilter) *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunsURL", arg0)
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// GetRunsURL indicates an expected call of GetRunsURL
func (mr *MockAPIMockRecorder) GetRunsURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunsURL", reflect.TypeOf((*MockAPI)(nil).GetRunsURL), arg0)
}

// GetServiceHostname mocks base method
func (m *MockAPI) GetServiceHostname(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceHostname", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceHostname indicates an expected call of GetServiceHostname
func (mr *MockAPIMockRecorder) GetServiceHostname(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceHostname", reflect.TypeOf((*MockAPI)(nil).GetServiceHostname), arg0)
}

// GetUploader mocks base method
func (m *MockAPI) GetUploader(arg0 string) (shared.Uploader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUploader", arg0)
	ret0, _ := ret[0].(shared.Uploader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUploader indicates an expected call of GetUploader
func (mr *MockAPIMockRecorder) GetUploader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUploader", reflect.TypeOf((*MockAPI)(nil).GetUploader), arg0)
}

// GetVersion mocks base method
func (m *MockAPI) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockAPIMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockAPI)(nil).GetVersion))
}

// GetVersionedHostname mocks base method
func (m *MockAPI) GetVersionedHostname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionedHostname")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersionedHostname indicates an expected call of GetVersionedHostname
func (mr *MockAPIMockRecorder) GetVersionedHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionedHostname", reflect.TypeOf((*MockAPI)(nil).GetVersionedHostname))
}

// IsAdmin mocks base method
func (m *MockAPI) IsAdmin() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdmin")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdmin indicates an expected call of IsAdmin
func (mr *MockAPIMockRecorder) IsAdmin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdmin", reflect.TypeOf((*MockAPI)(nil).IsAdmin))
}

// IsFeatureEnabled mocks base method
func (m *MockAPI) IsFeatureEnabled(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled
func (mr *MockAPIMockRecorder) IsFeatureEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockAPI)(nil).IsFeatureEnabled), arg0)
}

// IsLoggedIn mocks base method
func (m *MockAPI) IsLoggedIn() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoggedIn")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLoggedIn indicates an expected call of IsLoggedIn
func (mr *MockAPIMockRecorder) IsLoggedIn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoggedIn", reflect.TypeOf((*MockAPI)(nil).IsLoggedIn))
}

// LoginURL mocks base method
func (m *MockAPI) LoginURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginURL indicates an expected call of LoginURL
func (mr *MockAPIMockRecorder) LoginURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginURL", reflect.TypeOf((*MockAPI)(nil).LoginURL), arg0)
}

// ScheduleResultsTask mocks base method
func (m *MockAPI) ScheduleResultsTask(arg0 string, arg1, arg2 []string, arg3 map[string]string) (*taskqueue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleResultsTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*taskqueue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleResultsTask indicates an expected call of ScheduleResultsTask
func (mr *MockAPIMockRecorder) ScheduleResultsTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleResultsTask", reflect.TypeOf((*MockAPI)(nil).ScheduleResultsTask), arg0, arg1, arg2, arg3)
}

// UpdatePendingTestRun mocks base method
func (m *MockAPI) UpdatePendingTestRun(arg0 shared.PendingTestRun) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePendingTestRun", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePendingTestRun indicates an expected call of UpdatePendingTestRun
func (mr *MockAPIMockRecorder) UpdatePendingTestRun(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePendingTestRun", reflect.TypeOf((*MockAPI)(nil).UpdatePendingTestRun), arg0)
}

// UploadToGCS mocks base method
func (m *MockAPI) UploadToGCS(arg0 string, arg1 io.Reader, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadToGCS", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadToGCS indicates an expected call of UploadToGCS
func (mr *MockAPIMockRecorder) UploadToGCS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadToGCS", reflect.TypeOf((*MockAPI)(nil).UploadToGCS), arg0, arg1, arg2)
}
