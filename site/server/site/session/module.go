package session

import "github.com/4strodev/4stroblog/site/server/core"

var SiteSessionModule = core.Module{
	Singletons: []any{
		NewSiteSessionController,
	},
	Controllers: []core.Controller{
		&SiteSessionController{},
	},
	ExportSingletons: []any{
		NewSiteSessionController,
	},
}
