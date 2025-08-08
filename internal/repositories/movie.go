package repositories

import (
	"cinema-service/internal/models"
)

type MovieRepository interface {
	GetAll() ([]models.Movie, error)
	GetMovieByID(id int) (models.Movie, error)
	DeleteMovieByID(id int) (bool, error)
	CreateMovie(body models.Movie) (bool, error)
	UpdateMovie(id int, body models.Movie) (bool, error)
}

func (repo *RepositoryImpl) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	if err := repo.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (repo *RepositoryImpl) GetMovieByID(id int) (models.Movie, error) {
	var movie models.Movie
	if err := repo.db.First(&movie, id).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (repo *RepositoryImpl) DeleteMovieByID(id int) (bool, error) {
	if err := repo.db.Delete(&models.Movie{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *RepositoryImpl) CreateMovie(body models.Movie) (bool, error) {
	if err := repo.db.Create(&body).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *RepositoryImpl) UpdateMovie(id int, body models.Movie) (bool, error) {
	var movie models.Movie

	// Найти фильм по id
	if err := repo.db.First(&movie, id).Error; err != nil {
		return false, err
	}

	// Обновить поля
	if err := repo.db.Model(&movie).Updates(models.Movie{
		Title:       body.Title,
		Description: body.Description,
		Duration:    body.Duration,
		PosterUrl:   body.PosterUrl,
	}).Error; err != nil {
		return false, err
	}

	return true, nil
}
