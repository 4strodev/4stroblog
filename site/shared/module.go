package shared

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
)

var SharedModule = core.Module{
	Singletons: []any{
		db.GetDbInstance,
		config.GetConfig,
	},
	Transients: []any{
		application.NewSessionService,
	},
}
