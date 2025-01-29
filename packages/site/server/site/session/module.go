package session

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/server/core"
)

var SiteSessionModule = core.Module{
	Singletons: []any{
		application.NewSessionService,
	},
	Controllers: []core.Controller{
		&SiteSessionController{},
	},
}
