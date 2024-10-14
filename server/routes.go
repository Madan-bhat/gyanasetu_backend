package server

import (
	"context"
	"gyanasetu/backend/handlers"
	"gyanasetu/backend/middlewares"
	"gyanasetu/backend/services"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/joho/godotenv"
)

var (
	appHandlers    handlers.Handlers       = GetInitializedHandlers()
	appMiddlewares middlewares.Middlewares = GetInitializedMiddlewares()
)

func RegisterApiRoutes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Mount("/auth", RegisterAuthRoutes())
	mux.Mount("/org", RegisterOrganizationRoutes())
	return mux
}

func RegisterAuthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", appHandlers.Register)
	r.Group(func(r chi.Router) {
		r.Use(appMiddlewares.RestictedAccess(false, false)) // dont allow bdfl
		r.Put("/role", appHandlers.SelectRole)
	})
	return r
}
func RegisterOrganizationRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/list", appHandlers.ListOrganizations)

	// non bdfl group
	r.Group(func(r chi.Router) {
		r.Use(appMiddlewares.RestictedAccess(false, false))
		r.Post("/join", appHandlers.JoinOrganization)
	})

	// only bdfl group
	r.Group(func(r chi.Router) {
		r.Use(appMiddlewares.RestictedAccess(true, true))
		r.Post("/create", appHandlers.CreateOrganization)
	})

	r.Get("/list", appHandlers.ListOrganizations)
	return r
}

func RegisterServiceInfo() services.Services {
	godotenv.Load()
	ctx := context.Background()
	db, sqlDB, bdflID := RegisterDB(ctx)
	return services.Services{
		Db:     db,
		DbSQL:  sqlDB,
		Ctx:    ctx,
		Secret: []byte(os.Getenv("SIG_SECRET")),
		BDFLId: bdflID,
	}
}
func GetInitializedHandlers() handlers.Handlers {
	godotenv.Load()
	serviceInfo := RegisterServiceInfo()
	handlers := handlers.Handlers{
		Services:  serviceInfo,
		Validator: validator.New(),
	}
	return handlers
}

func GetInitializedMiddlewares() middlewares.Middlewares {
	return middlewares.Middlewares{
		Services: RegisterServiceInfo(),
	}
}
