package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender struct {
	MALE   string
	FEMALE string
}

type Images struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty" validate:"required"`
	Skin_diag   string             `json:"skin_diag,omitempty" bson:"skin_diag,omitempty" validate:"required"`
	Loc_disease string             `json:"loc_disease,omitempty" bson:"loc_disease,omitempty" validate:"required"`
	Dif_diag    string             `json:"dif_diag,omitempty" bson:"dif_diag,omitempty" validate:"required"`
	Gender      string             `json:"gender,omitempty" bson:"gender,omitempty" validate:"required, eq=MALE|eq=FEMALE"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty" validate:"required"`
	Country     string             `json:"country,omitempty" bson:"country,omitempty" validate:"required"`
	State       string             `json:"state,omitempty" bson:"state,omitempty" validate:"required"`
	Comp_desc   string             `json:"comp_desc,omitempty" bson:"comp_desc,omitempty" validate:"required"`
	Fav_count   int                `json:"fav_count,omitempty" bson:"fav_count,omitempty"`
	Img_url     string             `json:"img_url,omitempty" bson:"img_url,omitempty" validate:"required"`
	Created_At  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_At  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
