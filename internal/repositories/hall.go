package repositories

import (
	"cinema-service/internal/models"
)

type HallRepository interface {
	GetAllHalls() ([]models.Hall, error)
	GetHallByID(id int) (models.Hall, error)
	DeleteHallByID(id int) (bool, error)
	CreateHall(body models.Hall) (bool, error)
	UpdateHall(id int, body models.Hall) (bool, error)
}

func (repo *RepositoryImpl) GetAllHalls() ([]models.Hall, error) {
	var halls []models.Hall
	if err := repo.db.Find(&halls).Error; err != nil {
		return nil, err
	}
	return halls, nil
}

func (repo *RepositoryImpl) GetHallByID(id int) (models.Hall, error) {
	var hall models.Hall
	if err := repo.db.First(&hall, id).Error; err != nil {
		return models.Hall{}, err
	}
	return hall, nil
}

func (repo *RepositoryImpl) DeleteHallByID(id int) (bool, error) {
	if err := repo.db.Delete(&models.Movie{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *RepositoryImpl) CreateHall(body models.Hall) (bool, error) {
	if err := repo.db.Create(&body).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *RepositoryImpl) UpdateHall(id int, body models.Hall) (bool, error) {
	var hall models.Hall

	// Найти зал по id
	if err := repo.db.First(&hall, id).Error; err != nil {
		return false, err
	}

	// Обновить поля
	if err := repo.db.Model(&hall).Updates(models.Hall{
		Name:        body.Name,
		Rows:        body.Rows,
		SeatsPerRow: body.SeatsPerRow,
	}).Error; err != nil {
		return false, err
	}

	return true, nil
}
