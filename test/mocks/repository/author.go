package repository_test

import (
	"context"

	"github.com/otaviotech/quode/internal/application/repository"
	"github.com/stretchr/testify/mock"
)

type AuthorRepositoryMock struct {
	mock.Mock
}

func (m *AuthorRepositoryMock) Create(ctx context.Context, input repository.CreateAuthorData) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}
