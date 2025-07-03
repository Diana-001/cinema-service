package factories

import (
	"cinema-service/internal/controllers/auth"
	"cinema-service/internal/controllers/halls"
	"cinema-service/internal/controllers/movie"
	sessions2 "cinema-service/internal/controllers/sessions"
)

func NewHandlersFactory(repoFactory *MySqlRepositoryFactory) *HandlersFactory {
	return &HandlersFactory{
		MovieController:   movie.NewMovieController(repoFactory.MovieRepo, repoFactory.UserRepo),
		AuthController:    auth.NewAuthController(repoFactory.UserRepo),
		HallController:    halls.NewHallController(repoFactory.HallRepo, repoFactory.UserRepo),
		SessionController: sessions2.NewSessionController(repoFactory.SessionRepo),
	}
}

type HandlersFactory struct {
	MovieController   *movie.MovieController
	AuthController    *auth.AuthController
	HallController    *halls.HallController
	SessionController *sessions2.SessionController
}
