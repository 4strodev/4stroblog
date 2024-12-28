// core saves building blocks to create an web application, controllers, middlewares, etc.
package core

import "github.com/gofiber/fiber/v3"

type Controller interface {
	Init(router fiber.Router) error
}

func LoadNestedControllers(router fiber.Router, controllers []Controller) error {
	for _, controller := range controllers {
		err := controller.Init(router)
		if err != nil {
			return err
		}
	}

	return nil
}
