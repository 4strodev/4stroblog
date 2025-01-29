package session

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/shared/config"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type SessionController struct {
	Db *gorm.DB
}

func (c *SessionController) Init(container wiring.Container) error {
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	var router fiber.Router
	err = container.Resolve(&router)
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

		response, err := loginService.Create(body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
