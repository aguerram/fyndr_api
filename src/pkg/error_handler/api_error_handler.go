package error_handler

import (
	"errors"
	"fyndr.com/api/src/internal/api/response/api_error"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

type ApiErrorHandler struct {
}

func NewApiErrorHandler() *ApiErrorHandler {
	return &ApiErrorHandler{}
}

func (a ApiErrorHandler) ApiErrorHandler(c *fiber.Ctx, err error) error {
	var fiberError *fiber.Error
	if err != nil && errors.As(err, &fiberError) {
		log.Info().Msgf("Handling error: %v", fiberError)
		return handleApiStatusErrors(c, fiberError)
	}
	return nil
}

func (a ApiErrorHandler) HandleApiPanic(c *fiber.Ctx, r any) {
	log.Error().Msgf("Recovered from panic: %v", r)
	c.Status(fiber.StatusInternalServerError).JSON(api_error.InternalServerError())
}

func handleApiStatusErrors(c *fiber.Ctx, err *fiber.Error) error {
	c.Set("Content-Type", "application/json")
	switch err.Code {
	case fiber.StatusNotFound:
		return handleNotFound(c, err)
	case fiber.StatusBadRequest:

		return handleBadRequest(c, err)
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
