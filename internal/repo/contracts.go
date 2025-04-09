package repo

import (
	"context"

	"homework_crud/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type (
	UserRepo interface {
		FetchAll(ctx context.Context) ([]entity.User, error)
		Find(ctx context.Context, id int) (*entity.User, error)
		Delete(ctx context.Context, id int) error
		Create(ctx context.Context, user *entity.User) error
	}
)
