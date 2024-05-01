package http

import (
	"github.com/cilloparch/cillop/v2/i18np"
	"github.com/cilloparch/cillop/v2/log"
	"github.com/cilloparch/cillop/v2/middlewares/i18n"
	"github.com/cilloparch/cillop/v2/rescode"
	"github.com/cilloparch/cillop/v2/result"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler func(c *fiber.Ctx, err error) error

func NewErrorHandler(logger log.Service, translator *i18np.I18n) ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusBadRequest
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Status).JSON(e)
		}
		if e, ok := err.(*result.DetailResult); ok {
			return c.Status(e.Status).JSON(e.Detail)
		}
		if e, ok := err.(*i18np.Error); ok {
			msg := translator.TranslateFromError(*e, i18n.ParseLocale(c))
			return c.Status(code).JSON(e.JSON(msg))
		}
		if e, ok := err.(*rescode.RC); ok {
			msg := e.Message
			if e.Translateable {
				msg = translator.Translate(e.Message, i18n.ParseLocale(c))
			}
			return c.Status(e.HttpStatus).JSON(e.JSON(msg))
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		logger.Error(err)
		return c.Status(code).JSON(map[string]interface{}{})
	}
}
