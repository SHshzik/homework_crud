package amqprpc

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"homework_crud/internal/entity"
	"homework_crud/internal/usecase"
	"homework_crud/pkg/rabbitmq/rmq_rpc/server"
)

type userRoutes struct {
	userUseCase usecase.User
}

func newUserRoutes(routes map[string]server.CallHandler, t usecase.User) {
	r := &userRoutes{t}
	{
		routes["index"] = r.getHistory()
	}
}

type historyResponse struct {
	History []entity.User `json:"history"`
}

func (r *userRoutes) getHistory() server.CallHandler {
	return func(_ *amqp.Delivery) (any, error) {
		//translations, err := r.userUseCase.History(context.Background())
		//if err != nil {
		//	return nil, fmt.Errorf("amqp_rpc - translationRoutes - getHistory - r.translationUseCase.History: %w", err)
		//}
		//
		//response := historyResponse{translations}

		return nil, nil
	}
}
