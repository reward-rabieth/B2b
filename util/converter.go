package util

import users "github.com/reward-rabieth/b2b/db/sqlc"

type NewStruct struct {
	Website            string `json:"website,omitempty"`
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
}

func Convert(input users.BusinessParticular) NewStruct {
	// Check if Website field is empty, if so, set it to nil
	var website string
	if input.Website != "" {
		websiteValue := input.Website
		website = websiteValue
	}

	return NewStruct{
		Website:            website,
		RegBusinessName:    input.RegBusinessName,
		BrelaRegNumber:     input.BrelaRegNumber,
		UserID:             input.UserID,
		PoBox:              input.PoBox,
		OccupationLocation: input.OccupationLocation,
		Country:            input.Country,
		Region:             input.Region,
		Tin:                input.Tin,
		FirstName:          input.FirstName,
		LastName:           input.LastName,
		Contact:            input.Contact,
	}
}
