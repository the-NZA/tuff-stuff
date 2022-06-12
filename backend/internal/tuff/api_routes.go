package tuff

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	langParamLength = 2
)

var (
	ErrInvalidQueryParam = errors.New("invalid query parameter")
	ErrInvalidURLParam   = errors.New("invalid url parameter")
)

// apiRouter returns sub router for api.
func (app *App) apiRouter() http.Handler {
	router := chi.NewRouter()
	// TODO: enable auth middleware before go to production.
	// router.Use(app.authMiddleware) // require auth for api access

	router.Get("/", app.rootAPIHandler())

	// Shopping cards routes
	router.Route("/shopping-card", func(r chi.Router) {
		r.Post("/", app.shoppingCardCreateHandler())
		r.Get("/", app.shoppingCardGetByLangHandler())
		r.Get("/{id}", app.shoppingCardGetByIDHandler())
		r.Delete("/{id}", app.shoppingCardDeleteByIDHandler())
		r.Put("/{id}", app.shoppingCardUpdateByIDHandler())
	})

	// Commission cards routes
	router.Route("/commission-card", func(r chi.Router) {
		r.Post("/", app.commissionCardCreateHandler())
		r.Get("/", app.commissionCardGetByLangHandler())
		r.Get("/{id}", app.commissionCardGetByIDHandler())
		r.Delete("/{id}", app.commissionCardDeleteByIDHandler())
		r.Put("/{id}", app.commissionCardUpdateByIDHandler())
	})

	// How works cards routes
	router.Route("/how-works-card", func(r chi.Router) {
		r.Post("/", app.howWorksCardCreateHandler())
		r.Get("/", app.howWorksCardGetByLangHandler())
		r.Get("/{id}", app.howWorksCardGetByIDHandler())
		r.Delete("/{id}", app.howWorksCardDeleteByIDHandler())
		r.Put("/{id}", app.howWorksCardUpdateByIDHandler())
	})

	// Why us cards routes
	router.Route("/why-us-card", func(r chi.Router) {
		r.Post("/", app.whyUsCardCreateHandler())
		r.Get("/", app.whyUsCardGetByLangHandler())
		r.Get("/{id}", app.whyUsCardGetByIDHandler())
		r.Delete("/{id}", app.whyUsCardDeleteByIDHandler())
		r.Put("/{id}", app.whyUsCardUpdateByIDHandler())
	})

	// Options routes
	router.Route("/option", func(r chi.Router) {
		r.Post("/", app.optionCreateHandler())
		r.Get("/", app.optionGetAllHandler())
		r.Patch("/", app.optionUpdateValueHandler())
		r.Put("/", app.optionUpdateMultipleHandler())
	})

	// Homepages routes
	router.Route("/homepage", func(r chi.Router) {
		r.Post("/", app.homepageCreateHandler())
		r.Get("/", app.homepageGetByLangHandler())
		r.Put("/{id}", app.homepageUpdateByIDHandler())
	})

	// ImageGrid routes
	router.Route("/image-grid", func(r chi.Router) {
		r.Get("/", app.imageGridGetAllHandler())
		r.Put("/{id}", app.imageGridUpdateByIDHandler())
	})

	// Images routes
	router.Route("/image", func(r chi.Router) {
		r.Post("/", app.imageCreateHandler())
		r.Get("/{id}", app.imageGetByIDHandler())
	})

	return router
}

// rootAPIHandler sends response with api version.
func (app *App) rootAPIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: "API is running. Version 1.",
		})
	}
}
