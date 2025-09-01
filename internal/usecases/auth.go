package usecases

import (
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
}

func (u *UsecaseImpl) CreateUser(ctx context.Context, reqData *models.CreateUserRequest) (*models.User, error) {
	if reqData == nil {
		u.l.WarnCtx(ctx, "Ошибка при создании пользователя: данные пользователя пусты")
		return nil, errors.New("данные пользователя пусты")
	}
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
