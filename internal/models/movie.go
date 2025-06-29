package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Duration    uint      `gorm:"not null" json:"duration"`
	PosterUrl   string    `gorm:"not null" json:"posterUrl"`
	CreatedAt   time.Time `json:"created_at"`
}
