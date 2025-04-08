package repo

import "homework_crud/internal/entity"

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type (
	UserRepo interface {
		FetchAll() []entity.User
	}
	//// TranslationRepo -.
	//TranslationRepo interface {
	//	Store(context.Context, entity.Translation) error
	//	GetHistory(context.Context) ([]entity.Translation, error)
	//}
	//
	//// TranslationWebAPI -.
	//TranslationWebAPI interface {
	//	Translate(entity.Translation) (entity.Translation, error)
	//}
)
