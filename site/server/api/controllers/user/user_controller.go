package user

import (
	"github.com/4strodev/4stroblog/site/application/user/services"
	"github.com/4strodev/4stroblog/site/shared/db"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
}

func (init *UserController) Init(router fiber.Router) error {
	db, err := db.GetDbInstance()
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
