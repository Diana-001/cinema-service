package usecases

import (
	"cinema-service/internal/utils"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"cinema-service/internal/models"
)

const (
	defaultUserRole = "user"
)

type AuthUsecase interface {
	CreateUser(ctx context.Context, reqData *models.CreateUserRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CheckIsAdmin(id int) bool
	Login(payload models.LoginRequest) (*models.TokenResponse, error)
	Refresh(refreshToken string) (*models.TokenResponse, error)
}

func (u *UsecaseImpl) CreateUser(ctx context.Context, reqData *models.CreateUserRequest) (*models.User, error) {
	if reqData == nil {
		u.l.WarnCtx(ctx, "Ошибка при создании пользователя: данные пользователя пусты")
		return nil, errors.New("данные пользователя пусты")
	}
	// todo: здесь же можно добавить проверку на уникальность email и тому подобное
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)
	if err != nil {
		u.l.Error("Ошибка при хешировании пароля", err)
		return nil, fmt.Errorf("ошибка при хешировании пароля: %w", err)
	}

	userData := &models.User{
		Email:    reqData.Email,
		Password: string(hashedPassword),
		Role:     defaultUserRole,
	}

	return u.r.CreateUser(userData)
}

func (u *UsecaseImpl) GetUserByEmail(email string) (*models.User, error) {
	return u.r.GetUserByEmail(email)
}

func (u *UsecaseImpl) CheckIsAdmin(id int) bool {
	return u.r.CheckIsAdmin(id)
}

func (u *UsecaseImpl) Login(email, password string) (*models.TokenResponse, error) {
	user, err := u.GetUserByEmail(email)
	if err != nil {
		u.l.Error("Ошибка: пользователь не найден", err)
		return nil, errors.New("неверный email ")
	}

	// Сравниваем пароли
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("неверный пароль")
	}

	// Генерация токенов
	accessToken, err := utils.GenerateAccessToken(user.Email)
	if err != nil {
		u.l.Error("Ошибка при генерации access токена", err)
		return nil, fmt.Errorf("не удалось создать access token: %w", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Email)
	if err != nil {
		u.l.Error("Ошибка при генерации refresh токена", err)
		return nil, fmt.Errorf("не удалось создать refresh token: %w", err)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UsecaseImpl) Refresh(refreshToken string) (*models.TokenResponse, error) {
	claims, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("невалидный refresh токен")
	}

	user, err := u.GetUserByEmail(claims.Email)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	accessToken, err := utils.GenerateAccessToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации access токена: %w", err)
	}

	newRefreshToken, err := utils.GenerateRefreshToken(user.Email)
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации refresh токена: %w", err)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
