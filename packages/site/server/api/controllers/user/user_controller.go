package user

import (
	"github.com/4strodev/4stroblog/site/features/user/services"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func (c *UserController) Init(container wiring.Container) error {
	var err error
	var router fiber.Router
	err = container.Resolve(&router)
	if err != nil {
		return err
	}

	group := router.Group("/user")
	group.Post("/register", func(ctx fiber.Ctx) error {
		registerService := services.RegisterService{
			DB: c.Db,
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
