package site

import (
	wiring "github.com/4strodev/wiring/pkg"
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
	router.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/")
	})

	router.Get("*", func(ctx fiber.Ctx) error {
		return ctx.Redirect().To("/site/not-found")
	})
	return nil
}
