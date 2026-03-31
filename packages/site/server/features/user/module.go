package user

import (
	"github.com/4strodev/4stroblog/site/features/user/domain"
	"github.com/4strodev/4stroblog/site/features/user/infrastructure"
	"github.com/4strodev/4stroblog/site/features/user/application"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var UserFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) domain.UserRepository {
			return &infrastructure.GormUserRepository{
				DB: db,
			}
		},
		func(repository domain.UserRepository) *domain.UserService {
			return &domain.UserService{
				Repository: repository,
			}
		},
	},
	ExportSingletons: []any{
		func(service *domain.UserService) *application.RegisterService {
			return &application.RegisterService{
				UserService: service,
			}
		},
	},
}
