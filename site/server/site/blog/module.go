package blog

import "github.com/4strodev/4stroblog/site/server/core"

var SiteBlogModule = core.Module{
	Controllers: []core.Controller{
		&SiteBlogController{},
	},
}
