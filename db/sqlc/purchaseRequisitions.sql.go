// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: purchaseRequisitions.sql

package users

import (
	"context"

	"github.com/google/uuid"
)

const createPurchaseRequisition = `-- name: CreatePurchaseRequisition :one
INSERT INTO purchase_requisitions(
requisition_id,
requester_id,
description,
title,
status
)VALUES ($1, $2, $3, $4,$5
         ) RETURNING requisition_id, requester_id, description, title, status, date_submitted, date_approved, date_rejected
`

type CreatePurchaseRequisitionParams struct {
	RequisitionID uuid.UUID `json:"requisition_id"`
	RequesterID   string    `json:"requester_id"`
	Description   string    `json:"description"`
	Title         string    `json:"title"`
	Status        string    `json:"status"`
}

func (q *Queries) CreatePurchaseRequisition(ctx context.Context, arg CreatePurchaseRequisitionParams) (PurchaseRequisition, error) {
	row := q.db.QueryRow(ctx, createPurchaseRequisition,
		arg.RequisitionID,
		arg.RequesterID,
		arg.Description,
		arg.Title,
		arg.Status,
	)
	var i PurchaseRequisition
	err := row.Scan(
		&i.RequisitionID,
		&i.RequesterID,
		&i.Description,
		&i.Title,
		&i.Status,
		&i.DateSubmitted,
		&i.DateApproved,
		&i.DateRejected,
	)
	return i, err
}

const deletePurchaseRequisition = `-- name: DeletePurchaseRequisition :exec
DELETE FROM purchase_requisitions
WHERE requisition_id = $1
`

func (q *Queries) DeletePurchaseRequisition(ctx context.Context, requisitionID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePurchaseRequisition, requisitionID)
	return err
}

const getPurchaseRequisitionByID = `-- name: GetPurchaseRequisitionByID :one
SELECT requisition_id, requester_id, description, title, status, date_submitted, date_approved, date_rejected FROM purchase_requisitions
WHERE requisition_id = $1 LIMIT 1
`

func (q *Queries) GetPurchaseRequisitionByID(ctx context.Context, requisitionID uuid.UUID) (PurchaseRequisition, error) {
	row := q.db.QueryRow(ctx, getPurchaseRequisitionByID, requisitionID)
	var i PurchaseRequisition
	err := row.Scan(
		&i.RequisitionID,
		&i.RequesterID,
		&i.Description,
		&i.Title,
		&i.Status,
		&i.DateSubmitted,
		&i.DateApproved,
		&i.DateRejected,
	)
	return i, err
}

const listPurchaseRequisitions = `-- name: ListPurchaseRequisitions :many
SELECT requisition_id, requester_id, description, title, status, date_submitted, date_approved, date_rejected FROM purchase_requisitions
ORDER BY date_submitted
`

func (q *Queries) ListPurchaseRequisitions(ctx context.Context) ([]PurchaseRequisition, error) {
	rows, err := q.db.Query(ctx, listPurchaseRequisitions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PurchaseRequisition{}
	for rows.Next() {
		var i PurchaseRequisition
		if err := rows.Scan(
			&i.RequisitionID,
			&i.RequesterID,
			&i.Description,
			&i.Title,
			&i.Status,
			&i.DateSubmitted,
			&i.DateApproved,
			&i.DateRejected,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePurchaseRequisition = `-- name: UpdatePurchaseRequisition :exec
UPDATE purchase_requisitions
set status= $1
WHERE requisition_id = $2
RETURNING requisition_id, requester_id, description, title, status, date_submitted, date_approved, date_rejected
`

type UpdatePurchaseRequisitionParams struct {
	Status        string    `json:"status"`
	RequisitionID uuid.UUID `json:"requisition_id"`
}

func (q *Queries) UpdatePurchaseRequisition(ctx context.Context, arg UpdatePurchaseRequisitionParams) error {
	_, err := q.db.Exec(ctx, updatePurchaseRequisition, arg.Status, arg.RequisitionID)
	return err
}
