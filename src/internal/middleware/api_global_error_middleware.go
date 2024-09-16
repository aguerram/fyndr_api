package middleware

import (
	"errors"
	"fyndr.com/api/src/internal/api/response/api_error"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

func NewApiGlobalErrorMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Msgf("Recovered from panic: %v", r)
				c.Status(fiber.StatusInternalServerError)
				c.JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
		}()
		err := c.Next()
		var fiberError *fiber.Error
		if err != nil && errors.As(err, &fiberError) {
			return handleApiStatusErrors(c, fiberError)
		}
		return nil
	}
}

func handleApiStatusErrors(c *fiber.Ctx, err *fiber.Error) error {
	c.Set("Content-Type", "application/json")
	switch err.Code {
	case fiber.StatusNotFound:
		handleNotFound(c, err)
	case fiber.StatusBadRequest:
		handleBadRequest(c, err)
	}
	return handleInternalServer(c)
}

func handleNotFound(c *fiber.Ctx, err *fiber.Error) error {
	return c.Status(fiber.StatusNotFound).JSON(api_error.PageNotFound())
}

func handleBadRequest(c *fiber.Ctx, err *fiber.Error) error {
	return c.Status(fiber.StatusBadRequest).JSON(api_error.BadRequest())
}

func handleInternalServer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(api_error.InternalServerError())
}
