package app

import (
	"fmt"
	"log"

	usersPb "github.com/SHshzik/homework_crud/api/proto"
	"github.com/SHshzik/homework_crud/services/user-client/adapters"
	"github.com/SHshzik/homework_crud/services/user-client/config"
	"github.com/SHshzik/homework_crud/services/user-client/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cfg *config.Config) {
	var client adapters.Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	switch cfg.ClientType {
	case "grpc":
		conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", cfg.GRPC.PORT), opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		userServiceClient := usersPb.NewUsersServiceClient(conn)

		client = adapters.NewGRPCClient(userServiceClient)
	case "http":
		client = adapters.NewHTTPClient()
	}

	userCase := usecase.New(client)

	switch cfg.RequestType {
	case "index":
		users, err := userCase.Index()
		if err != nil {
			log.Fatalf("fail to index: %v", err)
		}
		for _, user := range users {
			fmt.Printf("%#v\n", user.Name)
		}
	}
}
