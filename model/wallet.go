package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wallet struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Balance  float32            `json:"balance,omitempty" bson:"balance,omitempty"`
	CreateAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
