package application

import (
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"github.com/google/uuid"
)

type SessionDeleteReq struct {
	ID uuid.UUID `json:"id"`
}

func (s *SessionService) Delete(req SessionDeleteReq) error {
	var session models.Session
	return s.DB.Delete(&session, req.ID).Error
}
