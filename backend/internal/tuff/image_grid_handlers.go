package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

// imageGridGetAllHandler handles getting all grid images from the database.
func (app *App) imageGridGetAllHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all grid images from the database.
		images, err := app.store.GridImages().GetAll()
		if err != nil {
			app.logger.Logf("[ERROR] failed to get all grid images: %v", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Respond with the grid images.
		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: images,
		})
	}
}

// imageGridGetAllWithURLsHandler( handles getting the URLs for the grid images.
func (app *App) imageGridGetAllWithURLsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get all grid images from the database.
		images, err := app.store.GridImages().GetAllWithURLs()
		if err != nil {
			app.logger.Logf("[ERROR] failed to get all grid images: %v", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Respond with the grid images.
		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: images,
		})
	}
}

// imageGridUpdateByIDHandler handles updating a grid image by ID.
func (app *App) imageGridUpdateByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err       error
			id        string
			gridImage model.GridImage
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] invalid grid image ID: %s", id)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Parse request body
		if err = json.NewDecoder(r.Body).Decode(&gridImage); err != nil {
			app.logger.Logf("[ERROR] failed to parse request body: %s", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Update grid image ID to id from request
		gridImage.ID = id

		// Update grid image in the database.
		gridImage, err = app.store.GridImages().Update(gridImage)
		if err != nil {
			app.logger.Logf("[ERROR] failed to update grid image: %s", err)

			app.respondJSON(w, response{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			})
			return
		}

		// Respond with updated grid image.
		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: gridImage,
		})
	}
}
