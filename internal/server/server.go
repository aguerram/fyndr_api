package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/phuslu/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GracefulShutdownHandler func(ctx context.Context)

func StartHttpServer() (*fiber.App, func(ctx context.Context)) {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/v1/health"
		},
	}))
	//serve static files
	app.Static("/", "./assets/static")
	return app, func(ctx context.Context) {
		log.Info().Msg("Shutting down server")
		if err := app.ShutdownWithContext(ctx); err != nil {
			log.Fatal().Msgf("Error shutting down server %v", err)
		} else {
			log.Info().Msg("Server successfully shutdown")
		}
	}
}

func HandleGracefulShutdowns(handlers ...GracefulShutdownHandler) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-c
		log.Info().Msg("Received shutdown signal")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for _, handler := range handlers {
			if handler == nil {
				continue
			}
			handler(timeoutCtx) // You can pass shutdownCtx to handlers that require context for graceful shutdown
		}

		log.Info().Msg("Service successfully shutdown")
		os.Exit(0)
	}()
}
