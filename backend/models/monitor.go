package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Monitor struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Keywords   string             `bson:"keywords"      json:"keywords"    validate:"required"`
	UserID     string             `bson:"user_id"       json:"user_id"     validate:"required"`
	PostalCode string             `bson:"postal_code"   json:"postal_code" validate:"required"`
	MaxPrice   float64            `bson:"max_price"     json:"max_price"`
	CreatedAt  time.Time          `bson:"created_at"    json:"created_at"  validate:"required"` // Time it was added to DB
}
