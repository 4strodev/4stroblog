package session

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/features/session/infrastructure"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var SessionFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) domain.SessionRepository {
			return &infrastructure.GormSessionRepository{
				DB: db,
			}
		},
	},
}
