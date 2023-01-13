// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	schedule "github.com/GavinMeteion/goharbor-client/v5/apiv2/internal/api/client/schedule"
)

// MockScheduleClientService is an autogenerated mock type for the ClientService type
type MockScheduleClientService struct {
	mock.Mock
}

// GetSchedulePaused provides a mock function with given fields: params, authInfo
func (_m *MockScheduleClientService) GetSchedulePaused(params *schedule.GetSchedulePausedParams, authInfo runtime.ClientAuthInfoWriter) (*schedule.GetSchedulePausedOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *schedule.GetSchedulePausedOK
	if rf, ok := ret.Get(0).(func(*schedule.GetSchedulePausedParams, runtime.ClientAuthInfoWriter) *schedule.GetSchedulePausedOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedule.GetSchedulePausedOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schedule.GetSchedulePausedParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSchedules provides a mock function with given fields: params, authInfo
func (_m *MockScheduleClientService) ListSchedules(params *schedule.ListSchedulesParams, authInfo runtime.ClientAuthInfoWriter) (*schedule.ListSchedulesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *schedule.ListSchedulesOK
	if rf, ok := ret.Get(0).(func(*schedule.ListSchedulesParams, runtime.ClientAuthInfoWriter) *schedule.ListSchedulesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedule.ListSchedulesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*schedule.ListSchedulesParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScheduleClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockScheduleClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockScheduleClientService creates a new instance of MockScheduleClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockScheduleClientService(t mockConstructorTestingTNewMockScheduleClientService) *MockScheduleClientService {
	mock := &MockScheduleClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
