// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "carshare-api/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// StoryGetter is an autogenerated mock type for the StoryGetter type
type StoryGetter struct {
	mock.Mock
}

// GetStory provides a mock function with given fields: carUUID
func (_m *StoryGetter) GetStory(carUUID string) ([]models.Point, error) {
	ret := _m.Called(carUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetStory")
	}

	var r0 []models.Point
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.Point, error)); ok {
		return rf(carUUID)
	}
	if rf, ok := ret.Get(0).(func(string) []models.Point); ok {
		r0 = rf(carUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Point)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(carUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStoryGetter creates a new instance of StoryGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStoryGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *StoryGetter {
	mock := &StoryGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
