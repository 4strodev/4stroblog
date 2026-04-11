package user

import (
	"github.com/4strodev/4stroblog/site/features/user/application"
	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/features/user/infrastructure"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var UserFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) domain.UserRepository {
			return &infrastructure.GormUserRepository{DB: db}
		},
		func(profileRepo domain.ProfileRepository) *domain.ProfileService {
			return &domain.ProfileService{ProfileRepository: profileRepo}
		},
	},
	ExportSingletons: []any{
		func(userRepo domain.UserRepository, profileService *domain.ProfileService) *application.RegisterService {
			return &application.RegisterService{
				UserRepository: userRepo,
				ProfileService: profileService,
			}
		},
		func(db *gorm.DB) domain.ProfileRepository {
			return &infrastructure.GormProfileRepository{DB: db}
		},
	},
}
