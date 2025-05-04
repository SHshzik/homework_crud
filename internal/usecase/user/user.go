package user

import (
	"context"

	"github.com/SHshzik/homework_crud/internal/entity"
	"github.com/SHshzik/homework_crud/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo *repo.UserRepo
}

// New -.
func New(r *repo.UserRepo) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) ReadAll(ctx context.Context) ([]*entity.User, error) {
	users, err := u.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UseCase) Create(ctx context.Context, user *entity.User) error {
	err := u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) Read(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UseCase) Update(ctx context.Context, user *entity.User) error {
	err := u.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCase) Delete(ctx context.Context, id int) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
