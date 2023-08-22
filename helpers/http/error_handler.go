package http

import (
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandlerConfig struct {
	DfMsgKey string
	I18n     *i18np.I18n
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
		if cfg.DfMsgKey != "" {
			l, a := i18n.GetLanguagesInContext(*cfg.I18n, c)
			return c.Status(code).JSON(result.Error(cfg.I18n.Translate(cfg.DfMsgKey, l, a), code))
		}
		err = c.Status(code).JSON(result.Error(err.Error(), code))
		return err
	}
}
