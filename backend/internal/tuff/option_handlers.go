package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// optionCreateHandler handles option creating request.
func (app *App) optionCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			option model.Option
			err    error
		)

		if err = json.NewDecoder(r.Body).Decode(&option); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		option, err = app.store.Options().Create(option)
		if err != nil {
			app.logger.Logf("[ERROR] during option creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: option,
		})
	}
}

// optionGetAllHandler handler all options getting.
func (app *App) optionGetAllHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		options, err := app.store.Options().GetAll()
		if err != nil {
			app.logger.Logf("[ERROR] during options getting %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: options,
		})
	}
}

// optionUpdateValueHandler updates value field in option with given ID.
func (app *App) optionUpdateValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err    error
			option model.Option
		)

		if err = json.NewDecoder(r.Body).Decode(&option); err != nil {
			app.logger.Logf("[ERROR] during body parsing %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		option, err = app.store.Options().UpdateValue(option)
		if err != nil {
			app.logger.Logf("[ERROR] during option value updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: option,
		})
	}
}

// optionUpdateMultipleHandler updates multiple options with given IDs.
func (app *App) optionUpdateMultipleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err     error
			options []model.Option = nil
		)

		if err = json.NewDecoder(r.Body).Decode(&options); err != nil {
			app.logger.Logf("[ERROR] during body parsing %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Check if no options are specified.
		if len(options) == 0 {
			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: "no options provided",
			})
			return
		}

		// Update options.
		options, err = app.store.Options().UpdateMultiple(options)
		if err != nil {
			app.logger.Logf("[ERROR] during multi options updating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: options,
		})
	}
}
