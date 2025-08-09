package core

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/4strodev/4stroblog/site/features/blog"
	"github.com/4strodev/4stroblog/site/shared/i18n"
	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/text/language"
)

type Server struct {
	Wiring      *container.Container
	fiber       *fiber.App
	middlewares []fiber.Handler
	modules     []*Module
	logger      *slog.Logger
	viewsEngine *html.Engine
}

func (s *Server) AddModule(module Module) {
	s.modules = append(s.modules, &module)
}

func (s *Server) AddMiddleware(handler fiber.Handler) {
	s.middlewares = append(s.middlewares, handler)
}

// Init initialize server dependencies and modules to be ready to start listening requests
func (s *Server) Init() error {
	if s.Wiring == nil {
		s.Wiring = container.New()
	}
	s.viewsEngine = html.New("./views", ".html")

	// setup dependencies
	err := s.Wiring.Singleton(func() fiber.Router {
		return s.fiber
	})
	if err != nil {
		return err
	}
	for _, module := range s.modules {
		err := module.initDependencies(s.Wiring)
		if err != nil {
			return err
		}
	}

	// setup logger
	s.logger, err = container.Resolve[*slog.Logger](s.Wiring)
	if err != nil {
		return err
	}

	// Views need dependencies to be initialized
	err = s.setupViews()
	if err != nil {
		return err
	}

	s.fiber = fiber.New(fiber.Config{
		Views: s.viewsEngine,
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			s.logger.Error(err.Error())
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		},
	})

	// recover middleware
	s.fiber.Use(func(ctx fiber.Ctx) error {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			debug.PrintStack()
			strContent := fmt.Sprint(r)

			err := ctx.Status(http.StatusInternalServerError).SendString(strContent)
			if err != nil {
				s.logger.Error("error caught", "stack", strContent)
			}
		}()
		return ctx.Next()
	})

	langMatcher := language.NewMatcher([]language.Tag{
		language.English,
		language.Spanish,
	})
	s.fiber.Use(func(ctx fiber.Ctx) error {
		lang := ctx.Cookies("lang")
		acceptLanguage := ctx.Get("Accept-Language")
		langTag, _ := language.MatchStrings(langMatcher, lang, acceptLanguage)
		language, _ := langTag.Base()
		ctx.Context().SetUserValue("lang", language)

		return ctx.Next()
	})

	s.fiber.Get("/assets/*", static.New("./assets"))
	s.fiber.Get("/uploads/*", static.New("./uploads"))

	for _, middleware := range s.middlewares {
		s.fiber.Use(middleware)
	}

	return nil
}

func (s *Server) setupViews() error {
	s.viewsEngine.AddFunc("renderPost", func(post string) string {
		content, err := blog.RenderPost(post)
		if err != nil {
			return ""
		}

		return string(content)
	})
	s.viewsEngine.AddFunc("unescape", func(s string) template.HTML {
		return template.HTML(s)
	})

	translationService, err := container.Resolve[*i18n.TranslationService](s.Wiring)
	if err != nil {
		return err
	}
	s.viewsEngine.AddFunc("translate", func(lang, key string, fallback ...string) string {
		log.Println("translating", lang, key)
		if len(fallback) > 0 {
			return translationService.TranslateOr(lang, key, fallback[0])
		}
		return translationService.Translate(lang, key)
	})
	s.viewsEngine.AddFunc("namedUrl", func(name string) string {
		return ""
	})

	return nil
}

func (s *Server) Start(port int) error {
	if s.fiber == nil {
		return errors.New("server not initialized")
	}

	for _, module := range s.modules {
		err := module.initControllers()
		if err != nil {
			return err
		}
	}

	s.fiber.Hooks().OnListen(func(data fiber.ListenData) error {
		s.logger.Info("server started", "host", data.Host, "port", data.Port)
		return nil
	})

	addr := fmt.Sprintf(":%d", port)
	return s.fiber.Listen(addr, fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
