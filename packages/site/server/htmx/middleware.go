package htmx

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// Prefix is the path under which every htmx endpoint lives.
const Prefix = "/htmx"

// RequireHtmx rejects requests to Prefix that were not sent by htmx.
// htmx endpoints return partial HTML that only makes sense inside a page.
// It also sets "Vary: HX-Request" so caches never mix htmx and regular responses.
func RequireHtmx(ctx fiber.Ctx) error {
	path := ctx.Path()
	if path != Prefix && !strings.HasPrefix(path, Prefix+"/") {
		return ctx.Next()
	}

	ctx.Vary("HX-Request")
	if ctx.Get("HX-Request") != "true" {
		return ctx.Status(http.StatusBadRequest).SendString("htmx request required")
	}

	return ctx.Next()
}
