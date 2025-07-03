package sessions

import (
	"cinema-service/internal/models"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (repo *SessionRepository) GetAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	if err := repo.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (repo *SessionRepository) GetSessionByID(id int) (*models.Session, error) {
	var session models.Session
	if err := repo.db.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *SessionRepository) CreateSession(session models.Session) (bool, error) {
	if err := repo.db.Create(&session).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *SessionRepository) UpdateSession(id int, body models.Session) (bool, error) {
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

func (repo *SessionRepository) DeleteByID(id int) (bool, error) {
	if err := repo.db.Delete(&models.Session{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
