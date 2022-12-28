package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AddressFrom string             `json:"address_from,omitempty" bson:"address_from,omitempty"`
	AddressTo   string             `json:"address_to,omitempty" bson:"address_to,omitempty"`
	Amount      float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
