package main

import (
	"auth-service/cmd/api/routes"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CreateApp() *chi.Mux {
	baseRouter := chi.NewRouter()
	baseRouter.Use(middleware.Logger)
	baseRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
	}))

	rootRouter := routes.RootRoute(baseRouter)

	log.Println("Auth service started on port", port)

	return rootRouter
}