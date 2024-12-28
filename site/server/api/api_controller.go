package api

import (
	"github.com/4strodev/4stroblog/site/server/api/controllers/session"
	"github.com/4strodev/4stroblog/site/server/api/controllers/user"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/gofiber/fiber/v3"
)

var controllers = []core.Controller{
	&user.UserController{},
	&session.SessionController{},
}

type ApiController struct {
}

func (c *ApiController) Init(router fiber.Router) error {
	api := router.Group("/api")
	for _, controller := range controllers {
		err := controller.Init(api)
		if err != nil {
			return err
		}
	}
	return nil
}
