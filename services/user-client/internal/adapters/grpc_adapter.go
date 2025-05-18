package adapters

import (
	"context"

	usersPb "github.com/SHshzik/homework_crud/api/proto"
	"github.com/SHshzik/homework_crud/services/user-client/internal/entity"
)

type GrpcClient struct {
	UserServiceClient usersPb.UsersServiceClient
}

func NewGRPCClient(userServiceClient usersPb.UsersServiceClient) *GrpcClient {
	return &GrpcClient{UserServiceClient: userServiceClient}
}

func (c *GrpcClient) Index() ([]*entity.User, error) {
	getUsersResponse, err := c.UserServiceClient.GetUsers(context.Background(), &usersPb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(getUsersResponse.Users))

	for i, user := range getUsersResponse.Users {
		users[i] = &entity.User{
			ID:    int(user.Id),
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return users, nil
}

func (c *GrpcClient) Create(name, email string) (*entity.User, error) {
	createUserResponse, err := c.UserServiceClient.CreateUser(context.Background(), &usersPb.CreateUserRequest{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:    int(createUserResponse.User.Id),
		Name:  createUserResponse.User.Name,
		Email: createUserResponse.User.Email,
	}, nil
}

func (c *GrpcClient) Read(id int) (*entity.User, error) {
	getUserResponse, err := c.UserServiceClient.GetUser(context.Background(), &usersPb.GetUserRequest{Id: int64(id)})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:    int(getUserResponse.User.Id),
		Name:  getUserResponse.User.Name,
		Email: getUserResponse.User.Email,
	}, nil
}

func (c *GrpcClient) Update(user *entity.User) (*entity.User, error) {
	updateUserResponse, err := c.UserServiceClient.UpdateUser(context.Background(), &usersPb.UpdateUserRequest{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:    int(updateUserResponse.User.Id),
		Name:  updateUserResponse.User.Name,
		Email: updateUserResponse.User.Email,
	}, nil
}

func (c *GrpcClient) Delete(id int) error {
	_, err := c.UserServiceClient.DeleteUser(context.Background(), &usersPb.DeleteUserRequest{Id: int64(id)})
	if err != nil {
		return err
	}

	return nil
}
