package uploads

import (
	"net/http"

	"github.com/4strodev/4stroblog/site/features/uploads/application"
	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

type SiteUploadsController struct {
	UploadService *application.UploadsService
}

func (c *SiteUploadsController) Init(cont *container.Container) error {
	router, err := container.Resolve[fiber.Router](cont)
	if err != nil {
		return err
	}

	group := router.Group("/site/uploads")

	group.Post("/", func(ctx fiber.Ctx) error {
		document, err := ctx.FormFile("document")
		if err != nil {
			return err
		}

		file, err := document.Open()
		if err != nil {
			return err
		}

		err = c.UploadService.UploadBlob(ctx.Context(), &domain.Upload{}, file)
		if err != nil {
			return err
		}

		return ctx.SendStatus(http.StatusCreated)
	})
	return nil
}
