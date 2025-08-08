package usecases

import (
	"cinema-service/internal/repositories"
	"cinema-service/pkg/logger"
)

type (
	// todo: сделай по аналогии с репозиториями
	Usecase interface {
	}

	UsecaseImpl struct {
		r repositories.Repository
		l logger.Logger
	}
)

func New(r repositories.Repository, l logger.Logger) *UsecaseImpl {
	return &UsecaseImpl{
		r: r,
		l: l,
	}
}
