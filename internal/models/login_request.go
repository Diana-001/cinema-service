package models

import (
	"errors"
	"net/mail"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() error {
	if l.Email == "" {
		return errors.New("email обязателен")
	}

	// Проверка валидности email
	if _, err := mail.ParseAddress(l.Email); err != nil {
		return errors.New("некорректный email")
	}

	if len(l.Password) < 6 {
		return errors.New("пароль должен быть не меньше 6 символов")
	}

	return nil
}

func (l *LoginRequest) ToCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Email:    l.Email,
		Password: l.Password,
	}
}
