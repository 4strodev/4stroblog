package site

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin"
	"github.com/4strodev/4stroblog/site/server/site/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/4strodev/4stroblog/site/server/site/session"
	"github.com/4strodev/4stroblog/site/server/site/uploads"
)

var SiteModule = core.Module{
	Controllers: []core.Controller{
		&page.SitePageController{
			Prefix: "/site",
		},
		&SiteController{},
	},
	Imports: []*core.Module{
		&session.SiteSessionModule,
		&blog.SiteBlogModule,
		&admin.SiteAdminModule,
		&uploads.SiteUploadsModule,
	},
	Name: "SiteModule",
}
