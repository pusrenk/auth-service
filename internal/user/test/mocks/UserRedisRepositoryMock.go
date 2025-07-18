// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/pusrenk/auth-service/internal/user/entities"
	mock "github.com/stretchr/testify/mock"
)

// UserRedisRepositoryMock is an autogenerated mock type for the UserRedisRepository type
type UserRedisRepositoryMock struct {
	mock.Mock
}

type UserRedisRepositoryMock_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRedisRepositoryMock) EXPECT() *UserRedisRepositoryMock_Expecter {
	return &UserRedisRepositoryMock_Expecter{mock: &_m.Mock}
}

// GetUserBySessionID provides a mock function with given fields: ctx, id
func (_m *UserRedisRepositoryMock) GetUserBySessionID(ctx context.Context, id string) (*entities.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserBySessionID")
	}

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRedisRepositoryMock_GetUserBySessionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserBySessionID'
type UserRedisRepositoryMock_GetUserBySessionID_Call struct {
	*mock.Call
}

// GetUserBySessionID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *UserRedisRepositoryMock_Expecter) GetUserBySessionID(ctx interface{}, id interface{}) *UserRedisRepositoryMock_GetUserBySessionID_Call {
	return &UserRedisRepositoryMock_GetUserBySessionID_Call{Call: _e.mock.On("GetUserBySessionID", ctx, id)}
}

func (_c *UserRedisRepositoryMock_GetUserBySessionID_Call) Run(run func(ctx context.Context, id string)) *UserRedisRepositoryMock_GetUserBySessionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRedisRepositoryMock_GetUserBySessionID_Call) Return(_a0 *entities.User, _a1 error) *UserRedisRepositoryMock_GetUserBySessionID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRedisRepositoryMock_GetUserBySessionID_Call) RunAndReturn(run func(context.Context, string) (*entities.User, error)) *UserRedisRepositoryMock_GetUserBySessionID_Call {
	_c.Call.Return(run)
	return _c
}

// StoreUserSession provides a mock function with given fields: ctx, user
func (_m *UserRedisRepositoryMock) StoreUserSession(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for StoreUserSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRedisRepositoryMock_StoreUserSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StoreUserSession'
type UserRedisRepositoryMock_StoreUserSession_Call struct {
	*mock.Call
}

// StoreUserSession is a helper method to define mock.On call
//   - ctx context.Context
//   - user *entities.User
func (_e *UserRedisRepositoryMock_Expecter) StoreUserSession(ctx interface{}, user interface{}) *UserRedisRepositoryMock_StoreUserSession_Call {
	return &UserRedisRepositoryMock_StoreUserSession_Call{Call: _e.mock.On("StoreUserSession", ctx, user)}
}

func (_c *UserRedisRepositoryMock_StoreUserSession_Call) Run(run func(ctx context.Context, user *entities.User)) *UserRedisRepositoryMock_StoreUserSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.User))
	})
	return _c
}

func (_c *UserRedisRepositoryMock_StoreUserSession_Call) Return(_a0 error) *UserRedisRepositoryMock_StoreUserSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRedisRepositoryMock_StoreUserSession_Call) RunAndReturn(run func(context.Context, *entities.User) error) *UserRedisRepositoryMock_StoreUserSession_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRedisRepositoryMock creates a new instance of UserRedisRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRedisRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRedisRepositoryMock {
	mock := &UserRedisRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
