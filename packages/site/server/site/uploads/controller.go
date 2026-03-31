package uploads

import (
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type SiteUploadsController struct {
}

func (c *SiteUploadsController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	router.Group("/site/uploads")
	return nil
}
