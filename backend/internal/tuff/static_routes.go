package tuff

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var staticFS = http.FileServer(http.Dir("static"))

// staticRouter returns sub router for static files.
func (app *App) staticRouter() http.Handler {
	router := chi.NewRouter()

	router.Handle("/*", http.StripPrefix("/static/", staticFS))

	return router
}
