package server

import (
	"encoding/json"
	"fmt"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"github.com/reward-rabieth/b2b/util"
	"net/http"
)

type SupplierRequest struct {
	//SupplierID    uuid.UUID `json:"supplier_id"`
	SupplierName  string `json:"supplier_name"`
	ContactPerson string `json:"contact_person"`
	ContactMail   string `json:"contact_mail"`
	SupplierType  string `json:"supplier_type"`
}

func (app *App) createSupplier(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody SupplierRequest
	)
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body %w", err), http.StatusInternalServerError, w)
		return
	}
	supID := util.GenerateUUID()
	arg := users.CreateSupplierParams{
		SupplierID:    supID,
		SupplierName:  reqBody.SupplierName,
		ContactPerson: reqBody.ContactPerson,
		ContactMail:   reqBody.ContactMail,
		SupplierType:  reqBody.SupplierType,
	}

	supplier, err := app.repos.CreateSupplier(r.Context(), arg)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create supplier %w", err), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(supplier)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response: %w", err), http.StatusInternalServerError, w)
	}

	return
}
