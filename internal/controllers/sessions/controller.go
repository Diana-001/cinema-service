package sessions

import (
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"cinema-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SessionController interface {
	GetAllSessions() gin.HandlerFunc
	CreateSession() gin.HandlerFunc
	GetSessionByID() gin.HandlerFunc
	UpdateSession() gin.HandlerFunc
	DeleteSession() gin.HandlerFunc
}

type SessionControllerImpl struct {
	usecase usecases.Usecase
	l       *logger.Logger
}

func NewSessionController(u usecases.Usecase, l logger.Logger) SessionController {
	return &SessionControllerImpl{
		usecase: u,
		l:       &l,
	}
}

// GetAllSessions
// @Summary      Получить все сеансы
// @Description  Возвращает список всех сеансов в системе
// @Tags         sessions
// @Produce      json
// @Success      200  {array}  models.Session
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/sessions [get]
func (s *SessionControllerImpl) GetAllSessions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		items, err := s.usecase.GetAllSessions()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении сеансов"})
			return
		}
		ctx.JSON(http.StatusOK, items)
	}
}

// GetSessionByID
// @Summary      Получить сеанс по ID
// @Description  Возвращает данные сеанса по его уникальному идентификатору
// @Tags         sessions
// @Produce      json
// @Param        id   path      int  true  "ID сеанса"
// @Success      200  {object}  models.Session
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/sessions/{id} [get]
func (s *SessionControllerImpl) GetSessionByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionIDStr := ctx.Param("id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID сеанса"})
			return
		}

		session, err := s.usecase.GetSessionByID(sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Сеанс не найден"})
			return
		}

		ctx.JSON(http.StatusOK, session)
	}
}

// CreateSession
// @Summary      Создать новый сеанс
// @Description  Добавляет новый сеанс в систему
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        session  body      models.Session  true  "Данные нового сеанса"
// @Success      200      {object}  bool
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /cinema-service/v1/api/sessions [post]
func (s *SessionControllerImpl) CreateSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var session models.Session
		if err := ctx.ShouldBindJSON(&session); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := s.usecase.CreateSession(session)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании сеанса"})
			return
		}

		ctx.JSON(http.StatusOK, isSuccess)
	}
}

// UpdateSession
// @Summary      Обновить сеанс по ID
// @Description  Изменяет данные сеанса по его идентификатору
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "ID сеанса"
// @Param        session  body      models.Session  true  "Новые данные сеанса"
// @Success      200      {object}  bool
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /cinema-service/v1/api/sessions/{id} [put]
func (s *SessionControllerImpl) UpdateSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionIDStr := ctx.Param("id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID сеанса"})
			return
		}

		var session models.Session
		if err := ctx.ShouldBindJSON(&session); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isSuccess, err := s.usecase.UpdateSession(sessionID, session)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении сеанса"})
			return
		}

		ctx.JSON(http.StatusOK, isSuccess)
	}
}

// DeleteSession
// @Summary      Удалить сеанс по ID
// @Description  Удаляет сеанс по его уникальному идентификатору
// @Tags         sessions
// @Produce      json
// @Param        id   path      int  true  "ID сеанса"
// @Success      200  {object}  bool
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /cinema-service/v1/api/sessions/{id} [delete]
func (s *SessionControllerImpl) DeleteSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionIDStr := ctx.Param("id")
		sessionID, err := strconv.Atoi(sessionIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID сеанса"})
			return
		}

		isSuccess, err := s.usecase.DeleteSessionByID(sessionID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении сеанса"})
			return
		}

		ctx.JSON(http.StatusOK, isSuccess)
	}
}
