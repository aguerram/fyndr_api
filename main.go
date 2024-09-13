package main

import (
	"context"
	"fmt"
	"fyndr.com/api/config"
	"fyndr.com/api/internal/server"
	"github.com/phuslu/log"
	"time"
)

var gracefulShutdowns []server.GracefulShutdownHandler

func main() {
	config.InitializeLogger()
	env := config.InitializeEnv()

	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Application panicked: %v", r)
			handlePanic()
		}
	}()

	databaseConnection, closeDatabase := config.NewDatabaseConnection(env)
	gracefulShutdowns = append(gracefulShutdowns, closeDatabase)

	httpServer, shutdownHttpServer := server.StartHttpServer()
	gracefulShutdowns = append(gracefulShutdowns, shutdownHttpServer)

	server.NewApiHttpServer(env, databaseConnection, httpServer)

	//handle graceful shutdowns
	server.HandleGracefulShutdowns(gracefulShutdowns...)

	log.Info().Msgf("Server started on port %s", env.HttpPort)

	//start http server
	serverErr := httpServer.Listen(fmt.Sprintf(":%s", env.HttpPort))
	if serverErr != nil {
		log.Fatal().Msgf("Error starting server %v", serverErr)
	}
}

func handlePanic() {
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	for _, handler := range gracefulShutdowns {
		if handler == nil {
			continue
		}
		handler(ctx)
	}
}
