package user

import (
	"app/models"
	"context"

	"github.com/stretchr/testify/mock"
)

// why this is here, well, let it on user.go looked weird
type repoMock struct {
	mock.Mock
}

func (m *repoMock) Create(ctx context.Context, in models.User) (err error) {
	args := m.Called(ctx, in)
	return args.Error(0)
}
func (m *repoMock) ExistID(ctx context.Context, id string) (ok bool, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(bool), args.Error(1)
}
func (m *repoMock) Get(ctx context.Context, id string) (out models.User, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *repoMock) Delete(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *repoMock) Patch(ctx context.Context, in models.User) (err error) {
	args := m.Called(ctx, in)
	return args.Error(0)
}

// / ====================
type uuidMock struct {
	mock.Mock
}

func (m *uuidMock) NewString() (str string) {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *uuidMock) IsValid(str string) (err error) {
	args := m.Called(str)
	return args.Error(0)
}
