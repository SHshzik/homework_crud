package adapters

import "github.com/SHshzik/homework_crud/services/user-client/entity"

type Client interface {
	Index() ([]*entity.User, error)
	Create(name, email string) (*entity.User, error)
	Read(id int) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id int) error
}
