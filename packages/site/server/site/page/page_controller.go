package page

import (
	"path/filepath"
	"strings"

	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SitePageController struct {
	Prefix      string `wire:",ignore"`
	PagesFolder string `wire:",ignore"`
}

func (c *SitePageController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}

	groupRouter := router.Group(c.Prefix)
	groupRouter.Get("*", func(ctx fiber.Ctx) error {
		routePath := ctx.Path()
		subPage := strings.TrimPrefix(routePath, c.Prefix)
		page := filepath.Join("pages", c.PagesFolder, subPage)
		err := ctx.Render(page, nil, "layouts/main")
		if !TemplateNotFound(err) {
			return err
		}

		// Try for index
		indexPage := filepath.Join(page, "index")
		err = ctx.Render(indexPage, nil, "layouts/main")
		if !TemplateNotFound(err) {
			return err
		}

		return ctx.Redirect().To("/site/not-found")
	})
	return nil
}

func TemplateNotFound(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "does not exist")
}
