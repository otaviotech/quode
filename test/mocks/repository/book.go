package repository_test

import (
	"context"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) Create(ctx context.Context, data repository.CreateBookData) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
