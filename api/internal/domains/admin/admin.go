package admin

import (
	"github.com/Corray333/portfolio/internal/domains/admin/transport"
	"github.com/go-chi/chi/v5"
)

func MustInit(router *chi.Mux) {
	router.Post("/api/admin/login", transport.LogIn())
	router.Get("/api/admin/refresh", transport.RefreshAccessToken())
}
