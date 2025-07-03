package models

import (
	"gorm.io/gorm"
	"time"
)

type Hall struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"uniqueIndex" json:"name"`
	Rows        int            `json:"rows"`
	SeatsPerRow int            `json:"seatsPerRow"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
