package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname      string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,min=2"`
	Lastname       string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,min=2"`
	Phone_Number   string             `json:"phone_number,omitempty" bson:"phone_number,omitempty" validate:"min=11"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	MDCN           string             `json:"mdcn,omitempty" bson:"mdcn,omitempty"`
	Anual_lics_num string             `json:"anual_lics_num,omitempty" bson:"anual_lics_num,omitempty"`
	Folio_num      string             `json:"folio_num,omitempty" bson:"folio_num,omitempty"`
	School_grad    string             `json:"school_grad,omitempty" bson:"school_grad,omitempty"`
	Token          string             `json:"token,omitempty" bson:"token,omitempty"`
	Refresh_token  string             `json:"refresh_token,omitempty" bson:"token,omitempty"`
	Year_grad      int                `json:"year_grad,omitempty" bson:"year_grad,omitempty"`
	Year_fellow    int                `json:"year_fellow,omitempty" bson:"year_fellow,omitempty"`
	Cert_url       string             `json:"cert_url,omitempty" bson:"cert_url,omitempty"`
	Created_At     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	IsActive       bool               `json:"isActive,omitempty" bson:"isActive,omitempty"`
	IsVerified     bool               `json:"isVerified,omitempty" bson:"isVerified,omitempty"`
}