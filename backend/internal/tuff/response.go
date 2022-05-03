package tuff

import (
	"encoding/json"
	"net/http"
)

// response defines application's response type.
type response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// respondJSON writes response to writer with json encoding.
func (app *App) respondJSON(w http.ResponseWriter, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.logger.Logf("[ERROR] during response writing %v\n", err)
	}
}
