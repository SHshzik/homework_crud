package persistent

import (
	"context"
	"fmt"
	"homework_crud/internal/entity"
	"homework_crud/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r UserRepo) FetchAll(ctx context.Context) ([]entity.User, error) {
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

	entities := make([]entity.User, 0, _defaultEntityCap)
	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.Id, &e.Name, &e.Email, &e.Phone)
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
	err = r.Pool.QueryRow(ctx, sql, id).Scan(&user.Id, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - Find - r.Pool.Query: %w", err)
	}
	return user, nil
}
