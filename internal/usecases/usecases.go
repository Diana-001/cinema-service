package usecases

import (
	"cinema-service/internal/repositories"
	"cinema-service/pkg/logger"
)

var _ Usecase = (*UsecaseImpl)(nil)

type (
	// todo: сделай по аналогии с репозиториями
	Usecase interface {
		AuthUsecase
		MovieUsecase
		SessionUsecase
		HallUsecase
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
