package routes

import (
	"github.com/go-chi/chi/v5"

	"broker/cmd/api/handlers"
)

func RootRoute(router *chi.Mux) *chi.Mux {
	router.HandleFunc("GET /", handlers.RootHandler)
	router.HandleFunc("GET /health", handlers.GetHealthHandler)
	return router
}