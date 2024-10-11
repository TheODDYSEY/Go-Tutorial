package handlers

import (
	"go-tutorial/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r * chi.Mux) {
	// Global middleware

	r.Use(chimiddle.StripSlashes)
	r.route("/account", func(router chi.Router){

		// Middleware for account route 
		router.Use(middleware.Authorization)

		router.Get("/coins",GetCoinBalance)

	})
}