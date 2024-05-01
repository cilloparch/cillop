package http

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
)

// ParseBody parses the request body into the given struct and validates it.
func ParseBody(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_body", l, a), fiber.StatusBadRequest))
	}
	validateStruct(c.UserContext(), d, v, i, l, a)
}

// ParseQuery parses the request query into the given struct and validates it.
func ParseQuery(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.QueryParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_query", l, a), fiber.StatusBadRequest))
	}
	validateStruct(c.UserContext(), d, v, i, l, a)
}

// ParseParams parses the request params into the given struct and validates it.
func ParseParams(c *fiber.Ctx, v validation.Validator, i i18np.I18n, d interface{}) {
	l, a := i18n.GetLanguagesInContext(i, c)
	if err := c.ParamsParser(d); err != nil {
		panic(result.Error(i.Translate("error_invalid_request_params", l, a), fiber.StatusBadRequest))
	}
	validateStruct(c.UserContext(), d, v, i, l, a)
}

func validateStruct(ctx context.Context, d interface{}, v validation.Validator, i i18np.I18n, l, a string) {
	if errors := v.ValidateStruct(ctx, d, l, a); len(errors) > 0 {
		panic(result.ErrorDetail(i.Translate("error_validation_failed", l, a), errors, fiber.StatusBadRequest))
	}
}
