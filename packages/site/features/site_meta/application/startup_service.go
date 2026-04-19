package application

import (
	"context"
	"log/slog"

	"github.com/4strodev/4stroblog/site/features/site_meta/domain"
	"github.com/4strodev/4stroblog/site/features/user/application"
	"github.com/4strodev/4stroblog/site/shared/config"
)

type StartupService struct {
	Logger              *slog.Logger
	SiteMetaRepository  domain.SiteMetaRepository
	UserRegisterService application.RegisterService
	Config              config.Config
}

func (s StartupService) Startup(ctx context.Context) error {
	s.Logger.Info("Running startup setup")
	siteMeta, err := s.SiteMetaRepository.Get(ctx)
	if err != nil {
		return err
	}

	if !siteMeta.FirstStart {
		return nil
	}

	userRegistered, err := s.UserRegisterService.Register(ctx, application.RegisterReqDTO{
		Email:    s.Config.Site.AdminEmail,
		Name:     "admin",
		Password: "1234",
	})

	if err != nil {
		return err
	}

	s.Logger.Info("admin user created",
		"email", s.Config.Site.AdminEmail,
		"id", userRegistered.UserID)

	return nil
}
