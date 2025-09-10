package usecases

import (
	"cinema-service/internal/mocks"
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"cinema-service/internal/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	reqPassword := "somePassword"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqPassword), bcrypt.DefaultCost)
	require.NoError(t, err)

	mockRepo.EXPECT().
		GetUserByEmail("test@example.com").
		Return(&models.User{
			Email:    "test@example.com",
			Password: string(hashedPassword),
		}, nil).Times(1)

	uc := usecases.New(mockRepo, nil)

	tokens, err := uc.Login(models.LoginRequest{
		Email:    "test@example.com",
		Password: reqPassword,
	})

	assert.NoError(t, err)
	assert.NotNil(t, tokens)
}

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	ctx := context.Background()

	newUser := &models.CreateUserRequest{
		Email:    "test@example.com",
		Password: "qwerty123",
	}

	expectedResponse := &models.User{
		ID:    1,
		Email: newUser.Email,
		Role:  "user",
	}

	mockRepo.EXPECT().
		CreateUser(gomock.AssignableToTypeOf(&models.User{})).
		Return(expectedResponse, nil).Times(1)

	user, err := uc.CreateUser(ctx, newUser)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, user)
}

func TestRefresh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	email := "test@example.com"

	refreshToken, err := utils.GenerateRefreshToken(email)
	require.NoError(t, err)
	require.NotEmpty(t, refreshToken)

	expectedUser := &models.User{
		ID:    1,
		Email: email,
		Role:  "user",
	}
	mockRepo.EXPECT().
		GetUserByEmail(email).
		Return(expectedUser, nil).Times(1).Times(1)

	resp, err := uc.Refresh(refreshToken)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.AccessToken)
	assert.NotEmpty(t, resp.RefreshToken)
}
