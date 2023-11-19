package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listing struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name"          json:"name"        validate:"required"`
	Make        string             `bson:"make"          json:"make"        validate:"required"`
	Model       string             `bson:"model"         json:"model"       validate:"required"`
	Description string             `bson:"description"   json:"description" validate:"required"`
	Picture     string             `bson:"picture"       json:"picture"     validate:"required"`
	Url         string             `bson:"url"           json:"url"         validate:"required"`
	Site        string             `bson:"site"          json:"site"        validate:"required"`
	Price       float64            `bson:"price"         json:"price"       validate:"required"`
	CreatedAt   time.Time          `bson:"created_at"    json:"created_at"  validate:"required"` // Time it was added to DB
}
