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
	UserId      primitive.ObjectID `json:"_userId,omitempty" bson:"_userId,omitempty"`
	Skin_diag   string             `json:"skin_diag,omitempty" bson:"skin_diag,omitempty"`
	Loc_disease string             `json:"loc_disease,omitempty" bson:"loc_disease,omitempty"`
	Dif_diag    string             `json:"dif_diag,omitempty" bson:"dif_diag,omitempty"`
	Gender      Gender             `json:"gender,omitempty" bson:"gender,omitempty"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty"`
	Country     string             `json:"country,omitempty" bson:"country,omitempty"`
	Comp_desc   string             `json:"comp_desc,omitempty" bson:"comp_desc,omitempty"`
	Fav_count   int                `json:"fav_count,omitempty" bson:"fav_count,omitempty"`
	Img_url     string             `json:"img_url,omitempty" bson:"img_url,omitempty"`
	Created_At  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
