// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/hanoys/sigma-music-core/domain"
	mock "github.com/stretchr/testify/mock"
)

// TokenProvider is an autogenerated mock type for the ITokenProvider type
type TokenProvider struct {
	mock.Mock
}

// CloseSession provides a mock function with given fields: ctx, refreshTokenString
func (_m *TokenProvider) CloseSession(ctx context.Context, refreshTokenString string) error {
	ret := _m.Called(ctx, refreshTokenString)

	if len(ret) == 0 {
		panic("no return value specified for CloseSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, refreshTokenString)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSession provides a mock function with given fields: ctx, payload
func (_m *TokenProvider) NewSession(ctx context.Context, payload domain.Payload) (domain.TokenPair, error) {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for NewSession")
	}

	var r0 domain.TokenPair
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Payload) (domain.TokenPair, error)); ok {
		return rf(ctx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Payload) domain.TokenPair); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Get(0).(domain.TokenPair)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Payload) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSession provides a mock function with given fields: ctx, refreshTokenString
func (_m *TokenProvider) RefreshSession(ctx context.Context, refreshTokenString string) (domain.TokenPair, error) {
	ret := _m.Called(ctx, refreshTokenString)

	if len(ret) == 0 {
		panic("no return value specified for RefreshSession")
	}

	var r0 domain.TokenPair
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.TokenPair, error)); ok {
		return rf(ctx, refreshTokenString)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.TokenPair); ok {
		r0 = rf(ctx, refreshTokenString)
	} else {
		r0 = ret.Get(0).(domain.TokenPair)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, refreshTokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyToken provides a mock function with given fields: ctx, accessTokenString
func (_m *TokenProvider) VerifyToken(ctx context.Context, accessTokenString string) (domain.Payload, error) {
	ret := _m.Called(ctx, accessTokenString)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 domain.Payload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Payload, error)); ok {
		return rf(ctx, accessTokenString)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Payload); ok {
		r0 = rf(ctx, accessTokenString)
	} else {
		r0 = ret.Get(0).(domain.Payload)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, accessTokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenProvider creates a new instance of TokenProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenProvider {
	mock := &TokenProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
