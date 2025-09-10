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

func TestGetAllMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	expectedMovies := []models.Movie{
		{ID: 1, Title: "Mortal Combat", Duration: 123},
		{ID: 2, Title: "Avatar", Duration: 360},
	}

	mockRepo.EXPECT().GetAllMovies().Return(expectedMovies, nil).Times(1)

	movies, err := uc.GetAllMovies()

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, movies)
}

func TestGetMoviesByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	expectedMovies := models.Movie{ID: 1, Title: "Mortal Combat", Duration: 123}

	mockRepo.EXPECT().GetMovieByID(1).Return(expectedMovies, nil).Times(1)

	movies, err := uc.GetMovieByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, movies)
}

func TestDeleteMoviesByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	mockRepo.EXPECT().DeleteMovieByID(1).Return(true, nil).Times(1)

	isDeleted, err := uc.DeleteMovieByID(1)

	assert.NoError(t, err)
	assert.Equal(t, true, isDeleted)
}

func TestCreateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	reqBody := models.Movie{ID: 1, Title: "Mortal Combat", Duration: 123}

	mockRepo.EXPECT().CreateMovie(reqBody).Return(true, nil).Times(1)

	isCreated, err := uc.CreateMovie(reqBody)

	assert.NoError(t, err)
	assert.Equal(t, true, isCreated)
}

func TestUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	uc := &usecases.UsecaseImpl{
		R: mockRepo,
	}

	ctx := context.Background()
	movie := models.Movie{ID: 1, Title: "Mortal Combat 3", Duration: 100}

	// Ожидания
	mockRepo.EXPECT().CheckIsAdmin(123).Return(true).Times(1)
	mockRepo.EXPECT().UpdateMovie(1, movie).Return(true, nil).Times(1)

	ok, err := uc.UpdateMovie(1, 123, movie, ctx)

	assert.NoError(t, err)
	assert.True(t, ok)
}
