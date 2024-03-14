package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"net/http"
)

type _createQuotationRequest struct {
	ProposalRequestID int64     `json:"proposal_request_id"`
	RequisitionID     uuid.UUID `json:"requisition_id"`
	SupplierID        uuid.UUID `json:"supplier_id"`
	Status            string    `json:"status"`
}

func (app *App) createQuotationRequest(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody _createQuotationRequest
	)

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body %w", err), http.StatusInternalServerError, w)
		return
	}

	requisitions, err := app.repos.ListPurchaseRequisitions(r.Context())
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to list purchase requisitions: %w", err), http.StatusInternalServerError, w)
		return
	}

	if len(requisitions) == 0 {
		app.HandleAPIError(errors.New("no purchase requisitions available"), http.StatusNotFound, w)
		return
	}

	//requisitionID := requisitions[0].RequisitionID

	supplier, err := app.repos.ListSupplier(r.Context())
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to list purchase requisitions: %w", err), http.StatusInternalServerError, w)
		return
	}

	if len(supplier) == 0 {
		app.HandleAPIError(errors.New("no p supplier available"), http.StatusNotFound, w)
		return
	}

	arg := users.CreateQuotationRequestParams{
		ProposalRequestID: reqBody.ProposalRequestID,
		RequisitionID:     reqBody.RequisitionID,
		SupplierID:        reqBody.SupplierID,
		Status:            "", // todo:You need to set the status appropriately
	}
	request, err := app.repos.CreateQuotationRequest(r.Context(), arg)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create quotation request: %w", err), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(request); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response: %w", err), http.StatusInternalServerError, w)
		return
	}
}
