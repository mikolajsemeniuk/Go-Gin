package inputs

import (
	e "github.com/mikolajsemeniuk/Supreme-Go/enums"
)

var ErrorMessages = map[string]string{
	"error_invalid_full_name":              "invalid full name",
	"error_invalid_email":                  "invalid email",
	"error_invalid_phone_number":           "invalid phone number",
	"error_invalid_personal_url":           "invalid personal url",
	"error_invalid_years_of_age":           "invalid years of age",
	"error_invalid_is_external_contractor": "invalid is external contractor",
	"error_invalid_relationship_status":    "invalid relationship status",
	"error_invalid_note":                   "invalid note",
}

type Account struct {
	FullName             string               `json:"full_name" binding:"required,min=3,max=100" msg:"error_invalid_full_name"`
	EmailAddress         string               `json:"email_address" binding:"required,email" msg:"error_invalid_email"`
	PhoneNumber          string               `json:"phone_number" binding:"required" msg:"error_invalid_phone_number"`
	PersonalUrl          string               `json:"personal_url" binding:"required" msg:"error_invalid_personal_url"`
	YearsOfAge           int                  `json:"years_of_age" binding:"required,numeric,min=18,max=40" msg:"error_invalid_years_of_age"`
	IsExternalContractor bool                 `json:"is_external_contractor" binding:"required" msg:"error_invalid_is_external_contractor"`
	RelationshipStatus   e.RelationshipStatus `json:"relationship_status" binding:"required,min=0,max=3" msg:"error_invalid_relationship_status"`
	Note                 string               `json:"note" binding:"required,min=5,max=500" msg:"error_invalid_note"`
}
