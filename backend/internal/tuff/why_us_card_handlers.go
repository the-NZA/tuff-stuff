package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// whyUsCardCreateHandler handles why us card creating request.
func (app *App) whyUsCardCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			card model.Card
			err  error
		)

		if err = json.NewDecoder(r.Body).Decode(&card); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		card, err = app.store.WhyUsCards().Create(card)
		if err != nil {
			app.logger.Logf("[ERROR] during why us card creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: card,
		})

	}
}

// whyUsCardGetByLangHandler handles get all why us cards by lang query param.
func (app *App) whyUsCardGetByLangHandler() http.HandlerFunc {
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

		cards, err := app.store.WhyUsCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during why us cards searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: cards,
		})
	}
}

// whyUsCardGetByIDHandler handles get why us card by ID.
func (app *App) whyUsCardGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		card, err := app.store.WhyUsCards().GetByID(id)
		if err != nil {
			app.logger.Logf("[ERROR] during why us card searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: card,
		})
	}
}

// whyUsCardDeleteByIDHandler handles delete why us card by ID.
func (app *App) whyUsCardDeleteByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := app.store.WhyUsCards().Delete(id)
		if err != nil {
			app.logger.Logf("[ERROR] during why us card deleting %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: "deleted",
		})
	}
}

// whyUsCardUpdateByIDHandler updates (replaces) why us card by ID.
func (app *App) whyUsCardUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id   string
			card model.Card
			err  error
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during why us card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(&card); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update why us card ID to id from request
		card.ID = id

		card, err = app.store.WhyUsCards().Update(card)
		if err != nil {
			app.logger.Logf("[ERROR] during why us card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: card,
		})
	}
}
