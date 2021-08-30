package mocks

import (
	context "context"
	"diffme.dev/diffme-api/server/modules/changes"
)
import mock "github.com/stretchr/testify/mock"

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ArticleRepository) Fetch(ctx context.Context, cursor string, num int64) ([]domain.domain, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Change
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Change); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Change)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ArticleRepository) GetByID(ctx context.Context, id int64) (domain.Change, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Change
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Change); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Change)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
