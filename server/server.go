package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LogServerStarted(addr string) {
	fmt.Printf("Gyanasetu Server listening at %s\n", addr)
}
func CheckPrereqs() {
	secret := os.Getenv("APP_SECRET")
	if secret == " " {
		panic("Cannot start server without registred secret")
	}
}
func StartServer(addr string) {
	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(middleware.Heartbeat("/health"))

	apiRouter := RegisterApiRoutes()
	mainRouter.Mount("/api", apiRouter)

	server := &http.Server{
		Addr:    addr,
		Handler: mainRouter,
	}

	CheckPrereqs()
	LogServerStarted(addr)
	server.ListenAndServe()
}
