package usecase

import (
	"github.com/SHshzik/homework_crud/services/user-client/internal/adapters"
	"github.com/SHshzik/homework_crud/services/user-client/internal/entity"
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

func (u *UserCase) Create(name, email string) (*entity.User, error) {
	return u.Client.Create(name, email)
}

func (u *UserCase) Read(id int) (*entity.User, error) {
	return u.Client.Read(id)
}

func (u *UserCase) Update(id int, name, email string) (*entity.User, error) {
	user := &entity.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return u.Client.Update(user)
}

func (u *UserCase) Delete(id int) error {
	return u.Client.Delete(id)
}
