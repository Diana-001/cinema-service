package repositories

import (
	"cinema-service/internal/models"
)

type UserRepository interface {
	CheckIsAdmin(id int) bool
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

func (repo *RepositoryImpl) CheckIsAdmin(id int) bool {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return false
	}

	if user.Role != "admin" {
		return false
	}

	return true
}

func (repo *RepositoryImpl) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *RepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
