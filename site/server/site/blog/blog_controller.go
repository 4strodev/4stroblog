package blog

import (
	"github.com/gofiber/fiber/v3"
)

type SiteBlogController struct {
}

func (c *SiteBlogController) Init(router fiber.Router) error {
	blogRouter := router.Group("blog")

	blogRouter.Get("/:title", func(ctx fiber.Ctx) error {
		// Get post metadata
		// Validate publish date
		// Check if user can access to this post
		// Get post Markdown
		// Render post
		// Render template
		return nil
	})

	return nil
}
