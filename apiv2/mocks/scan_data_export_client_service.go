// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	io "io"

	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scan_data_export "github.com/GavinMeteion/goharbor-client/v5/apiv2/internal/api/client/scan_data_export"
)

// MockScan_data_exportClientService is an autogenerated mock type for the ClientService type
type MockScan_data_exportClientService struct {
	mock.Mock
}

// DownloadScanData provides a mock function with given fields: params, authInfo, writer
func (_m *MockScan_data_exportClientService) DownloadScanData(params *scan_data_export.DownloadScanDataParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*scan_data_export.DownloadScanDataOK, error) {
	ret := _m.Called(params, authInfo, writer)

	var r0 *scan_data_export.DownloadScanDataOK
	if rf, ok := ret.Get(0).(func(*scan_data_export.DownloadScanDataParams, runtime.ClientAuthInfoWriter, io.Writer) *scan_data_export.DownloadScanDataOK); ok {
		r0 = rf(params, authInfo, writer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_data_export.DownloadScanDataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_data_export.DownloadScanDataParams, runtime.ClientAuthInfoWriter, io.Writer) error); ok {
		r1 = rf(params, authInfo, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportScanData provides a mock function with given fields: params, authInfo
func (_m *MockScan_data_exportClientService) ExportScanData(params *scan_data_export.ExportScanDataParams, authInfo runtime.ClientAuthInfoWriter) (*scan_data_export.ExportScanDataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_data_export.ExportScanDataOK
	if rf, ok := ret.Get(0).(func(*scan_data_export.ExportScanDataParams, runtime.ClientAuthInfoWriter) *scan_data_export.ExportScanDataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_data_export.ExportScanDataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_data_export.ExportScanDataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanDataExportExecution provides a mock function with given fields: params, authInfo
func (_m *MockScan_data_exportClientService) GetScanDataExportExecution(params *scan_data_export.GetScanDataExportExecutionParams, authInfo runtime.ClientAuthInfoWriter) (*scan_data_export.GetScanDataExportExecutionOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_data_export.GetScanDataExportExecutionOK
	if rf, ok := ret.Get(0).(func(*scan_data_export.GetScanDataExportExecutionParams, runtime.ClientAuthInfoWriter) *scan_data_export.GetScanDataExportExecutionOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_data_export.GetScanDataExportExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_data_export.GetScanDataExportExecutionParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanDataExportExecutionList provides a mock function with given fields: params, authInfo
func (_m *MockScan_data_exportClientService) GetScanDataExportExecutionList(params *scan_data_export.GetScanDataExportExecutionListParams, authInfo runtime.ClientAuthInfoWriter) (*scan_data_export.GetScanDataExportExecutionListOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_data_export.GetScanDataExportExecutionListOK
	if rf, ok := ret.Get(0).(func(*scan_data_export.GetScanDataExportExecutionListParams, runtime.ClientAuthInfoWriter) *scan_data_export.GetScanDataExportExecutionListOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_data_export.GetScanDataExportExecutionListOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_data_export.GetScanDataExportExecutionListParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScan_data_exportClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockScan_data_exportClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockScan_data_exportClientService creates a new instance of MockScan_data_exportClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockScan_data_exportClientService(t mockConstructorTestingTNewMockScan_data_exportClientService) *MockScan_data_exportClientService {
	mock := &MockScan_data_exportClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
