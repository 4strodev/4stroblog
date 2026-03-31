package api

import (
	"github.com/4strodev/4stroblog/site/server/api/controllers/session"
	"github.com/4strodev/4stroblog/site/server/api/controllers/user"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

var controllers = []core.Controller{
	&user.UserController{},
	&session.SessionController{},
}

type ApiController struct {
}

func (c *ApiController) Init(cont *container.Container) error {
	// Replace router by group router
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}
	api := router.Group("/api")

	extendedContainer := cont.Derived()
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
