package user

import (
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
