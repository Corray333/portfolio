package post

import (
	post_storage "github.com/Corray333/portfolio/internal/domains/post/storage"
	"github.com/Corray333/portfolio/internal/domains/post/transport"
	"github.com/Corray333/portfolio/internal/storage"
	"github.com/Corray333/portfolio/pkg/server/auth"
	"github.com/go-chi/chi/v5"
)

func MustInit(router chi.Router, storage storage.Storage) {
	store := post_storage.New(storage)

	router.Get("/api/posts", transport.GetPosts(store))
	router.Get("/api/posts/{id}", transport.GetPost(store))
	router.Post("/api/posts/{id}/reactions", transport.UpdateReaction(store))
	router.With(auth.NewMiddleware()).Post("/api/posts", transport.CreatePost(store))
	router.With(auth.NewMiddleware()).Post("/api/upload/image", transport.UploadImage())
}
