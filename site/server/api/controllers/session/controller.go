package session

import (
	"github.com/4strodev/4stroblog/site/features/session/application"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SessionController struct {
}

func (init *SessionController) Init(container wiring.Container) error {
	db, err := db.GetDbInstance()
	if err != nil {
		return err
	}
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
			DB:     db,
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
