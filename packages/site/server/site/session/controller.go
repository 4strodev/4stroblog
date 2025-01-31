package session

import (
	"net/http"

	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/server"
	"github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteSessionController struct {
	sessionService *application.SessionService
}

func (c *SiteSessionController) Init(container pkg.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	group := router.Group("/site/session")

	group.Post("/", func(ctx fiber.Ctx) error {
		var body struct {
			User     string `form:"user"`
			Password string `form:"password"`
		}
		var req application.SessionCreateReq
		if err := ctx.Bind().Body(&body); err != nil {
			return err
		}

		req = application.SessionCreateReq{
			User:     body.User,
			Password: body.Password,
		}

		res, err := c.sessionService.Create(ctx.Context(), req)
		if err != nil {
			return err
		}

		ctx.Response().Header.Add("HX-Location", "/")
		ctx.Cookie(&fiber.Cookie{
			HTTPOnly: true,
			Name:     server.SESSION_HEADER,
			Value:    res.ID.String(),
			SameSite: "Strict",
		})

		return ctx.SendStatus(http.StatusCreated)
	})

	group.Delete("/:id", func(ctx fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	return nil
}
