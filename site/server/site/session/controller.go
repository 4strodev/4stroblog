package session

import (
	"net/http"

	"github.com/4strodev/4stroblog/site/application/session/services"
	"github.com/gofiber/fiber/v3"
)

type SiteSessionController struct {
	loginService services.LoginService
}

func NewSiteSessionController(loginService *services.LoginService) *SiteSessionController {
	return &SiteSessionController{
		loginService: *loginService,
	}
}

func (c *SiteSessionController) Init(router fiber.Router) error {
	group := router.Group("/session")

	group.Post("/login", func(ctx fiber.Ctx) error {
		var req services.LoginReqDTO
		if err := ctx.Bind().Body(&req); err != nil {
			return err
		}

		res, err := c.loginService.Login(req)
		if err != nil {
			return err
		}

		ctx.Response().Header.Add("HX-Location", "/site")
		return ctx.Status(http.StatusCreated).JSON(res)
	})

	return nil
}
