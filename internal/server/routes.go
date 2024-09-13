package server

import (
	"fyndr.com/api/config"
	"fyndr.com/api/internal/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApiV1Routes(router fiber.Router, env *config.AppEnv, db *gorm.DB) {
	//home handlers
	handler.NewHomeHandler(env, db).RegisterRoutes(router)
}
