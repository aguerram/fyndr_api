package middleware

import (
	"fyndr.com/api/src/pkg/error_handler"
	"github.com/gofiber/fiber/v2"
)

func NewApiGlobalErrorMiddleware(apiErrorHandler *error_handler.ApiErrorHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				apiErrorHandler.HandleApiPanic(c, r)
			}
		}()
		err := c.Next()
		return apiErrorHandler.ApiErrorHandler(c, err)
	}
}
