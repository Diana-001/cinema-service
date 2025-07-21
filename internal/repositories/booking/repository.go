package booking

import (
	"cinema-service/internal/models"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (repo *BookingRepository) GetAllBooking() ([]models.Booking, error) {
	var booking []models.Booking
	if err := repo.db.Find(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (repo *BookingRepository) CreateBooking(booking models.Booking) (int, error) {
	if err := repo.db.Create(&booking).Error; err != nil {
		return 0, err
	}
	return booking.ID, nil
}
