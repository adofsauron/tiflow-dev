// Code generated by MockGen. DO NOT EDIT.
// Source: cdc/processor/sourcemanager/engine/engine.go

// Package mock_engine is a generated GoMock package.
package mock_engine

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/pingcap/tiflow/cdc/model"
	engine "github.com/pingcap/tiflow/cdc/processor/sourcemanager/engine"
)

// MockSortEngine is a mock of SortEngine interface.
type MockSortEngine struct {
	ctrl     *gomock.Controller
	recorder *MockSortEngineMockRecorder
}

// MockSortEngineMockRecorder is the mock recorder for MockSortEngine.
type MockSortEngineMockRecorder struct {
	mock *MockSortEngine
}

// NewMockSortEngine creates a new mock instance.
func NewMockSortEngine(ctrl *gomock.Controller) *MockSortEngine {
	mock := &MockSortEngine{ctrl: ctrl}
	mock.recorder = &MockSortEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSortEngine) EXPECT() *MockSortEngineMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockSortEngine) Add(tableID model.TableID, events ...*model.PolymorphicEvent) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tableID}
	for _, a := range events {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Add", varargs...)
}

// Add indicates an expected call of Add.
func (mr *MockSortEngineMockRecorder) Add(tableID interface{}, events ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tableID}, events...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSortEngine)(nil).Add), varargs...)
}

// AddTable mocks base method.
func (m *MockSortEngine) AddTable(tableID model.TableID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTable", tableID)
}

// AddTable indicates an expected call of AddTable.
func (mr *MockSortEngineMockRecorder) AddTable(tableID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTable", reflect.TypeOf((*MockSortEngine)(nil).AddTable), tableID)
}

// CleanAllTables mocks base method.
func (m *MockSortEngine) CleanAllTables(upperBound engine.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanAllTables", upperBound)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanAllTables indicates an expected call of CleanAllTables.
func (mr *MockSortEngineMockRecorder) CleanAllTables(upperBound interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanAllTables", reflect.TypeOf((*MockSortEngine)(nil).CleanAllTables), upperBound)
}

// CleanByTable mocks base method.
func (m *MockSortEngine) CleanByTable(tableID model.TableID, upperBound engine.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanByTable", tableID, upperBound)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanByTable indicates an expected call of CleanByTable.
func (mr *MockSortEngineMockRecorder) CleanByTable(tableID, upperBound interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanByTable", reflect.TypeOf((*MockSortEngine)(nil).CleanByTable), tableID, upperBound)
}

// Close mocks base method.
func (m *MockSortEngine) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSortEngineMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSortEngine)(nil).Close))
}

// FetchAllTables mocks base method.
func (m *MockSortEngine) FetchAllTables(lowerBound engine.Position) engine.EventIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllTables", lowerBound)
	ret0, _ := ret[0].(engine.EventIterator)
	return ret0
}

// FetchAllTables indicates an expected call of FetchAllTables.
func (mr *MockSortEngineMockRecorder) FetchAllTables(lowerBound interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllTables", reflect.TypeOf((*MockSortEngine)(nil).FetchAllTables), lowerBound)
}

// FetchByTable mocks base method.
func (m *MockSortEngine) FetchByTable(tableID model.TableID, lowerBound, upperBound engine.Position) engine.EventIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByTable", tableID, lowerBound, upperBound)
	ret0, _ := ret[0].(engine.EventIterator)
	return ret0
}

// FetchByTable indicates an expected call of FetchByTable.
func (mr *MockSortEngineMockRecorder) FetchByTable(tableID, lowerBound, upperBound interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByTable", reflect.TypeOf((*MockSortEngine)(nil).FetchByTable), tableID, lowerBound, upperBound)
}

// GetResolvedTs mocks base method.
func (m *MockSortEngine) GetResolvedTs(tableID model.TableID) model.Ts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvedTs", tableID)
	ret0, _ := ret[0].(model.Ts)
	return ret0
}

// GetResolvedTs indicates an expected call of GetResolvedTs.
func (mr *MockSortEngineMockRecorder) GetResolvedTs(tableID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvedTs", reflect.TypeOf((*MockSortEngine)(nil).GetResolvedTs), tableID)
}

// GetStatsByTable mocks base method.
func (m *MockSortEngine) GetStatsByTable(tableID model.TableID) engine.TableStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatsByTable", tableID)
	ret0, _ := ret[0].(engine.TableStats)
	return ret0
}

// GetStatsByTable indicates an expected call of GetStatsByTable.
func (mr *MockSortEngineMockRecorder) GetStatsByTable(tableID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatsByTable", reflect.TypeOf((*MockSortEngine)(nil).GetStatsByTable), tableID)
}

// IsTableBased mocks base method.
func (m *MockSortEngine) IsTableBased() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTableBased")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTableBased indicates an expected call of IsTableBased.
func (mr *MockSortEngineMockRecorder) IsTableBased() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTableBased", reflect.TypeOf((*MockSortEngine)(nil).IsTableBased))
}

// OnResolve mocks base method.
func (m *MockSortEngine) OnResolve(action func(model.TableID, model.Ts)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnResolve", action)
}

// OnResolve indicates an expected call of OnResolve.
func (mr *MockSortEngineMockRecorder) OnResolve(action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnResolve", reflect.TypeOf((*MockSortEngine)(nil).OnResolve), action)
}

// ReceivedEvents mocks base method.
func (m *MockSortEngine) ReceivedEvents() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedEvents")
	ret0, _ := ret[0].(int64)
	return ret0
}

// ReceivedEvents indicates an expected call of ReceivedEvents.
func (mr *MockSortEngineMockRecorder) ReceivedEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedEvents", reflect.TypeOf((*MockSortEngine)(nil).ReceivedEvents))
}

// RemoveTable mocks base method.
func (m *MockSortEngine) RemoveTable(tableID model.TableID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveTable", tableID)
}

// RemoveTable indicates an expected call of RemoveTable.
func (mr *MockSortEngineMockRecorder) RemoveTable(tableID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTable", reflect.TypeOf((*MockSortEngine)(nil).RemoveTable), tableID)
}

// MockEventIterator is a mock of EventIterator interface.
type MockEventIterator struct {
	ctrl     *gomock.Controller
	recorder *MockEventIteratorMockRecorder
}

// MockEventIteratorMockRecorder is the mock recorder for MockEventIterator.
type MockEventIteratorMockRecorder struct {
	mock *MockEventIterator
}

// NewMockEventIterator creates a new mock instance.
func NewMockEventIterator(ctrl *gomock.Controller) *MockEventIterator {
	mock := &MockEventIterator{ctrl: ctrl}
	mock.recorder = &MockEventIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventIterator) EXPECT() *MockEventIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockEventIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventIterator)(nil).Close))
}

// Next mocks base method.
func (m *MockEventIterator) Next() (*model.PolymorphicEvent, engine.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*model.PolymorphicEvent)
	ret1, _ := ret[1].(engine.Position)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Next indicates an expected call of Next.
func (mr *MockEventIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockEventIterator)(nil).Next))
}