package transport

import (
	"encoding/json"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Corray333/portfolio/internal/domains/post/types"
	"github.com/go-chi/chi/v5"
)

const MaxFileSize = 5 << 20

type Storage interface {
	SelectPosts(post_type string, title string, id string, lang string, tags []string, user_agent string, offset uint64) ([]types.Post, error)
	SelectPost(post_id string, lang string) (*types.Post, error)
	InsertPost(langs []types.Post) (int, error)
	UpdateReaction(post_id string, increment, decrement int) error
}

func GetPosts(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		lang := r.URL.Query().Get("lang")
		tags := r.URL.Query()["tags"]
		post_id := r.URL.Query().Get("id")
		post_type := r.URL.Query().Get("type")
		offset := r.URL.Query().Get("offset")

		var uintOffset uint64
		var err error
		if offset != "" {
			if uintOffset, err = strconv.ParseUint(offset, 10, 64); err != nil {
				slog.Error("error parsing offset: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		posts, err := store.SelectPosts(post_type, title, post_id, lang, tags, r.UserAgent(), uintOffset)
		if err != nil {
			slog.Error("error getting posts", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(struct {
			Posts []types.Post `json:"posts"`
		}{
			Posts: posts,
		}); err != nil {
			slog.Error("error encoding or sending posts: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func GetPost(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post_id := chi.URLParam(r, "id")
		lang := r.URL.Query().Get("lang")

		post, err := store.SelectPost(post_id, lang)
		if err != nil {
			slog.Error("error getting posts", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(struct {
			Post *types.Post `json:"post"`
		}{
			Post: post,
		}); err != nil {
			slog.Error("error encoding or sending posts: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func CreatePost(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var langs []types.Post
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("error reading body: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(body, &langs); err != nil {
			slog.Error("error decoding post: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		post_id, err := store.InsertPost(langs)
		if err != nil {
			slog.Error("error inserting post: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(struct {
			ID int `json:"id"`
		}{
			ID: post_id,
		}); err != nil {
			slog.Error("error encoding or sending posts: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func UploadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(MaxFileSize); err != nil {
			slog.Error("error parsing multipart form: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			slog.Error("error getting file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		randomStr := ""
		for i := 0; i < 10; i++ {
			randomStr += strconv.Itoa(rand.Intn(10))
		}
		fileName := randomStr + filepath.Ext(header.Filename)
		newFile, err := os.Create("../files/images/" + fileName)
		if err != nil {
			slog.Error("error creating file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer newFile.Close()
		if _, err := io.Copy(newFile, file); err != nil {
			slog.Error("error copying file: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(struct {
			URL string `json:"url"`
		}{
			URL: "/api/images/" + fileName,
		}); err != nil {
			slog.Error("error encoding or sending file name: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func UpdateReaction(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post_id := chi.URLParam(r, "id")

		var request struct {
			Increment int `json:"increment"`
			Decrement int `json:"decrement"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			slog.Error("error decoding request: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := store.UpdateReaction(post_id, request.Increment, request.Decrement); err != nil {
			slog.Error("error updating reaction: ", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
