package redirect

import (
	"github.com/gofiber/fiber/v3"
)

type SiteRedirectController struct {
}

func (c *SiteRedirectController) Init(router fiber.Router) error {
	// If no route is matched then go to not found
	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/home")
	})
	router.All("*", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/not-found")
	})
	return nil
}
