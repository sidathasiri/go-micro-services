package routes

import (
	"auth-service/cmd/api/handlers"

	"github.com/go-chi/chi/v5"
)

func RootRoute(router *chi.Mux) *chi.Mux {
	router.HandleFunc("POST /auth/login", handlers.AuthHandler)
	return router
}