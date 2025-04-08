package v1

import (
	"homework_crud/internal/entity"
	"homework_crud/internal/usecase"
	"homework_crud/pkg/logger"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
	v *validator.Validate
}

func NewUserRoutes(apiV1Group fiber.Router, t usecase.User, l logger.Interface) {
	r := &userRoutes{t, l, validator.New(validator.WithRequiredStructEnabled())}

	translationGroup := apiV1Group.Group("/users")
	{
		translationGroup.Get("/", r.index)
		//translationGroup.Post("/do-translate", r.doTranslate)
	}
}

type indexResponse struct {
	Users []entity.User `json:"users"`
}

// @Summary     Show users
// @Description Show all users
// @ID          index
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     200 {object} indexResponse
// @Failure     500 {object} response
// @Router      /users [get]
func (r *userRoutes) index(ctx *fiber.Ctx) error {
	users, err := r.t.ReadAll(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "http - v1 - index")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(indexResponse{users})
}
