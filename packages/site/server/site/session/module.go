package session

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/features/session"
	"github.com/4strodev/4stroblog/site/server/features/user"
	"github.com/4strodev/4stroblog/site/shared"
)

var SiteSessionModule = core.Module{
	Imports: []*core.Module{
		&session.SessionFeatureModule,
		&user.UserFeatureModule,
		&shared.SharedModule,
	},
	Singletons: []any{
		domain.NewJwtVerify,
	},
	Controllers: []core.Controller{
		&SiteSessionController{},
	},
}
