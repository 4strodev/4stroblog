package site

import (
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type SiteController struct {
}

func (c *SiteController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
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
