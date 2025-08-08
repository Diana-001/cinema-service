package repositories

import (
	"cinema-service/internal/models"
)

type SessionRepository interface {
	GetAllSessions() ([]models.Session, error)
	GetSessionByID(id int) (*models.Session, error)
	CreateSession(session models.Session) (bool, error)
	UpdateSession(id int, body models.Session) (bool, error)
	DeleteSessionByID(id int) (bool, error)
}

func (repo *RepositoryImpl) GetAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	if err := repo.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (repo *RepositoryImpl) GetSessionByID(id int) (*models.Session, error) {
	var session models.Session
	if err := repo.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *RepositoryImpl) CreateSession(session models.Session) (bool, error) {
	if err := repo.db.Create(&session).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *RepositoryImpl) UpdateSession(id int, body models.Session) (bool, error) {
	var session models.Session
	if err := repo.db.First(&session, id).Error; err != nil {
		return false, err
	}

	if err := repo.db.Model(&session).Updates(models.Session{
		MovieID:   body.MovieID,
		HallID:    body.HallID,
		StartTime: body.StartTime,
		Price:     body.Price,
	}).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (repo *RepositoryImpl) DeleteSessionByID(id int) (bool, error) {
	if err := repo.db.Delete(&models.Session{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
