package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// howWorksCardCreateHandler handles how works card creating request.
func (app *App) howWorksCardCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			howWorksCard model.Card
			err          error
		)

		if err = json.NewDecoder(r.Body).Decode(&howWorksCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		howWorksCard, err = app.store.HowWorksCards().Create(howWorksCard)
		if err != nil {
			app.logger.Logf("[ERROR] during how works card creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: howWorksCard,
		})

	}
}

// howWorksCardGetByLangHandler handles get all how works cards by lang query param.
func (app *App) howWorksCardGetByLangHandler() http.HandlerFunc {
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

		howWorksCards, err := app.store.HowWorksCards().GetByLang(lang)
		if err != nil {
			app.logger.Logf("[ERROR] during how works cards searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: howWorksCards,
		})
	}
}

// howWorksCardGetByIDHandler handles get how works card by ID.
func (app *App) howWorksCardGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		howWorksCard, err := app.store.HowWorksCards().GetByID(id)
		if err != nil {
			app.logger.Logf("[ERROR] during how works card searching %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: howWorksCard,
		})
	}
}

// howWorksCardDeleteByIDHandler handles delete how works card by ID.
func (app *App) howWorksCardDeleteByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := app.store.HowWorksCards().Delete(id)
		if err != nil {
			app.logger.Logf("[ERROR] during how works card deleting %v\n", err)

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

// howWorksCardUpdateByIDHandler updates (replaces) how works card by ID.
func (app *App) howWorksCardUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			id           string
			howWorksCard model.Card
			err          error
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during how works card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(&howWorksCard); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update how works card ID to id from request
		howWorksCard.ID = id

		howWorksCard, err = app.store.HowWorksCards().Update(howWorksCard)
		if err != nil {
			app.logger.Logf("[ERROR] during how works card updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: howWorksCard,
		})
	}
}
