package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func (app *App) NewHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/procurement", app.CreateRequisition).Methods(http.MethodPost)
	corsHandler := cors.AllowAll()
	return corsHandler.Handler(r)
}

func (app *App) HandleAPIError(err error, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": err.Error(),
	}); err != nil {
		w.Write([]byte("failed to encode error message"))
		return
	}
}
