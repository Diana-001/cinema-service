package routes

import (
	_ "cinema-service/docs"
	"cinema-service/internal/factories"
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

	//первая версия CRUD для проекта
	v1 := router.Group("/cinema-service/v1/api")
	{
		v1.GET("movies", handler.MovieController.GetAllMovies())
		v1.GET("movies/:id", handler.MovieController.GetMovieByID())
		v1.DELETE("movies/:id", handler.MovieController.DeleteByID())
		v1.POST("movies", handler.MovieController.CreateMovie())
		v1.PUT("movies/:id", handler.MovieController.UpdateMovie())
	}
}
