package usecases

import (
	"cinema-service/internal/mocks"
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestGetAllSessions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	expectedSessions := []models.Session{
		{ID: 1,
			MovieID:   10,
			Movie:     models.Movie{ID: 10, Title: "Inception", Duration: 148},
			HallID:    5,
			Hall:      models.Hall{ID: 5, Name: "IMAX"},
			StartTime: time.Date(2025, 9, 5, 19, 0, 0, 0, time.UTC),
			Price:     1500.00,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()},
	}

	mockRepo.EXPECT().GetAllSessions().Return(expectedSessions, nil)

	sessions, err := uc.GetAllSessions()

	assert.NoError(t, err)
	assert.Equal(t, expectedSessions, sessions)
}
