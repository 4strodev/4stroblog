package admin

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin/blog"
)

var SiteAdminModule = core.Module{
	Singletons: []any{
		NewSiteAdminController,
		blog.NewSiteAdminBlogController,
	},
	Controllers: []core.Controller{
		&SiteAdminController{},
		&blog.SiteAdminBlogController{},
	},
	ExportSingletons: []any{
		NewSiteAdminController,
		blog.NewSiteAdminBlogController,
	},
}
