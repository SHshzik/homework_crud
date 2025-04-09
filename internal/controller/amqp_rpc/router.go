package amqprpc

import (
	"homework_crud/internal/usecase"
	"homework_crud/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.User) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newUserRoutes(routes, t)
	}

	return routes
}
