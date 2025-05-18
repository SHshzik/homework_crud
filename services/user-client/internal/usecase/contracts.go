package usecase

import (
	"github.com/SHshzik/homework_crud/services/user-client/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test
type (
	// User -.
	User interface {
		Index() ([]*entity.User, error)
		Create(name, email string) (*entity.User, error)
		Read(id int) (*entity.User, error)
		Update(id int, name, email string) (*entity.User, error)
		Delete(id int) error
	}
)
