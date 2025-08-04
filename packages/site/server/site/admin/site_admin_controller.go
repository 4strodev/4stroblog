package admin

import (
	"log/slog"
	"strings"

	"github.com/4strodev/4stroblog/site/features/session/domain"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type SiteAdminController struct {
	JwtVerifier *domain.JwtVerify
	Logger    *slog.Logger
}

type SessionHeaders struct {
	Authorization string `header:"Authorization"`
}

func (c SiteAdminController) isLoggedOut(ctx fiber.Ctx) (bool, error) {
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

func (c SiteAdminController) isLoggedIn(ctx fiber.Ctx) (bool, error) {
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
	return c.isTokenValid(token), nil
}

func (c SiteAdminController) isTokenValid(token string) bool {
	err := c.JwtVerifier.Verify(token)

	if err != nil {
		c.Logger.Warn("invalid token provided")
		return false
	}
	return true
}

func (c *SiteAdminController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	router.Use("/site/admin", func(ctx fiber.Ctx) error {
		if strings.HasPrefix(ctx.Path(), "/site/admin/session") {
			logOut, err := c.isLoggedOut(ctx)
			if err != nil {
				return err
			}

			if logOut {
				return ctx.Next()
			}

			return ctx.Redirect().To("/site/admin")
		}
		logIn, err := c.isLoggedIn(ctx)
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
