package repositories

import (
	"cinema-service/pkg/logger"
	"gorm.io/gorm"
)

var _ Repository = (*RepositoryImpl)(nil)

type (
	Repository interface {
		UserRepository
		SessionRepository
		MovieRepository
		HallRepository
	}

	RepositoryImpl struct {
		db *gorm.DB
		l  logger.Logger
	}
)

func NewRepository(db *gorm.DB, l logger.Logger) Repository {
	return &RepositoryImpl{db: db, l: l}
}
