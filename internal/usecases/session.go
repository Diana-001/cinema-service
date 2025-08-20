package usecases

import "cinema-service/internal/models"

type SessionUsecase interface {
	GetAllSessions() ([]models.Session, error)
	GetSessionByID(id int) (*models.Session, error)
	CreateSession(session models.Session) (bool, error)
	UpdateSession(id int, body models.Session) (bool, error)
	DeleteSessionByID(id int) (bool, error)
}

func (u *UsecaseImpl) GetAllSessions() ([]models.Session, error) {
	return u.r.GetAllSessions()
}

func (u *UsecaseImpl) GetSessionByID(id int) (*models.Session, error) {
	return u.r.GetSessionByID(id)
}

func (u *UsecaseImpl) CreateSession(session models.Session) (bool, error) {
	return u.r.CreateSession(session)
}

func (u *UsecaseImpl) UpdateSession(id int, body models.Session) (bool, error) {
	return u.r.UpdateSession(id, body)
}

func (u *UsecaseImpl) DeleteSessionByID(id int) (bool, error) {
	return u.r.DeleteSessionByID(id)
}
