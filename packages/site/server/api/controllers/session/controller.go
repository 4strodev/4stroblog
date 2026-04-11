package session

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type SessionController struct {
	SessionService *application.SessionAppService
}

func (c *SessionController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	err = cont.Fill(c)
	if err != nil {
		return err
	}

	group := router.Group("/session")
	group.Post("/login", func(ctx fiber.Ctx) error {
		body := application.SessionCreateReq{}
		if err := ctx.Bind().Body(&body); err != nil {
			return err
		}

		response, err := c.SessionService.Create(ctx.Context(), body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
