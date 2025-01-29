package admin

import (
	"strings"

	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteAdminController struct {
}

type SessionHeaders struct {
	Authorization string `header:"Authorization"`
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

	return nil
}
