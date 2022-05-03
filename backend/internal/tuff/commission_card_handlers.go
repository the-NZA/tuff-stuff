package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// commissionCardCreateHandler handles commission card creating request.
func (app *App) commissionCardCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			commissionCard model.Card
			err            error
		)

		if err = json.NewDecoder(r.Body).Decode(&commissionCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		commissionCard, err = app.store.CommissionCards().Create(commissionCard)
		if err != nil {
			app.logger.Logf("[ERROR] during commission card creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: commissionCard,
		})
	}
}

// commissionCardGetByLangHandler handles get all commission cards by lang query param.
func (app *App) commissionCardGetByLangHandler() http.HandlerFunc {
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

		commissionCards, err := app.store.CommissionCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during commission cards searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: commissionCards,
		})
	}
}

// commissionCardGetByIDHandler handles get commission card by ID.
func (app *App) commissionCardGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		commissionCard, err := app.store.CommissionCards().GetByID(id)
		if err != nil {
			app.logger.Logf("[ERROR] during commission card searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: commissionCard,
		})
	}
}

// commissionCardDeleteByIDHandler handles delete commission card by ID.
func (app *App) commissionCardDeleteByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := app.store.CommissionCards().Delete(id)
		if err != nil {
			app.logger.Logf("[ERROR] during commission card deleting %v\n", err)

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

// commissionCardUpdateByIDHandler updates (replaces) commission card by ID.
func (app *App) commissionCardUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id             string
			commissionCard model.Card
			err            error
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during commission card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(&commissionCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update commission card ID to id from request
		commissionCard.ID = id

		commissionCard, err = app.store.CommissionCards().Update(commissionCard)
		if err != nil {
			app.logger.Logf("[ERROR] during commission card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: commissionCard,
		})
	}
}
