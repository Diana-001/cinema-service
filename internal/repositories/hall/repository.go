package hall

import (
	"cinema-service/internal/models"
	"gorm.io/gorm"
)

type HallRepository struct {
	db *gorm.DB
}

func NewHallRepository(db *gorm.DB) *HallRepository {
	return &HallRepository{
		db: db,
	}
}

func (repo *HallRepository) GetAllHalls() ([]models.Hall, error) {
	var halls []models.Hall
	if err := repo.db.Find(&halls).Error; err != nil {
		return nil, err
	}
	return halls, nil
}

func (repo *HallRepository) GetHallByID(id int) (models.Hall, error) {
	var hall models.Hall
	if err := repo.db.First(&hall, id).Error; err != nil {
		return models.Hall{}, err
	}
	return hall, nil
}

func (repo *HallRepository) DeleteByID(id int) (bool, error) {
	if err := repo.db.Delete(&models.Movie{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *HallRepository) CreateHall(body models.Hall) (bool, error) {
	if err := repo.db.Create(&body).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *HallRepository) UpdateHall(id int, body models.Hall) (bool, error) {
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
