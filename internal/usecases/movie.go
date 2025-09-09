package usecases

import (
	"cinema-service/internal/models"
	"context"
	"errors"
)

type MovieUsecase interface {
	GetAll() ([]models.Movie, error)
	GetMovieByID(id int) (models.Movie, error)
	DeleteMovieByID(id int) (bool, error)
	CreateMovie(body models.Movie) (bool, error)
	UpdateMovie(id, userId int, body models.Movie, ctx context.Context) (bool, error)
}

func (u *UsecaseImpl) GetAll() ([]models.Movie, error) {
	return u.R.GetAll()
}

func (u *UsecaseImpl) GetMovieByID(id int) (models.Movie, error) {
	return u.R.GetMovieByID(id)
}

func (u *UsecaseImpl) DeleteMovieByID(id int) (bool, error) {
	return u.R.DeleteMovieByID(id)
}

func (u *UsecaseImpl) CreateMovie(body models.Movie) (bool, error) {
	return u.R.CreateMovie(body)
}

func (u *UsecaseImpl) UpdateMovie(id, userId int, body models.Movie, ctx context.Context) (bool, error) {
	isAdmin := u.R.CheckIsAdmin(userId)
	if isAdmin != true {
		u.L.WarnCtx(ctx, "Ошибка при обновлений данных о фильме: oперация не доступна для пользователя")
		return false, errors.New("oперация не доступна для пользователя")
	}

	return u.R.UpdateMovie(id, body)
}
