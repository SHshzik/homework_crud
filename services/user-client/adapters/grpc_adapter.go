package adapters

import (
	"context"

	usersPb "github.com/SHshzik/homework_crud/api/proto"
	"github.com/SHshzik/homework_crud/services/user-server/entity"
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
