package usecases

import (
	"cinema-service/internal/mocks"
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUpdateHall_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	ctx := context.Background()
	hall := models.Hall{ID: 1, Name: "IMAX"}

	// Ожидания
	mockRepo.EXPECT().CheckIsAdmin(123).Return(true)
	mockRepo.EXPECT().UpdateHall(1, hall).Return(true, nil)

	ok, err := uc.UpdateHall(1, 123, hall, ctx)

	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestGetAllHalls(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	expectedHalls := []models.Hall{
		{ID: 1, Name: "Games of thrones"},
		{ID: 2, Name: "300 spartans"},
	}

	mockRepo.EXPECT().GetAllHalls().Return(expectedHalls, nil)

	halls, err := uc.GetAllHalls()

	assert.NoError(t, err)
	assert.Equal(t, expectedHalls, halls)
}
