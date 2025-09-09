package usecases

import (
	"cinema-service/internal/models"
	"context"
	"errors"
)

type HallUsecase interface {
	GetAllHalls() ([]models.Hall, error)
	GetHallByID(id int) (models.Hall, error)
	DeleteHallByID(id int) (bool, error)
	CreateHall(body models.Hall) (bool, error)
	UpdateHall(id, userID int, body models.Hall, ctx context.Context) (bool, error)
}

func (u *UsecaseImpl) GetAllHalls() ([]models.Hall, error) {
	return u.r.GetAllHalls()
}

func (u *UsecaseImpl) GetHallByID(id int) (models.Hall, error) {
	return u.r.GetHallByID(id)
}

func (u *UsecaseImpl) DeleteHallByID(id int) (bool, error) {
	return u.r.DeleteHallByID(id)
}

func (u *UsecaseImpl) CreateHall(body models.Hall) (bool, error) {
	return u.r.CreateHall(body)
}

func (u *UsecaseImpl) UpdateHall(id, userId int, body models.Hall, ctx context.Context) (bool, error) {
	isAdmin := u.r.CheckIsAdmin(userId)
	if isAdmin != true {
		u.l.WarnCtx(ctx, "Ошибка при обновлений данных зала: oперация не доступна для пользователя")
		return false, errors.New("oперация не доступна для пользователя")
	}

	return u.r.UpdateHall(id, body)
}
