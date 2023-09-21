// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github/kunhou/virtual-filesystem-cli/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// CreateFile mocks base method.
func (m *MockIRepository) CreateFile(username, folderName string, file *entity.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", username, folderName, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFile indicates an expected call of CreateFile.
func (mr *MockIRepositoryMockRecorder) CreateFile(username, folderName, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockIRepository)(nil).CreateFile), username, folderName, file)
}

// CreateFolder mocks base method.
func (m *MockIRepository) CreateFolder(username string, folder *entity.Folder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", username, folder)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFolder indicates an expected call of CreateFolder.
func (mr *MockIRepositoryMockRecorder) CreateFolder(username, folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockIRepository)(nil).CreateFolder), username, folder)
}

// CreateUser mocks base method.
func (m *MockIRepository) CreateUser(user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIRepository)(nil).CreateUser), user)
}

// DeleteFile mocks base method.
func (m *MockIRepository) DeleteFile(username, folderName, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", username, folderName, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockIRepositoryMockRecorder) DeleteFile(username, folderName, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockIRepository)(nil).DeleteFile), username, folderName, fileName)
}

// DeleteFolder mocks base method.
func (m *MockIRepository) DeleteFolder(username, folderName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFolder", username, folderName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFolder indicates an expected call of DeleteFolder.
func (mr *MockIRepositoryMockRecorder) DeleteFolder(username, folderName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFolder", reflect.TypeOf((*MockIRepository)(nil).DeleteFolder), username, folderName)
}

// GetUserByName mocks base method.
func (m *MockIRepository) GetUserByName(username string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", username)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockIRepositoryMockRecorder) GetUserByName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockIRepository)(nil).GetUserByName), username)
}

// ListFiles mocks base method.
func (m *MockIRepository) ListFiles(username, folderName string, opt entity.ListFileOption) ([]*entity.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", username, folderName, opt)
	ret0, _ := ret[0].([]*entity.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockIRepositoryMockRecorder) ListFiles(username, folderName, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockIRepository)(nil).ListFiles), username, folderName, opt)
}

// ListFolders mocks base method.
func (m *MockIRepository) ListFolders(username string, opt entity.ListFolderOption) ([]*entity.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFolders", username, opt)
	ret0, _ := ret[0].([]*entity.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFolders indicates an expected call of ListFolders.
func (mr *MockIRepositoryMockRecorder) ListFolders(username, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFolders", reflect.TypeOf((*MockIRepository)(nil).ListFolders), username, opt)
}

// RenameFolder mocks base method.
func (m *MockIRepository) RenameFolder(username, oldName, newName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameFolder", username, oldName, newName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameFolder indicates an expected call of RenameFolder.
func (mr *MockIRepositoryMockRecorder) RenameFolder(username, oldName, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameFolder", reflect.TypeOf((*MockIRepository)(nil).RenameFolder), username, oldName, newName)
}
