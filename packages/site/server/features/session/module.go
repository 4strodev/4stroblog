package session

import (
	sessionapp "github.com/4strodev/4stroblog/site/features/session/application"
	sessiondomain "github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/features/session/infrastructure"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var SessionFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) sessiondomain.SessionRepository {
			return &infrastructure.GormSessionRepository{DB: db}
		},
		func(repo sessiondomain.SessionRepository) sessiondomain.SessionService {
			return sessiondomain.SessionService{Repository: repo}
		},
		sessionapp.NewSessionAppService,
	},
	Name: "SessionFeatureModule",
}
