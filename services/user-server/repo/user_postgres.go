package repo

import (
	"context"
	"fmt"

	"github.com/SHshzik/homework_crud/pkg/postgres"
	"github.com/SHshzik/homework_crud/services/user-server/entity"
)

const defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r UserRepo) FetchAll(ctx context.Context) ([]*entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, name, email, phone").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - FetchAll - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - FetchAll - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]*entity.User, 0, defaultEntityCap)

	for rows.Next() {
		e := &entity.User{}

		err = rows.Scan(&e.ID, &e.Name, &e.Email, &e.Phone)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - FetchAll - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (r UserRepo) Find(ctx context.Context, id int) (*entity.User, error) {
	sql, _, err := r.Builder.
		Select("id, name, email, phone").
		From("users").
		Where("id = $1").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - Find - r.Builder: %w", err)
	}

	user := &entity.User{}

	err = r.Pool.QueryRow(ctx, sql, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - Find - r.Pool.Query: %w", err)
	}

	return user, nil
}

func (r UserRepo) Delete(ctx context.Context, id int) error {
	var deletedID string

	sql, _, err := r.Builder.
		Delete("users").
		Where("id = $1").
		Suffix("returning id").
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Delete - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, id).Scan(&deletedID)
	if err != nil {
		return fmt.Errorf("UserRepo - Delete - r.Pool.Exec %w", err)
	}

	return nil
}

func (r UserRepo) Create(ctx context.Context, user *entity.User) error {
	sql, _, err := r.Builder.
		Insert("users").
		Columns("name", "email", "phone").
		Values("$1", "$2", "$3").
		Suffix("returning id").
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, user.Name, user.Email, user.Phone).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Pool.Exec %w", err)
	}

	return nil
}

func (r UserRepo) Update(ctx context.Context, user *entity.User) error {
	sql, _, err := r.Builder.
		Update("users").
		Set("name", "$1").
		Set("email", "$2").
		Set("phone", "$3").
		Where("id = $4").
		Suffix("returning id").
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Update - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, user.Name, user.Email, user.Phone, user.ID).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("UserRepo - Update - r.Pool.Exec %w", err)
	}

	return nil
}
