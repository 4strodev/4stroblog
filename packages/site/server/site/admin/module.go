package admin

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
)

var SiteAdminModule = core.Module{
	Controllers: []core.Controller{
		&SiteAdminController{},
		&blog.SiteAdminBlogController{},
		&page.SitePageController{
			Prefix:      "/site/admin",
			PagesFolder: "/admin",
		},
	},
}
