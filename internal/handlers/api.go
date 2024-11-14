package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/", func(router chi.Router) {
		router.Get("/", Health)

		router.Get("/users", GetAllUsers)

		router.Post("/users", PostSignUp)

		router.Post("/email", PostSendEmail)
	})
}
