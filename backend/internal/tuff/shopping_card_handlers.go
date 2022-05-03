package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// shoppingCardCreateHandler handles shopping card creating request.
func (app *App) shoppingCardCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			shoppingCard model.Card
			err          error
		)

		if err = json.NewDecoder(r.Body).Decode(&shoppingCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		shoppingCard, err = app.store.ShoppingCards().Create(shoppingCard)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping card creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: shoppingCard,
		})
	}
}

// shoppingCardGetByLangHandler handles get all shopping cards by lang query param.
func (app *App) shoppingCardGetByLangHandler() http.HandlerFunc {
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

		shoppingCards, err := app.store.ShoppingCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping cards searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: shoppingCards,
		})
	}
}

// shoppingCardGetByIDHandler handles get shopping card by ID.
func (app *App) shoppingCardGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		shoppingCard, err := app.store.ShoppingCards().GetByID(id)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping card searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: shoppingCard,
		})
	}
}

// shoppingCardDeleteByIDHandler handles delete shopping card by ID.
func (app *App) shoppingCardDeleteByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := app.store.ShoppingCards().Delete(id)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping card deleting %v\n", err)

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

// shoppingCardUpdateByIDHandler updates (replaces) shopping card by ID.
func (app *App) shoppingCardUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id           string
			shoppingCard model.Card
			err          error
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during shopping card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(&shoppingCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update shopping card ID to id from request
		shoppingCard.ID = id

		shoppingCard, err = app.store.ShoppingCards().Update(shoppingCard)
		if err != nil {
			app.logger.Logf("[ERROR] during shopping card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: shoppingCard,
		})
	}
}
