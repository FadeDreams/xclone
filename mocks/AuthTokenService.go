// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	xclone "github.com/fadedreams/xclone"
)

// AuthTokenService is an autogenerated mock type for the AuthTokenService type
type AuthTokenService struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: ctx, user
func (_m *AuthTokenService) CreateAccessToken(ctx context.Context, user xclone.User) (string, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, xclone.User) (string, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, xclone.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, xclone.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: ctx, user, tokenID
func (_m *AuthTokenService) CreateRefreshToken(ctx context.Context, user xclone.User, tokenID string) (string, error) {
	ret := _m.Called(ctx, user, tokenID)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, xclone.User, string) (string, error)); ok {
		return rf(ctx, user, tokenID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, xclone.User, string) string); ok {
		r0 = rf(ctx, user, tokenID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, xclone.User, string) error); ok {
		r1 = rf(ctx, user, tokenID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseToken provides a mock function with given fields: ctx, payload
func (_m *AuthTokenService) ParseToken(ctx context.Context, payload string) (xclone.AuthToken, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for ParseToken")
	}

	var r0 xclone.AuthToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (xclone.AuthToken, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) xclone.AuthToken); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(xclone.AuthToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseTokenFromRequest provides a mock function with given fields: ctx, r
func (_m *AuthTokenService) ParseTokenFromRequest(ctx context.Context, r *http.Request) (xclone.AuthToken, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for ParseTokenFromRequest")
	}

	var r0 xclone.AuthToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) (xclone.AuthToken, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) xclone.AuthToken); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(xclone.AuthToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthTokenService creates a new instance of AuthTokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthTokenService {
	mock := &AuthTokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
