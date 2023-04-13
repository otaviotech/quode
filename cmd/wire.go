//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	repository "github.com/otaviotech/quode/internal/application/repository"
	usecase "github.com/otaviotech/quode/internal/application/usecase"
	"github.com/otaviotech/quode/internal/infra/db/sql/repositories"
	handlers "github.com/otaviotech/quode/internal/infra/web/webserver/handlers"
)

var authorRepositorySet = wire.NewSet(
	repositories.NewAuthorRepository,
	wire.Bind(new(repository.AuthorRepositoryInterface), new(*repositories.AuthorRepository)),
)

var subjectRepositorySet = wire.NewSet(
	repositories.NewSubjectRepository,
	wire.Bind(new(repository.SubjectRepositoryInterface), new(*repositories.SubjectRepository)),
)

func NewCreateAuthorHandler(db *sql.DB) *handlers.CreateAuthorHandler {
	wire.Build(
		authorRepositorySet,
		wire.Bind(new(usecase.CreateAuthorUseCaseInterface), new(*usecase.CreateAuthorUseCase)),
		usecase.NewCreateAuthorUseCase,
		handlers.NewCreateAuthorHandler,
	)

	return &handlers.CreateAuthorHandler{}
}

func NewCreateSubjectHandler(db *sql.DB) *handlers.CreateSubjectHandler {
	wire.Build(
		subjectRepositorySet,
		wire.Bind(new(usecase.CreateSubjectUseCaseInterface), new(*usecase.CreateSubjectUseCase)),
		usecase.NewCreateSubjectUseCase,
		handlers.NewCreateSubjectHandler,
	)

	return &handlers.CreateSubjectHandler{}
}
