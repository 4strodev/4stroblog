package infrastructure

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/shared/db/models"
)

func sessionModelToSession(sessionModel models.Session) domain.Session {
	return domain.Session{
		ID:             sessionModel.ID,
		UserID:         sessionModel.UserID,
		ExpriationTime: sessionModel.ExpirationTime,
		ProfileID:      sessionModel.ProfileID,
	}
}
