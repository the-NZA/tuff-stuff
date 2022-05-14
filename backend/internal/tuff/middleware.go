package tuff

import (
	"net/http"

	"github.com/the-NZA/tuff-stuff/backend/internal/jwt"
)

func (app *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(authCookieName)
		if err != nil {
			app.logger.Logf("[ERROR] during cookie parse %v\n", err)

			switch err {
			case http.ErrNoCookie:
				app.respondJSON(w, response{
					Code:  http.StatusUnauthorized,
					Error: err.Error(),
				})
			default:
				app.respondJSON(w, response{
					Code:  http.StatusInternalServerError,
					Error: err.Error(),
				})
			}

			return
		}

		parsedToken, err := jwt.Parse(cookie.Value, app.config.SecretKey)
		if err != nil {
			app.logger.Logf("[ERROR] during token parsing %v", err)
			app.respondJSON(w, response{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			return
		}

		if !parsedToken.Valid {
			app.logger.Logf("[ERROR] during token validation %v", jwt.ErrInvalidToken)
			app.respondJSON(w, response{
				Code:  http.StatusUnauthorized,
				Error: jwt.ErrInvalidToken.Error(),
			})
			return
		}

		needUpdate, err := jwt.NeedUpdate(parsedToken)
		if err != nil {
			app.logger.Logf("[ERROR] during token checking %v", err)
			app.respondJSON(w, response{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			})
			return
		}

		if needUpdate {
			// Get claims
			claims, ok := parsedToken.Claims.(*jwt.Claims)
			if !ok {
				app.logger.Logf("[ERROR] during token update %v\n", err)
				app.respondJSON(w, response{
					Code:  http.StatusUnauthorized,
					Error: jwt.ErrInvalidTokenClaims.Error(),
				})
				return
			}

			// Get login from claims
			login := claims.Login
			token, expireTime, err := jwt.Create(login, app.config.SecretKey)
			if err != nil {
				app.logger.Logf("[ERROR] during token update %v\n", err)
				app.respondJSON(w, response{
					Code:  http.StatusUnauthorized,
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

		}

		next.ServeHTTP(w, r)
	})
}
