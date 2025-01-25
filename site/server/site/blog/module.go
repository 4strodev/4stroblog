package blog

import "github.com/4strodev/4stroblog/site/server/core"

var SiteBlogModule = core.Module{
	Singletons: []any{
		NewSiteBlogController,
	},
	Controllers: []core.Controller{
		&SiteBlogController{},
	},
	ExportSingletons: []any{
		NewSiteBlogController,
	},
}
