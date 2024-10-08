package server

import (
	"fyndr.com/api/src/config"
	"fyndr.com/api/src/internal/middleware"
	"fyndr.com/api/src/pkg/error_handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func registerGlobalMiddlewares(app *fiber.App, apiErrorHandler *error_handler.ApiErrorHandler) {
	app.Use(middleware.NewApiGlobalErrorMiddleware(apiErrorHandler))
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/health"
		},
	}))
	//serve static files
	app.Static("/", "./assets/static")

	app.Use(helmet.New())
}

func NewApiHttpServer(env *config.AppEnv, db *gorm.DB, app *fiber.App, apiErrorHandler *error_handler.ApiErrorHandler) {
	registerGlobalMiddlewares(app, apiErrorHandler)

	healthGroup := app.Group("/health")
	RegisterHealthRoutes(healthGroup, db)

	authGroup := app.Group("/auth")
	RegisterAuthRoutes(authGroup, env, db)

	apiV1Group := app.Group("/api/v1")
	RegisterApiV1Routes(apiV1Group, env, db)

}
