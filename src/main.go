package main

import (
	"context"
	"fmt"
	"fyndr.com/api/src/config"
	"fyndr.com/api/src/internal/server"
	"fyndr.com/api/src/pkg/error_handler"
	"github.com/phuslu/log"
	"time"
)

var gracefulShutdowns []server.GracefulShutdownHandler

func main() {
	config.InitializeLogger()
	env := config.InitializeEnv()
	//initialize oauth2

	//to handle panics in the application
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Application panicked: %v", r)
			handlePanic()
		}
	}()

	databaseConnection, closeDatabase := config.NewDatabaseConnection(env)
	gracefulShutdowns = append(gracefulShutdowns, closeDatabase)

	//global api error handler
	apiErrorHandler := error_handler.NewApiErrorHandler()
	//-----------------start http server-----------------
	httpServer, shutdownHttpServer := server.StartHttpServer(env, apiErrorHandler)
	gracefulShutdowns = append(gracefulShutdowns, shutdownHttpServer)

	//initialize api http server
	server.NewApiHttpServer(env, databaseConnection, httpServer, apiErrorHandler)

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
	ctx, closeCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer closeCtx()
	for _, handler := range gracefulShutdowns {
		if handler == nil {
			continue
		}
		handler(ctx)
	}
}
