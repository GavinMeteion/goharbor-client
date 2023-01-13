// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scanner "github.com/GavinMeteion/goharbor-client/v5/apiv2/internal/api/client/scanner"
)

// MockScannerClientService is an autogenerated mock type for the ClientService type
type MockScannerClientService struct {
	mock.Mock
}

// CreateScanner provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) CreateScanner(params *scanner.CreateScannerParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.CreateScannerCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.CreateScannerCreated
	if rf, ok := ret.Get(0).(func(*scanner.CreateScannerParams, runtime.ClientAuthInfoWriter) *scanner.CreateScannerCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.CreateScannerCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.CreateScannerParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScanner provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) DeleteScanner(params *scanner.DeleteScannerParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.DeleteScannerOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.DeleteScannerOK
	if rf, ok := ret.Get(0).(func(*scanner.DeleteScannerParams, runtime.ClientAuthInfoWriter) *scanner.DeleteScannerOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.DeleteScannerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.DeleteScannerParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanner provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) GetScanner(params *scanner.GetScannerParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.GetScannerOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.GetScannerOK
	if rf, ok := ret.Get(0).(func(*scanner.GetScannerParams, runtime.ClientAuthInfoWriter) *scanner.GetScannerOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.GetScannerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.GetScannerParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScannerMetadata provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) GetScannerMetadata(params *scanner.GetScannerMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.GetScannerMetadataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.GetScannerMetadataOK
	if rf, ok := ret.Get(0).(func(*scanner.GetScannerMetadataParams, runtime.ClientAuthInfoWriter) *scanner.GetScannerMetadataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.GetScannerMetadataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.GetScannerMetadataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScanners provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) ListScanners(params *scanner.ListScannersParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.ListScannersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.ListScannersOK
	if rf, ok := ret.Get(0).(func(*scanner.ListScannersParams, runtime.ClientAuthInfoWriter) *scanner.ListScannersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.ListScannersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.ListScannersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingScanner provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) PingScanner(params *scanner.PingScannerParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.PingScannerOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.PingScannerOK
	if rf, ok := ret.Get(0).(func(*scanner.PingScannerParams, runtime.ClientAuthInfoWriter) *scanner.PingScannerOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.PingScannerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.PingScannerParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetScannerAsDefault provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) SetScannerAsDefault(params *scanner.SetScannerAsDefaultParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.SetScannerAsDefaultOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.SetScannerAsDefaultOK
	if rf, ok := ret.Get(0).(func(*scanner.SetScannerAsDefaultParams, runtime.ClientAuthInfoWriter) *scanner.SetScannerAsDefaultOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.SetScannerAsDefaultOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.SetScannerAsDefaultParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScannerClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateScanner provides a mock function with given fields: params, authInfo
func (_m *MockScannerClientService) UpdateScanner(params *scanner.UpdateScannerParams, authInfo runtime.ClientAuthInfoWriter) (*scanner.UpdateScannerOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scanner.UpdateScannerOK
	if rf, ok := ret.Get(0).(func(*scanner.UpdateScannerParams, runtime.ClientAuthInfoWriter) *scanner.UpdateScannerOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.UpdateScannerOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scanner.UpdateScannerParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockScannerClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockScannerClientService creates a new instance of MockScannerClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockScannerClientService(t mockConstructorTestingTNewMockScannerClientService) *MockScannerClientService {
	mock := &MockScannerClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
