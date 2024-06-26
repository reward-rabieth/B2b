package server

import (
	"encoding/json"
	"github.com/rs/cors"
	"net/http"
)

func (app *App) NewHandler() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("POST /api/procurement", app.AuthMiddleware(app.CreateRequisition))
	r.HandleFunc("POST /register", app.Register)
	r.HandleFunc("POST /login", app.Login)
	r.HandleFunc("POST /api/business", app.AddBusinessParticulars)
	r.HandleFunc("POST /api/quotation", app.createQuotationRequest)
	r.HandleFunc("POST /api/supplier", app.createSupplier)
	r.HandleFunc("/api/procurement/{reqID}", app.UpdatePurchaseRequisitionByID)

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
