// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	member "github.com/GavinMeteion/goharbor-client/v5/apiv2/internal/api/client/member"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockMemberClientService is an autogenerated mock type for the ClientService type
type MockMemberClientService struct {
	mock.Mock
}

// CreateProjectMember provides a mock function with given fields: params, authInfo
func (_m *MockMemberClientService) CreateProjectMember(params *member.CreateProjectMemberParams, authInfo runtime.ClientAuthInfoWriter) (*member.CreateProjectMemberCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *member.CreateProjectMemberCreated
	if rf, ok := ret.Get(0).(func(*member.CreateProjectMemberParams, runtime.ClientAuthInfoWriter) *member.CreateProjectMemberCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*member.CreateProjectMemberCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*member.CreateProjectMemberParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProjectMember provides a mock function with given fields: params, authInfo
func (_m *MockMemberClientService) DeleteProjectMember(params *member.DeleteProjectMemberParams, authInfo runtime.ClientAuthInfoWriter) (*member.DeleteProjectMemberOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *member.DeleteProjectMemberOK
	if rf, ok := ret.Get(0).(func(*member.DeleteProjectMemberParams, runtime.ClientAuthInfoWriter) *member.DeleteProjectMemberOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*member.DeleteProjectMemberOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*member.DeleteProjectMemberParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectMember provides a mock function with given fields: params, authInfo
func (_m *MockMemberClientService) GetProjectMember(params *member.GetProjectMemberParams, authInfo runtime.ClientAuthInfoWriter) (*member.GetProjectMemberOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *member.GetProjectMemberOK
	if rf, ok := ret.Get(0).(func(*member.GetProjectMemberParams, runtime.ClientAuthInfoWriter) *member.GetProjectMemberOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*member.GetProjectMemberOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*member.GetProjectMemberParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjectMembers provides a mock function with given fields: params, authInfo
func (_m *MockMemberClientService) ListProjectMembers(params *member.ListProjectMembersParams, authInfo runtime.ClientAuthInfoWriter) (*member.ListProjectMembersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *member.ListProjectMembersOK
	if rf, ok := ret.Get(0).(func(*member.ListProjectMembersParams, runtime.ClientAuthInfoWriter) *member.ListProjectMembersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*member.ListProjectMembersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*member.ListProjectMembersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockMemberClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateProjectMember provides a mock function with given fields: params, authInfo
func (_m *MockMemberClientService) UpdateProjectMember(params *member.UpdateProjectMemberParams, authInfo runtime.ClientAuthInfoWriter) (*member.UpdateProjectMemberOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *member.UpdateProjectMemberOK
	if rf, ok := ret.Get(0).(func(*member.UpdateProjectMemberParams, runtime.ClientAuthInfoWriter) *member.UpdateProjectMemberOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*member.UpdateProjectMemberOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*member.UpdateProjectMemberParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockMemberClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMemberClientService creates a new instance of MockMemberClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMemberClientService(t mockConstructorTestingTNewMockMemberClientService) *MockMemberClientService {
	mock := &MockMemberClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
