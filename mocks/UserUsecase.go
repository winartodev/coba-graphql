// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "winartodev/coba-graphql/entity"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserUsecase) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) *entity.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserByID provides a mock function with given fields: ctx, ID
func (_m *UserUsecase) DeleteUserByID(ctx context.Context, ID int) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: ctx, ID
func (_m *UserUsecase) GetUserByID(ctx context.Context, ID int) (*entity.User, error) {
	ret := _m.Called(ctx, ID)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.User); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *UserUsecase) GetUsers(ctx context.Context) ([]entity.User, error) {
	ret := _m.Called(ctx)

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func(context.Context) []entity.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: ctx, ID, user
func (_m *UserUsecase) UpdateUserByID(ctx context.Context, ID int, user entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, ID, user)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int, entity.User) *entity.User); ok {
		r0 = rf(ctx, ID, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, entity.User) error); ok {
		r1 = rf(ctx, ID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a cleanup function to assert the mocks expectations.
func NewUserUsecase(t testing.TB) *UserUsecase {
	mock := &UserUsecase{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
