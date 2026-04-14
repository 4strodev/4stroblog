package session

import (
	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/features/session"
	"github.com/4strodev/4stroblog/site/server/features/user"
)

var SiteSessionModule = core.Module{
	Imports: []*core.Module{
		&session.SessionFeatureModule,
		&user.UserFeatureModule,
	},
	Singletons: []any{
		domain.NewJwtVerify,
	},
	Controllers: []core.Controller{
		&SiteSessionController{},
	},
	Name: "SiteSessionModule",
}
