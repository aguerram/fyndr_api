package server

import (
	"fyndr.com/api/src/config"
	"fyndr.com/api/src/internal/handler"
	"fyndr.com/api/src/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApiV1Routes(router fiber.Router, env *config.AppEnv, db *gorm.DB) {
	//home handlers
	handler.NewHomeHandler(env, db).RegisterRoutes(router)
}

func RegisterAuthRoutes(router fiber.Router, env *config.AppEnv, db *gorm.DB) {
	//auth handlers
	authService := service.NewAuthService(env, db)
	handler.NewAuthHandler(authService).RegisterRoutes(router)
}

func RegisterHealthRoutes(router fiber.Router, db *gorm.DB) {
	//health handlers
	healthService := service.NewHealthService(db)
	handler.NewHealthHandler(healthService).RegisterRoutes(router)
}
