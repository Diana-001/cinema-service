package routes

import (
	_ "cinema-service/docs"
	"cinema-service/internal/factories"
	"cinema-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(router *gin.Engine, handler *factories.HandlersFactory) {
	//проверка что сервис доступен
	healthcheck := router.Group("/cinema-service/healthcheck")
	{
		healthcheck.GET("/liveness", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "cinema-service ok",
			})
		})
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.AuthController.Register())
		auth.POST("/login", handler.AuthController.Login())
		auth.POST("/refresh", handler.AuthController.Refresh())
	}

	//первая версия CRUD для проекта
	v1 := router.Group("/cinema-service/v1/api")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("movies", handler.MovieController.GetAllMovies())
		v1.GET("movies/:id", handler.MovieController.GetMovieByID())
		v1.DELETE("movies/:id", handler.MovieController.DeleteByID())
		v1.POST("movies", handler.MovieController.CreateMovie())
		v1.PUT("movies/:id", handler.MovieController.UpdateMovie())

		v1.GET("halls", handler.HallController.GetAllHalls())
		v1.GET("halls/:id", handler.HallController.GetHallsByID())
		v1.DELETE("halls/:id", handler.HallController.DeleteByID())
		v1.POST("halls", handler.HallController.CreateHall())
		v1.PUT("halls/:id", handler.HallController.UpdateHAll())

		v1.GET("sessions", handler.SessionController.GetAllSessions())
		v1.GET("sessions/:id", handler.SessionController.GetSessionByID())
		v1.POST("sessions", handler.SessionController.CreateSession())
		v1.PUT("sessions/:id", handler.SessionController.UpdateSession())
		v1.DELETE("sessions/:id", handler.SessionController.DeleteSession())

		v1.GET("bookings", handler.BookingController.GetAllBooking())
		v1.POST("bookings", handler.BookingController.CreateBooking())
	}
}
