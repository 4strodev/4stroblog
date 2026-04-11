package application

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	userdomain "github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/shared/config"
)

func NewSessionAppService(
	sessionService domain.SessionService,
	profileRepo userdomain.ProfileRepository,
	userRepo userdomain.UserRepository,
	cfg config.Config,
) *SessionAppService {
	return &SessionAppService{
		SessionService:    sessionService,
		ProfileRepository: profileRepo,
		UserRepository:    userRepo,
		Config:            cfg,
	}
}

type SessionAppService struct {
	SessionService    domain.SessionService
	ProfileRepository userdomain.ProfileRepository
	UserRepository    userdomain.UserRepository
	Config            config.Config
}
