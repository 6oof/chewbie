package app

import (
	"github.com/6oof/miniweb-base/app/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Get("/", handlers.HandleIndex)
}
