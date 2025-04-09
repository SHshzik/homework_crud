package usecase

import (
	"context"

	"homework_crud/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// User -.
	User interface {
		ReadAll(ctx context.Context) ([]entity.User, error)
		Create()
		Read(ctx context.Context, id int) (*entity.User, error)
		Update()
		Delete(ctx context.Context, id int) error
	}
)
