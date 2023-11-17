package http

import (
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandlerConfig struct {
	I18n *i18np.I18n
}

func NewErrorHandler(cfg ErrorHandlerConfig) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Status).JSON(e)
		}
		if e, ok := err.(*result.DetailResult); ok {
			return c.Status(e.Status).JSON(e.Detail)
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		return c.Status(code).JSON(map[string]interface{}{})
	}
}
