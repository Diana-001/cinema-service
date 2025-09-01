package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) Validate() error {
	// todo: валидация входных параметров
	return nil
}

func (l *LoginRequest) ToCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Email:    l.Email,
		Password: l.Password,
	}
}
