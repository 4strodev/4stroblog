package page

import (
	"path/filepath"
	"strings"

	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

const defaultLayout = "layouts/main"

type SitePageController struct {
	Config      config.Config
	Prefix      string `wiring:",omit"`
	PagesFolder string `wiring:",omit"`
	// pagesMeta holds the front matter of each page, loaded once on Init
	pagesMeta map[string]PageMeta `wiring:",omit"`
}

func (c *SitePageController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	c.pagesMeta, err = LoadPagesMeta(c.Config.Views.Folder)
	if err != nil {
		return err
	}

	groupRouter := router.Group(c.Prefix)
	groupRouter.Get("*", func(ctx fiber.Ctx) error {
		routePath := ctx.Path()
		subPage := strings.TrimPrefix(routePath, c.Prefix)
		page := filepath.Join("pages", c.PagesFolder, subPage)
		err := c.render(ctx, page)
		if !TemplateNotFound(err) {
			return err
		}

		// Try for index
		indexPage := filepath.Join(page, "index")
		err = c.render(ctx, indexPage)
		if !TemplateNotFound(err) {
			return err
		}

		return ctx.Redirect().To("/site/not-found")
	})
	return nil
}

// render renders a page with the layout set in its front matter, or the main layout.
// The front matter is available in templates as .Meta
func (c *SitePageController) render(ctx fiber.Ctx, page string) error {
	meta := c.pagesMeta[page]
	layout := meta.Layout
	if layout == "" {
		layout = defaultLayout
	}
	return ctx.Render(page, fiber.Map{"Meta": meta}, layout)
}

func TemplateNotFound(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "does not exist")
}
