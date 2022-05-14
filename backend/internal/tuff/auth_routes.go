package tuff

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/tuff-stuff/backend/internal/jwt"
	"github.com/the-NZA/tuff-stuff/backend/internal/model"
)

const (
	authCookieName = "tuff_token"
)

// authRouter returns sub router for auth.
func (app *App) authRouter() http.Handler {
	router := chi.NewRouter()

	router.Post("/login", app.loginHandler())
	router.Post("/register", app.registerHandler())

	return router
}

func (app *App) loginHandler() http.HandlerFunc {
	type LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  LoginRequest
			err  error
			user model.User
		)
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Data:  req,
				Error: err.Error(),
			})
			return
		}

		user, err = app.store.Users().Login(req.Login, req.Password)
		if err != nil {
			app.logger.Logf("[ERROR] during user login %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Data:  req,
				Error: err.Error(),
			})
			return
		}

		token, expireTime, err := jwt.Create(user.Login, app.config.SecretKey)
		if err != nil {
			app.logger.Logf("[ERROR] during user login %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Data:  req,
				Error: err.Error(),
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     authCookieName,
			Value:    token,
			HttpOnly: true,
			Expires:  expireTime,
			Path:     "/",
			Domain:   app.config.Domain,
		})

		app.respondJSON(w, response{
			Code: http.StatusOK,
			Data: map[string]string{
				"login": user.Login,
				"token": token,
			},
		})
	}
}

func (app *App) registerHandler() http.HandlerFunc {
	type RegisterRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			req  RegisterRequest
			err  error
			user model.User
		)

		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			app.logger.Logf("[ERROR] during body decoding %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Data:  req,
				Error: err.Error(),
			})
			return
		}

		user, err = app.store.Users().Register(req.Login, req.Password)
		if err != nil {
			app.logger.Logf("[ERROR] during user register %v\n", err)

			app.respondJSON(w, response{
				Code:  http.StatusBadRequest,
				Data:  req,
				Error: err.Error(),
			})
			return
		}

		app.respondJSON(w, response{
			Code: http.StatusCreated,
			Data: user,
		})
	}
}
