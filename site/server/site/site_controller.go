package site

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin"
	"github.com/4strodev/4stroblog/site/server/site/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/4strodev/4stroblog/site/server/site/redirect"
	"github.com/4strodev/4stroblog/site/server/site/session"
	"github.com/gofiber/fiber/v3"
)

type SiteController struct {
}

func (c *SiteController) Init(router fiber.Router) error {
	redirectController := &redirect.SiteRedirectController{}
	site := router.Group("site")

	err := core.LoadNestedControllers(site, []core.Controller{
		&session.SiteSessionController{},
		&blog.SiteBlogController{},
		&admin.SiteAdminController{},

		// Page controllers has to be loaded the lasts ones
		// to avoid routes collisions
		&page.SitePageController{},
	})
	if err != nil {
		return err
	}

	return redirectController.Init(router)
}
