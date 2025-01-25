package core

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/4strodev/4stroblog/site/features/blog"
	wiring "github.com/4strodev/wiring/pkg"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

type Server struct {
	Wiring      wiring.Container
	fiber       *fiber.App
	middlewares []fiber.Handler
	modules     []Module
}

func (s *Server) AddModule(module Module) {
	s.modules = append(s.modules, module)
}

func (s *Server) AddMiddleware(handler fiber.Handler) {
	s.middlewares = append(s.middlewares, handler)
}

func (s *Server) Start(port int) error {
	if s.Wiring == nil {
		s.Wiring = wiring.New()
	}
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
	s.fiber.Use(func(ctx fiber.Ctx) error {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			debug.PrintStack()
			s := fmt.Sprint(r)
			log.Println(s)

			err := ctx.Status(http.StatusInternalServerError).SendString(s)
			if err != nil {
				log.Println(err)
			}
		}()
		return ctx.Next()
	})
	s.fiber.Get("/assets/*", static.New("./assets"))
	s.fiber.Get("/uploads/*", static.New("./uploads"))

	err := s.Wiring.Singleton(func() fiber.Router {
		return s.fiber
	})
	if err != nil {
		return err
	}

	for _, middleware := range s.middlewares {
		s.fiber.Use(middleware)
	}

	for _, module := range s.modules {
		err := module.init(s.Wiring)
		if err != nil {
			return err
		}
	}
	return s.fiber.Listen(addr)
}
