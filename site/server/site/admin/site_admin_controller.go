package admin

import (
	"strings"

	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site/admin/blog"
	"github.com/4strodev/4stroblog/site/server/site/page"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/4strodev/wiring/pkg/extended"
	"github.com/gofiber/fiber/v3"
)

func NewSiteAdminController(siteAdminBlogController *blog.SiteAdminBlogController) *SiteAdminController {
	return &SiteAdminController{
		SiteAdminBlogController: siteAdminBlogController,
	}
}

type SiteAdminController struct {
	SiteAdminBlogController *blog.SiteAdminBlogController
}

type SessionHeaders struct {
	Authorization string `header:"Authorization"`
}

var resolvers = []any{
	blog.NewSiteAdminBlogController,
}

func isLoggedOut(ctx fiber.Ctx) (bool, error) {
	headers := SessionHeaders{}
	err := ctx.Bind().Header(&headers)
	if err != nil {
		return false, err
	}

	if headers.Authorization == "" {
		return true, nil
	}

	return false, nil
}

func isLoggedIn(ctx fiber.Ctx) (bool, error) {
	headers := SessionHeaders{}
	err := ctx.Bind().Header(&headers)
	if err != nil {
		return false, err
	}

	if headers.Authorization == "" {
		return false, nil
	}

	splitHeader := strings.Split(headers.Authorization, " ")
	if splitHeader[0] != "Bearer" {
		return false, nil
	}

	token := splitHeader[1]
	// Check token
	return isTokenValid(token), nil
}

func isTokenValid(token string) bool {
	return true
}

func (c *SiteAdminController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	adminRouter := router.Group("/admin")
	derivedContainer := extended.Derived(container)

	adminRouter.Use(func(ctx fiber.Ctx) error {
		if strings.HasPrefix(ctx.Path(), "/site/admin/session") {
			logOut, err := isLoggedOut(ctx)
			if err != nil {
				return err
			}

			if logOut {
				return ctx.Next()
			}

			return ctx.Redirect().To("/site/admin")
		}
		logIn, err := isLoggedIn(ctx)
		if err != nil {
			return err
		}

		if logIn || true {
			return ctx.Next()
		}

		return ctx.Redirect().To("/site/not-found")
	})

	err = derivedContainer.Singleton(func() fiber.Router {
		return adminRouter
	})
	if err != nil {
		return err
	}

	for _, resolver := range resolvers {
		err = derivedContainer.Singleton(resolver)
		if err != nil {
			return err
		}
	}

	err = derivedContainer.Fill(c)
	if err != nil {
		return err
	}
	return core.LoadNestedControllers(derivedContainer, []core.Controller{
		c.SiteAdminBlogController,
		&page.SitePageController{
			Prefix:      "/site/admin",
			PagesFolder: "/admin",
		},
	})
}
