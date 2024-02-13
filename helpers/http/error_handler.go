package http

import (
	"fmt"

	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler func(c *fiber.Ctx, err error) error

func NewErrorHandler(log bool) ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusBadRequest
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Status).JSON(e)
		}
		if e, ok := err.(*result.DetailResult); ok {
			return c.Status(e.Status).JSON(e.Detail)
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if log {
			fmt.Printf("Error: %v\n", err.Error())
		}
		return c.Status(code).JSON(map[string]interface{}{})
	}
}
