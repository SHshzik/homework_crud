package usecase

import (
	"context"

	"github.com/SHshzik/homework_crud/services/user-server/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test
type (
	// User -.
	User interface {
		ReadAll(ctx context.Context) ([]*entity.User, error)
		Create(ctx context.Context, user *entity.User) error
		Read(ctx context.Context, id int) (*entity.User, error)
		Update(ctx context.Context, user *entity.User) error
		Delete(ctx context.Context, id int) error
	}
)
