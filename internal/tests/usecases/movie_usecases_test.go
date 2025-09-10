package usecases

import (
	"cinema-service/internal/mocks"
	"cinema-service/internal/models"
	"cinema-service/internal/usecases"
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

	mockRepo.EXPECT().GetAllMovies().Return(expectedMovies, nil)

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

	mockRepo.EXPECT().GetMovieByID(1).Return(expectedMovies, nil)

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

	mockRepo.EXPECT().DeleteMovieByID(1).Return(true, nil)

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

	mockRepo.EXPECT().CreateMovie(reqBody).Return(true, nil)

	isCreated, err := uc.CreateMovie(reqBody)

	assert.NoError(t, err)
	assert.Equal(t, true, isCreated)
}
