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

	mockRepo.EXPECT().GetAllSessions().Return(expectedSessions, nil).Times(1)

	sessions, err := uc.GetAllSessions()

	assert.NoError(t, err)
	assert.Equal(t, expectedSessions, sessions)
}

func TestGetSessionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	exceptedSession := &models.Session{
		ID:        5,
		MovieID:   10,
		Movie:     models.Movie{ID: 10, Title: "Inception", Duration: 148},
		HallID:    5,
		Hall:      models.Hall{ID: 5, Name: "IMAX"},
		StartTime: time.Date(2025, 9, 5, 19, 0, 0, 0, time.UTC),
		Price:     1500.00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	mockRepo.EXPECT().GetSessionByID(5).Return(exceptedSession, nil).Times(1)

	session, err := uc.GetSessionByID(5)

	assert.NoError(t, err)
	assert.Equal(t, exceptedSession, session)
}

func TestDeleteSessionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	mockRepo.EXPECT().DeleteSessionByID(13).Return(true, nil).Times(1)

	isDeleted, err := uc.DeleteSessionByID(13)

	assert.NoError(t, err)
	assert.Equal(t, true, isDeleted)
}

func TestCreateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	reqBody := models.Session{
		ID:        5,
		MovieID:   10,
		Movie:     models.Movie{ID: 10, Title: "Inception", Duration: 148},
		HallID:    5,
		Hall:      models.Hall{ID: 5, Name: "IMAX"},
		StartTime: time.Date(2025, 9, 5, 19, 0, 0, 0, time.UTC),
		Price:     1500.00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	mockRepo.EXPECT().CreateSession(reqBody).Return(true, nil).Times(1)

	isCreated, err := uc.CreateSession(reqBody)

	assert.NoError(t, err)
	assert.Equal(t, true, isCreated)
}

func TestUpdateSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	session := models.Session{
		ID:        5,
		MovieID:   15,
		Movie:     models.Movie{ID: 1, Title: "Mortal Combat 3", Duration: 100},
		HallID:    5,
		Hall:      models.Hall{ID: 3, Name: "IMAX"},
		StartTime: time.Date(2025, 9, 5, 19, 0, 0, 0, time.UTC),
		Price:     900.00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	mockRepo.EXPECT().UpdateSession(1, session).Return(true, nil).Times(1)

	ok, err := uc.UpdateSession(1, session)

	assert.NoError(t, err)
	assert.True(t, ok)
}
