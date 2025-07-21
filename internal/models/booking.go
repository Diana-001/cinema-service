package models

import "time"

type Seat struct {
	Row  int `json:"row"`
	Seat int `json:"seat"`
}

// Booking представляет бронирование
type Booking struct {
	ID        int       `json:"id"`        // уникальный идентификатор
	UserID    int       `json:"userId"`    // внешний ключ на пользователя
	SessionID int       `json:"sessionId"` // внешний ключ на сеанс
	Seats     []Seat    `json:"seats"`     // массив сидений
	CreatedAt time.Time `json:"createdAt"` // время создания
}
