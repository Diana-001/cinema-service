package user

import (
	"cinema-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CheckIsAdmin(id int) bool {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return false
	}

	if user.Role != "admin" {
		return false
	}

	return true
}
