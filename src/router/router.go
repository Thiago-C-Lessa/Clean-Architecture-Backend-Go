package router

import "github.com/go-chi/chi/v5"

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api", setupRoutes())

	return r
}

func setupRoutes() chi.Router {
	r := chi.NewRouter()

	r.Mount("/notes", NoteRoutes())
	return r
}
