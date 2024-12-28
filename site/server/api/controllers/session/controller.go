package session

import (
	"github.com/4strodev/4stroblog/site/application/session/services"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	"github.com/gofiber/fiber/v3"
)

type SessionController struct {
}

func (init *SessionController) Init(router fiber.Router) error {
	db, err := db.GetDbInstance()
	if err != nil {
		return err
	}
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	group := router.Group("/session")
	group.Post("/login", func(ctx fiber.Ctx) error {
		loginService := services.LoginService{
			DB:     db,
			Config: config,
		}
		body := services.LoginReqDTO{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		response, err := loginService.Login(body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
