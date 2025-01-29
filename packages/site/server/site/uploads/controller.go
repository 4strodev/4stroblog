package uploads

import (
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteUploadsController struct {
}

func (c *SiteUploadsController) Init(container wiring.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}

	router.Group("/site/uploads")
	return nil
}
