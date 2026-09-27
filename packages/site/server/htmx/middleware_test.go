package htmx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func newTestApp() *fiber.App {
	app := fiber.New()
	app.Use(RequireHtmx)
	app.Get("/htmx/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	app.Get("/site/page", func(ctx fiber.Ctx) error {
		return ctx.SendString("page")
	})
	return app
}

func TestRequireHtmx(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		hxRequest  string
		wantStatus int
		wantBody   string
		wantVary   string
	}{
		{"htmx request passes", "/htmx/ping", "true", http.StatusOK, "pong", "HX-Request"},
		{"missing header rejected", "/htmx/ping", "", http.StatusBadRequest, "htmx request required", "HX-Request"},
		{"wrong header value rejected", "/htmx/ping", "1", http.StatusBadRequest, "htmx request required", "HX-Request"},
		{"other paths untouched", "/site/page", "", http.StatusOK, "page", ""},
		{"similar prefix untouched", "/htmxfoo", "", http.StatusNotFound, "", ""},
	}

	app := newTestApp()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			if tt.hxRequest != "" {
				req.Header.Set("HX-Request", tt.hxRequest)
			}

			res, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			body, _ := io.ReadAll(res.Body)

			if res.StatusCode != tt.wantStatus {
				t.Errorf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}
			if tt.wantBody != "" && string(body) != tt.wantBody {
				t.Errorf("body = %q, want %q", body, tt.wantBody)
			}
			if got := res.Header.Get("Vary"); got != tt.wantVary {
				t.Errorf("Vary = %q, want %q", got, tt.wantVary)
			}
		})
	}
}
