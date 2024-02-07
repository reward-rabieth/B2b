package server

import (
	"context"
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"github.com/reward-rabieth/b2b/util"
	"net/http"
)

type RequisitionRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	RequesterId   string `json:"requester_id"`
	RequisitionId string `json:"requisition_id"`
	Status        string `json:"status"`
}

func (c RequisitionRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Description, validation.Required),
		validation.Field(&c.Title, validation.Required),
	)
}

func (app *App) CreateRequisition(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody RequisitionRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body: %w", err), http.StatusInternalServerError, w)
		return
	}
	status := "submitted"
	arg := users.CreatePurchaseRequisitionParams{
		Requisitionid: util.GenerateUUID(),
		Requesterid:   util.GenerateUUID(),
		Title:         reqBody.Title,
		Description:   reqBody.Description,
		Status:        status,
	}
	fmt.Println(arg)

	if err := reqBody.Validate(); err != nil {
		fmt.Println(err)
		app.HandleAPIError(err, http.StatusBadRequest, w)
		return
	}

	requisition, err := app.repos.CreatePurchaseRequisition(context.Background(), arg)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create requisition: %w", err), http.StatusInternalServerError, w)
		return
	}
	app.logger.Info("Requisition created:", "requisitionId", "requesterid", requisition.Requesterid, "Description", requisition.Description, "Status", requisition.Status, "Title", requisition.Title, "requisitionID")
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create requisition: %w", err), http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(requisition); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response:%w", err), http.StatusInternalServerError, w)
	}
	return
}
