package app

import (
	"github.com/6oof/miniweb-base/app/handlers"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up the routing for the MiniWeb application by defining the routes and associating them with their respective handlers.
func registerRoutes(r *chi.Mux) {
	// Define routes and associate them with handlers
	r.Get("/", handlers.HandleIndex)

	r.Get("/*", handlers.HandleNotFound)
}
