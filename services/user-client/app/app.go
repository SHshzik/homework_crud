package app

import (
	"fmt"

	usersPb "github.com/SHshzik/homework_crud/api/proto"
	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/services/user-client/adapters"
	"github.com/SHshzik/homework_crud/services/user-client/config"
	"github.com/SHshzik/homework_crud/services/user-client/controller"
	"github.com/SHshzik/homework_crud/services/user-client/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	var client adapters.Client

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	switch cfg.ClientType {
	case "grpc":
		conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", cfg.GRPC.PORT), opts...)
		if err != nil {
			l.Fatal("fail to dial: %v", err)
		}
		defer conn.Close()

		userServiceClient := usersPb.NewUsersServiceClient(conn)

		client = adapters.NewGRPCClient(userServiceClient)
	case "http":
		client = adapters.NewHTTPClient(&cfg.HTTP, l)
	}

	userCase := usecase.New(client)
	route := controller.NewRoute(l, userCase, cfg)
	route.Run()
}
