package grpc

import (
	"context"

	usersService "github.com/SHshzik/homework_crud/services/user-server/internal/controller/grpc/api/proto"
	"github.com/SHshzik/homework_crud/services/user-server/internal/entity"
	"github.com/SHshzik/homework_crud/services/user-server/internal/usecase"
)

type RouteUserService struct {
	usersService.UnimplementedUsersServiceServer
	t usecase.User
}

func New(t usecase.User) *RouteUserService {
	return &RouteUserService{t: t}
}

func (s *RouteUserService) GetUsers(ctx context.Context, _ *usersService.GetUsersRequest) (*usersService.GetUsersResponse, error) {
	users, err := s.t.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	usersProto := make([]*usersService.User, len(users))
	for i, user := range users {
		usersProto[i] = &usersService.User{
			Id:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}
	}

	return &usersService.GetUsersResponse{Users: usersProto}, nil
}

func (s *RouteUserService) GetUser(ctx context.Context, req *usersService.GetUserRequest) (*usersService.GetUserResponse, error) {
	user, err := s.t.Read(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &usersService.GetUserResponse{User: &usersService.User{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}}, nil
}

func (s *RouteUserService) DeleteUser(ctx context.Context, req *usersService.DeleteUserRequest) (*usersService.DeleteUserResponse, error) {
	err := s.t.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &usersService.DeleteUserResponse{}, nil
}

func (s *RouteUserService) CreateUser(ctx context.Context, req *usersService.CreateUserRequest) (*usersService.CreateUserResponse, error) {
	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	err := s.t.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &usersService.CreateUserResponse{User: &usersService.User{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}}, nil
}

func (s *RouteUserService) UpdateUser(ctx context.Context, req *usersService.UpdateUserRequest) (*usersService.UpdateUserResponse, error) {
	user := &entity.User{
		ID:    int(req.Id),
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	err := s.t.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &usersService.UpdateUserResponse{User: &usersService.User{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}}, nil
}
