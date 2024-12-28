package blog

import (
	"log"

	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/gofiber/fiber/v3"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

type SiteAdminBlogController struct {
}

type RenderMDReqDTO struct {
	Content string `form:"content"`
}

func (c *SiteAdminBlogController) Init(router fiber.Router) error {
	blogRouter := router.Group("/blog")
	blogRouter.Post("/md-render", func(ctx fiber.Ctx) error {
		log.Println("rendering markdown")
		body := RenderMDReqDTO{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		md := body.Content

		extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
		p := parser.NewWithExtensions(extensions)
		doc := p.Parse([]byte(md))
		htmlFlags := html.CommonFlags | html.HrefTargetBlank
		opts := html.RendererOptions{Flags: htmlFlags}
		renderer := html.NewRenderer(opts)

		rawHtml := markdown.Render(doc, renderer)
		html := bluemonday.UGCPolicy().SanitizeBytes(rawHtml)

		return ctx.SendString(string(html))
	})

	return core.LoadNestedControllers(blogRouter, []core.Controller{
		&page.SitePageController{
			Prefix: "admin/blog",
		},
	})
}
