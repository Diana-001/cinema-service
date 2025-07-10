package auth

import (
	"cinema-service/internal/models"
	"cinema-service/internal/repositories/user"
	"cinema-service/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct {
	userRepo *user.UserRepository
}

func NewAuthController(userRepo *user.UserRepository) *AuthController {
	return &AuthController{
		userRepo: userRepo,
	}
}
func (a *AuthController) Register() gin.HandlerFunc {
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

		if err := a.userRepo.CreateUser(userData); err != nil {
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

func (a *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		user, err := a.userRepo.GetUserByEmail(req.Email) // userRepo — экземпляр UserRepository
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

func (a *AuthController) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
