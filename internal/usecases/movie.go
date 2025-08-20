package usecases

import "cinema-service/internal/models"

type MovieUsecase interface {
	GetAll() ([]models.Movie, error)
	GetMovieByID(id int) (models.Movie, error)
	DeleteMovieByID(id int) (bool, error)
	CreateMovie(body models.Movie) (bool, error)
	UpdateMovie(id int, body models.Movie) (bool, error)
}

func (u *UsecaseImpl) GetAll() ([]models.Movie, error) {
	return u.r.GetAll()
}

func (u *UsecaseImpl) GetMovieByID(id int) (models.Movie, error) {
	return u.r.GetMovieByID(id)
}

func (u *UsecaseImpl) DeleteMovieByID(id int) (bool, error) {
	return u.r.DeleteMovieByID(id)
}

func (u *UsecaseImpl) CreateMovie(body models.Movie) (bool, error) {
	return u.r.CreateMovie(body)
}

func (u *UsecaseImpl) UpdateMovie(id int, body models.Movie) (bool, error) {
	return u.r.UpdateMovie(id, body)
}
