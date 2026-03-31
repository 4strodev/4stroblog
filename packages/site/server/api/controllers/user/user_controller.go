package user

import (
	"github.com/4strodev/4stroblog/site/features/user/application"
	"github.com/4strodev/4stroblog/site/shared/domain/domainerrors"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	UserRegister *application.RegisterService
}

func (c *UserController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return domainerrors.WrapError(domainerrors.RUNTIME, err)
	}

	err = cont.Fill(c)
	if err != nil {
		return domainerrors.WrapError(domainerrors.RUNTIME, err)
	}

	userRouter := router.Group("/user")
	userRouter.Post("/register", func(ctx fiber.Ctx) error {
		body := application.RegisterReqDTO{}
		err := ctx.Bind().Body(&body)
		if err != nil {
			return err
		}

		response, err := c.UserRegister.Register(ctx.Context(), body)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	})
	return nil
}
