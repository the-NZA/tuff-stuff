package tuff

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var uploadsFS = http.FileServer(http.Dir("uploads"))

// uploadsRouter returns sub router for uploads.
func (app *App) uploadsRouter() http.Handler {
	router := chi.NewRouter()

	router.Handle("/*", http.StripPrefix("/uploads/", uploadsFS))

	return router
}
