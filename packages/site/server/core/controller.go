// core saves building blocks to create an web application, controllers, middlewares, etc.
package core

import (
	"github.com/4strodev/wiring_graphs/pkg/container"
)

type Controller interface {
	Init(container *container.Container) error
}

func LoadNestedControllers(container *container.Container, controllers []Controller) error {
	for _, controller := range controllers {
		err := controller.Init(container)
		if err != nil {
			return err
		}
	}

	return nil
}
