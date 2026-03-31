package application

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	userDomain "github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/config"
)

func NewSessionAppService(sessionService domain.SessionService, cfg config.Config) *SessionAppService {
	return &SessionAppService{
		SessionService: sessionService,
		Config:         cfg,
	}
}

type SessionAppService struct {
	SessionService domain.SessionService
	ProfileService userDomain
	Config         config.Config
}
