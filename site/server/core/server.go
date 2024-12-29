package core

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/4strodev/4stroblog/site/application/blog"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

type Server struct {
	fiber       *fiber.App
	controllers []Controller
	middlewares []fiber.Handler
}

func (s *Server) AddController(controller Controller) {
	s.controllers = append(s.controllers, controller)
}

func (s *Server) AddMiddleware(handler fiber.Handler) {
	s.middlewares = append(s.middlewares, handler)
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf(":%d", port)
	viewsEngine := html.New("./views", ".html")
	viewsEngine.AddFunc("renderPost", func(post string) string {
		content, err := blog.RenderPost(post)
		if err != nil {
			return ""
		}

		return string(content)
	})
	viewsEngine.AddFunc("unescape", func(s string) template.HTML {
		return template.HTML(s)
	})

	s.fiber = fiber.New(fiber.Config{
		Views: viewsEngine,
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			log.Println(err)
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		},
	})
	s.fiber.Use(recover.New())
	s.fiber.Get("/assets/*", static.New("./assets"))
	s.fiber.Get("/uploads/*", static.New("./uploads"))

	for _, middleware := range s.middlewares {
		s.fiber.Use(middleware)
	}

	for _, controller := range s.controllers {
		err := controller.Init(s.fiber)
		if err != nil {
			return err
		}
	}
	return s.fiber.Listen(addr)
}
