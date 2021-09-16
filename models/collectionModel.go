package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collections struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Img_Id     primitive.ObjectID `json:"_img_id,omitempty" bson:"_img_id,omitempty"`
	User_Id    primitive.ObjectID `json:"_user_id,omitempty" bson:"_user_id,omitempty"`
	Created_At time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
