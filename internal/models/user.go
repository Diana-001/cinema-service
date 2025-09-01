package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint   `gorm:"primaryKey"`
		Email     string `gorm:"uniqueIndex"`
		Password  string
		Role      string `gorm:"default:user"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	CreateUserRequest struct {
		Email    string
		Password string
	}
)
