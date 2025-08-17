package auth

import (
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"cinema-service/internal/utils"
	"cinema-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

		// Хешируем пароль
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хешировании пароля"})
			return
		}

		userData := &models.User{
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     "user",
		}

		if err := a.usecase.CreateUser(userData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
			return
		}

		// Возвращаем успешный ответ
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Пользователь успешно зарегистрирован",
			"user_id": userData.ID,
		})
	}
}

func (a *AuthControllerImpl) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user, err := a.usecase.GetUserByEmail(req.Email) // userRepo — экземпляр UserRepositoryImpl
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Проверка пароля
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}

		// Если всё ок — выдаём токены
		accessToken, err := utils.GenerateAccessToken(req.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(req.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
			return
		}

		ctx.JSON(http.StatusOK, models.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})

	}
}

func (a *AuthControllerImpl) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
