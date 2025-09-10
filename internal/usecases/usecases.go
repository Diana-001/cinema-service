package usecases

import (
	"cinema-service/internal/repositories"
	"cinema-service/pkg/logger"
)

var _ Usecase = (*UsecaseImpl)(nil)

type (
	Usecase interface {
		AuthUsecase
		MovieUsecase
		SessionUsecase
		HallUsecase
	}

	UsecaseImpl struct {
		R repositories.Repository
		L logger.Logger
	}
)

func New(r repositories.Repository, l logger.Logger) *UsecaseImpl {
	return &UsecaseImpl{
		R: r,
		L: l,
	}
}
