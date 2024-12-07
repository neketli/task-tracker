package task

import (
	"projectly-server/internal/domain/task/delivery"
	"projectly-server/internal/domain/task/repository"
	"projectly-server/internal/domain/task/usecase"
	"projectly-server/pkg/logger"
	"projectly-server/pkg/postgres"

	"github.com/labstack/echo/v4"
)

type Dependency struct {
	Logger   *logger.Logger
	Postgres *postgres.Postgres
	Router   *echo.Group
}

func New(dependency Dependency) usecase.TaskUseCase {
	repo := repository.New(dependency.Postgres)

	taskUseCase := usecase.New(repo, dependency.Logger)

	delivery.New(dependency.Router, taskUseCase)

	return taskUseCase
}
