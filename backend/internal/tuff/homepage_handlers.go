package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// homepageCreateHandler handles creating new homepage.
func (app *App) homepageCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			homepage = &model.Homepage{}
			err      error
		)

		if err = json.NewDecoder(r.Body).Decode(homepage); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Save homepage
		homepage, err = app.store.Homepages().Create(homepage)
		if err != nil {
			app.logger.Logf("[ERROR] during homepage creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: homepage,
		})
	}
}

// homepageGetByLangHandler handles getting homepage by given lang.
func (app *App) homepageGetByLangHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")

		if len(lang) != langParamLength {
			app.logger.Logf("[ERROR] during query params getting\n")

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidQueryParam.Error(),
			})
			return
		}

		homepage, err := app.store.Homepages().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during homepage searching %v \n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: homepage,
		})
	}
}

// homepageUpdateByIDHandler handles updating homepage with given ID.
func (app *App) homepageUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id       string
			homepage = &model.Homepage{}
			err      error
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during homepage updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(homepage); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update homepage ID to id from request
		homepage.ID = id

		// Update homepage
		homepage, err = app.store.Homepages().Update(homepage)
		if err != nil {
			app.logger.Logf("[ERROR] during homepage updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: homepage,
		})
	}
}
