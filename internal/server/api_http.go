package server

import (
	"fyndr.com/api/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewApiHttpServer(env *config.AppEnv, db *gorm.DB, app *fiber.App) {
	apiV1Group := app.Group("/api/v1")
	RegisterApiV1Routes(apiV1Group, env, db)
}
