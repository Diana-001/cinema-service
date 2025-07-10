package models

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	UserEmail string `gorm:"not null"`
	Token     string `gorm:"not null;unique"`
}
