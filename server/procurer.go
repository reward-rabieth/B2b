package server

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/reward-rabieth/b2b/core/components/Procurer/models"
	"net/http"
)

type RequisitionRequest struct {
	Title   string `json:"title"`
	Status  string `json:"status"`
	Details string `json:"details"`
}

func (c RequisitionRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Details, validation.Required),
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Status, validation.Required),
	)
}

func (app *App) CreateRequisition(w http.ResponseWriter, r *http.Request) {
	var reqBody RequisitionRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body: %w", err), http.StatusInternalServerError, w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		app.HandleAPIError(err, http.StatusBadRequest, w)
		return
	}

	createRequisition, err := app.procurerComponent.Create(r.Context(), &models.Requisition{
		Status:  reqBody.Status,
		Title:   reqBody.Title,
		Details: reqBody.Details,
	})
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create requisition: %w", err), http.StatusInternalServerError, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createRequisition); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response:%w", err), http.StatusInternalServerError, w)
	}
	return
}
