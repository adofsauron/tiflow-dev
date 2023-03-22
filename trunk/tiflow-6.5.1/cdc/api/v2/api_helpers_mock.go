// Code generated by MockGen. DO NOT EDIT.
// Source: cdc/api/v2/api_helpers.go

// Package v2 is a generated GoMock package.
package v2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kv "github.com/pingcap/tidb/kv"
	model "sdbflow/cdc/model"
	owner "sdbflow/cdc/owner"
	config "sdbflow/pkg/config"
	security "sdbflow/pkg/security"
	client "github.com/tikv/pd/client"
)

// MockAPIV2Helpers is a mock of APIV2Helpers interface.
type MockAPIV2Helpers struct {
	ctrl     *gomock.Controller
	recorder *MockAPIV2HelpersMockRecorder
}

// MockAPIV2HelpersMockRecorder is the mock recorder for MockAPIV2Helpers.
type MockAPIV2HelpersMockRecorder struct {
	mock *MockAPIV2Helpers
}

// NewMockAPIV2Helpers creates a new mock instance.
func NewMockAPIV2Helpers(ctrl *gomock.Controller) *MockAPIV2Helpers {
	mock := &MockAPIV2Helpers{ctrl: ctrl}
	mock.recorder = &MockAPIV2HelpersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIV2Helpers) EXPECT() *MockAPIV2HelpersMockRecorder {
	return m.recorder
}

// createTiStore mocks base method.
func (m *MockAPIV2Helpers) createTiStore(pdAddrs []string, credential *security.Credential) (kv.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "createTiStore", pdAddrs, credential)
	ret0, _ := ret[0].(kv.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createTiStore indicates an expected call of createTiStore.
func (mr *MockAPIV2HelpersMockRecorder) createTiStore(pdAddrs, credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createTiStore", reflect.TypeOf((*MockAPIV2Helpers)(nil).createTiStore), pdAddrs, credential)
}

// getPDClient mocks base method.
func (m *MockAPIV2Helpers) getPDClient(ctx context.Context, pdAddrs []string, credential *security.Credential) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPDClient", ctx, pdAddrs, credential)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getPDClient indicates an expected call of getPDClient.
func (mr *MockAPIV2HelpersMockRecorder) getPDClient(ctx, pdAddrs, credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPDClient", reflect.TypeOf((*MockAPIV2Helpers)(nil).getPDClient), ctx, pdAddrs, credential)
}

// getVerfiedTables mocks base method.
func (m *MockAPIV2Helpers) getVerfiedTables(replicaConfig *config.ReplicaConfig, storage kv.Storage, startTs uint64) ([]model.TableName, []model.TableName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getVerfiedTables", replicaConfig, storage, startTs)
	ret0, _ := ret[0].([]model.TableName)
	ret1, _ := ret[1].([]model.TableName)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// getVerfiedTables indicates an expected call of getVerfiedTables.
func (mr *MockAPIV2HelpersMockRecorder) getVerfiedTables(replicaConfig, storage, startTs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getVerfiedTables", reflect.TypeOf((*MockAPIV2Helpers)(nil).getVerfiedTables), replicaConfig, storage, startTs)
}

// verifyCreateChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyCreateChangefeedConfig(ctx context.Context, cfg *ChangefeedConfig, pdClient client.Client, statusProvider owner.StatusProvider, ensureGCServiceID string, kvStorage kv.Storage) (*model.ChangeFeedInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyCreateChangefeedConfig", ctx, cfg, pdClient, statusProvider, ensureGCServiceID, kvStorage)
	ret0, _ := ret[0].(*model.ChangeFeedInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// verifyCreateChangefeedConfig indicates an expected call of verifyCreateChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyCreateChangefeedConfig(ctx, cfg, pdClient, statusProvider, ensureGCServiceID, kvStorage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyCreateChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyCreateChangefeedConfig), ctx, cfg, pdClient, statusProvider, ensureGCServiceID, kvStorage)
}

// verifyResumeChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyResumeChangefeedConfig(ctx context.Context, pdClient client.Client, gcServiceID string, changefeedID model.ChangeFeedID, checkpointTs uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyResumeChangefeedConfig", ctx, pdClient, gcServiceID, changefeedID, checkpointTs)
	ret0, _ := ret[0].(error)
	return ret0
}

// verifyResumeChangefeedConfig indicates an expected call of verifyResumeChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyResumeChangefeedConfig(ctx, pdClient, gcServiceID, changefeedID, checkpointTs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyResumeChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyResumeChangefeedConfig), ctx, pdClient, gcServiceID, changefeedID, checkpointTs)
}

// verifyUpdateChangefeedConfig mocks base method.
func (m *MockAPIV2Helpers) verifyUpdateChangefeedConfig(ctx context.Context, cfg *ChangefeedConfig, oldInfo *model.ChangeFeedInfo, oldUpInfo *model.UpstreamInfo, kvStorage kv.Storage, checkpointTs uint64) (*model.ChangeFeedInfo, *model.UpstreamInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyUpdateChangefeedConfig", ctx, cfg, oldInfo, oldUpInfo, kvStorage, checkpointTs)
	ret0, _ := ret[0].(*model.ChangeFeedInfo)
	ret1, _ := ret[1].(*model.UpstreamInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// verifyUpdateChangefeedConfig indicates an expected call of verifyUpdateChangefeedConfig.
func (mr *MockAPIV2HelpersMockRecorder) verifyUpdateChangefeedConfig(ctx, cfg, oldInfo, oldUpInfo, kvStorage, checkpointTs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyUpdateChangefeedConfig", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyUpdateChangefeedConfig), ctx, cfg, oldInfo, oldUpInfo, kvStorage, checkpointTs)
}

// verifyUpstream mocks base method.
func (m *MockAPIV2Helpers) verifyUpstream(ctx context.Context, changefeedConfig *ChangefeedConfig, cfInfo *model.ChangeFeedInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "verifyUpstream", ctx, changefeedConfig, cfInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// verifyUpstream indicates an expected call of verifyUpstream.
func (mr *MockAPIV2HelpersMockRecorder) verifyUpstream(ctx, changefeedConfig, cfInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "verifyUpstream", reflect.TypeOf((*MockAPIV2Helpers)(nil).verifyUpstream), ctx, changefeedConfig, cfInfo)
}
