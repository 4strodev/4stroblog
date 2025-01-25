package site

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
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
	SiteSessionController *session.SiteSessionController
	SiteBlogController    *blog.SiteBlogController
	SiteAdminController   *admin.SiteAdminController
	container             wiring.Container `wire:",ignore"`
}

func (c *SiteController) Init(container wiring.Container) error {
	var err error
	var sessionService *application.SessionService

	// Creating derived container
	c.container = extended.Derived(container)
	err = c.container.Resolve(&sessionService)
	if err != nil {
		return err
	}

	// Setting up routes
	var router fiber.Router
	err = container.Resolve(&router)
	if err != nil {
		return err
	}

	// If no route is matched then go to not found
	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/")
	})

	siteRouter := router.Group("site")
	err = c.container.Singleton(func() fiber.Router {
		return siteRouter
	})
	if err != nil {
		return err
	}

	err = c.container.Fill(c)
	if err != nil {
		return err
	}
	err = core.LoadNestedControllers(c.container, []core.Controller{
		//c.SiteSessionController,
		//c.SiteBlogController,
		//c.SiteAdminController,

		//// Page controllers has to be loaded the lasts ones
		//// to avoid routes collisions
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
