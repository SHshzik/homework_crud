// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	//amqprpc "github.com/evrone/go-clean-template/internal/controller/amqp_rpc"
	//"github.com/evrone/go-clean-template/pkg/httpserver"
	//"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"

	"homework_crud/config"
	v1 "homework_crud/internal/controller/http"
	"homework_crud/internal/repo/persistent"
	"homework_crud/internal/usecase/user"
	"homework_crud/pkg/httpserver"
	"homework_crud/pkg/logger"
	"homework_crud/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	userUseCase := user.New(persistent.New(pg))

	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)

	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	v1.NewRouter(httpServer.App, cfg, l, userUseCase)

	//// Start servers
	//rmqServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	//
	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
