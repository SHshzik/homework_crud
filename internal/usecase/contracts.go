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
		Read()
		Update()
		Delete()

		//Translate(context.Context, entity.Translation) (entity.Translation, error)
		//History(context.Context) ([]entity.Translation, error)
	}
)
