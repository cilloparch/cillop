package http

import (
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_body", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseQuery(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.QueryParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_query", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func ParseParams(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.ParamsParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_params", l, a), fiber.StatusBadRequest))
	}
	validateStruct(d, v, i, l, a)
}

func validateStruct(d interface{}, v validation.Validator, i i18np.I18n, l, a string) {
	if errors := v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorDetail(i.Translate("error_validation_failed", l, a), errors, fiber.StatusBadRequest))
	}
}
