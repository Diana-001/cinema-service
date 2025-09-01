package repositories

import (
	"fmt"

	"cinema-service/internal/models"
)

type UserRepository interface {
	CheckIsAdmin(id int) bool
	CreateUser(user *models.User) (*models.User, error)
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

func (repo *RepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	return user, nil
}

func (repo *RepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
