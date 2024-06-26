// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: businessParticulars.sql

package users

import (
	"context"
)

const createBusinessParticular = `-- name: CreateBusinessParticular :one
INSERT INTO business_particulars(REG_BUSINESS_NAME,
                                 BRELA_REG_NUMBER,
                                 USER_ID,
                                 PO_BOX,
                                 OCCUPATION_LOCATION,
                                 COUNTRY,
                                 REGION,
                                 TIN,
                                 FIRST_NAME,
                                 LAST_NAME,
                                 CONTACT,
                                 WEBSITE
                                 ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
returning reg_business_name, brela_reg_number, user_id, po_box, occupation_location, country, region, tin, first_name, last_name, contact, website
`

type CreateBusinessParticularParams struct {
	RegBusinessName    string `json:"reg_business_name"`
	BrelaRegNumber     string `json:"brela_reg_number"`
	UserID             int32  `json:"user_id"`
	PoBox              string `json:"po_box"`
	OccupationLocation string `json:"occupation_location"`
	Country            string `json:"country"`
	Region             string `json:"region"`
	Tin                string `json:"tin"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Contact            string `json:"contact"`
	Website            string `json:"website"`
}

func (q *Queries) CreateBusinessParticular(ctx context.Context, arg CreateBusinessParticularParams) (BusinessParticular, error) {
	row := q.db.QueryRow(ctx, createBusinessParticular,
		arg.RegBusinessName,
		arg.BrelaRegNumber,
		arg.UserID,
		arg.PoBox,
		arg.OccupationLocation,
		arg.Country,
		arg.Region,
		arg.Tin,
		arg.FirstName,
		arg.LastName,
		arg.Contact,
		arg.Website,
	)
	var i BusinessParticular
	err := row.Scan(
		&i.RegBusinessName,
		&i.BrelaRegNumber,
		&i.UserID,
		&i.PoBox,
		&i.OccupationLocation,
		&i.Country,
		&i.Region,
		&i.Tin,
		&i.FirstName,
		&i.LastName,
		&i.Contact,
		&i.Website,
	)
	return i, err
}

const updateBusinessParticulars = `-- name: UpdateBusinessParticulars :exec
UPDATE business_particulars
SET reg_business_name=$1,
    brela_reg_number=$2,
    po_box=$3,
    occupation_location=$4,
    country=$5,
    region=$6,
    tin=$7,
    first_name=$8,
    last_name=$9,
    contact=$10,
    website=$11
where brela_reg_number=$2
`

type UpdateBusinessParticularsParams struct {
	RegBusinessName    string `json:"reg_business_name"`
	BrelaRegNumber     string `json:"brela_reg_number"`
	PoBox              string `json:"po_box"`
	OccupationLocation string `json:"occupation_location"`
	Country            string `json:"country"`
	Region             string `json:"region"`
	Tin                string `json:"tin"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Contact            string `json:"contact"`
	Website            string `json:"website"`
}

func (q *Queries) UpdateBusinessParticulars(ctx context.Context, arg UpdateBusinessParticularsParams) error {
	_, err := q.db.Exec(ctx, updateBusinessParticulars,
		arg.RegBusinessName,
		arg.BrelaRegNumber,
		arg.PoBox,
		arg.OccupationLocation,
		arg.Country,
		arg.Region,
		arg.Tin,
		arg.FirstName,
		arg.LastName,
		arg.Contact,
		arg.Website,
	)
	return err
}
