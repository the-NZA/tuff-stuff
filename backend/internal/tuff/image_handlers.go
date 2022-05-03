package tuff

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	maxFileSize int64 = 32 << 20 // Max upload file size 32MB
)

var (
	ErrUnsupportedMIMEType = errors.New("unsupported MIME type")

	allowedMimeTypes = []string{
		"image/png", "image/vnd.mozilla.png", "image/jpeg", "image/gif", "image/svg+xml",
	}
)

// imageCreateHandler handles saving new image.
func (app *App) imageCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err   error
			image model.Image
		)

		// Parse multipart form data
		if err = r.ParseMultipartForm(maxFileSize); err != nil {
			app.logger.Logf("[ERROR] parse multipart: %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Get file from form
		file, header, err := r.FormFile("tuff_image")
		if err != nil {
			app.logger.Logf("[ERROR] during getting form data %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		defer func(file multipart.File) {
			internalError := file.Close()
			if internalError != nil {
				app.logger.Logf("[ERROR] form file closing: %v\n", internalError)
			}
		}(file)

		// Get MIME type
		mimeType, err := mimetype.DetectReader(file)
		if err != nil {
			app.logger.Logf("[ERROR] during MIME type checking %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Check MIME type
		if !mimetype.EqualsAny(mimeType.String(), allowedMimeTypes...) {
			app.logger.Logf("[ERROR] during MIME type checking %v\n", ErrUnsupportedMIMEType)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrUnsupportedMIMEType.Error(),
			})
			return
		}

		// Generate save path
		savePath := fmt.Sprintf("uploads/images/%d_%s", time.Now().Unix(), header.Filename)

		// Restore file cursor
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			app.logger.Logf("[ERROR] during file seeking %v\n", ErrUnsupportedMIMEType)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrUnsupportedMIMEType.Error(),
			})
			return
		}

		dst, err := os.Create(savePath)
		if err != nil {
			app.logger.Logf("[ERROR] during dst file creating %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}
		defer func(dst *os.File) {
			internalError := dst.Close()
			if internalError != nil {
				app.logger.Logf("[ERROR] during dst file closing %v\n", internalError)
			}
		}(dst)

		// Copy form file to dst
		_, err = io.Copy(dst, file)
		if err != nil {
			app.logger.Logf("[ERROR] during file copying %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		// Save image info
		image, err = app.store.Images().Create(model.Image{URL: savePath})
		if err != nil {
			app.logger.Logf("[ERROR] during file saving %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: image,
		})
	}
}

// imageGetByIDHandler handles getting image by giving ID.
func (app *App) imageGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err   error
			image model.Image
			id    string
		)

		id = chi.URLParam(r, "id")
		if len(id) < 1 {
			app.logger.Logf("[ERROR] during image getting\n")

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: ErrInvalidURLParam.Error(),
			})
			return
		}

		// Get image by given ID
		image, err = app.store.Images().GetByID(id)
		if err != nil {
			app.logger.Logf("[ERROR] during image getting %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: image,
		})
	}
}
