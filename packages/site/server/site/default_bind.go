package site

import (
	"github.com/gofiber/fiber/v3"
)

type DefaultBind[T any] struct {
	Lang string
	Data T
}

func GetTemplateBind[T any](ctx fiber.Ctx) DefaultBind[T] {
	language, ok := ctx.Context().UserValue("lang").(string)
	if !ok {
		language = "en"
	}
	return DefaultBind[T]{
		Lang: language,
	}
}
