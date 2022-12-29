package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AddressFrom string             `json:"from,omitempty" bson:"from,omitempty"`
	AddressTo   string             `json:"to,omitempty" bson:"to,omitempty"`
	Amount      float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type CreateTransactionRequest struct {
	AddressFrom string    `json:"from,omitempty" bson:"from,omitempty"`
	AddressTo   string    `json:"to,omitempty" bson:"to,omitempty"`
	Amount      float32   `json:"amount,omitempty" bson:"amount,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
