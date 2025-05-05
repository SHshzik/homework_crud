package grpc

import (
	"context"
	"strconv"

	"github.com/SHshzik/homework_crud/internal/usecase"
	usersService "github.com/SHshzik/homework_crud/pkg/api/proto"
)

type routeUserService struct {
	usersService.UnimplementedUsersServiceServer
	t usecase.User
}

func New(t usecase.User) *routeUserService {
	return &routeUserService{t: t}
}

func (s *routeUserService) GetUsers(ctx context.Context, req *usersService.GetUsersRequest) (*usersService.GetUsersResponse, error) {
	users, err := s.t.ReadAll(ctx)
	if err != nil {
		// r.l.Error(err, "http - v1 - index")

		return nil, err
	}

	usersProto := make([]*usersService.User, len(users))
	for i, user := range users {
		usersProto[i] = &usersService.User{
			Id:    strconv.Itoa(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}
	}

	return &usersService.GetUsersResponse{Users: usersProto}, nil
}
