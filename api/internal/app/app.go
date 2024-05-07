package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Corray333/portfolio/internal/domains/admin"
	"github.com/Corray333/portfolio/internal/domains/post"
	"github.com/Corray333/portfolio/internal/storage"
	"github.com/go-chi/chi/v5"
)

type App struct {
	router chi.Router
	store  storage.Storage
}

func New() *App {
	router := chi.NewMux()
	store := storage.New()

	fs := http.FileServer(http.Dir("../files/images"))
	router.Handle("/api/images/*", http.StripPrefix("/api/images", fs))

	post.MustInit(router, store)
	admin.MustInit(router)

	return &App{
		router: router,
		store:  store,
	}
}
func (a *App) Run() {
	slog.Info("Server started on " + os.Getenv("HOST_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("HOST_PORT"), a.router); err != nil {
		slog.Error("error while running server: " + err.Error())
		panic(err)
	}
}
