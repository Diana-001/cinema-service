package usecases

import "cinema-service/internal/models"

type AuthUsecase interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	CheckIsAdmin(id int) bool
}

func (u *UsecaseImpl) CreateUser(user *models.User) error {
	return u.r.CreateUser(user)
}

func (u *UsecaseImpl) GetUserByEmail(email string) (*models.User, error) {
	return u.r.GetUserByEmail(email)
}

func (u *UsecaseImpl) CheckIsAdmin(id int) bool {
	return u.r.CheckIsAdmin(id)
}
