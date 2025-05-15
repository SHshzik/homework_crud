package usecase

import (
	"github.com/SHshzik/homework_crud/services/user-client/adapters"
	"github.com/SHshzik/homework_crud/services/user-server/entity"
)

type UserCase struct {
	Client adapters.Client
}

func New(client adapters.Client) *UserCase {
	return &UserCase{Client: client}
}

func (u *UserCase) Index() ([]*entity.User, error) {
	return u.Client.Index()
}
