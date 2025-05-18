// Package app configures and runs application.
package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	usersService "github.com/SHshzik/homework_crud/api/proto"
	"github.com/SHshzik/homework_crud/pkg/httpserver"
	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/pkg/postgres"
	"github.com/SHshzik/homework_crud/services/user-server/config"
	usersServer "github.com/SHshzik/homework_crud/services/user-server/internal/controller/grpc"
	v1 "github.com/SHshzik/homework_crud/services/user-server/internal/controller/http"
	"github.com/SHshzik/homework_crud/services/user-server/internal/repo"
	"github.com/SHshzik/homework_crud/services/user-server/internal/usecase/user"
	"google.golang.org/grpc"
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
	userUseCase := user.New(repo.New(pg))

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	v1.NewRouter(httpServer.App, cfg, l, userUseCase)

	// Start servers
	httpServer.Start()

	// GRPC Server
	usersServer := usersServer.New(userUseCase)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPC.PORT))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - net.Listen: %w", err))
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	usersService.RegisterUsersServiceServer(grpcServer, usersServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - grpcServer.Serve: %w", err))
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
