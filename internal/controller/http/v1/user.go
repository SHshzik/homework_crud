package v1

import (
	"homework_crud/internal/entity"
	"homework_crud/internal/usecase"
	"homework_crud/pkg/logger"
	"net/http"
	"strconv"

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

	userGroup := apiV1Group.Group("/users")
	{
		userGroup.Get("/", r.index)
		userGroup.Put("/:user_id", r.update)
		userGroup.Post("/", r.create)
		userGroup.Get("/:user_id", r.show)
		userGroup.Delete("/:user_id", r.delete)
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

type userResponse struct {
	*entity.User
}

// @Summary     Show user by id
// @Description Show user detail
// @ID          show
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     200 {object} userResponse
// @Failure     500 {object} response
// @Router      /users/:id [get]
func (r *userRoutes) show(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		r.l.Error(err, "http - v1 - show")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "wrong user id")
	}

	user, err := r.t.Read(ctx.UserContext(), userId)
	if err != nil {
		r.l.Error(err, "http - v1 - index")

		return errorResponse(ctx, http.StatusNotFound, "user not found")
	}

	return ctx.Status(http.StatusOK).JSON(userResponse{user})
}

// @Summary     Delete user by id
// @Description Delete user from db
// @ID          delete
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     204
// @Failure     500 {object} response
// @Router      /users/:id [delete]
func (r *userRoutes) delete(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		r.l.Error(err, "http - v1 - delete")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "wrong user id")
	}

	err = r.t.Delete(ctx.UserContext(), userId)
	if err != nil {
		r.l.Error(err, "http - v1 - delete")

		return errorResponse(ctx, http.StatusNotFound, "user not found")
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

type userForm struct {
	Name  string `form:"name"  validate:"required"`
	Email string `form:"email" validate:"required"`
	Phone string `form:"phone" validate:"required"`
}

// @Summary     Create new user
// @Description Create new user
// @ID          create
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     201 {object} userResponse
// @Failure     500 {object} response
// @Router      /users [post]
func (r *userRoutes) create(ctx *fiber.Ctx) error {
	formUser := userForm{}
	err := ctx.BodyParser(&formUser)
	if err != nil {
		r.l.Error(err, "http - v1 - create")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "Bad user params")
	}

	err = r.v.Struct(formUser)
	if err != nil {
		r.l.Error(err, "http - v1 - create")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	user := entity.NewUser(formUser.Name, formUser.Email, formUser.Phone)
	err = r.t.Create(ctx.UserContext(), user)
	if err != nil {
		r.l.Error(err, "http - v1 - create")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "user not created")
	}

	return ctx.Status(http.StatusCreated).JSON(userResponse{user})
}

// @Summary     Update user
// @Description Update exists user
// @ID          update
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     200 {object} userResponse
// @Failure     500 {object} response
// @Router      /users [put]
func (r *userRoutes) update(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		r.l.Error(err, "http - v1 - delete")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "wrong user id")
	}

	formUser := userForm{}
	err = ctx.BodyParser(&formUser)
	if err != nil {
		r.l.Error(err, "http - v1 - update")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "Bad user params")
	}

	err = r.v.Struct(formUser)
	if err != nil {
		r.l.Error(err, "http - v1 - update")

		return errorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	user := entity.NewUser(formUser.Name, formUser.Email, formUser.Phone)
	user.Id = userId

	err = r.t.Update(ctx.UserContext(), user)
	if err != nil {
		r.l.Error(err, "http - v1 - update")

		return errorResponse(ctx, http.StatusUnprocessableEntity, "user not updated")
	}

	return ctx.Status(http.StatusCreated).JSON(userResponse{user})
}
