package usecases

import (
	"cinema-service/internal/mocks"
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	mockRepo.EXPECT().
		GetUserByEmail("test@example.com").
		Return(&models.User{
			Email:    "test@example.com",
			Password: "$2a$10$hashПароля",
		}, nil)

	uc := usecases.New(mockRepo, nil)

	tokens, err := uc.Login(models.LoginRequest{
		Email:    "test@example.com",
		Password: "правильныйПароль",
	})

	assert.NoError(t, err)
	assert.NotNil(t, tokens)
}
