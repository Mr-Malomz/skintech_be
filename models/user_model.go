package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID `json:"_id,omitempty"`
	Firstname      string             `json:"firstname,omitempty" validate:"required,min=2"`
	Lastname       string             `json:"lastname,omitempty" validate:"required,min=2"`
	Phone_Number   string             `json:"phone_number,omitempty"`
	Email          string             `json:"email,omitempty" validate:"required,email"`
	MDCN           string             `json:"mdcn,omitempty"`
	Anual_lics_num string             `json:"anual_lics_num,omitempty"`
	Folio_num      string             `json:"folio_num,omitempty"`
	School_grad    string             `json:"school_grad,omitempty"`
	Token          string             `json:"token,omitempty"`
	OTP            string             `json:"otp,omitempty"`
	Refresh_token  string             `json:"refresh_token,omitempty"`
	Year_grad      int                `json:"year_grad,omitempty"`
	Year_fellow    int                `json:"year_fellow,omitempty"`
	Cert_url       string             `json:"cert_url,omitempty"`
	Created_At     time.Time          `json:"created_at,omitempty"`
	Password       string             `json:"password,omitempty" validate:"required,min=8"`
	IsActive       bool               `json:"isActive,omitempty"`
	IsVerified     bool               `json:"isVerified,omitempty"`
}

type UpdateUser struct {
	MDCN           string `json:"mdcn,omitempty"`
	Anual_lics_num string `json:"anual_lics_num,omitempty"`
	Folio_num      string `json:"folio_num,omitempty"`
	School_grad    string `json:"school_grad,omitempty"`
	Year_grad      int    `json:"year_grad,omitempty"`
	Year_fellow    int    `json:"year_fellow,omitempty"`
	Cert_url       string `json:"cert_url,omitempty"`
}
