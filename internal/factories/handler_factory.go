package factories

import (
	"cinema-service/internal/controllers/auth"
	"cinema-service/internal/controllers/halls"
	"cinema-service/internal/controllers/movie"
	"cinema-service/internal/controllers/sessions"
	"cinema-service/internal/usecases"
	"cinema-service/pkg/logger"
)

func NewHandlersFactory(repoFactory *MySqlRepositoryFactory, l logger.Logger) *HandlersFactory {
	usecase := usecases.New(repoFactory.r, l)

	authController := auth.NewAuthController(usecase, l)
	movieController := movie.NewMovieController(usecase, l)
	hallController := halls.NewHallController(usecase, l)
	sessionController := sessions.NewSessionController(usecase, l)

	return &HandlersFactory{
		Auth:              authController,
		MovieController:   movieController,
		HallController:    hallController,
		SessionController: sessionController,
	}
}

type HandlersFactory struct {
	Auth              auth.AuthController
	MovieController   movie.MovieController
	HallController    halls.HallController
	SessionController sessions.SessionController
}
