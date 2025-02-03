package blog

import (
	"github.com/4strodev/4stroblog/site/features/blog"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteAdminBlogController struct {
}

type RenderMDReqDTO struct {
	Content string `form:"content"`
}

func (c *SiteAdminBlogController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}

	router.Post("/site/admin/blog/md-render", func(ctx fiber.Ctx) error {
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
