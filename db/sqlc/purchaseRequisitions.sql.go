// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: purchaseRequisitions.sql

package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPurchaseRequisition = `-- name: CreatePurchaseRequisition :one
INSERT INTO purchaserequisitions (
    requisitionid, requesterid,title, description,status
) VALUES (
             $1, $2,$3,$4,$5
         )
    RETURNING requisitionid, requesterid, title, description, status, datesubmitted, dateapproved, daterejected
`

type CreatePurchaseRequisitionParams struct {
	Requisitionid string `json:"requisitionid"`
	Requesterid   string `json:"requesterid"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Status        string `json:"status"`
}

func (q *Queries) CreatePurchaseRequisition(ctx context.Context, arg CreatePurchaseRequisitionParams) (Purchaserequisition, error) {
	row := q.db.QueryRow(ctx, createPurchaseRequisition,
		arg.Requisitionid,
		arg.Requesterid,
		arg.Title,
		arg.Description,
		arg.Status,
	)
	var i Purchaserequisition
	err := row.Scan(
		&i.Requisitionid,
		&i.Requesterid,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Datesubmitted,
		&i.Dateapproved,
		&i.Daterejected,
	)
	return i, err
}

const deletePurchaseRequisition = `-- name: DeletePurchaseRequisition :exec
DELETE FROM purchaserequisitions
WHERE requisitionid = $1
`

func (q *Queries) DeletePurchaseRequisition(ctx context.Context, requisitionid string) error {
	_, err := q.db.Exec(ctx, deletePurchaseRequisition, requisitionid)
	return err
}

const getPurchaseRequisition = `-- name: GetPurchaseRequisition :one
SELECT requisitionid, requesterid, title, description, status, datesubmitted, dateapproved, daterejected FROM purchaserequisitions
WHERE requisitionid = $1 LIMIT 1
`

func (q *Queries) GetPurchaseRequisition(ctx context.Context, requisitionid string) (Purchaserequisition, error) {
	row := q.db.QueryRow(ctx, getPurchaseRequisition, requisitionid)
	var i Purchaserequisition
	err := row.Scan(
		&i.Requisitionid,
		&i.Requesterid,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Datesubmitted,
		&i.Dateapproved,
		&i.Daterejected,
	)
	return i, err
}

const listPurchaseRequisitions = `-- name: ListPurchaseRequisitions :many
SELECT requisitionid, requesterid, title, description, status, datesubmitted, dateapproved, daterejected FROM purchaserequisitions
ORDER BY datesubmitted
`

func (q *Queries) ListPurchaseRequisitions(ctx context.Context) ([]Purchaserequisition, error) {
	rows, err := q.db.Query(ctx, listPurchaseRequisitions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Purchaserequisition{}
	for rows.Next() {
		var i Purchaserequisition
		if err := rows.Scan(
			&i.Requisitionid,
			&i.Requesterid,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Datesubmitted,
			&i.Dateapproved,
			&i.Daterejected,
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
UPDATE purchaserequisitions
set requesterid = $2,
    description = $3,
    status= $4,
    datesubmitted=$5,
    dateapproved=$6,
    daterejected=$7

WHERE requisitionid = $1
`

type UpdatePurchaseRequisitionParams struct {
	Requisitionid string             `json:"requisitionid"`
	Requesterid   string             `json:"requesterid"`
	Description   string             `json:"description"`
	Status        string             `json:"status"`
	Datesubmitted pgtype.Timestamptz `json:"datesubmitted"`
	Dateapproved  pgtype.Timestamptz `json:"dateapproved"`
	Daterejected  pgtype.Timestamptz `json:"daterejected"`
}

func (q *Queries) UpdatePurchaseRequisition(ctx context.Context, arg UpdatePurchaseRequisitionParams) error {
	_, err := q.db.Exec(ctx, updatePurchaseRequisition,
		arg.Requisitionid,
		arg.Requesterid,
		arg.Description,
		arg.Status,
		arg.Datesubmitted,
		arg.Dateapproved,
		arg.Daterejected,
	)
	return err
}