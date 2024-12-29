package blog

import (
	"github.com/4strodev/4stroblog/site/application/blog"
	"github.com/gofiber/fiber/v3"
)

type SiteAdminBlogController struct {
}

type RenderMDReqDTO struct {
	Content string `form:"content"`
}

func (c *SiteAdminBlogController) Init(router fiber.Router) error {
	blogRouter := router.Group("/blog")
	blogRouter.Post("/md-render", func(ctx fiber.Ctx) error {
		body := RenderMDReqDTO{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		md := body.Content

		html := blog.RenderMarkdown([]byte(md))

		return ctx.SendString(string(html))
	})

	return nil
}
