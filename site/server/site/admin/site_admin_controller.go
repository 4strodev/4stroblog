package admin

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	"github.com/gofiber/fiber/v3"
)

type SiteAdminController struct {
}

func (c *SiteAdminController) Init(router fiber.Router) error {
	adminRouter := router.Group("/admin")

	return core.LoadNestedControllers(adminRouter, []core.Controller{
		&blog.SiteAdminBlogController{},
		&page.SitePageController{
			Prefix:       "admin",
		},
	})
}
