package factories

import (
	"cinema-service/internal/repositories/movie"
	"cinema-service/internal/repositories/user"
	"gorm.io/gorm"
)

func NewSqlRepositoryFactory(dbClient *gorm.DB) *MySqlRepositoryFactory {
	return &MySqlRepositoryFactory{
		MovieRepo: movie.NewMovieRepository(dbClient),
		UserRepo:  user.NewUserRepository(dbClient),
	}
}

type MySqlRepositoryFactory struct {
	MovieRepo *movie.MovieRepository
	UserRepo  *user.UserRepository
}
