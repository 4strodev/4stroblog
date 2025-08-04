package session

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type SessionController struct {
	Db *gorm.DB
}

func (c *SessionController) Init(cont *container.Container) error {
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	group := router.Group("/session")
	group.Post("/login", func(ctx fiber.Ctx) error {
		loginService := application.SessionService{
			DB:     c.Db,
			Config: config,
		}
		body := application.SessionCreateReq{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		response, err := loginService.Create(ctx.Context(), body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
