package movie

import (
	"cinema-service/internal/models"
	"cinema-service/internal/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieController struct {
	movieRepo *repositories.MovieRepository
	userRepo  *repositories.UserRepository
}

func NewMovieController(repo *repositories.MovieRepository, userRepo *repositories.UserRepository) *MovieController {
	return &MovieController{
		movieRepo: repo,
		userRepo:  userRepo,
	}
}

// GetAllMovies
// @Summary      Получить все фильмы
// @Description  Возвращает список всех фильмов
// @Tags         movies
// @Produce      json
// @Success      200 {array} models.Movie
// @Failure      500 {object} map[string]string
// @Router       /movies [get]
func (c *MovieController) GetAllMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movies, err := c.movieRepo.GetAll()
		if err != nil {
			// Возвращаем 500 и сообщение об ошибке клиенту
			ctx.JSON(500, gin.H{"error": "Ошибка при получении всех фильмов"})
			return
		}

		// Возвращаем фильмы в формате JSON с кодом 200
		ctx.JSON(200, movies)
	}
}

// GetMovieByID
// @Summary      Получить фильм по ID
// @Description  Возвращает фильм по его ID
// @Tags         movies
// @Produce      json
// @Param        id   path      int  true  "ID фильма"
// @Success      200  {object}  models.Movie
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /movies/{id} [get]
func (c *MovieController) GetMovieByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем id из параметра запроса
		movieIDStr := ctx.Param("id")

		// Конвертируем string в int
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		movie, err := c.movieRepo.GetMovieByID(movieID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при получении фильма с ID %d", movieID)})
			return
		}

		ctx.JSON(200, movie)
	}
}

// DeleteByID
// @Summary      Удаление фильма по ID
// @Description  Удаление фильма по его ID
// @Tags         movies
// @Produce      json
// @Param        id   path      int  true  "ID фильма"
// @Success      200  {boolean}  bool
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /movies/{id} [delete]
func (c *MovieController) DeleteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем id из параметра запроса
		movieIDStr := ctx.Param("id")

		// Конвертируем string в int
		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		isSuccess, err := c.movieRepo.DeleteByID(movieID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при удалений фильма с ID %d", movieID)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}

// CreateMovie
// @Summary      Создание фильма
// @Description  Создание фильма
// @Tags         movies
// @Produce      json
// @Param        movie  body   models.Movie  true  "Данные фильма"
// @Success      200  {boolean}  bool
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /movies [post]
func (c *MovieController) CreateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var movie models.Movie

		if err := ctx.ShouldBindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := c.movieRepo.CreateMovie(movie)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при созданий фильма с ID %d", movie)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}

// UpdateMovie
// @Summary      Обновление фильма по ID
// @Description   Обновление фильма по ID
// @Tags         movies
// @Produce      json
// @Param        movie  body   models.Movie  true  "Данные фильма"
// @Success      200  {boolean}  bool
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /movies/{id} [put]
func (c *MovieController) UpdateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdStr := ctx.GetHeader("user-id")

		userID, err := strconv.Atoi(userIdStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID пользователя"})
			return
		}

		isAdmin := c.userRepo.CheckIsAdmin(userID)
		if isAdmin != true {
			ctx.JSON(403, gin.H{"error": "Операция не доступна для пользователя"})
			return
		}

		movieIDStr := ctx.Param("id")

		movieID, err := strconv.Atoi(movieIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		var movie models.Movie
		if err := ctx.ShouldBindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := c.movieRepo.UpdateMovie(movieID, movie)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при созданий фильма с ID %d", movie)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}
