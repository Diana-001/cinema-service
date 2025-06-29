package factories

import "cinema-service/internal/controllers/movie"

func NewHandlersFactory(repoFactory *MySqlRepositoryFactory) *HandlersFactory {
	return &HandlersFactory{
		MovieController: movie.NewMovieController(repoFactory.MovieRepo, repoFactory.UserRepo),
	}
}

type HandlersFactory struct {
	MovieController *movie.MovieController
}
