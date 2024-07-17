package handlers

import (
	"github.com/Vinay-786/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capitalization of the func name makes it public so it can be imported
func Handler(r *chi.Mux) {
	//Global middleware

	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
