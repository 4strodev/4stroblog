package user

import (
	"github.com/4strodev/4stroblog/site/features/user/services"
	"github.com/4strodev/4stroblog/site/shared/db"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
}

func (init *UserController) Init(container wiring.Container) error {
	db, err := db.GetDbInstance()
	if err != nil {
		return err
	}

	var router fiber.Router
	err = container.Resolve(&router)
	if err != nil {
		return err
	}

	group := router.Group("/user")
	group.Post("/register", func(ctx fiber.Ctx) error {
		registerService := services.RegisterService{
			DB: db,
		}
		body := services.RegisterReqDTO{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		response, err := registerService.Register(body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
