package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NoteRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get Notes"))
	})

	r.Get("/{ID}", func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		w.Write([]byte("Get Notes by id: " + ID))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Post Notes"))
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Put Notes"))
	})

	r.Delete("/{ID}", func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		w.Write([]byte("Delete Notes,id :" + ID))
	})

	return r
}
