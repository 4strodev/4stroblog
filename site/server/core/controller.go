// core saves building blocks to create an web application, controllers, middlewares, etc.
package core

import (
	wiring "github.com/4strodev/wiring/pkg"
)

type Controller interface {
	Init(container wiring.Container) error
}

func LoadNestedControllers(container wiring.Container, controllers []Controller) error {
	for _, controller := range controllers {
		err := controller.Init(container)
		if err != nil {
			return err
		}
	}

	return nil
}
