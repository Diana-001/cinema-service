package auth

import (
	"cinema-service/internal/models"
	"cinema-service/internal/repositories/user"
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
		var req struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=6"`
		}

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
