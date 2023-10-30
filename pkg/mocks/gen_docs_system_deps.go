// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: gen_docs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// GenDocsSystemDeps is a mock of GenDocsSystemDeps interface.
type GenDocsSystemDeps struct {
	ctrl     *gomock.Controller
	recorder *GenDocsSystemDepsMockRecorder
}

// GenDocsSystemDepsMockRecorder is the mock recorder for GenDocsSystemDeps.
type GenDocsSystemDepsMockRecorder struct {
	mock *GenDocsSystemDeps
}

// NewGenDocsSystemDeps creates a new mock instance.
func NewGenDocsSystemDeps(ctrl *gomock.Controller) *GenDocsSystemDeps {
	mock := &GenDocsSystemDeps{ctrl: ctrl}
	mock.recorder = &GenDocsSystemDepsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *GenDocsSystemDeps) EXPECT() *GenDocsSystemDepsMockRecorder {
	return m.recorder
}

// Pipe mocks base method.
func (m *GenDocsSystemDeps) Pipe() (*os.File, *os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipe")
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(*os.File)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Pipe indicates an expected call of Pipe.
func (mr *GenDocsSystemDepsMockRecorder) Pipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*GenDocsSystemDeps)(nil).Pipe))
}

// SetStdout mocks base method.
func (m *GenDocsSystemDeps) SetStdout(arg0 *os.File) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStdout", arg0)
}

// SetStdout indicates an expected call of SetStdout.
func (mr *GenDocsSystemDepsMockRecorder) SetStdout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStdout", reflect.TypeOf((*GenDocsSystemDeps)(nil).SetStdout), arg0)
}

// Stdout mocks base method.
func (m *GenDocsSystemDeps) Stdout() *os.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stdout")
	ret0, _ := ret[0].(*os.File)
	return ret0
}

// Stdout indicates an expected call of Stdout.
func (mr *GenDocsSystemDepsMockRecorder) Stdout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stdout", reflect.TypeOf((*GenDocsSystemDeps)(nil).Stdout))
}