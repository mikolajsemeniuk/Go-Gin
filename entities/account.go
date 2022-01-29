package entities

import (
	e "github.com/mikolajsemeniuk/Supreme-Go/enums"
)

type Account struct {
	Entity
	FullName             string               `json:"full_name"`
	EmailAddress         string               `json:"email_address"`
	PhoneNumber          string               `json:"phone_number"`
	PersonalUrl          string               `json:"personal_url"`
	YearsOfAge           int                  `json:"years_of_age"`
	IsExternalContractor bool                 `json:"is_external_contractor"`
	RelationshipStatus   e.RelationshipStatus `json:"relationship_status"`
	Note                 string               `json:"note"`
}
