package handler

import (
	"fyndr.com/api/src/config"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"gorm.io/gorm"
)

type HomeHandler struct {
	env *config.AppEnv
	db  *gorm.DB
}

func NewHomeHandler(env *config.AppEnv, db *gorm.DB) Handler {
	return &HomeHandler{
		env: env,
		db:  db,
	}
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {
	log.Info().Msg("Incoming requests to HomeHandler.Home")
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func (h *HomeHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.Home)
}
