package booking

import (
	"cinema-service/internal/models"
	"cinema-service/internal/repositories/booking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingController struct {
	BookingRepository *booking.BookingRepository
}

func NewBookingController(bookingRepo *booking.BookingRepository) *BookingController {
	return &BookingController{
		BookingRepository: bookingRepo,
	}
}

// GetAllBooking
// @Summary      Получить все бронирования
// @Description  Возвращает  список бронирования
// @Tags         sessions
// @Produce      json
// @Success      200  {array}  models.Booking
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/booking [get]
func (s *BookingController) GetAllBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		items, err := s.BookingRepository.GetAllBooking()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении бронирования"})
			return
		}
		ctx.JSON(http.StatusOK, items)
	}
}

func (s *BookingController) CreateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var booking models.Booking
		if err := ctx.ShouldBindJSON(&booking); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := s.BookingRepository.CreateBooking(booking)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании сеанса"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":   "Бронирование успешно создано",
			"bookingId": id,
		})
	}

}
