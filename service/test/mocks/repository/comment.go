// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/hanoys/sigma-music-core/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CommentRepository is an autogenerated mock type for the ICommentRepository type
type CommentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, comment
func (_m *CommentRepository) Create(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	ret := _m.Called(ctx, comment)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Comment) (domain.Comment, error)); ok {
		return rf(ctx, comment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Comment) domain.Comment); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Get(0).(domain.Comment)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Comment) error); ok {
		r1 = rf(ctx, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTrackID provides a mock function with given fields: ctx, trackID
func (_m *CommentRepository) GetByTrackID(ctx context.Context, trackID uuid.UUID) ([]domain.Comment, error) {
	ret := _m.Called(ctx, trackID)

	if len(ret) == 0 {
		panic("no return value specified for GetByTrackID")
	}

	var r0 []domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]domain.Comment, error)); ok {
		return rf(ctx, trackID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []domain.Comment); ok {
		r0 = rf(ctx, trackID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, trackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: ctx, userID
func (_m *CommentRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Comment, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 []domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]domain.Comment, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []domain.Comment); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCommentRepository creates a new instance of CommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentRepository {
	mock := &CommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
