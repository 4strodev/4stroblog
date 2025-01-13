package site

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin"
	"github.com/4strodev/4stroblog/site/server/site/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/4strodev/4stroblog/site/server/site/session"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/4strodev/wiring/pkg/extended"
	"github.com/gofiber/fiber/v3"
)

type SiteController struct {
}

func (c *SiteController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	siteRouter := router.Group("site")

	// If no route is matched then go to not found
	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/")
	})

	derivedContainer := extended.Derived(container)
	err = derivedContainer.Singleton(func() fiber.Router {
		return siteRouter
	})
	if err != nil {
		return err
	}

	err = core.LoadNestedControllers(derivedContainer, []core.Controller{
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
