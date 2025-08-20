package factories

import (
	"cinema-service/internal/repositories"
	"cinema-service/pkg/logger"
	"gorm.io/gorm"
)

func NewSqlRepositoryFactory(dbClient *gorm.DB, l logger.Logger) *MySqlRepositoryFactory {
	repo := repositories.NewRepository(dbClient, l)
	return &MySqlRepositoryFactory{
		r: repo,
	}
}

type MySqlRepositoryFactory struct {
	r repositories.Repository
}
