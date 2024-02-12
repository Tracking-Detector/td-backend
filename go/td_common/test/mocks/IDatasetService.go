// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"

	payload "github.com/Tracking-Detector/td-backend/go/td_common/payload"
)

// IDatasetService is an autogenerated mock type for the IDatasetService type
type IDatasetService struct {
	mock.Mock
}

// CreateDataset provides a mock function with given fields: ctx, datasetPayload
func (_m *IDatasetService) CreateDataset(ctx context.Context, datasetPayload *payload.CreateDatasetPayload) (*model.Dataset, error) {
	ret := _m.Called(ctx, datasetPayload)

	if len(ret) == 0 {
		panic("no return value specified for CreateDataset")
	}

	var r0 *model.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *payload.CreateDatasetPayload) (*model.Dataset, error)); ok {
		return rf(ctx, datasetPayload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *payload.CreateDatasetPayload) *model.Dataset); ok {
		r0 = rf(ctx, datasetPayload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *payload.CreateDatasetPayload) error); ok {
		r1 = rf(ctx, datasetPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDataset provides a mock function with given fields: ctx, id
func (_m *IDatasetService) DeleteDataset(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDataset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDatasets provides a mock function with given fields:
func (_m *IDatasetService) GetAllDatasets() []*model.Dataset {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllDatasets")
	}

	var r0 []*model.Dataset
	if rf, ok := ret.Get(0).(func() []*model.Dataset); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Dataset)
		}
	}

	return r0
}

// GetDatasetByID provides a mock function with given fields: ctx, id
func (_m *IDatasetService) GetDatasetByID(ctx context.Context, id string) (*model.Dataset, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetDatasetByID")
	}

	var r0 *model.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Dataset, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Dataset); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsLabelValid provides a mock function with given fields: label
func (_m *IDatasetService) IsLabelValid(label string) bool {
	ret := _m.Called(label)

	if len(ret) == 0 {
		panic("no return value specified for IsLabelValid")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(label)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsValidDataset provides a mock function with given fields: ctx, id
func (_m *IDatasetService) IsValidDataset(ctx context.Context, id string) bool {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for IsValidDataset")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReloadCache provides a mock function with given fields: ctx
func (_m *IDatasetService) ReloadCache(ctx context.Context) {
	_m.Called(ctx)
}

// Save provides a mock function with given fields: ctx, dataset
func (_m *IDatasetService) Save(ctx context.Context, dataset *model.Dataset) (*model.Dataset, error) {
	ret := _m.Called(ctx, dataset)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Dataset) (*model.Dataset, error)); ok {
		return rf(ctx, dataset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Dataset) *model.Dataset); ok {
		r0 = rf(ctx, dataset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Dataset) error); ok {
		r1 = rf(ctx, dataset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: ctx, datasets
func (_m *IDatasetService) SaveAll(ctx context.Context, datasets []*model.Dataset) ([]*model.Dataset, error) {
	ret := _m.Called(ctx, datasets)

	if len(ret) == 0 {
		panic("no return value specified for SaveAll")
	}

	var r0 []*model.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.Dataset) ([]*model.Dataset, error)); ok {
		return rf(ctx, datasets)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*model.Dataset) []*model.Dataset); ok {
		r0 = rf(ctx, datasets)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*model.Dataset) error); ok {
		r1 = rf(ctx, datasets)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIDatasetService creates a new instance of IDatasetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDatasetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDatasetService {
	mock := &IDatasetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
