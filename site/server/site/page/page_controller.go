package page

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type SitePageController struct {
	Prefix string
	NestedRouter bool
}

func (c *SitePageController) Init(router fiber.Router) error {
	router.Get("/:page", func(ctx fiber.Ctx) error {
		subPage := ctx.Params("page")
		page := filepath.Join("pages", c.Prefix, subPage)
		err := ctx.Render(page, nil, "layouts/main")
		if err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				return ctx.Redirect().To("/site/not-found")
			}
		}

		return nil
	})
	return nil
}
