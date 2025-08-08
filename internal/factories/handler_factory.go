package factories

import (
	"cinema-service/internal/controllers/auth"
	"cinema-service/internal/usecases"
	"cinema-service/pkg/logger"
)

func NewHandlersFactory(repoFactory *MySqlRepositoryFactory, l logger.Logger) *HandlersFactory {
	usecase := usecases.New(repoFactory.r, l)

	authController := auth.NewAuthController(usecase, l)

	return &HandlersFactory{
		auth: authController,
	}
}

type HandlersFactory struct {
	auth auth.AuthController
}
