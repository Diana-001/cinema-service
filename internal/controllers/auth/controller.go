package auth

import (
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"cinema-service/pkg/logger"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}

type AuthControllerImpl struct {
	usecase usecases.Usecase
	l       *logger.Logger
}

func NewAuthController(u usecases.Usecase, l logger.Logger) AuthController {
	return &AuthControllerImpl{
		usecase: u,
		l:       &l,
	}
}
func (a *AuthControllerImpl) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *models.LoginRequest

		// Парсим JSON из тела запроса
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := a.usecase.CreateUser(ctx, req.ToCreateUserRequest())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Возвращаем успешный ответ
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Пользователь успешно зарегистрирован",
			"user_id": user.ID,
		})
	}
}

func (a *AuthControllerImpl) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидаций"})
			return
		}

		tokens, err := a.usecase.Login(req)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка данных"})
			return
		}

		ctx.JSON(http.StatusOK, models.TokenResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		})
	}
}

func (a *AuthControllerImpl) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.RefreshRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
			return
		}

		tokens, err := a.usecase.Refresh(req.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, tokens)
	}
}
