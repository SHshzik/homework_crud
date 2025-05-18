package http

import (
	"net/http"

	_ "github.com/SHshzik/homework_crud/docs" // Swagger docs.
	"github.com/SHshzik/homework_crud/pkg/logger"
	"github.com/SHshzik/homework_crud/services/user-server/config"
	"github.com/SHshzik/homework_crud/services/user-server/internal/controller/http/middleware"
	v1 "github.com/SHshzik/homework_crud/services/user-server/internal/controller/http/v1"
	"github.com/SHshzik/homework_crud/services/user-server/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// NewRouter -.
// Swagger spec:
// @title       V1 API
// @description User CRUD
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *fiber.App, cfg *config.Config, l logger.Interface, t usecase.User) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewUserRoutes(apiV1Group, t, l)
	}
}
