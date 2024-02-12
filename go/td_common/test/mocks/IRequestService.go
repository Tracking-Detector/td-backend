// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"
)

// IRequestService is an autogenerated mock type for the IRequestService type
type IRequestService struct {
	mock.Mock
}

// CountDocumentsForUrlFilter provides a mock function with given fields: ctx, url
func (_m *IRequestService) CountDocumentsForUrlFilter(ctx context.Context, url string) (int64, error) {
	ret := _m.Called(ctx, url)

	if len(ret) == 0 {
		panic("no return value specified for CountDocumentsForUrlFilter")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPagedRequestsFilterdByUrl provides a mock function with given fields: ctx, url, page, pageSize
func (_m *IRequestService) GetPagedRequestsFilterdByUrl(ctx context.Context, url string, page int, pageSize int) ([]*model.RequestData, error) {
	ret := _m.Called(ctx, url, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for GetPagedRequestsFilterdByUrl")
	}

	var r0 []*model.RequestData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) ([]*model.RequestData, error)); ok {
		return rf(ctx, url, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []*model.RequestData); ok {
		r0 = rf(ctx, url, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RequestData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, url, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequestById provides a mock function with given fields: ctx, id
func (_m *IRequestService) GetRequestById(ctx context.Context, id string) (*model.RequestData, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetRequestById")
	}

	var r0 *model.RequestData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.RequestData, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.RequestData); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RequestData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertManyRequests provides a mock function with given fields: ctx, requests
func (_m *IRequestService) InsertManyRequests(ctx context.Context, requests []*model.RequestData) error {
	ret := _m.Called(ctx, requests)

	if len(ret) == 0 {
		panic("no return value specified for InsertManyRequests")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.RequestData) error); ok {
		r0 = rf(ctx, requests)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRequest provides a mock function with given fields: ctx, request
func (_m *IRequestService) SaveRequest(ctx context.Context, request *model.RequestData) (*model.RequestData, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SaveRequest")
	}

	var r0 *model.RequestData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.RequestData) (*model.RequestData, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.RequestData) *model.RequestData); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RequestData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.RequestData) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamAll provides a mock function with given fields: ctx
func (_m *IRequestService) StreamAll(ctx context.Context) (<-chan *model.RequestData, <-chan error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StreamAll")
	}

	var r0 <-chan *model.RequestData
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan *model.RequestData, <-chan error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.RequestData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.RequestData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) <-chan error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// NewIRequestService creates a new instance of IRequestService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRequestService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRequestService {
	mock := &IRequestService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
