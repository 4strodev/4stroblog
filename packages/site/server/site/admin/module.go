package admin

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/4strodev/4stroblog/site/shared"
)

var SiteAdminModule = core.Module{
	Imports: []*core.Module{
		&shared.SharedModule,
	},
	Controllers: []core.Controller{
		&SiteAdminController{},
		&blog.SiteAdminBlogController{},
		&page.SitePageController{
			Prefix:      "/site/admin",
			PagesFolder: "/admin",
		},
	},
}
