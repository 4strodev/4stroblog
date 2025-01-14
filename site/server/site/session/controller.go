package session

import (
	"net/http"

	"github.com/4strodev/4stroblog/site/modules/session/application"
	"github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
)

type SiteSessionController struct {
	loginService application.SessionService
}

func NewSiteSessionController(loginService *application.SessionService) *SiteSessionController {
	return &SiteSessionController{
		loginService: *loginService,
	}
}

func (c *SiteSessionController) Init(container pkg.Container) error {
	var router fiber.Router
	err := container.Resolve(&router)
	if err != nil {
		return err
	}
	group := router.Group("/session")

	group.Post("/login", func(ctx fiber.Ctx) error {
		var req application.SessionCreateReq
		if err := ctx.Bind().Body(&req); err != nil {
			return err
		}

		res, err := c.loginService.Create(req)
		if err != nil {
			return err
		}

		ctx.Response().Header.Add("HX-Location", "/site")
		return ctx.Status(http.StatusCreated).JSON(res)
	})

	return nil
}
