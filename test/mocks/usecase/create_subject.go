package usecase_test

import (
	"context"

	"github.com/otaviotech/quode/internal/application/usecase"
	"github.com/stretchr/testify/mock"
)

type CreateSubjectUseCaseMock struct {
	mock.Mock
}

func (m *CreateSubjectUseCaseMock) Execute(ctx context.Context, input *usecase.CreateSubjectInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}
