package uploads

import (
	"net/http"

	"github.com/4strodev/4stroblog/site/features/uploads/application"
	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

		err = c.UploadService.SaveUpload(ctx.Context(), &domain.Upload{}, file)
		if err != nil {
			return err
		}

		return ctx.SendStatus(http.StatusCreated)
	})

	group.Delete("/:id", func (ctx fiber.Ctx) error {
		id := ctx.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		return c.UploadService.DeleteUpload(ctx.Context(), uuid)
	})

	return nil
}
