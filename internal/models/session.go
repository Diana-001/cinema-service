package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	MovieID   uint           `json:"movieId"`
	Movie     Movie          `gorm:"foreignKey:MovieID" json:"-"`
	HallID    uint           `json:"hallId"`
	Hall      Hall           `gorm:"foreignKey:HallID" json:"-"`
	StartTime time.Time      `json:"startTime"`
	Price     float64        `json:"price"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
