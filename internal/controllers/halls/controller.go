package halls

import (
	"cinema-service/internal/models"
	"cinema-service/internal/repositories/hall"
	"cinema-service/internal/repositories/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HallController struct {
	HallRepository *hall.HallRepository
	userRepo       *user.UserRepository
}

func NewHallController(hallRepo *hall.HallRepository, userRepo *user.UserRepository) *HallController {
	return &HallController{
		HallRepository: hallRepo,
		userRepo:       userRepo,
	}
}

// GetAllHalls
// @Summary      Получить все залы
// @Description  Возвращает список всех залов
// @Tags         halls
// @Produce      json
// @Success      200  {array}  models.Hall
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/halls [get]
func (h *HallController) GetAllHalls() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		halls, err := h.HallRepository.GetAllHalls()
		if err != nil {
			// Возвращаем 500 и сообщение об ошибке клиенту
			ctx.JSON(500, gin.H{"error": "Ошибка при получении всех фильмов"})
			return
		}

		// Возвращаем фильмы в формате JSON с кодом 200
		ctx.JSON(200, halls)
	}

}

// GetHallsByID
// @Summary      Получить зал по ID
// @Description  Возвращает зал по его уникальному идентификатору
// @Tags         halls
// @Produce      json
// @Param        id   path      int  true  "ID зала"
// @Success      200  {object}  models.Hall
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/halls/{id} [get]
func (h *HallController) GetHallsByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем id из параметра запроса
		hallIDStr := ctx.Param("id")

		// Конвертируем string в int
		hallID, err := strconv.Atoi(hallIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		data, err := h.HallRepository.GetHallByID(hallID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при получении фильма с ID %d", hallID)})
			return
		}

		ctx.JSON(200, data)
	}
}

// DeleteByID
// @Summary      Удалить зал по ID
// @Description  Удаляет зал по его уникальному идентификатору
// @Tags         halls
// @Produce      json
// @Param        id   path      int  true  "ID зала"
// @Success      200  {object}  bool
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/halls/{id} [delete]
func (h *HallController) DeleteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получаем id из параметра запроса
		hallIDStr := ctx.Param("id")

		// Конвертируем string в int
		hallID, err := strconv.Atoi(hallIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		isSuccess, err := h.HallRepository.DeleteByID(hallID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при удалений фильма с ID %d", hallID)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}

// CreateHall
// @Summary      Создать новый зал
// @Description  Добавляет новый зал в систему
// @Tags         halls
// @Accept       json
// @Produce      json
// @Param        hall  body      models.Hall  true  "Данные нового зала"
// @Success      200   {object}  bool
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /cinema-service/v1/api/halls [post]
func (h *HallController) CreateHall() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Hall

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := h.HallRepository.CreateHall(input)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при созданий фильма с ID %d", input)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}

// UpdateHall
// @Summary      Обновить зал по ID
// @Description  Изменяет данные зала по его идентификатору (доступно только администратору)
// @Tags         halls
// @Accept       json
// @Produce      json
// @Param        user-id  header    string        true  "ID пользователя (для проверки прав)"
// @Param        id       path      int           true  "ID зала"
// @Param        hall     body      models.Hall   true  "Новые данные зала"
// @Success      200      {object}  bool
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /cinema-service/v1/api/halls/{id} [put]
func (h *HallController) UpdateHAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdStr := ctx.GetHeader("user-id")

		userID, err := strconv.Atoi(userIdStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID пользователя"})
			return
		}

		isAdmin := h.userRepo.CheckIsAdmin(userID)
		if isAdmin != true {
			ctx.JSON(403, gin.H{"error": "Операция не доступна для пользователя"})
			return
		}

		hallIDStr := ctx.Param("id")

		hallID, err := strconv.Atoi(hallIDStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Некорректный ID фильма"})
			return
		}

		var input models.Hall
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := h.HallRepository.UpdateHall(hallID, input)
		if err != nil {
			ctx.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка при созданий фильма с ID %d", input)})
			return
		}

		ctx.JSON(200, isSuccess)
	}
}
