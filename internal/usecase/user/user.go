package user

import (
	"context"
	"homework_crud/internal/entity"
	"homework_crud/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.UserRepo
}

// New -.
func New(r repo.UserRepo) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) ReadAll(ctx context.Context) ([]entity.User, error) {
	users, err := u.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UseCase) Create() {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) Read(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UseCase) Update() {
	//TODO implement me
	panic("implement me")
}

func (u *UseCase) Delete() {
	//TODO implement me
	panic("implement me")
}

// History - getting translate history from store.
//func (uc *UseCase) History(ctx context.Context) ([]entity.Translation, error) {
//	translations, err := uc.repo.GetHistory(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
//	}
//
//	return translations, nil
//}
//
//// Translate -.
//func (uc *UseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
//	translation, err := uc.webAPI.Translate(t)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
//	}
//
//	err = uc.repo.Store(ctx, translation)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
//	}
//
//	return translation, nil
//}
