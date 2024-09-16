package handler

import (
	"fyndr.com/api/src/internal/service"
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	healthService service.HealthService
}

func NewHealthHandler(healthService service.HealthService) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

func (h *HealthHandler) GetStatus(ctx *fiber.Ctx) error {
	status := h.healthService.GetStatus()
	if status.IsSuccess() {
		return ctx.Status(fiber.StatusOK).JSON(status)
	}
	return ctx.Status(fiber.StatusServiceUnavailable).JSON(status)
}

func (h *HealthHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetStatus)
}
