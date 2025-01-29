package api

import (
	"github.com/4strodev/4stroblog/site/server/api/controllers/session"
	"github.com/4strodev/4stroblog/site/server/api/controllers/user"
	"github.com/4strodev/4stroblog/site/server/core"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/4strodev/wiring/pkg/extended"
	"github.com/gofiber/fiber/v3"
)

var controllers = []core.Controller{
	&user.UserController{},
	&session.SessionController{},
}

type ApiController struct {
}

func (c *ApiController) Init(container wiring.Container) error {
	// Replace router by group router
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	api := router.Group("/api")

	extendedContainer := extended.Derived(container)
	err = extendedContainer.Singleton(func() fiber.Router {
		return api
	})
	if err != nil {
		return err
	}

	for _, controller := range controllers {
		err := controller.Init(extendedContainer)
		if err != nil {
			return err
		}
	}
	return nil
}
