package blog

import (
	"github.com/4strodev/4stroblog/site/features/blog"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteBlogController struct {
}

func (c *SiteBlogController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	blogRouter := router.Group("/site/blog")

	blogRouter.Get("/render/post/:title", func(ctx fiber.Ctx) error {
		// Get post metadata
		// Validate publish date
		// Check if user can access to this post
		// Get post Markdown
		// Render post
		// Render template

		title := ctx.Params("title")

		html, err := blog.RenderPost(title)
		if err != nil {
			return err
		}
		return ctx.Status(200).Send(html)
	})

	blogRouter.Get("/post/:title", func(ctx fiber.Ctx) error {
		// Get post metadata
		// Validate publish date
		// Check if user can access to this post
		// Get post Markdown
		// Render post
		// Render template

		title := ctx.Params("title")

		return ctx.Render("scaffolds/post", fiber.Map{
			"Title": title,
		}, "layouts/main")
	})

	return nil
}
