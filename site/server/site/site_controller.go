package site

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin"
	"github.com/4strodev/4stroblog/site/server/site/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/4strodev/4stroblog/site/server/site/session"
	"github.com/gofiber/fiber/v3"
)

type SiteController struct {
}

func (c *SiteController) Init(router fiber.Router) error {
	siteRouter := router.Group("site")

	// If no route is matched then go to not found
	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/")
	})

	err := core.LoadNestedControllers(siteRouter, []core.Controller{
		&session.SiteSessionController{},
		&blog.SiteBlogController{},
		&admin.SiteAdminController{},

		// Page controllers has to be loaded the lasts ones
		// to avoid routes collisions
		&page.SitePageController{
			Prefix: "/site",
		},
	})

	if err != nil {
		return err
	}

	router.Get("*", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/not-found")
	})

	return nil
}
