package server

import (
	"context"
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"github.com/reward-rabieth/b2b/util"
	"net/http"
)

type RequisitionRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	RequesterId   string `json:"requester_id"`
	RequisitionId string `json:"requisition_id"`
	//Status        string `json:"status"`
}

type RequisitionRequestUpdate struct {
	Status string `json:"status"`
}

type CreateBusinessParticularsRequest struct {
	RegBusinessName    string `json:"reg_business_name"`
	BrelaRegNumber     string `json:"brela_reg_number"`
	PoBox              string `json:"po_box"`
	OccupationLocation string `json:"occupation_location"`
	Country            string `json:"country"`
	Region             string `json:"region"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Contact            string `json:"contact"`
	Website            string `json:"website,omitempty"`
	Tin                string `json:"tin"`
}

func (c RequisitionRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Description, validation.Required),
		validation.Field(&c.Title, validation.Required),
	)
}

func (c CreateBusinessParticularsRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.RegBusinessName, validation.Required),
		validation.Field(&c.BrelaRegNumber, validation.Required),
		validation.Field(&c.Country, validation.Required),
		validation.Field(&c.RegBusinessName, validation.Required),
		validation.Field(&c.Contact, validation.Required),
		validation.Field(&c.FirstName, validation.Required),
		validation.Field(&c.LastName, validation.Required),
		validation.Field(&c.PoBox, validation.Required),
		validation.Field(&c.Tin, validation.Required),
		validation.Field(&c.Region, validation.Required),
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

	user, err := app.repos.ListUsers(r.Context())
	reqID := util.GenerateUUID()

	if err != nil {
		fmt.Printf("error in fetching user %s", err)
	}
	arg := users.CreatePurchaseRequisitionParams{
		RequisitionID: reqID,
		RequesterID:   user[0].Username,
		Title:         reqBody.Title,
		Description:   reqBody.Description,
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
	app.logger.Info("Requisition created:", "requisitionId", "requesterid", requisition.RequesterID, "Description", requisition.Description, "Status", requisition.Status, "Title", requisition.Title, "requisitionID")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(requisition); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode the response:%w", err), http.StatusInternalServerError, w)
	}
	return
}

func (app *App) AddBusinessParticulars(w http.ResponseWriter, r *http.Request) {
	var (
		reqBody CreateBusinessParticularsRequest
	)

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body %w", err), http.StatusInternalServerError, w)
		return
	}

	arg := users.CreateBusinessParticularParams{
		RegBusinessName:    reqBody.RegBusinessName,
		BrelaRegNumber:     reqBody.BrelaRegNumber,
		PoBox:              reqBody.PoBox,
		OccupationLocation: reqBody.OccupationLocation,
		Country:            reqBody.Country,
		Region:             reqBody.Region,
		Tin:                reqBody.Tin,
		FirstName:          reqBody.FirstName,
		LastName:           reqBody.LastName,
		Contact:            reqBody.Contact,
		Website:            reqBody.Website,
	}

	if err := reqBody.Validate(); err != nil {
		app.HandleAPIError(err, http.StatusBadRequest, w)
		return
	}

	particular, err := app.repos.CreateBusinessParticular(context.Background(), arg)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to create business particulars %w", err), http.StatusInternalServerError, w)
		return
	}
	// Convert to NewStruct such that the optional field "website" won't be returned in the response when not passed in the request
	newParticular := util.Convert(particular)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newParticular); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode response:%w", err), http.StatusInternalServerError, w)
	}

}

func (app *App) UpdatePurchaseRequisitionByID(w http.ResponseWriter, r *http.Request) {

	var reqBody RequisitionRequestUpdate
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to decode body %w", err), http.StatusInternalServerError, w)
	}

	reqID, err := uuid.Parse(r.PathValue("reqID"))
	fmt.Println(reqID)
	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to parse request ID: %w", err), http.StatusBadRequest, w)
		return
	}
	arg := users.UpdatePurchaseRequisitionParams{
		Status:        reqBody.Status,
		RequisitionID: reqID,
	}
	purchases, err := app.repos.GetPurchaseRequisitionByID(r.Context(), arg.RequisitionID)

	if err != nil {

		app.HandleAPIError(fmt.Errorf("failed to get requisition %w", err), http.StatusInternalServerError, w)

	}
	err = app.repos.UpdatePurchaseRequisition(r.Context(), arg)

	if err != nil {
		app.HandleAPIError(fmt.Errorf("failed to update requisition %w", err), http.StatusInternalServerError, w)
		return
	}

	purchases.Status = arg.Status

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(purchases); err != nil {
		app.HandleAPIError(fmt.Errorf("failed to encode response:%w", err), http.StatusInternalServerError, w)
		return
	}

}
