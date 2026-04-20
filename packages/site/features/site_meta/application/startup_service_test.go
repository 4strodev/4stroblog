package application_test

import (
	"testing"

	"github.com/4strodev/4stroblog/site/features/site_meta/application"
	"github.com/4strodev/4stroblog/site/features/site_meta/infrastructure"
	userapplication "github.com/4strodev/4stroblog/site/features/user/application"
	userdomain "github.com/4strodev/4stroblog/site/features/user/domain"
	userinfrastructure "github.com/4strodev/4stroblog/site/features/user/infrastructure"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/logger"
)

func TestCreateAdminOnFirstStartup(t *testing.T) {
	startupService := application.StartupService{
		Logger: logger.NewLogger(),
		Config: config.Config{
			Site: config.SiteConfig{
				AdminEmail: "admin@example.com",
			},
		},
		SiteMetaRepository: &infrastructure.InMemorySiteMetaRepository{},
		UserRegisterService: userapplication.RegisterService{
			UserRepository: &userinfrastructure.InMemoryUserRepository{},
			ProfileService: &userdomain.ProfileService{
				ProfileRepository: &userinfrastructure.InMemoryProfileRepository{},
			},
		},
	}
}
