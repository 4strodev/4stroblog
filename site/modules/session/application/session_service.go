package application

import (
	"github.com/4strodev/4stroblog/site/shared/config"
	"gorm.io/gorm"
)

func NewSessionService(db *gorm.DB, cfg config.Config) *SessionService {
	return &SessionService{
		DB:     db,
		Config: cfg,
	}
}

type SessionService struct {
	DB     *gorm.DB
	Config config.Config
}
